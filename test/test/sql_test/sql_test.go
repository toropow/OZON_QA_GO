//go:build sql_test
// +build sql_test

package homework14

import (
	"context"
	"fmt"
	configProject "github.com/ozonmp/act-device-api/internal/config"
	"github.com/ozonmp/act-device-api/test/client_db"
	"github.com/ozonmp/act-device-api/test/config"
	"github.com/ozonmp/act-device-api/test/internal/steps"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"strconv"
	"testing"
)

func TestCRUDSql(t *testing.T) {
	cfg, _ := config.GetConfig()
	host := cfg.GrpcURI
	configPath := "../../../config.yml"
	if err := configProject.ReadConfigYML(configPath); err != nil {
		log.Fatal().Err(err).Msg("Err read config yaml")
	}

	cfgDb := configProject.GetConfigInstance()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfgDb.Database.Host,
		cfgDb.Database.Port,
		cfgDb.Database.User,
		cfgDb.Database.Password,
		cfgDb.Database.Name,
		cfgDb.Database.SslMode,
	)

	db, err := client_db.NewPostgres(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Database init err")
	}

	defer db.DB.Close()

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("grpc.Dial err")
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("conn.Close err")
		}
	}(conn)

	runner.Run(t, "create device - check database", func(t provider.T) {
		t.Epic("SQL test")
		//Arrange
		const platform = " "
		const userId = uint64(1)

		//Act
		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		if err != nil {
			log.Fatal().Err(err).Msg("Request error")
		}
		t.Require().NotNil(createDevicesV1Response.DeviceId, "Device Id is null")
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		data, err := db.ByDeviceId(context.Background(), int(createDevicesV1Response.DeviceId))
		if err != nil {
			log.Fatal().Err(err).Msg("Query error")
		}
		t.WithNewAttachment("Database response", allure.Text, []byte(strconv.Itoa(int(data.DeviceId))))

		//Assert
		t.Require().Equal(data.DeviceId, createDevicesV1Response.DeviceId, "Get %v, want %v", data.DeviceId, createDevicesV1Response.DeviceId)
		t.Require().Equal(data.Device.Platform, platform, "Get %v, want %v", data.Device.Platform, platform)
		t.Require().Equal(data.Device.UserID, userId, "Get %v, want %v", data.Device.UserID, userId)
	})

	runner.Run(t, "update device - check database", func(t provider.T) {
		t.Epic("SQL test")
		//Arrange
		const platform = "Linux"
		const userId = uint64(1)
		const newUserId = uint64(5)
		const updateDeviceExpected = true

		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response create device:", allure.Text, []byte(createDevicesV1Response.String()))

		//Act
		updateDevicesV1Response, err := steps.UpdateDevice(conn, newUserId, platform, createDevicesV1Response.DeviceId)
		t.Require().Equal(updateDevicesV1Response.Success, updateDeviceExpected, "Get: %v, want: %v", updateDevicesV1Response.Success, updateDeviceExpected)
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		data, err := db.ByDeviceId(context.Background(), int(createDevicesV1Response.DeviceId))
		if err != nil {
			log.Fatal().Err(err).Msg("Query error")
		}
		t.WithNewAttachment("Database response", allure.Text, []byte(strconv.Itoa(int(data.DeviceId))))

		//Assert
		t.Require().Equal(data.Device.UserID, userId, "Get: %v, want: %v", data.Device.UserID, userId)
		t.Require().Equal(data.Device.Platform, platform, "Get: %v, want: %v", data.Device.Platform, platform)
		t.Require().Equal(createDevicesV1Response.DeviceId, data.DeviceId, "Get: %v, want: %v", createDevicesV1Response.DeviceId, data.DeviceId)

	})

	runner.Run(t, "delete device - check database", func(t provider.T) {
		t.Epic("SQL test")
		//Arrange
		const platform = "Ubuntu"
		const userId = uint64(1)
		const foundDeviceExpected = true

		createDevicesV1Response, err := steps.CreateDevice(conn, platform, userId)
		t.WithNewAttachment("Response create device:", allure.Text, []byte(createDevicesV1Response.String()))

		//Act
		removeDevicesV1Response, err := steps.RemoveDevice(conn, createDevicesV1Response.DeviceId)
		t.Require().Equal(removeDevicesV1Response.Found, foundDeviceExpected, "Get %v, want %v", removeDevicesV1Response.Found, foundDeviceExpected)
		if err != nil {
			log.Fatal().Err(err).Msg("Error remove device")
		}
		t.WithNewAttachment("Response", allure.Text, []byte(createDevicesV1Response.String()))

		data, err := db.ByDeviceId(context.Background(), int(createDevicesV1Response.DeviceId))
		if err != nil {
			log.Fatal().Err(err).Msg("Query error")
		}
		t.WithNewAttachment("Database response", allure.Text, []byte(strconv.Itoa(int(data.DeviceId))))

		//Assert
		t.Require().Equal(data.Device.UserID, userId, "Get: %v, want: %v", data.Device.UserID, userId)
		t.Require().Equal(data.Device.Platform, platform, "Get: %v, want: %v", data.Device.Platform, platform)
		t.Require().Equal(data.DeviceId, createDevicesV1Response.DeviceId, "Get: %v, want: %v", data.DeviceId, createDevicesV1Response.DeviceId)
	})

}
