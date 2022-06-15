package repetier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RestClient blah
type RestClient struct {
	Proto   string
	Host    string
	Port    int
	Apikey  string
}

// NewRestClient blah
func NewRestClient(r, h string, p int, ak string) *RestClient {
	return &RestClient{
		Proto:  r,
		Host:   h,
		Port:   p,
		Apikey: ak,
	}
}

func (obj *RestClient) request(api, act string, data map[string]interface{}, slug string) []byte {
	jdata, err := json.Marshal(data)
	panicCheck(err)
	udata := url.QueryEscape(string(jdata))
	fullURL := fmt.Sprintf("%s/%s/%s?apikey=%s&a=%s&data=%s", obj.BaseURL(), api, slug, obj.Apikey, act, udata)
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

// BaseURL blah
func (obj *RestClient) BaseURL() string {
	return fmt.Sprintf("%s://%s:%d/printer", obj.Proto, obj.Host, obj.Port)
}

// Action blah
func (obj *RestClient) Action(act string, data map[string]interface{}, slug string) []byte {
	return obj.request("api", act, data, slug)
}
