package homework13

import (
	"github.com/ozonmp/act-device-api/homeworks/homework13/config"
	. "github.com/ozonmp/act-device-api/homeworks/homework13/internal/expects"
	"github.com/ozonmp/act-device-api/homeworks/homework13/internal/steps"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
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

	runner.Run(t, "describe device - successful", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		//Prepare device
		const platform = "Mac"
		const userId = uint64(3)

		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response create device:", allure.Text, []byte(createDevicesV1Response.String()))

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, createDevicesV1Response.DeviceId)
		t.WithNewAttachment("Response", allure.Text, []byte(describeDevicesV1Response.String()))

		//Assert
		CheckDescribeDeviceSuccessful(t, err, describeDevicesV1Response, platform, userId)
	})

	runner.Run(t, "describe device - not exists", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		const errorMsg = "rpc error: code = NotFound desc = device not found"

		//Prepare device
		const platform = "Mac"
		const userId = uint64(3)
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response create device", allure.Text, []byte(createDevicesV1Response.String()))

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, createDevicesV1Response.DeviceId+1)
		t.WithNewAttachment("Response", allure.Text, []byte(describeDevicesV1Response.String()))

		//Assert
		CheckDescribeDeviceError(t, describeDevicesV1Response, err, errorMsg)
	})

	runner.Run(t, "describe device - zero device", func(t provider.T) {
		t.Epic("GRPS test")
		//Arrange
		const errorMsg = "rpc error: code = InvalidArgument desc = invalid DescribeDeviceV1Request.DeviceId: value must be greater than 0"
		const userId = uint64(0)

		//Act
		describeDevicesV1Response, err := steps.DescribeDevice(conn, userId)
		t.WithNewAttachment("Response", allure.Text, []byte(describeDevicesV1Response.String()))

		//Assert
		CheckDescribeDeviceError(t, describeDevicesV1Response, err, errorMsg)
	})

}
