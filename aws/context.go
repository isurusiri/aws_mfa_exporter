package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workdocs"
)

// GetAWSSession returns a session to access aws
// resources and services.
// TODO: this code must further optimized.
func GetAWSSession() *workdocs.WorkDocs {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	svc := workdocs.New(awsSession)
	
	return svc
}
