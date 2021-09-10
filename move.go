package repetier

// Move blah
func (obj *RestClient) Move(slug string, x, y, z, e, s float64, r bool) []byte {
	return obj.request("api", "move", map[string]interface{}{
		"z": z, "x": x, "y": y, "e": e, "speed": s, "relative": r}, slug)
}
