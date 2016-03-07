package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Upload a new EC2 Keypair with the given name.
func UploadEC2KeyPair(awsRegion string, name string, publicKey string) error {
	svc := ec2.New(session.New())

	params := &ec2.ImportKeyPairInput{
		KeyName: aws.String(name), // Required
		PublicKeyMaterial: []byte(publicKey),
		DryRun:  aws.Bool(false),
	}

	_, err := svc.ImportKeyPair(params)
	if err != nil {
		return fmt.Errorf("Failed to import EC2 keypair: %s\n", err)
	}

	return nil
}
