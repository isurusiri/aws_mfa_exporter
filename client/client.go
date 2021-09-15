package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/isurusiri/aws_mfa_exporter/aws"
)

// AWSMFA specifies the
type AWSMFA struct {
	DisabledUsers []string
	count         uint32
}

var orgID = ""

// getAllAWSUsers returns all users in an
// aws account.
func getAllAWSUsers() {
	svc := aws.GetAWSSession()

	input := new(workdocs.DescribeUsersInput)
	input.OrganizationId = &orgID

	users, err := svc.DescribeUsers(input)
	if err != nil {
		fmt.Println("Error getting user info", err)
    	return
	}

	for _, user := range users.Users {
		fmt.Println("Username:   " + *user.Username)
	
		fmt.Println("Firstname:  " + *user.GivenName)
		fmt.Println("Lastname:   " + *user.Surname)
		fmt.Println("Email:      " + *user.EmailAddress)
		fmt.Println("Root folder " + *user.RootFolderId)
	
		fmt.Println("")
	}
}

// getMFADisabled returns AWSMFA struct containing
// a list of users with mfa disabled.
func getMFADisabled() AWSMFA {
  users := AWSMFA {
	DisabledUsers: []string{"isuru", "siriwardana"},
	count: 2,
  }

  return users
}

// GetMFADisabled is used to
func GetMFADisabled() AWSMFA {
	getAllAWSUsers()
	return getMFADisabled()
}
