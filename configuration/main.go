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

	// Load YAML config
	l1 := onion.NewFileLayer("config.yaml")
	err = cfg.AddLayer(l1)
	check(err)

	// Load TOML config
	l2 := onion.NewFileLayer("config.toml")
	err = cfg.AddLayer(l2)
	check(err)

	// Get ENV variables
	l3 := onion.NewEnvLayer("HOME")
	err = cfg.AddLayer(l3)
	check(err)

	// Get command line flags
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	l4 := flagslayer.NewFlagLayer(fs)
	l4.SetString("home", "home", "none", "Set home directory")
	l4.SetBool("enable", "enable", false, "Enable demo")
	err = cfg.AddLayer(l4)
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

	fmt.Println(strings.Repeat("-", 80))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
