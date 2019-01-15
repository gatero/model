package util

func Assign(args ...map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for _, item := range args {
		for key, value := range item {
			output[key] = value
		}
	}

	return output
}

func Contains(slice []string, query string) bool {
	for _, pointer := range slice {
		if query == pointer {
			return true
		}
	}
	return false
}

func Find(slice []string, query string) int {
	for index, pointer := range slice {
		if query == pointer {
			return index
		}
	}
	return len(slice)
}
