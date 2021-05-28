package connect

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestAliasesRequests(t *testing.T) {
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
	app := &Application{
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
	domains, err := conn.DomainsGet(NewSearchQuery(nil, nil, 0, -1, nil))
	if err != nil {
		t.Error(err)
	}
	if len(domains) > 0 {
		_, err := conn.AliasesGet(domains[0].ID, NewSearchQuery(nil, nil, 0, -1, nil))
		if err != nil {
			t.Error(err)
		}

	}
	err = conn.Logout()
	if err != nil {
		t.Error(err)
	}
}
