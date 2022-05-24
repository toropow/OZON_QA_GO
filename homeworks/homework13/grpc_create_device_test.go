package homework13

import (
	"github.com/ozonmp/act-device-api/homeworks/homework13/config"
	. "github.com/ozonmp/act-device-api/homeworks/homework13/internal/expects"
	"github.com/ozonmp/act-device-api/homeworks/homework13/internal/steps"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"google.golang.org/grpc"
	"math/rand"
	"testing"
)

func TestCreateDevice(t *testing.T) {
	cfg, _ := config.GetConfig()
	host := cfg.GrpcURI
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

	runner.Run(t, "create device - successful", func(s provider.T) {
		s.Epic("GRPS test")
		//Arrange
		type CreateDeviceTest struct {
			name     string
			platform string
			userid   uint64
		}

		tests := []CreateDeviceTest{
			{name: "Linux platform", platform: "Ubuntu", userid: uint64(rand.Int())},
			{name: "Windows platform", platform: "Windows10", userid: uint64(rand.Int())},
			{name: "Mac platform", platform: "Mac", userid: uint64(rand.Int())},
		}

		for _, tt := range tests {

			runner.Run(t, tt.name, func(s provider.T) {
				//Act
				createDevicesV1Response, err := steps.CreateDevice(conn, tt.platform, tt.userid)

				s.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

				//Assert
				CheckCreateDeviceSuccessful(s, err, createDevicesV1Response)
			})
		}

	})

	runner.Run(t, "create device - platform is one space", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		const platform = " "
		const userId = uint64(1)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		//Assert
		CheckCreateDeviceSuccessful(t, err, createDevicesV1Response)
	})

	runner.Run(t, "create device - UserId is 0", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.UserId:" +
			" value must be greater than 0"
		const platform = "Mint 10"
		const userId = uint64(0)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		//Assert
		CheckCreateDeviceError(t, createDevicesV1Response, err, errorMsg)
	})

	runner.Run(t, "create device - Platform is empty", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.Platform: value length" +
			" must be at least 1 runes"
		const platform = ""
		const userId = uint64(10)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		//Assert
		CheckCreateDeviceError(t, createDevicesV1Response, err, errorMsg)
	})

}
