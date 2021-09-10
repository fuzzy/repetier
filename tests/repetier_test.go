package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/fuzzy/repetier"
)

var config = readConfig()

type TestConfig struct {
	Proto   string `json:"proto"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	APIKey  string `json:"apikey"`
	Printer string `json:"printer"`
}

func readJSONConfig() *TestConfig {
	retv := &TestConfig{}
	fn := "./test-fixture-data.json"
	data, e := ioutil.ReadFile(fn)
	if e != nil {
		panic(e)
	}
	json.Unmarshal(data, retv)
	return retv
}

func readEnvConfig() *TestConfig {
	retv := &TestConfig{}
	retv.Proto = os.Getenv("TEST_PROTO")
	retv.Host = os.Getenv("TEST_HOST")
	retv.Port, _ = strconv.Atoi(os.Getenv("TEST_PORT"))
	retv.APIKey = os.Getenv("TEST_API_KEY")
	retv.Printer = os.Getenv("TEST_SLUG")
	return retv
}

func readConfig() *TestConfig {
	fn := "./test-fixture-data.json"
	if _, err := os.Stat(fn); err == nil {
		return readJSONConfig()
	}
	return readEnvConfig()
}

func TestClientProto(t *testing.T) {

	if config.Proto != "http" && config.Proto != "https" {
		t.Log("Protocol is unsupported or not reading correctly.")
		t.FailNow()
	}
}

func TestClientHost(t *testing.T) {
	if config.Host != "10.0.0.99" {
		t.Log("Host is not what was expected.")
		t.FailNow()
	}
}

func TestClientPort(t *testing.T) {
	if config.Port != 3344 {
		t.Log("Port is unsupported or not reading correctly.")
		t.FailNow()
	}
}

func TestClientListPrinter(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	flag := false
	for _, v := range api.ListPrinter() {
		if v.Slug == config.Printer {
			flag = true
		}
	}
	if !flag {
		t.FailNow()
	}
}

func TestClientStateList(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	data := api.StateList(config.Printer, false)
	if !data[config.Printer].PowerOn {
		t.FailNow()
	}
}

func TestClientMove(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	ret := api.Move(config.Printer, 25.0, 25.0, 25.0, 0.0, 10.0, false)
	if string(ret) != "{}" {
		t.FailNow()
	}
}

func TestClientMessages(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.Messages()
}

func TestClientRemoveMessage(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	for _, msg := range api.Messages() {
		if msg.Slug == config.Printer {
			if string(api.RemoveMessage(msg.ID, "")) != "{}" {
				t.FailNow()
			}
		}
	}
}

func TestClientListModels(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.ListModels("*", config.Printer)
}

func TestClientCopyModel(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	models := api.ListModels("*", config.Printer)
	if len(models) > 0 {
		api.CopyModel(models["data"][0].ID, true, config.Printer)
	}
}

func TestClientListJobs(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.ListJobs(config.Printer)
}

func TestClientGetPrinterConfig(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.GetPrinterConfig(config.Printer)
}
