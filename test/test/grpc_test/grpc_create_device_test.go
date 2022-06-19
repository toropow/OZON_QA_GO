//go:build grpc_test
// +build grpc_test

package grpc_test

import (
	"github.com/ozonmp/act-device-api/test/config"
	. "github.com/ozonmp/act-device-api/test/internal/expects"
	"github.com/ozonmp/act-device-api/test/internal/steps"

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

	t.Run("create device - successful", func(t *testing.T) {
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

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				//Act
				createDevicesV1Response, err := steps.CreateDevice(conn, tt.platform, tt.userid)

				//Assert
				CheckCreateDeviceSuccessful(t, err, createDevicesV1Response)
			})
		}

	})

	t.Run("create device - platform is one space", func(t *testing.T) {
		//Arrange
		const platform = " "
		const userId = uint64(1)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)

		//Assert
		CheckCreateDeviceSuccessful(t, err, createDevicesV1Response)
	})

	t.Run("create device - UserId is 0", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.UserId:" +
			" value must be greater than 0"
		const platform = "Mint"
		const userId = uint64(0)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)

		//Assert
		CheckCreateDeviceError(t, createDevicesV1Response, err, errorMsg)
	})

	t.Run("create device - Platform is empty", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.Platform: value length" +
			" must be at least 1 runes"
		const platform = ""
		const userId = uint64(10)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)

		//Assert
		CheckCreateDeviceError(t, createDevicesV1Response, err, errorMsg)
	})

}
