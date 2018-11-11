package utils

func assign(args ...map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for _, item := range args {
		for key, value := range item {
			output[key] = value
		}
	}

	return output
}
