package repetier

func panicCheck(e error) {
	if e != nil {
		panic(e)
	}
}
