package homework11

import (
	"context"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestDevice(t *testing.T) {
	ctx := context.Background()
	host := "localhost:8082"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err:%v", err)
		}
	}(conn)
	t.Run("empty result", func(t *testing.T) {

		// grpc клиент act_device_api
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		request := &act_device_api.ListDevicesV1Request{}

		listDevicesV1Response, err := deviceApiClient.ListDevicesV1(ctx, request)
		t.Logf("status.Code: %v", status.Code(err).String())
		assert.Equal(t, codes.OK.String(), status.Code(err).String())
		require.NoErrorf(t, err, "Error while ListDevicesV1(): %s", err)
		require.NotNil(t, listDevicesV1Response)
		t.Logf("listDevicesV1Response: %v", listDevicesV1Response)
		// где структура listDevicesV1Response
		assert.Emptyf(t, listDevicesV1Response.GetItems(), "listDevicesV1Response.Items - не пустой")

	})
	t.Run("first grpc test", func(t *testing.T) {

		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		pages := uint64(3)
		request := &act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: pages,
		}

		listDevicesV1Response, err := deviceApiClient.ListDevicesV1(ctx, request)
		require.NoErrorf(t, err, "Error while ListDevicesV1(): %s", err)
		require.NotNil(t, listDevicesV1Response, "Response is empty")
		assert.Equal(t, len(listDevicesV1Response.Items), int(pages))
		for _, value := range listDevicesV1Response.Items {
			require.NotEmpty(t, value.Platform, "Platform is empty")
			t.Log(value.Platform)
		}

	})

	t.Run("second page - grpc test", func(t *testing.T) {

		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		pages := uint64(1)
		request := &act_device_api.ListDevicesV1Request{
			Page:    2,
			PerPage: pages,
		}

		listDevicesV1Response, err := deviceApiClient.ListDevicesV1(ctx, request)

		require.NoErrorf(t, err, "Error while ListDevicesV1(): %s", err)
		require.NotNil(t, listDevicesV1Response, "Response is empty")
		require.Equal(t, len(listDevicesV1Response.Items), int(pages))
		for _, value := range listDevicesV1Response.Items {
			require.NotEmpty(t, value.UserId, "UserId is empty")
			t.Log(value.UserId)
		}

	})

}
