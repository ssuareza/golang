/* go run list_clusters.go */
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

func main() {
	// session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "dev",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Println("Error creating session ", err)
		return
	}

	// Create eks service client
	svc := eks.New(sess)
	result, err := svc.ListClusters(nil)
	if err != nil {
		fmt.Println("Unable to list clusters, %v", err)
		os.Exit(1)
	}

	// Printing clusters
	for _, c := range result.Clusters {
		fmt.Println(*c)
	}
}
