package terratest

// ApplyAndDestroy wraps all setup and teardown required for a Terraform Apply operation. It returns the output of the terraform operations.
func ApplyAndDestroy(options *TerratestOptions) (string, error) {
	defer destroyHelper(options, options.getTfStateFileName())
	return Apply(options)
}