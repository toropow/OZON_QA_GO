//go:build http_test
// +build http_test

package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	routeclient "github.com/ozonmp/act-device-api/test/client"
	. "github.com/ozonmp/act-device-api/test/internal/expects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

type ListOfItemsResponse struct {
	Items []struct {
		ID        string     `json:"id"`
		Platform  string     `json:"platform"`
		UserID    string     `json:"userId"`
		EnteredAt *time.Time `json:"enteredAt"`
	} `json:"items"`
}

type ItemRequest struct {
	Platform string `json:"platform"`
	UserID   string `json:"userId"`
}

func Test_HttpServer_List(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	client := routeclient.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
	ctx := context.TODO()

	t.Run("GET on list return 200", func(t *testing.T) {
		response, err := http.Get("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=1")
		if err != nil {
			panic(err)
		}

		if response.StatusCode != http.StatusOK {
			t.Errorf("Got %v, but want %v", response.StatusCode, http.StatusOK)
		}
	})

	t.Run("GET on list return devices list", func(t *testing.T) {
		countOfItems := 8
		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=%d", countOfItems))
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		list := new(ListOfItemsResponse)
		err = json.Unmarshal(data, &list)
		if err != nil {
			panic(err)
		}

		if len(list.Items) != countOfItems {
			t.Errorf("Want %d, get %d items", countOfItems, len(list.Items))
		}
	})

	t.Run("GET on list return devices list if zeroed", func(t *testing.T) {
		countOfItems := 0
		response, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/api/v1/devices?page=1&perPage=%d", countOfItems))
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		list := new(ListOfItemsResponse)
		err = json.Unmarshal(data, &list)
		if err != nil {
			panic(err)
		}

		if len(list.Items) != countOfItems {
			t.Errorf("Want %d, get %d items", countOfItems, len(list.Items))
		}
	})

	t.Run("POST on creating device", func(t *testing.T) {
		data := []byte(`{"platform": "Android", "userId": "123456"}`)
		r := bytes.NewReader(data)
		contentType := "application/json"

		_, err := http.Post("http://127.0.0.1:8080/api/v1/devices", contentType, r)
		if err != nil {
			panic(err)
		}

		payload := ItemRequest{Platform: "Android", UserID: "123456"}
		payloadJSON, _ := json.Marshal(payload)

		_, err = http.Post("http://127.0.0.1:8080/api/v1/devices", contentType, bytes.NewBuffer(payloadJSON))
		if err != nil {
			panic(err)
		}
	})

	t.Run("Why do we need a client?", func(t *testing.T) {
		t.Skip()
		// nc -lp 9090
		_, err := http.Get("http://127.0.0.1:9090")
		if err != nil {
			panic(err)
		}
	})

	t.Run("POST with client", func(t *testing.T) {
		// arrange
		payload := ItemRequest{Platform: "Android", UserID: "666"}
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(payload)
		if err != nil {
			panic(err)
		}
		client := routeclient.NewHTTPClient("http://127.0.0.1:8080", 5, 1*time.Second)
		// action
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/v1/devices", b)
		if err != nil {
			panic(err)
		}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				t.Log(err)
			}
		}(res.Body)
		//assert

		if res.StatusCode != http.StatusOK {
			t.Errorf("Got %v, but want %v", res.StatusCode, http.StatusOK)
		}
		data, _ := ioutil.ReadAll(res.Body)
		if len(data) != 0 {
			t.Log(string(data))
		}
	})

	t.Run("Create device via client API", func(t *testing.T) {
		const platform = "Ubuntu"
		const userId = "701"

		device := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   userId,
		}
		id, _, _ := client.CreateDevice(ctx, device)
		t.Logf("New device is %d", id.DeviceId)

		assert.GreaterOrEqual(t, id.DeviceId, 0)
	})

	t.Run("List devices via client API", func(t *testing.T) {
		const page = "1"
		const perPage = "100"
		opts := url.Values{}
		opts.Add("page", page)
		opts.Add("perPage", perPage)
		items, _, _ := client.ListDevices(ctx, opts)
		assert.GreaterOrEqual(t, len(items.Items), 1)
	})

	t.Run("Remove devices via client API", func(t *testing.T) {
		//Arrange
		const platform = "Ubuntu"
		const userId = "701"

		//Create device for remove
		device := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   userId,
		}

		id, _, _ := client.CreateDevice(ctx, device)
		deviceId := strconv.Itoa(id.DeviceId)

		//Act
		responseRemoveDevice, code, err := client.RemoveDevice(ctx, deviceId)

		//Assert
		CheckRemoveResponse(t, err, code, responseRemoveDevice, 200, true)
	})

	t.Run("Remove devices with empty Id via client API", func(t *testing.T) {
		//Arrange
		const deviceId = ""

		//Act
		responseRemoveDevice, code, err := client.RemoveDevice(ctx, deviceId)

		//Assert
		CheckRemoveResponse(t, err, code, responseRemoveDevice, 400, false)
	})

	t.Run("Remove isn't exist devices via client API", func(t *testing.T) {
		//Arrange
		const deviceId = "999"

		//Act
		responseRemoveDevice, code, err := client.RemoveDevice(ctx, deviceId)

		//Assert
		CheckRemoveResponse(t, err, code, responseRemoveDevice, 200, false)
	})

	t.Run("Update devices via client API", func(t *testing.T) {
		//Arrange
		const platform = "Ubuntu"
		const userId = "701"
		const newUserId = "1"

		//Create device for update
		device := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   userId,
		}

		id, _, errDevice := client.CreateDevice(ctx, device)
		require.Nilf(t, errDevice, "Error while executing the method: %s", errDevice)
		deviceId := strconv.Itoa(id.DeviceId)

		deviceUpd := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   newUserId,
		}

		//Act
		responseUpdateDevice, code, err := client.UpdateDevice(ctx, deviceId, deviceUpd)

		//Assert
		CheckUpdateResponse(t, err, code, responseUpdateDevice, 200, true)
	})

	t.Run("Update isn't exists devices via client API", func(t *testing.T) {
		//Arrange
		const platform = "Ubuntu"
		const userId = "701"
		const newUserId = "1"

		//Create device for update
		device := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   userId,
		}

		id, _, errDevice := client.CreateDevice(ctx, device)
		require.Nilf(t, errDevice, "Error while executing the method: %s", errDevice)
		deviceId := strconv.Itoa(id.DeviceId + 1)

		deviceUpd := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   newUserId,
		}

		//Act
		responseUpdateDevice, code, err := client.UpdateDevice(ctx, deviceId, deviceUpd)

		//Assert
		CheckUpdateResponse(t, err, code, responseUpdateDevice, 200, false)
	})

	t.Run("Update devices only Platform via client API", func(t *testing.T) {
		//Arrange
		const platform = "Ubuntu"
		const userId = "701"
		const newUserId = ""

		//Create device for update
		device := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   userId,
		}
		id, _, errDevice := client.CreateDevice(ctx, device)
		require.Nilf(t, errDevice, "Error while executing the method: %s", errDevice)

		deviceId := strconv.Itoa(id.DeviceId)

		deviceUpd := routeclient.DeviceRequest{
			Platform: platform,
			UserId:   newUserId,
		}

		//Act
		response, code, err := client.UpdateDevice(ctx, deviceId, deviceUpd)

		//Assert
		CheckUpdateResponseError(t, err, code, response)
	})

}
