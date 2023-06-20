package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gssh/pkg/aws"
	"gssh/pkg/config"
	"gssh/pkg/ssh"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
)

func main() {
	// args
	var filter string
	if len(os.Args) == 2 {
		filter = os.Args[1]
	}

	// get config
	config, err := config.Get()
	if err != nil {
		log.Panic(err)
	}

	// get instances
	profiles := strings.Split(config.AWS.Profile, ",")
	var instances []*ec2.DescribeInstancesOutput
	for k := range profiles {
		svc, err := aws.NewService(profiles[k], config.AWS.Region)
		if err != nil {
			log.Fatal(err)
		}

		list, err := aws.Get(svc, filter)
		if err != nil {
			log.Fatal(err)
		}
		instances = append(instances, list)
	}

	i := aws.Metadata(instances)
	if err != nil {
		log.Panic(err)
	}

	var instanceID string
	switch {
	// no instances
	case len(i) == 0:
		fmt.Println("no instances found")
		os.Exit(0)
	// if there is only 1 instance
	case len(i) == 1:
		instanceID = i[0].Values["instance-id"]
	default:
		// filter
		printTable(i)

		// select instance
		fmt.Print("Select InstanceID: ")
		_, err = fmt.Scanln(&instanceID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// get ip
	iptype := "private"
	if len(config.SSH.Bastion) == 0 {
		iptype = "public"
	}

	ip, err := (aws.GetIP(instanceID, i, iptype))
	if err != nil {
		log.Fatal(err)
	}

	// and connect
	ssh.Shell(ip, config)
}

func printTable(i []aws.Server) {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow(color.YellowString("InstanceID"), color.YellowString("Name"), color.YellowString("PrivateIP"), color.YellowString("PublicIP"))
	for _, instance := range i {
		table.AddRow(color.GreenString(instance.Values["instance-id"]), instance.Name, instance.Values["private-ip"], instance.Values["public-ip"])
	}
	fmt.Printf("%s\n\n", table)
}
