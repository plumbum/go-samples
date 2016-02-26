package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("./appname")
	viper.AddConfigPath("./appname2")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(viper.Get("name"))
	fmt.Println(viper.Get("val1"))
	fmt.Println(viper.Get("val2"))
	fmt.Println(viper.Get("home"))
	fmt.Println(viper.Get("user"))
	fmt.Println(viper.Get("term"))

}
