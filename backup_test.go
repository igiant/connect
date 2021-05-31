package connect

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestBackupRequests(t *testing.T) {
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
	status, err := conn.BackupGetStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", status)
	err = conn.Logout()
	if err != nil {
		t.Error(err)
	}
}
