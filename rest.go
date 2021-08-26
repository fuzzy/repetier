package repetier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	Proto  string
	Host   string
	Port   int
	Apikey string
}

func newClient(r, h string, p int, ak string) *client {
	return &client{
		Proto:  r,
		Host:   h,
		Port:   p,
		Apikey: ak,
	}
}

func (obj *client) baseURL() string {
	return fmt.Sprintf("%s://%s:%d/printer", obj.Proto, obj.Host, obj.Port)
}

func (obj *client) request(api, act string, data map[string]interface{}, slug string) []byte {
	jdata, err := json.Marshal(data)
	panicCheck(err)
	udata := url.QueryEscape(string(jdata))
	fullURL := fmt.Sprintf("%s/%s/%s?apikey=%s&a=%s&data=%s", obj.baseURL(), api, slug, obj.Apikey, act, udata)
	retv, err := http.Get(fullURL)
	panicCheck(err)
	// log.Printf("%s: %s", fullURL, retv.Status)
	if retv.StatusCode == 200 {
		retn, err := ioutil.ReadAll(retv.Body)
		panicCheck(err)
		return retn
	}
	return []byte{}
}

func (obj *client) action(act string, data map[string]interface{}, slug string) []byte {
	return obj.request("api", act, data, slug)
}
