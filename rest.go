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
	Proto  string
	Host   string
	Port   int
	Apikey string
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

// ListPrinter blah
func (obj *RestClient) ListPrinter() []byte {
	return obj.request("api", "listPrinter", map[string]interface{}{}, "")
}

// StateList blah
func (obj *RestClient) StateList(s string, ih bool) []byte {
	return obj.request("api", "stateList", map[string]interface{}{"includeHistory": ih}, s)
}

// Move blah
func (obj *RestClient) Move(slug string, x, y, z, e, s float64, r bool) []byte {
	return obj.request("api", "move", map[string]interface{}{
		"z": z, "x": x, "y": y, "e": e, "speed": s, "relative": r}, slug)
}

// Messages blah
func (obj *RestClient) Messages() []byte {
	return obj.request("api", "messages", map[string]interface{}{}, "")
}

// RemoveMessage blah
func (obj *RestClient) RemoveMessage(id int64, a string) []byte {
	if a != "" && a != "unpause" {
		panic("Invalid action passed to RestClient.RemoveMessage")
	}
	return obj.request("api", "removeMessage", map[string]interface{}{"id": id, "a": a}, "")
}

// RemoveModel blah
func (obj *RestClient) RemoveModel(id int64, s string) []byte {
	return obj.request("api", "removeModel", map[string]interface{}{"id": id}, s)
}
