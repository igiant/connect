package connect

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDomainRequests(t *testing.T) {
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
	domains, num, err := conn.DomainsGet(SearchQuery{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(num)
	if len(domains) == 0 || domains[0].Id == "" {
		t.Error("Domains information not received")
	}
	dSettings, err := conn.DomainsGetSettings()
	if err != nil {
		t.Error(err)
	}
	if dSettings.ServerId == "" {
		t.Error("Domains settings information not received")
	}
	err = conn.Logout()
	if err != nil {
		t.Error(err)
	}
}
