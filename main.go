package main

import (
	"fmt"
	"os"

	"github.com/usthooz/oozgconf"
	ozlog "github.com/usthooz/oozlog/go"
)

// Config config
type Config struct {
	Pokers []string `yaml:"pokers"`
}

var (
	firstPlayer  = make(chan []string, 20)
	secondPlayer = make(chan []string, 20)
	mainPlayer   = make(chan []string, 20)
)

var (
	conf *Config
)

// init init
func init() {
	ozconf := oozgconf.NewConf(&oozgconf.OozGconf{
		ConfPath: "./config/config.yaml",
		Subffix:  "yaml",
	})
	err := ozconf.GetConf(&conf)
	if err != nil {
		ozlog.Errorf("GetConf Err: %v", err.Error())
	}
}

func main() {
	var (
		isStart string
	)
	// 开始执行
	ozlog.Infof("oozdou 斗地主启动成功.")
	ozlog.Infof("You want start this game? y/n")
	fmt.Scanln(&isStart)
	if isStart == "n" {
		os.Exit(0)
	}
}
