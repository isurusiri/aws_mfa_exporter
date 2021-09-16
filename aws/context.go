package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workdocs"
)

// GetAWSSession returns a session to access aws
// resources and services.
// TODO: this code must further optimized.
func GetAWSSession() *workdocs.WorkDocs {
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	
	svc := workdocs.New(awsSession)
	
	return svc
}
