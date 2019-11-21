package main

import "fmt"

// interface (here we define what methods other types must have)
type Cloud interface {
	Instances() string
}

// lets create a struct for each cloud provider
type AWS struct {
	AccountID string
}

type GCP struct {
	ID string
}

// methods for each type
// AWS
func (a AWS) Instances() string {
	// put your code here using a.AccountID to obtain your instances list in AWS
	// I'm returning some text
	return fmt.Sprintf("AccountID: %s, InstanceID: i-adfasdfadf, PrivateIP: 1.1.1.1", a.AccountID)
}

// GCP
func (g GCP) Instances() string {
	// put your code here using g.ID to obtain your instances list in Google Cloud
	// I'm returning some text
	return fmt.Sprintf("ID: %s, InstanceID: adfadgaetaet, PrivateIP: 1.1.1.1", g.ID)
}

// Now for the important part. Let's satisfy the interface. This is the method we will use at the end :-)
func GetInstances(c Cloud) {
	switch c.(type) {
	case AWS:
		fmt.Println("This is type AWS with AccountID: ", c.(AWS))
	case GCP:
		fmt.Println("This is type GCP with ID: ", c.(GCP))
	}

	fmt.Println(c.Instances())
}

// main
func main() {
	victor := AWS{AccountID: "325235235"}
	GetInstances(victor)

	sebas := GCP{ID: "99893543"}
	GetInstances(sebas)
}
