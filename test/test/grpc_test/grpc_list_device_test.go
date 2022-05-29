package grpc_test

import (
	"github.com/ozonmp/act-device-api/test/config"
	. "github.com/ozonmp/act-device-api/test/internal/expects"
	"github.com/ozonmp/act-device-api/test/internal/steps"
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
	t.Run("list device - successful", func(t *testing.T) {
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
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				//Act
				listDevicesV1Response, err := steps.ListDevice(conn, tt.page, tt.perPage)

				//Assert
				CheckListDeviceSuccessful(t, err, listDevicesV1Response, tt.perPage)
			})
		}
	})

}
