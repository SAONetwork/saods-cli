package config

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

type Config struct {
	AppId  string `json:"appId"`
	ApiKey string `json:"apiKey"`
	ServiceUrl string `json:"serviceUrl"`
}

func GetConfig() (Config, error) {
	var cliConfig Config
	hDir, err := homedir.Dir()
	if err != nil {
		return cliConfig, err
	}
	cfgPath := hDir + "/.saods/ds-cli.json"

	f, err := os.Open(cfgPath)
	if err != nil {
		return cliConfig, err
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return cliConfig, err
	}

	json.Unmarshal(byteValue, &cliConfig)

	return cliConfig, nil
}

func SetConfig(config Config) error {
	hDir, err := homedir.Dir()
	if err != nil {
		return err
	}
	cfgPath := hDir + "/.saods/ds-cli.json"

	out, err := os.Create(cfgPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer out.Close()

 	configContent, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_, err = out.WriteString(string(configContent))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
