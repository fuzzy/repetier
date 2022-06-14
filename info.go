package repetier

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type RepetierInfo struct {
	Apikey   string `json:"apikey"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	ServerName string `json:"servername"`
	ServerUUID string `json:"serveruuid"`
	Printers []struct {
		Active bool   `json:"active"`
		Name   string `json:"name"`
		Online int    `json:"online"`
		Slug   string `json:"slug"`
	} `json:"printers"`
}

func (obj *RestClient) Info() *RepetierInfo {
	retv := &RepetierInfo{}
	fUrl := fmt.Sprintf("%s/info", obj.BaseURL())
	retn, err := http.Get(fUrl)
	panicCheck(err)
	if retn.StatusCode == 200 {
		data, err := ioutil.ReadAll(retn.Body)
		panicCheck(err)
		json.Unmarshal(data, &retv)
	}
	json.Unmarshal([]byte{}, &retv)
	return retv
}
