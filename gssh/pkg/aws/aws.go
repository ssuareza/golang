package aws

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// NewService creates a service connection with ec2
func NewService(profile string, region string) (ec2iface.EC2API, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	return ec2.New(sess), nil
}

// Get gets the instances from aws
func Get(svc ec2iface.EC2API, filter string) (*ec2.DescribeInstancesOutput, error) {
	if len(filter) == 0 {
		filter = "*"
	}

	// Only grab instances that are running or just started
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String("*" + filter + "*"),
				},
			},
		},
	}

	res, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Instance represents an aws instance.
type Instance struct {
	Name   string
	Values map[string]string
}

// ByName is used to sort instance by name.
type ByName []Instance

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Metadata returns the instances with only a few metadata fields.
func Metadata(i []*ec2.DescribeInstancesOutput) []Instance {
	var instances []Instance

	for _, list := range i {
		for _, reservation := range list.Reservations {
			for _, instance := range reservation.Instances {
				record := make(map[string]string)
				record["instance-id"] = nilGuard(instance.InstanceId)
				record["public-hostname"] = nilGuard(instance.PublicDnsName)
				record["public-ip"] = nilGuard(instance.PublicIpAddress)
				record["private-hostname"] = nilGuard(instance.PrivateDnsName)
				record["private-ip"] = nilGuard(instance.PrivateIpAddress)

				for _, tag := range instance.Tags {
					record[fmt.Sprintf("tag:%s", *tag.Key)] = *tag.Value
				}

				instances = append(instances, Instance{
					Name:   record["tag:Name"],
					Values: record,
				})
			}
		}
	}

	sort.Sort(ByName(instances))

	return instances
}

func nilGuard(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
