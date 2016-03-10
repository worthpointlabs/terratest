package terraform

import (
	"log"
	"strings"

	"github.com/gruntwork-io/terraform-test/shell"
)

// terraform apply
// - pass in vars
// - read a template from a known location
// - terraform apply

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// NOTE: Copy & pasted from Jim's original terraform-modules/test/terraform_helper.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// To debug or report Terraform bugs, you need to add a TF_LOG environment variable and set it to a log level like
// TRACE (https://www.terraform.io/docs/internals/debugging.html). However, it seems that whenever you enable any
// logging with TF_LOG, Terraform spams a TON of logging at you, which makes it hard to read the output. Therefore,
// we pass an empty env var map for now, but if you need to debug, just uncomment the version with TF_LOG.
//var terraformDebugEnv = map[string]string{"TF_LOG": "INFO"}
var terraformDebugEnv = map[string]string{}

func ConfigureRemoteState(templatePath string, s3BucketName string, tfStateFileName string, awsRegion string, logger *log.Logger) error {
	logger.Println("Setting up Terraform remote state storage in S3 bucket", s3BucketName, "with tfstate file name", tfStateFileName, "for folder", templatePath)
	args := []string{"remote", "config", "-backend=s3", "-backend-config=bucket=" + s3BucketName, "-backend-config=key=" + tfStateFileName, "-backend-config=encrypt=true", "-backend-config=region=" + awsRegion}
	return shell.RunCommand(shell.Command{ Command: "terraform", Args: args, WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}

func Apply(templatePath string, vars map[string]string, logger *log.Logger) error {
	return shell.RunCommand(shell.Command { Command: "terraform", Args: FormatArgs(vars, "apply", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}

func ApplyAndGetOutput(terraformPath string, vars map[string]string, logger *log.Logger) (string, error) {
	logger.Println("Applying Terraform templates in folder", terraformPath)
	return shell.RunCommandAndGetOutput( shell.Command { Command: "terraform", Args: FormatArgs(vars, "apply", "-input=false"), WorkingDir: terraformPath, Env: terraformDebugEnv }, logger)
}

// Regrettably Terraform has many bugs. Often, it's sufficient to just re-run a Terraform apply.
// This function declares which Terraform error messages warrant an automatic retry. All other errors will not be retried automatically.
func ApplyWithRetry(terraformPath string, vars map[string]string, logger *log.Logger) error {
	output, err := ApplyAndGetOutput(terraformPath, vars, logger)
	if err != nil {
		logger.Printf("Terraform apply failed with error: %s\n", err.Error())

		// Check for all Terraform errors
		if strings.Contains(output, TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY) {
			logger.Printf("Terraform apply failed with the error '%s'. %s\n", TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY, TF_ERROR_DIFFS_DIDNT_MATCH_DURING_APPLY_MSG)
			return Apply(terraformPath, vars, logger)
		} else {
			return err
		}
	}

	return nil
}

func Destroy(templatePath string, vars map[string]string, logger *log.Logger) error {
	logger.Println("Destroy Terraform changes in folder", templatePath)
	return shell.RunCommand(shell.Command { Command: "terraform", Args: FormatArgs(vars, "destroy", "-force", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}