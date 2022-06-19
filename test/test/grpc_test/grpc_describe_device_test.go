//go:build grpc_test
// +build grpc_test

package grpc_test

import (
	"github.com/ozonmp/act-device-api/test/config"
	. "github.com/ozonmp/act-device-api/test/internal/expects"
	"github.com/ozonmp/act-device-api/test/internal/steps"
	"google.golang.org/grpc"
	"testing"
)

func TestDescribeDevice(t *testing.T) {
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

	t.Run("describe device - successful", func(t *testing.T) {
		//Arrange
		//Prepare device
		const platform = "Mac"
		const userId = uint64(3)
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, createDevicesV1Response.DeviceId)

		//Assert
		CheckDescribeDeviceSuccessful(t, err, describeDevicesV1Response, platform, userId)
	})

	t.Run("describe device - not exists", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = NotFound desc = device not found"

		//Prepare device
		const platform = "Mac"
		const userId = uint64(3)
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, createDevicesV1Response.DeviceId+1)

		//Assert
		CheckDescribeDeviceError(t, describeDevicesV1Response, err, errorMsg)
	})

	t.Run("describe device - zero device", func(t *testing.T) {
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid DescribeDeviceV1Request.DeviceId: value must be greater than 0"
		const userId = uint64(0)

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, userId)

		//Assert
		CheckDescribeDeviceError(t, describeDevicesV1Response, err, errorMsg)
	})

}
