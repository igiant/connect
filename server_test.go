package connect

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestServerRequests(t *testing.T) {
	param := struct {
		Server   string `yaml:"server"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}{}
	file, err := os.ReadFile("secret.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	err = yaml.Unmarshal(file, &param)
	if err != nil {
		t.Error(err)
		return
	}
	conf := NewConfig(param.Server)
	app := &ApiApplication{
		Name:    "MyApp",
		Vendor:  "Me",
		Version: "v0.0.1",
	}
	conn, err := conf.NewConnection()
	if err != nil {
		t.Error(err)
		return
	}
	err = conn.Login(param.User, param.Password, app)
	if err != nil {
		t.Error(err)
		return
	}
	version, err := conn.ServerGetVersion()
	if err != nil {
		t.Error(err)
	}
	if !strings.HasPrefix(version.Product, "Kerio Connect") {
		t.Error("server version not received")
	}
	info, err := conn.ServerGetProductInfo()
	if err != nil {
		t.Error(err)
	}
	if info.ProductName != "Kerio Connect" {
		t.Error("product info not received")
	}
	sessionList, num, err := conn.ServerGetWebSessions(SearchQuery{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(num)
	if len(sessionList) == 0 || sessionList[0].Id == "" {
		t.Error("domains information not received")
	}
	dirs, err := conn.ServerGetDirs("/")
	if err != nil {
		t.Error(err)
	}
	if len(dirs) == 0 || dirs[0].Name == "" {
		t.Error("dirs information not received")
	}
	extensions, err := conn.ServerGetLicenseExtensionsList()
	if err != nil {
		t.Error(err)
	}
	if len(extensions) == 0 {
		t.Error("extension list not received")
	}
	addresses, err := conn.ServerGetServerIpAddresses()
	if err != nil {
		t.Error(err)
	}
	if len(addresses) == 0 {
		t.Error("addresses list not received")
	}
	timeInfo, err := conn.ServerGetServerTime()
	if err != nil {
		t.Error(err)
	}
	if timeInfo.CurrentTime == 0 {
		t.Error("server time info not received")
	}
	_, err = conn.ServerGetClientStatistics()
	if err != nil {
		t.Error(err)
	}
	result, err := conn.ServerPathExists("/", Credentials{"", ""})
	if err != nil {
		t.Error(err)
	}
	if result == "" {
		t.Error("status path not received")
	}
	admin, err := conn.ServerGetRemoteAdministration()
	if err != nil {
		t.Error(err)
	}
	if admin != nil {
		err = conn.ServerValidateRemoteAdministration(*admin)
		if err != nil {
			t.Error(err)
		}
	}
	err = conn.Logout()
	if err != nil {
		t.Error(err)
	}
}
