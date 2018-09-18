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
	for i := 0; i < 2; i++ {
		index := genRandCountForDiff(pockers)
		lordTokers = append(lordTokers, pockers[index])
		// 从牌库移除地主牌
		pockers = append(pockers[:index], pockers[index+1:]...)
	}
	// 发放普通牌
	for i := 0; i < len(pockers); i++ {
		// 随机湖区一张牌
		index := genRandCountForDiff(pockers)
		if 0 <= index && index < len(pockers) {
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
	}
	ozlog.Infof("发牌任务完成.")
	return
}

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

// genRandCountForDiff 生成指定范围内的指定个数(不同的数字)
func genRandCountForDiff(keys []string) int {
	return r.Intn(len(keys))
}
