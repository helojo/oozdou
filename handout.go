package main

import (
	"math/rand"
	"time"

	"github.com/usthooz/oozlog/go"
)

var (
	pockers []string
)

// InitPockers 初始化得到所有牌
func InitPockers(confPockers []string) {
	pockers = confPockers
}

// Handout 发牌
func Handout() {
	// 取出三张地主牌
	for i := 0; i < 3; i++ {
		index := genRandCountForDiff(pockers)
		lordTokers = append(lordTokers, pockers[index])
		// 从牌库移除地主牌
		pockers = append(pockers[:index], pockers[index+1:]...)
	}
	length := len(pockers)
	// 发放普通牌
	for i := 0; i < length; i++ {
		ozlog.Infof("len: %d", len(pockers))
		// 随机获取一张牌
		index := genRandCountForDiff(pockers)
		ozlog.Infof("pock: %v", pockers)
		if i%3 == 0 {
			// 主玩家
			MainPlayer <- pockers[index]
		} else if i%2 == 0 {
			// 玩家2
			SecondPlayer <- pockers[index]
		} else {
			// 玩家1
			FirstPlayer <- pockers[index]
		}
		// 从牌库移除地主牌
		pockers = append(pockers[:index], pockers[index+1:]...)
	}
	ozlog.Infof("发牌任务完成.")
	return
}

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

// genRandCountForDiff 获取数组长度内的数字
func genRandCountForDiff(keys []string) int {
	return r.Intn(len(keys))
}
