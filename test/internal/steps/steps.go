package steps

import (
	"context"
	"github.com/ozonmp/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
)

func CreateDevice(conn *grpc.ClientConn, platform string, userid uint64) (*act_device_api.CreateDeviceV1Response, error) {
	deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
	request := &act_device_api.CreateDeviceV1Request{
		Platform: platform,
		UserId:   userid,
	}
	createDevicesV1Response, err := deviceApiClient.CreateDeviceV1(context.Background(), request)

	return createDevicesV1Response, err
}

func DescribeDevice(conn *grpc.ClientConn, deviceId uint64) (*act_device_api.DescribeDeviceV1Response, error) {
	deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
	request := &act_device_api.DescribeDeviceV1Request{DeviceId: deviceId}
	describeDevicesV1Response, err := deviceApiClient.DescribeDeviceV1(context.Background(), request)

	return describeDevicesV1Response, err
}

func ListDevice(conn *grpc.ClientConn, page uint64, perPage uint64) (*act_device_api.ListDevicesV1Response, error) {
	deviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
	request := &act_device_api.ListDevicesV1Request{
		PerPage: perPage,
		Page:    page,
	}
	listDevicesV1Response, err := deviceApiClient.ListDevicesV1(context.Background(), request)

	return listDevicesV1Response, err
}
