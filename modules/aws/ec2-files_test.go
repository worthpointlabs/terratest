package aws

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/ssh"
)

func TestSshOptions(t *testing.T) {
	// Only one should be set (exclusively), ensure it errors appropriately
	valid := []SshOptions{
		SshOptions{
			UserName: "a",
			KeyPair:  new(Ec2Keypair),
		},
		SshOptions{
			UserName: "a",
			SshAgent: true,
		},
		SshOptions{
			UserName:         "a",
			OverrideSshAgent: new(ssh.SshAgent),
		},
	}
	for _, v := range valid {
		err := v.Validate()
		if err != nil {
			t.Error("Expected nil error, got ", err)
		}
	}
	invalid := []SshOptions{
		SshOptions{}, // none should also error
		SshOptions{
			// No username
			KeyPair:          new(Ec2Keypair),
			SshAgent:         true,
			OverrideSshAgent: new(ssh.SshAgent),
		},
		SshOptions{
			UserName:         "a",
			KeyPair:          new(Ec2Keypair),
			SshAgent:         true,
			OverrideSshAgent: new(ssh.SshAgent),
		},
		SshOptions{
			UserName: "a",
			KeyPair:  new(Ec2Keypair),
			SshAgent: true,
		},
		SshOptions{
			UserName:         "a",
			SshAgent:         true,
			OverrideSshAgent: new(ssh.SshAgent),
		},
		SshOptions{
			UserName:         "a",
			KeyPair:          new(Ec2Keypair),
			OverrideSshAgent: new(ssh.SshAgent),
		},
	}
	for _, v := range invalid {
		err := v.Validate()
		if err == nil {
			t.Error("Expected error, got nil")
		}
	}
}
