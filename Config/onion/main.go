package main

import (
	"github.com/fzerorubigd/onion"
	_ "github.com/fzerorubigd/onion/tomlloader"
	_ "github.com/fzerorubigd/onion/yamlloader"
	_ "github.com/fzerorubigd/onion/flagslayer"
	"fmt"
	"strings"
	"log"
	"github.com/fzerorubigd/onion/flagslayer"
	"flag"
	"os"
)

type User struct {
	Name string
	Password string
}

func main() {

	var err error

	cfg := onion.New()

	// Get command line flags
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	c_cmd := flagslayer.NewFlagLayer(fs)
	c_cmd.SetString("home", "home", "none", "Set home directory")
	c_cmd.SetBool("enable", "enable", false, "Enable demo")
	c_cmd.SetBool("verbose", "v", false, "More verbose output")
	c_cmd.SetInt64("int", "i", 0, "Integer")
	err = cfg.AddLayer(c_cmd)
	check(err)

	// Load YAML config
	c_yaml := onion.NewFileLayer("config.yaml")
	err = cfg.AddLayer(c_yaml)
	check(err)

	// Load TOML config
	c_toml := onion.NewFileLayer("config.toml")
	err = cfg.AddLayer(c_toml)
	check(err)

	// Get ENV variables
	c_env := onion.NewEnvLayer("HOME", "USER")
	err = cfg.AddLayer(c_env)
	check(err)

	// Print results
	fmt.Println(strings.Repeat("=", 80))

	fmt.Println(cfg.GetBool("enable"))

	fmt.Println(cfg.GetInt("first"))
	fmt.Println(cfg.GetInt("second"))
	fmt.Println(cfg.GetStringSlice("list"))

	fmt.Println(cfg.GetInt("section.first"))
	fmt.Println(cfg.GetInt("section.second"))

	// Case insensitive
	fmt.Println(cfg.GetString("home"))
	fmt.Println(cfg.GetString("Home"))
	fmt.Println(cfg.GetString("HOME"))
	fmt.Println(cfg.GetString("user"))

	fmt.Println(strings.Repeat("-", 80))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
