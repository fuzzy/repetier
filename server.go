package repetier

// Server defines a given Repetier-Server instance
type Server struct {
	Name     string
	Software string
	Version  string
	Printers map[string]*Printer
	api      *client
}

// NewServer returns a new Server object
func NewServer(r, h string, p int, a string) *Server {
	retv := &Server{
		Printers: map[string]*Printer{},
		api:      newClient(r, h, p, a),
	}
	retv.Update()
	return retv
}

// ServerInfo represents some basic information about the server instance
type ServerInfo struct {
	Name     string `json:"name"`
	Printers []struct {
		Active bool   `json:"active"`
		Name   string `json:"name"`
		Online int    `json:"online"`
		Slug   string `json:"slug"`
	} `json:"printers"`
	Servername string `json:"servername"`
	Serveruuid string `json:"serveruuid"`
	Version    string `json:"version"`
}
