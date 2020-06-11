/**
 * Project Name:configs
 * File Name:configs.go
 * Package Name:configs
 * Date:2019年07月16日 17:10
 * Function:
 * Copyright (c) 2019, Jason.Wang All Rights Reserved.
 */
package configs

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

type Error struct {
	Message string
}

func (h *Error) Error() string {
	return h.Message
}

//判断当前的文件是否存在
func fileExists(filename string) bool {
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

//获取配置文件的相对路径
func GetConfigAbsolutePath(file string) string {
	app := path.Dir(os.Args[0])
	if strings.HasPrefix(app, os.TempDir()) {
		return getConfigAbsolutePathForTest(file)
	}
	return getConfigAbsolutePathForBase(file)
}

//获取正式文件的路径，相对于自己本身、config文件夹、/etc目录
func getConfigAbsolutePathForBase(file string) string {
	app := path.Base(os.Args[0])
	for _, dir := range []string{
		"",
		"config",
		"/etc/" + app,
		path.Join(os.Getenv("HOME"), "."+app),
	} {
		cf := path.Join(dir, file)
		if fileExists(cf) {
			return cf
		}
	}
	return ""
}

//获取文件给测试用例去使用
func getConfigAbsolutePathForTest(file string) string {
	_, filename, _, _ := runtime.Caller(2)
	dir := path.Dir(filename)
	for {
		for _, d := range []string{"", "config"} {
			cf := path.Join(dir, d, file)
			if fileExists(cf) {
				return cf
			}
		}
		dir = path.Dir(strings.TrimRight(dir, "/"))
		if dir == "/" {
			break
		}
	}
	return file
}

func Parse(cfg interface{}, file string) error {
	err := load(cfg, file)
	if h, ok := err.(*Error); ok {
		fmt.Println(h.Error())
		os.Exit(1)
	}
	return err
}

func load(cfg interface{}, file string) error {
	err := parseFile(cfg, file)
	if err != nil {
		return err
	}

	parser := flags.NewParser(cfg, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)
	if _, err := parser.Parse(); err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			return &Error{e.Message}
		}
		return err
	}
	return nil
}

//解析文件
func parseFile(cfg interface{}, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	switch path.Ext(file) {
	case ".json":
		return json.NewDecoder(f).Decode(cfg)
	case ".yaml", ".yml":
		in, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(in, cfg)
	case ".toml":
		_, err := toml.DecodeReader(f, cfg)
		return err
	default:
		return fmt.Errorf("unsupported config file format: %s", file)
	}
}
