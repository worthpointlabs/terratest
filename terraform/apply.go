package terraform

import "log"

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

func Apply(templatePath string, vars map[string]string, logger *log.Logger) error {
	return RunShellCommand(ShellCommand { Command: "terraform", Args: FormatArgs(vars, "apply", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}

func Destroy(templatePath string, vars map[string]string, logger *log.Logger) error {
	logger.Println("Destroy Terraform changes in folder", templatePath)
	return RunShellCommand(ShellCommand { Command: "terraform", Args: FormatArgs(vars, "destroy", "-force", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}