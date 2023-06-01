package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	path := os.Getenv("HOME") + "/.viper"
	_ = os.Mkdir(path, 744)

	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	// create file if doesn't exists
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("aws.profile", "default")
		viper.Set("aws.region", "us-east-1")
		viper.Set("ssh.user", "sebastian")
		viper.Set("ssh.port", 8822)
		viper.Set("ssh.bastion", true)

		if err := viper.WriteConfigAs(path + "/config.yaml"); err != nil {
			log.Fatal(err)
		}
	}

	// print values
	fmt.Println(viper.Get("aws.profile"))
	fmt.Println(viper.Get("ssh.user"))
}
