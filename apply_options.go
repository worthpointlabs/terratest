package terratest

// The options to be passed into any terratest.Apply or Destroy function
type ApplyOptions struct {
	TestName                    string            // the name of the test to run, for logging purposes.
	TemplatePath                string            // the relative or absolute path to the terraform template to be applied.
	Vars                        map[string]string // the vars to pass to the terraform template.
	AttemptTerraformRetry       bool              // if true, if a known error message occurs, automatically attempt a retry.
	RetryableTerraformErrors    map[string]string // a map of error messages which we expect on this template and which should invoke a second terraform apply, along with an additional message offering details on this error text.
	TfRemoteStateS3BucketName   string            // S3 bucket name where terraform.tfstate files should be stored for running any terraform runs. Bucket should already exist.
	TfRemoteStateS3BucketRegion string            // AWS Region where the TfRemoteStateS3BucketName exists.
}

// Initialize an ApplyOptions struct with default values
func NewApplyOptions() *ApplyOptions {
	return &ApplyOptions{
		TfRemoteStateS3BucketName: defaultTfRemoteStateS3BucketName,
		TfRemoteStateS3BucketRegion: defaultTfRemoteStateS3BuckeRegion,
	}
}

// generateTfStateFileName creates a path and filename used to reference a terraform tfstate file. E.g. this is
// useful with S3 for deciding where the tfstate file should be within a given bucket.
func (ao *ApplyOptions) generateTfStateFileName(r *RandomResourceCollection) string {
	return r.UniqueId + "/terraform.tfstate"
}