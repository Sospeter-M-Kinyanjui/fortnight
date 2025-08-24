package checknumber

func checkNumber(arg string) bool {
	if len(arg) == 0 {
		return false
	}

	for _, v := range arg {
		if v >= 48 && v <= 57 {
			return true
		}
	}
	return false
}
