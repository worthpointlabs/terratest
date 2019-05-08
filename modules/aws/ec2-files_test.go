package aws

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/ssh"
)

func TestSshAuth(t *testing.T) {
	// Only one should be set (exclusively), ensure it errors appropriately
	valid := []SshAuth{
		SshAuth{
			KeyPair: new(Ec2Keypair),
		},
		SshAuth{
			SshAgent: true,
		},
		SshAuth{
			OverrideSshAgent: new(ssh.SshAgent),
		},
	}
	for _, v := range valid {
		err := v.Validate()
		if err != nil {
			t.Error("Expected nil error, got ", err)
		}
	}
	invalid := []SshAuth{
		SshAuth{}, // none should also error
		SshAuth{
			KeyPair:          new(Ec2Keypair),
			SshAgent:         true,
			OverrideSshAgent: new(ssh.SshAgent),
		},
		SshAuth{
			KeyPair:  new(Ec2Keypair),
			SshAgent: true,
		},
		SshAuth{
			SshAgent:         true,
			OverrideSshAgent: new(ssh.SshAgent),
		},
		SshAuth{
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
