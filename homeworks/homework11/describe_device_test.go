package homework11

import (
	"context"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

func TestDescribeDevice(t *testing.T) {
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

	t.Run("describe device - successful", func(t *testing.T) {
		//Arrange
		const platform = "Mac"
		const userId = uint64(3)
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)

		//Prepare device
		requestDevice := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userId,
		}
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, requestDevice)
		request := &act_device_api.DescribeDeviceV1Request{DeviceId: createDevicesV1Response.DeviceId}

		//Act
		describeDevicesV1Response, err := deviceApiClient.DescribeDeviceV1(ctx, request)

		//Assert
		require.NoErrorf(t, err, "Error while DescribeDevicesV1(): %s", err)
		require.Equal(t, describeDevicesV1Response.Value.Platform, requestDevice.Platform,
			"want: %v get: %v", requestDevice.Platform, describeDevicesV1Response.Value.Platform)
		require.Equal(t, describeDevicesV1Response.Value.UserId, requestDevice.UserId,
			"want: %v get: %v", requestDevice.UserId, describeDevicesV1Response.Value.UserId)
	})

	t.Run("describe device - not exists", func(t *testing.T) {
		//Arrange
		const platform = "Mac"
		const userId = uint64(3)
		const errorMsg = "rpc error: code = NotFound desc = device not found"
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)

		//Prepare device
		requestDevice := &act_device_api.CreateDeviceV1Request{
			Platform: platform,
			UserId:   userId,
		}
		createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(ctx, requestDevice)

		request := &act_device_api.DescribeDeviceV1Request{DeviceId: createDevicesV1Response.DeviceId + 1}

		//Act
		describeDevicesV1Response, err := deviceApiClient.DescribeDeviceV1(ctx, request)

		//Assert
		require.Nil(t, describeDevicesV1Response, "Response is not empty")
		require.EqualErrorf(t, err, errorMsg, "want: %v, get: %v", errorMsg, err)
	})

	t.Run("describe device - zero device", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid DescribeDeviceV1Request.DeviceId: value must be greater than 0"
		deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)

		request := &act_device_api.DescribeDeviceV1Request{DeviceId: uint64(0)}

		//Act
		describeDevicesV1Response, err := deviceApiClient.DescribeDeviceV1(ctx, request)

		//Assert
		require.Nil(t, describeDevicesV1Response, "Response is not empty")
		require.EqualErrorf(t, err, errorMsg, "want: %v, get: %v", errorMsg, err)
	})

}
