package disk

import (
	"encoding/json"
	"github.com/just1689/silent-k8s-cluster/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func GenerateRouterConfigToFile() {
	c := model.RouterConfig{
		Address:  "192.168.0.1:8728",
		Username: "admin",
		Password: "password",
	}
	b, err := json.Marshal(c)
	if err != nil {
		logrus.Panicln(err)
	}
	err = ioutil.WriteFile("router-config.json", b, 0644)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.Println("Complete!")
	os.Exit(0)

}

func LoadRouterConfig() model.RouterConfig {
	result := model.RouterConfig{}
	b, err := ioutil.ReadFile("router-config.json")
	if err != nil {
		logrus.Panicln(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Panicln(err)
	}
	return result
}
