package repetier

// Move blah
func (obj *RestClient) Move(slug string, x, y, z, e, s float64, r bool) []byte {
	return obj.request("api", "move", map[string]interface{}{
		"z": z, "x": x, "y": y, "e": e, "speed": s, "relative": r}, slug)
}

func (obj *RestClient) MoveX(slug string, x, s float64, r bool) []byte {
	return obj.Move(slug, x, 0.0, 0.0, 0.0, s, r)
}

func (obj *RestClient) MoveY(slug string, y, s float64, r bool) []byte {
	return obj.Move(slug, 0.0, y, 0.0, 0.0, s, r)
}

func (obj *RestClient) MoveZ(slug string, z, s float64, r bool) []byte {
	return obj.Move(slug, 0.0, 0.0, z, 0.0, s, r)
}

func (obj *RestClient) Extrude(slug string, e, s float64, r bool) []byte {
	return obj.Move(slug, 0.0, 0.0, 0.0, e, s, r)
}
