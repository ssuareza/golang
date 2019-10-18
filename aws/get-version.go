/* go run ssm.go security admin */
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	// arguments
	profile := "dev"
	environment := os.Args[1]
	service := os.Args[2]

	// session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Println("Error creating session ", err)
		return
	}

	// Create ssm service client
	svc := ssm.New(sess)

	keyname := "/" + environment + "/services/" + service + "/version"
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: &keyname,
	})

	value := *param.Parameter.Value
	fmt.Println(value)
}

