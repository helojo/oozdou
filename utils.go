package main

import (
	"strconv"

	"github.com/usthooz/oozlog/go"
)

/*
	ComparePocker: 比较单牌的大小
	params: a1,a2
	if a1>a2 return 1
	if a1<a2 return -1
	else return 0

	if return 2 -> error
*/
func ComparePocker(a1, a2 string) int {
	// a1为大王
	if a1 == BigKing {
		return 1
	}
	// a2为大王
	if a2 == BigKing {
		return -1
	}
	// a1为小王
	if a1 == SmallKing {
		return 1
	}
	// a2为小王
	if a2 == SmallKing {
		return -1
	}
	// 获取牌面数字
	a1num, err := getPockerNumber(a1)
	if err != nil {
		ozlog.Fatalf("ComparePocker: get pocker num err->%v", err)
		return 2
	}
	a2num, err := getPockerNumber(a2)
	if err != nil {
		ozlog.Fatalf("ComparePocker: get pocker num err->%v", err)
		return 2
	}
	if a1num > a2num {
		return 1
	} else if a1num == a2num {
		return 0
	} else {
		return -1
	}
	return 2
}

// getPockerNumber 获取牌面数字
func getPockerNumber(a string) (int, error) {
	res, err := strconv.Atoi(a[1:len(a)])
	if err != nil {
		ozlog.Errorf("getPockerNumber: get pocker num err->%v", err)
		return 0, err
	}
	return res, nil
}
