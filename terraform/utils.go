package terraform

// Convert the inputs to a format palatable to terraform
func FormatArgs(customVars map[string]string, args ...string) []string {
	customArgs := []string{}

	for key, value := range customVars {
		customArgs = append(customArgs, "-var", key + "=\"" + value + "\"")
	}

	return append(args, customArgs...)
}