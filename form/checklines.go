package ascii

func Check(newlines []string) bool {
	for _, word := range newlines {
		if word != "" {
			return false
		}
	}
	return true
}