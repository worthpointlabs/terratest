package packer

import (
	"log"
	"github.com/gruntwork-io/terratest/shell"
	"regexp"
	"errors"
	"github.com/gruntwork-io/terratest/util"
)

type PackerOptions struct {
	Template string				// The path to the Packer template
	Vars     map[string]string  // The custom vars to pass when running the build command
}

// Build the given Packer template and return the generated AMI ID
func BuildAmi(options PackerOptions, logger *log.Logger) (string, error) {
	logger.Printf("Running Packer to generate AMI for template %s", options.Template)

	output, err := shell.RunCommandAndGetOutput(shell.Command {Command: "packer", Args: util.FormatArgs(options.Vars, "build", "-machine-readable", options.Template)}, logger)
	if err != nil {
		return "", err
	}

	return extractAmiId(output)
}

// The Packer machine-readable log output should contain an entry of this format:
//
// <timestamp>,<builder>,artifact,<index>,id,<region>:<ami_id>
//
// For example:
//
// 1456332887,amazon-ebs,artifact,0,id,us-east-1:ami-b481b3de
func extractAmiId(packerLogOutput string) (string, error) {
	re := regexp.MustCompile(".+artifact,\\d+?,id,.+?:(.+)")
	matches := re.FindStringSubmatch(packerLogOutput)

	if len(matches) == 2 {
		return matches[1], nil
	} else {
		return "", errors.New("Could not find AMI ID pattern in Packer output")
	}
}
