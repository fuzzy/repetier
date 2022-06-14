package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
	"strconv"
	"testing"
	"fmt"
	"github.com/fuzzy/repetier"
)

var (
	config = readConfig()
	red_c = "\033[1;31m"
	green_c = "\033[1;32m"
	yellow_c = "\033[1;33m"
	end_c = "\033[0m"
)

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

func colorTemp(v float64) string {
	if v < 40 {
		return fmt.Sprintf("%s%.02f%s", green_c, v, end_c)
	} else if v >=40 && v < 45 {
		return fmt.Sprintf("%s%.02f%s", yellow_c, v, end_c)
	} else {
		return fmt.Sprintf("%s%.02f%s", red_c, v, end_c)
	}
}

func TestClientProto(t *testing.T) {

	if config.Proto != "http" && config.Proto != "https" {
		t.Log("Protocol is unsupported or not reading correctly.")
		t.FailNow()
	}
}

func TestClientPort(t *testing.T) {
	if config.Port != 3344 {
		t.Log("Port is unsupported or not reading correctly.")
		t.FailNow()
	}
}

func TestInfo(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	data := api.Info()
	if len(data.ServerName) > 0 {
		fmt.Printf("%sServer%s: %s\n", green_c, end_c, data.ServerName)
		fmt.Printf("%sVersion%s: %s %s\n", green_c, end_c, data.Name, data.Version)
		fmt.Printf("%s# of Printers%s: %d\n", green_c, end_c, len(data.Printers))
		time.Sleep(500 * time.Millisecond)
	}
}

func TestClientListPrinter(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	flag := false
	for _, v := range api.ListPrinter() {
		fmt.Printf("%s>%s %s\n", green_c, end_c, v.Slug)
		if v.Slug == config.Printer {
			flag = true
		}
	}
	if !flag {
		t.FailNow()
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientStateList(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	for k, v := range api.StateList(config.Printer, false) {
		fmt.Printf("%s>%s %s Ext:[", green_c, end_c, k)
		for _, e := range v.Extruder {
			fmt.Printf("%s ", colorTemp(e.TempRead))
		}
		fmt.Printf("] Bed: %s\n", colorTemp(v.HeatedBeds[0].TempRead))
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientMove(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)

	ret := api.Move(config.Printer, 25.0, 25.0, 25.0, 0.0, 10.0, false)
	if string(ret) != "{}" {
		t.FailNow()
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientMessages(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	for _, v := range api.Messages() {
		fmt.Printf("%+v\n", v)
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientRemoveMessage(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	for _, msg := range api.Messages() {
		if string(api.RemoveMessage(msg.ID, "")) != "{}" {
			t.FailNow()
		} else {
			fmt.Printf("Removed message %d from printer %s\n", msg.ID, msg.Slug)
		}
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientListModels(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.ListModels("*", config.Printer)
	time.Sleep(500 * time.Millisecond)
}

func TestClientCopyModel(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	models := api.ListModels("*", config.Printer)
	if len(models["data"]) > 0 {
		api.CopyModel(models["data"][0].ID, true, config.Printer)
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientListJobs(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	for _, p := range api.ListPrinter() {
		fmt.Printf("\033[1;32m>\033[0m %s\n", p.Slug)
		for _, v := range api.ListJobs(p.Slug)["data"] {
			fmt.Printf("Name: %s -- State: %s\n", v.Name, v.State)
		}
	}
	time.Sleep(500 * time.Millisecond)
}

func TestClientGetPrinterConfig(t *testing.T) {
	api := repetier.NewRestClient(config.Proto, config.Host, config.Port, config.APIKey)
	api.GetPrinterConfig(config.Printer)
	time.Sleep(500 * time.Millisecond)
}
