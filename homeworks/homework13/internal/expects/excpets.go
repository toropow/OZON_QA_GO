package expects

import (
	routeclient "github.com/ozonmp/act-device-api/homeworks/homework13/client"
	pb "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"net/http"
)

func CheckCreateDeviceError(t provider.T, createDevicesV1Response *pb.CreateDeviceV1Response, err error, errorMsg string) {
	t.Require().Nil(createDevicesV1Response)
	t.Require().Error(err, errorMsg, "Error should be: %v, got: %v", errorMsg, err)
}

func CheckCreateDeviceSuccessful(t provider.T, err error, createDevicesV1Response *pb.CreateDeviceV1Response) {
	t.Require().NoError(err)
	t.Require().NotNil(createDevicesV1Response)
	t.Require().NotEmpty(createDevicesV1Response.DeviceId)
}

func CheckDescribeDeviceError(t provider.T, describeDevicesV1Response *pb.DescribeDeviceV1Response, err error, errorMsg string) {
	t.Require().Nil(describeDevicesV1Response)
	t.Require().Error(err, errorMsg, "want: %v, get: %v", errorMsg, err)
}

func CheckDescribeDeviceSuccessful(t provider.T, err error, describeDevicesV1Response *pb.DescribeDeviceV1Response, platform string, userId uint64) {
	t.Require().NoError(err)
	t.Require().Equal(describeDevicesV1Response.Value.Platform, platform,
		"want: %v get: %v", platform, describeDevicesV1Response.Value.Platform)
	t.Require().Equal(describeDevicesV1Response.Value.UserId, userId,
		"want: %v get: %v", userId, describeDevicesV1Response.Value.UserId)
}

func CheckUpdateResponseError(t provider.T, err error, code *http.Response, response routeclient.UpdateDeviceResponse) {
	t.Require().Nil(err, "Error while executing the method: %s", err)
	t.Require().Equal(code.StatusCode, 400)
	t.Require().NotNil(response, "Error while executing the method: %s", response)
}

func CheckUpdateResponse(t provider.T, err error, code *http.Response, responseUpdateDevice routeclient.UpdateDeviceResponse, expectedCode int, expectedFound bool) {
	t.Require().Nil(err, "Error while executing the method: %s", err)
	t.Require().Equal(code.StatusCode, expectedCode)
	t.Require().Equal(responseUpdateDevice.Success, expectedFound)
}

func CheckRemoveResponse(t provider.T, err error, code *http.Response, responseRemoveDevice routeclient.RemoveDeviceResponse, expectedCode int, expectedFound bool) {
	t.Require().Nil(err, "Error while executing the method: %s", err)
	t.Require().Equal(code.StatusCode, expectedCode)
	t.Require().Equal(responseRemoveDevice.Found, expectedFound)
}
