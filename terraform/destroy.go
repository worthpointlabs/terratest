// Wrapper functions for "terraform destroy"
package terraform

import (
	"log"

	"github.com/gruntwork-io/terratest/shell"
)

func Destroy(templatePath string, vars map[string]string, logger *log.Logger) error {
	logger.Println("Destroy Terraform changes in folder", templatePath)
	return shell.RunCommand(shell.Command { Command: "terraform", Args: FormatArgs(vars, "destroy", "-force", "-input=false"), WorkingDir: templatePath, Env: terraformDebugEnv }, logger)
}
