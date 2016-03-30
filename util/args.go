package util

// Convert the given custom vars and command line args into the format used by Terraform and Packer.
// For example, a terraform command might look like:
//
// terraform apply -input=false -no-color -var foo=bar
//
// To create the command line arguments in this form, you would call FormatArgs as follows:
//
// FormatArgs(map[string]string{"foo": "bar"}, "-input=false", "-no-color")
func FormatArgs(customVars map[string]string, args ...string) []string {
	customArgs := []string{}

	for key, value := range customVars {
		customArgs = append(customArgs, "-var", key + "=" + value)
	}

	return append(args, customArgs...)
}