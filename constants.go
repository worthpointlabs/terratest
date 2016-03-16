package terratest

// S3 Location where terraform.tfstate files should be stored for running all terraform runs as part of this test suite.
// - As of 3/10/16, this S3 bucket deletes all data after 30 days. This means tfstate hangs around for 30 days in case you need to debug anything.
const TF_REMOTE_STATE_S3_BUCKET_NAME = "gruntwork-terraform-test-remote-state"
const TF_REMOTE_STATE_S3_BUCKET_REGION = "us-west-2"