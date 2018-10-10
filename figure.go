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
	one, two, three, four := trimBoardGroup(playersPockers)
	weight += len(one) + len(two)*2 + len(three)*3 + len(four)*4
	weight += getKingBombNum(playersPockers)
	return weight
}

// getKingBombNum 王炸权重计算
func getKingBombNum(playersPockers []string) (weight int) {
	b := binarySearch(playersPockers, BigKing)
	s := binarySearch(playersPockers, SmallKing)
	if b >= 0 && s >= 0 {
		weight = KingBombCount
		return
	}
	if b >= 0 {
		weight = BigKingCount
		return
	}
	if s >= 0 {
		weight = SmallKingCount
		return
	}
	return
}

// getTwinsNum 时间复杂度 2logn output: 单个对子的总量
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

// trimBoardGroup 洗牌
func trimBoardGroup(boards []string) (map[int][]int, map[int][]int, map[int][]int, map[int][]int) {
	// 查找是否存在大小王
	b := binarySearch(boards, BigKing)
	if b >= 0 {
		boards = append(boards[:b], boards[b+1:]...)
	}
	s := binarySearch(boards, SmallKing)
	if s >= 0 {
		boards = append(boards[:s], boards[s+1:]...)
	}
	bs := sortBoard(boards)
	return trimTwinsBoards(bs)
}

// sortBoard 洗牌-按照牌面大小将所有牌从小到大排列
func sortBoard(boards []string) []int {
	var (
		pss []int
	)
	// 取出牌面
	for _, a := range boards {
		n, err := getPockerNumber(a)
		if err != nil {
			ozlog.Infof("getTwinsNum: get pocker num err->%v", err)
			continue
		}
		pss = append(pss, n)
	}
	return sectionSort(pss)
}

// trimTwinsBoards 整理连续的牌(2、3、4张)及单牌
func trimTwinsBoards(boards []int) (map[int][]int, map[int][]int, map[int][]int, map[int][]int) {
	var (
		curLen                                           int
		singleBoards, twoBoards, threeBoards, fourBoards map[int][]int
	)
	singleBoards = make(map[int][]int)
	twoBoards = make(map[int][]int)
	threeBoards = make(map[int][]int)
	fourBoards = make(map[int][]int)

	length := len(boards)
	for i := 0; i < length; i++ {
		if i < curLen {
			continue
		}
		if i+1 >= length {
			break
		}
		if boards[i] == boards[i+1] {
			if i+2 >= length {
				break
			}
			if boards[i] == boards[i+2] {
				if i+3 >= length {
					break
				}
				if boards[i] == boards[i+3] {
					fourBoards[4] = append(fourBoards[4], boards[i], boards[i+1], boards[i+2], boards[1+3])
					curLen += 4
					continue
				}
				threeBoards[3] = append(threeBoards[3], boards[i], boards[i+1], boards[i+2])
				curLen += 3
				continue
			}
			twoBoards[2] = append(twoBoards[3], boards[i], boards[i+1])
			curLen += 2
			continue
		}
		singleBoards[1] = append(singleBoards[1], boards[i])
		curLen += 1
	}
	return singleBoards, twoBoards, threeBoards, fourBoards
}
