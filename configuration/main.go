package main

import (
	"github.com/fzerorubigd/onion"
	_ "github.com/fzerorubigd/onion/tomlloader"
	_ "github.com/fzerorubigd/onion/yamlloader"
	"fmt"
	"strings"
)

type User struct {
	Name string
	Password string
}

func main() {

	cfg := onion.New()
	l1 := onion.NewFileLayer("configuration/config.yaml")
	l2 := onion.NewFileLayer("configuration/config.toml")
	l3 := onion.NewEnvLayer("HOME")

	cfg.AddLayer(l1)
	cfg.AddLayer(l2)
	cfg.AddLayer(l3)

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println(cfg.GetInt("first"))
	fmt.Println(cfg.GetInt("second"))
	fmt.Println(cfg.GetStringSlice("list"))

	fmt.Println(cfg.GetInt("section.first"))
	fmt.Println(cfg.GetInt("section.second"))

	fmt.Println(cfg.GetString("HOME"))

	fmt.Println(strings.Repeat("-", 80))
}
