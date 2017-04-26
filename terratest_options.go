package terratest

// The options to be passed into any terratest.Apply or Destroy function
type TerratestOptions struct {
	UniqueId		    string	           // A unique identifier for this terraform run.
	TestName                    string                 // the name of the test to run, for logging purposes.
	TemplatePath                string                 // the relative or absolute path to the terraform template to be applied.
	Vars                        map[string]interface{} // the vars to pass to the terraform template.
	RetryableTerraformErrors    map[string]string      // a map of error messages which we expect on this template and which should invoke a second terraform apply, along with an additional message offering details on this error text.
}

// Initialize a TerratestOptions struct with default values
func NewTerratestOptions() *TerratestOptions {
	return &TerratestOptions{}
}

// getTfStateFileName creates a path and filename used to reference a terraform tfstate file. E.g. this is
// useful with S3 for deciding where the tfstate file should be within a given bucket.
func (options *TerratestOptions) GetTfStateFileName() string {
	return options.UniqueId + "/terraform.tfstate"
}