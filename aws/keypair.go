package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Create a new EC2 Keypair with the given name.
func CreateEC2KeyPair(awsRegion string, name string, publicKey string) error {
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	params := &ec2.ImportKeyPairInput{
		KeyName: aws.String(name), // Required
		PublicKeyMaterial: []byte(publicKey), // Required
		DryRun:  aws.Bool(false),
	}

	_, err := svc.ImportKeyPair(params)
	if err != nil {
		return fmt.Errorf("Failed to import EC2 keypair: %s\n", err)
	}

	return nil
}

// Delete an EC2 Keypair
func DeleteEC2KeyPair(awsRegion string, name string) error {
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

	params := &ec2.DeleteKeyPairInput{
		KeyName: aws.String(name), // Required
		DryRun:  aws.Bool(false),
	}

	_, err := svc.DeleteKeyPair(params)
	if err != nil {
		return fmt.Errorf("Failed to delete EC2 keypair: %s\n", err)
	}

	return nil
}
