package terratest

// The options to be passed into any terratest.Apply or Destroy function
type TerratestOptions struct {
	UniqueId		    string	      // A unique identifier for this terraform run.
	TestName                    string            // the name of the test to run, for logging purposes.
	TemplatePath                string            // the relative or absolute path to the terraform template to be applied.
	Vars                        map[string]string // the vars to pass to the terraform template.
	RetryableTerraformErrors    map[string]string // a map of error messages which we expect on this template and which should invoke a second terraform apply, along with an additional message offering details on this error text.
	TfRemoteStateS3BucketName   string            // S3 bucket name where terraform.tfstate files should be stored for running any terraform runs. Bucket should already exist.
	TfRemoteStateS3BucketRegion string            // AWS Region where the TfRemoteStateS3BucketName exists.
}

// Initialize an ApplyOptions struct with default values
func NewTerratestOptions() *TerratestOptions {
	return &TerratestOptions{
		TfRemoteStateS3BucketName: defaultTfRemoteStateS3BucketName,
		TfRemoteStateS3BucketRegion: defaultTfRemoteStateS3BuckeRegion,
	}
}

// getTfStateFileName creates a path and filename used to reference a terraform tfstate file. E.g. this is
// useful with S3 for deciding where the tfstate file should be within a given bucket.
func (options *TerratestOptions) getTfStateFileName() string {
	return options.UniqueId + "/terraform.tfstate"
}