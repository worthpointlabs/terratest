package terratest

// ApplyAndDestroy wraps all setup and teardown required for a Terraform Apply operation. It returns the output of the terraform operations.
func ApplyAndDestroy(ao *ApplyOptions) (string, error) {
	defer destroyHelper(ao, ao.getTfStateFileName())
	return Apply(ao)
}