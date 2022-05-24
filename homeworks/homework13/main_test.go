package homework13

import (
	_ "fmt"
	"github.com/ozonmp/act-device-api/homeworks/homework13/config"
	"github.com/ozonmp/act-device-api/homeworks/homework13/internal/helpers"

	"net/url"
	"testing"
)

func TestMain(m *testing.M) {
	cfg, _ := config.GetConfig()

	helpers.IsAlive(url.URL{
		Scheme: "http",
		Host:   cfg.ApiHost,
		Path:   cfg.LivecheckURI,
	})
	m.Run()
}
