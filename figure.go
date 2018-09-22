package main

import (
	"github.com/usthooz/oozlog/go"
)

// GetPlayerWeight 获取用户牌面权重
func GetPlayerWeight(playersPockers []string) int {
	/*
		1. 计算对子权重
		2. 计算顺子权重
		3. 计算飞机权重
		4. 计算三对权重(三带一)
		5. 计算炸弹权重
	*/
	var (
		// 权重
		weight int
	)
	twins := getTwinsNum(playersPockers)
	weight += twins * 2
	return weight
}

// getTwinsNum 时间复杂度 2logn output: 对子的总量
func getTwinsNum(as []string) int {
	var (
		pss []int
		w   int
	)
	// 取出牌面
	for _, a := range as {
		n, err := getPockerNumber(a)
		if err != nil {
			ozlog.Infof("getTwinsNum: get pocker num err->%v", err)
			continue
		}
		pss = append(pss, n)
	}
	for i := 0; i < len(pss); i++ {
		// 用于存储成对的数
		tran := make(map[int]int)
		// 用于减去成对重复的数
		sub := make(map[int]int)
		for _, a := range pss {
			// 找到对子，首先判断之前是否已经成对了，如果已经成了，那么需要拆散这一对
			_, ok := tran[a]
			if pss[i] == a && !ok {
				tran[i] = i
				w++
			} else if pss[i] == a && ok {
				_, ok := sub[i]
				// 如果已经减去了，则不再重复减少
				if !ok {
					sub[i] = i
					w--
				}
			}
		}
	}
	return w
}
