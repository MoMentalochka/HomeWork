package inventory

func makeSet(values []string) map[string]bool {
	set := make(map[string]bool, len(values))

	for _, value := range values {
		set[value] = true
	}

	return set
}

func hasAny(values []string, allowed map[string]bool) bool {
	for _, value := range values {
		if allowed[value] {
			return true
		}
	}

	return false
}
