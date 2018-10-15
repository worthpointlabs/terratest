package aws

import (
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// GetAccountId gets the Account ID for the currently logged in IAM User.
func GetAccountId(t *testing.T) string {
	id, err := GetAccountIdE(t)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

// GetAccountIdE gets the Account ID for the currently logged in IAM User.
func GetAccountIdE(t *testing.T) (string, error) {
	svc := sts.New(session.New())
	callerIdentity, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	return *callerIdentity.Account, nil
}

// An IAM arn is of the format arn:aws:iam::123456789012:user/test. The account id is the number after arn:aws:iam::,
// so we split on a colon and return the 5th item.
func extractAccountIDFromARN(arn string) (string, error) {
	arnParts := strings.Split(arn, ":")

	if len(arnParts) < 5 {
		return "", errors.New("Unrecognized format for IAM ARN: " + arn)
	}

	return arnParts[4], nil
}
