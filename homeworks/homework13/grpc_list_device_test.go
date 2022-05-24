package homework13

import (
	"github.com/ozonmp/act-device-api/homeworks/homework13/config"
	"github.com/ozonmp/act-device-api/homeworks/homework13/internal/steps"
	act_device_api "github.com/ozonmp/act-device-api/pkg/act-device-api"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"google.golang.org/grpc"
	"testing"
)

func TestListDevice(t *testing.T) {
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
	runner.Run(t, "list device - successful", func(s provider.T) {
		s.Epic("GRPS test")
		//Arrange
		type ListDeviceTest struct {
			name    string
			page    uint64
			perPage uint64
		}

		tests := []ListDeviceTest{
			{name: "First page", page: 1, perPage: 2},
			{name: "Second page", page: 2, perPage: 2},
			{name: "Third page", page: 3, perPage: 1},
		}

		for _, tt := range tests {
			runner.Run(t, tt.name, func(t provider.T) {
				//Act
				listDevicesV1Response, err := steps.ListDevice(conn, tt.page, tt.perPage)
				t.WithNewAttachment("Response", allure.Text, []byte(listDevicesV1Response.String()))

				//Assert
				CheckListDeviceSuccessful(t, err, listDevicesV1Response, tt.perPage)
			})
		}
	})

}
func CheckListDeviceSuccessful(t provider.T, err error, listDevicesV1Response *act_device_api.ListDevicesV1Response, perPage uint64) {
	t.Require().NoError(err)
	t.Require().NotNil(listDevicesV1Response)
	t.Require().Equal(len(listDevicesV1Response.Items), int(perPage))
	for _, value := range listDevicesV1Response.Items {
		t.Require().NotEmpty(value.UserId)
		t.Log(value.UserId)
	}
}
