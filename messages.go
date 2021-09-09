package repetier

import (
	"encoding/json"
	"time"
)

// MessagesReturn blah
type MessagesReturn struct {
	Date  time.Time `json:"date"`
	ID    int       `json:"id"`
	Link  string    `json:"link"`
	Msg   string    `json:"msg"`
	Pause bool      `json:"pause"`
	Slug  string    `json:"slug"`
}

// Messages blah
func (obj *RestClient) Messages() []*MessagesReturn {
	retv := []*MessagesReturn{}
	data := obj.Action("messages", map[string]interface{}{}, "")
	json.Unmarshal(data, &retv)
	return retv
}

// RemoveMessage blah
func (obj *RestClient) RemoveMessage(id int, a string) []byte {
	if a != "" && a != "unpause" {
		panic("Invalid action passed to RestClient.RemoveMessage")
	}
	return obj.request("api", "removeMessage", map[string]interface{}{"id": id, "a": a}, "")
}
