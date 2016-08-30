package terraform

// Convert the inputs to a format palatable to terraform. This includes converting the given vars to the format the
// Terraform CLI expects (-var key=value).
func FormatArgs(customVars map[string]interface{}, args ...string) []string {
	varsAsArgs := FormatTerraformVarsAsArgs(customVars)
	return append(args, varsAsArgs...)
}

