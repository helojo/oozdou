package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/usthooz/oozgconf"
	ozlog "github.com/usthooz/oozlog/go"
)

// Config config
type Config struct {
	Pokers []string `json:"pokers"`
}

var (
	FirstPlayer  = make(chan string)
	SecondPlayer = make(chan string)
	MainPlayer   = make(chan string)
)

var (
	// 玩家1、玩家2、主玩家、地主牌
	firstPokers, secondPokers, mainPokers, lordTokers []string
)

var (
	conf *Config
)

// init init
func init() {
	ozconf := oozgconf.NewConf(&oozgconf.OozGconf{
		ConfPath: "./config/config.json",
		Subffix:  "json",
	})
	err := ozconf.GetConf(&conf)
	if err != nil {
		ozlog.Errorf("GetConf Err: %v", err.Error())
	}
	ozlog.Infof("Pockers len: %d", len(conf.Pokers))
	InitPockers(conf.Pokers)
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
	// 发牌
	go func() {
		for {
			ozlog.Infof("开始发牌.")
			Handout()
			ozlog.Infof("发牌完成.")
			ozlog.Infof("地主牌: %v", lordTokers)

			ozlog.Infof("MainPlayer Pockers: %v", mainPokers)
			ozlog.Infof("我的牌数: %d", len(mainPokers))
			weight := GetPlayerWeight(mainPokers)
			ozlog.Warnf("Main Weight: %d", weight)

			ozlog.Infof("FirstPlayer Pockers: %v", firstPokers)
			ozlog.Infof("我的牌数: %d", len(firstPokers))
			weight = GetPlayerWeight(firstPokers)
			ozlog.Warnf("First Weight: %d", weight)

			ozlog.Infof("SecondPlayer Pockers: %v", secondPokers)
			ozlog.Infof("我的牌数: %d", len(secondPokers))
			weight = GetPlayerWeight(secondPokers)
			ozlog.Warnf("Sconde Weight: %d", weight)

			time.Sleep(time.Duration(1) * time.Hour)
		}
		wg.Done()
	}()
	wg.Add(1)

	// 主玩家
	go func() {
		for {
			pocker, ok := <-MainPlayer
			if !ok {
				ozlog.Infof("main exit.")
				break
			}
			mainPokers = append(mainPokers, pocker)
		}
		wg.Done()
	}()
	wg.Add(1)

	// 玩家1
	go func() {
		for {
			pocker, ok := <-FirstPlayer
			if !ok {
				ozlog.Infof("first exit.")
				break
			}
			firstPokers = append(firstPokers, pocker)
		}
		wg.Done()
	}()
	wg.Add(1)

	// 玩家2
	go func() {
		for {
			pocker, ok := <-SecondPlayer
			if !ok {
				ozlog.Infof("second exit.")
				break
			}
			secondPokers = append(secondPokers, pocker)
		}
		wg.Done()
	}()
	wg.Add(1)
	wg.Wait()
}
