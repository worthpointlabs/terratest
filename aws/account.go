package aws
import (
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/aws/session"
	"strings"
	"errors"
)

// Get the Account ID for the currently logged in IAM User.
func GetAccountId() (string, error) {
	svc := iam.New(session.New())
	user, err := svc.GetUser(&iam.GetUserInput{})
	if err != nil {
		return "", err
	}

	return extractAccountIdFromArn(*user.User.Arn)
}

// An IAM arn is of the format arn:aws:iam::123456789012:user/test. The account id is the number after arn:aws:iam::,
// so we split on a colon and return the 5th item.
func extractAccountIdFromArn(arn string) (string, error) {
	arnParts := strings.Split(arn, ":")

	if len(arnParts) < 5 {
		return "", errors.New("Unrecognized format for IAM ARN: " + arn)
	}

	return arnParts[4], nil
}
