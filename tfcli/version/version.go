package version

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
)

type keyname struct {
	name string
}

func (this keyname) getValue() string {
	// variables
	environment := os.Args[2]
	service := os.Args[3]
	aws_profile := "dev"
	aws_region := "us-east-1"

	if environment == "tfprod" {
		aws_profile = "prod"
	}

	// session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: aws_profile,
		Config: aws.Config{
			Region: aws.String(aws_region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Println("Error creating session ", err)
		os.Exit(1)
	}

	// create ssm service client
	svc := ssm.New(sess)

	// get value from ssm
	keyname := "/" + environment + "/services/" + service + "/" + this.name
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: &keyname,
	})
	if err != nil {
		fmt.Println("KeynameNotFound")
		os.Exit(1)
	}

	return *param.Parameter.Value
}

func getVersion() {
	version := keyname{name: "version"}.getValue()
	fmt.Println(version)
}

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Get service current version running",
		Long:  `Get the version running for a specific environment/service.`,
		Run: func(cmd *cobra.Command, args []string) {
			getVersion()
		},
	}

	return cmd
}
