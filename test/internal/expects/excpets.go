package expects

import (
	"net/http"
	"testing"

	pb "github.com/ozonmp/act-device-api/pkg/act-device-api"
	routeClient "github.com/ozonmp/act-device-api/test/client"
	"github.com/stretchr/testify/require"
)

func CheckCreateDeviceError(t *testing.T, createDevicesV1Response *pb.CreateDeviceV1Response, err error, errorMsg string) {
	t.Helper()
	require.Nil(t, createDevicesV1Response, "Response not nil: %v", createDevicesV1Response)
	require.EqualErrorf(t, err, errorMsg, "Error should be: %v, got: %v", errorMsg, err)
}

func CheckCreateDeviceSuccessful(t *testing.T, err error, createDevicesV1Response *pb.CreateDeviceV1Response) {
	t.Helper()
	require.NoError(t, err, "Error: %v", err)
	require.NotNil(t, createDevicesV1Response, "Response is nil: %v", createDevicesV1Response)
	require.NotEmpty(t, createDevicesV1Response.DeviceId, "DeviceId is empty")
}

func CheckDescribeDeviceError(t *testing.T, describeDevicesV1Response *pb.DescribeDeviceV1Response, err error, errorMsg string) {
	t.Helper()
	require.Nil(t, describeDevicesV1Response, "Response not nil: %v", describeDevicesV1Response)
	require.EqualErrorf(t, err, errorMsg, "want: %v, get: %v", errorMsg, err)
}

func CheckDescribeDeviceSuccessful(t *testing.T, err error, describeDevicesV1Response *pb.DescribeDeviceV1Response, platform string, userId uint64) {
	t.Helper()
	require.NoError(t, err, "Error: %v", err)
	require.Equal(t, describeDevicesV1Response.Value.Platform, platform,
		"want: %v get: %v", platform, describeDevicesV1Response.Value.Platform)
	require.Equal(t, describeDevicesV1Response.Value.UserId, userId,
		"want: %v get: %v", userId, describeDevicesV1Response.Value.UserId)
}

func CheckUpdateResponseError(t *testing.T, err error, code *http.Response, response routeClient.UpdateDeviceResponse) {
	t.Helper()
	require.Nilf(t, err, "Error while executing the method: %s", err)
	require.Equal(t, code.StatusCode, 400, "Get: %v, want: %v", code.StatusCode, 400)
	require.NotNil(t, response, "Error while executing the method: %s", response)
}

func CheckUpdateResponse(t *testing.T, err error, code *http.Response, responseUpdateDevice routeClient.UpdateDeviceResponse, expectedCode int, expectedFound bool) {
	t.Helper()
	require.Nilf(t, err, "Error while executing the method: %s", err)
	require.Equal(t, code.StatusCode, expectedCode, "Get: %v, want: %v", code.StatusCode, expectedCode)
	require.Equal(t, responseUpdateDevice.Success, expectedFound, "Get: %v, want: %v", responseUpdateDevice.Success, expectedFound)
}

func CheckRemoveResponse(t *testing.T, err error, code *http.Response, responseRemoveDevice routeClient.RemoveDeviceResponse, expectedCode int, expectedFound bool) {
	t.Helper()
	require.Nilf(t, err, "Error while executing the method: %s", err)
	require.Equal(t, code.StatusCode, expectedCode, "Get: %v, want: %v", code.StatusCode, expectedCode)
	require.Equal(t, responseRemoveDevice.Found, expectedFound, "Get: %v, want: %v", responseRemoveDevice.Found, expectedFound)
}

func CheckListDeviceSuccessful(t *testing.T, err error, listDevicesV1Response *pb.ListDevicesV1Response, perPage uint64) {
	t.Helper()
	require.NoError(t, err, "Error: %v", err)
	require.NotNil(t, listDevicesV1Response, "Response is empty")
	require.Equal(t, len(listDevicesV1Response.Items), int(perPage), "Get: %v, want: %v", len(listDevicesV1Response.Items), int(perPage))
	for _, value := range listDevicesV1Response.Items {
		require.NotEmpty(t, value.UserId, "UserId is empty")
		t.Log(value.UserId)
	}
}
