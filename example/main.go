/**
 * Project Name:configs
 * File Name:main.go
 * Package Name:example
 * Date:2019年07月16日 17:32
 * Function:
 * Copyright (c) 2019, Jason.Wang All Rights Reserved.
 */
package main

import (
	"fmt"
	"github.com/kanyways/configs"
)

type ymlConf struct {
	Server []struct {
		Name string `yaml:"name" toml:"name" json:"name"`
	}
}

func main() {
	var conf ymlConf

	fmt.Println(configs.GetConfigAbsolutePath("example/config.json"))
	configs.Parse(&conf, configs.GetConfigAbsolutePath("example/config.json"))
	fmt.Printf("%#v\n", conf)

	fmt.Println(configs.GetConfigAbsolutePath("example/config.yml"))
	configs.Parse(&conf, configs.GetConfigAbsolutePath("example/config.yml"))
	fmt.Printf("%#v\n", conf)

	fmt.Println(configs.GetConfigAbsolutePath("example/config.yaml"))
	configs.Parse(&conf, configs.GetConfigAbsolutePath("example/config.yaml"))
	fmt.Printf("%#v\n", conf)

	fmt.Println(configs.GetConfigAbsolutePath("example/config.toml"))
	configs.Parse(&conf, configs.GetConfigAbsolutePath("example/config.toml"))
	fmt.Printf("%#v\n", conf)
}
