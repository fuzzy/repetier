package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"

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

func TestServerObjectInstantiation(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	if api.Name == "" || api.Version == "" {
		t.Errorf("API Client not updating properly on instantiation.")
	}
}

func TestServerObjectHasSlugs(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	if len(api.Slugs()) == 0 {
		t.Errorf("API Client not getting all the slugs.")
	}
}

func TestServerObjectHasExtruderTemps(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	for _, v := range api.Printers {
		temp := v.State.Extruder[0].TempRead
		if temp <= 1 {
			t.Errorf("Could not read extruder temperature (got %.02f)", temp)
		}
	}
}

func TestServerObjectHasHeatedBedTemps(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	for _, v := range api.Printers {
		temp := v.State.HeatedBeds[0].TempRead
		if temp <= 1 {
			t.Errorf("Could not read HeatedBed temperature (got %.02f)", temp)
		}
	}
}

func TestServerObjectCanMoveAxis(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	api.Printers[config.Printer].AbsMove(25.0, 25.0, 25.0, 0.0)
	api.Printers[config.Printer].RelMove(25.0, 25.0, 25.0, 0.0)
}

func TestServerObjectCanSetExtruderTemp(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	api.Printers[config.Printer].Extruders[0].SetTemp(215.0)
	// TODO check that the temp gets set
	time.Sleep(10 * time.Second)
	api.Printers[config.Printer].Extruders[0].SetTemp(0.0)
}

func TestServerObjectCanSetBedTemp(t *testing.T) {
	api := repetier.NewServer(config.Proto, config.Host, config.Port, config.APIKey)
	api.Printers[config.Printer].HeatedBed.SetTemp(55.0)
	api.Printers[config.Printer].HeatedBed.Update()
	time.Sleep(10 * time.Second)
	api.Printers[config.Printer].HeatedBed.SetTemp(0.0)
}
