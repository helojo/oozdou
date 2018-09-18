package main

import (
	"fmt"
	"os"
	"sync"

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

	var (
		wg sync.WaitGroup
	)
	// 玩家1
	go func() {
		for {
			pockers, ok := <-firstPlayer
			if !ok {
				break
			}
			ozlog.Infof("我拿到牌了，我的牌是: %v", pockers)
		}
		wg.Done()
	}()

	// 玩家2
	go func() {
		for {
			pockers, ok := <-secondPlayer
			if !ok {
				break
			}
			ozlog.Infof("我拿到牌了，我的牌是: %v", pocker)
		}
		wg.Done()
	}()

	// 主玩家
	go func() {
		for {
			pockers, ok := <-secondPlayer
			if !ok {
				break
			}
			ozlog.Infof("我拿到牌了，我的牌是: %v", pockers)
		}
		wg.Done()
	}()
	wg.Add(1)
	wg.Wait()
}
