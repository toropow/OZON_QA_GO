package homework11

import (
	"context"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateDevice(t *testing.T) {
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

	t.Run("create device - successful", func(t *testing.T) {
		//Arrange
		const platform = "Linux"
		const userId = uint64(1)
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		request := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userId,
		}

		//Act
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, request)

		//Assert
		require.NoErrorf(t, err, "Error while CreateDevicesV1(): %s", err)
		require.NotNil(t, createDevicesV1Response, "Response is empty")
		require.NotEmpty(t, createDevicesV1Response.DeviceId, "Empty value for DeviceID")
	})

	t.Run("create device - platform is one space", func(t *testing.T) {
		//Arrange
		const platform = " "
		const userId = uint64(1)
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		request := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userId,
		}

		//Act
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, request)

		//Assert
		require.NoErrorf(t, err, "Error while CreateDevicesV1(): %s", err)
		require.NotNil(t, createDevicesV1Response, "Response is empty")
		require.NotEmpty(t, createDevicesV1Response.DeviceId, "Empty value for DeviceID")
	})

	t.Run("create device - UserId is 0", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.UserId:" +
			" value must be greater than 0"
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		request := &act_device_api.CreateDeviceV1Request{
			Platform: "Mint",
			UserId:   0,
		}

		//Act
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, request)

		//Assert
		require.Nil(t, createDevicesV1Response, "Response is not empty")
		require.EqualErrorf(t, err, errorMsg, "Error should be: %v, got: %v", errorMsg, err)
	})

	t.Run("create device - Platform is empty", func(t *testing.T) {
		//Arrange
		const platform = ""
		const userId = uint64(10)
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.Platform: value length" +
			" must be at least 1 runes"
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
		request := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userId,
		}

		//Act
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, request)

		//Assert
		require.Nil(t, createDevicesV1Response, "Response is not empty")
		require.EqualErrorf(t, err, errorMsg, "Error should be: %v, got: %v", errorMsg, err)
	})

}
