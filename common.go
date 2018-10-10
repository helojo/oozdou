package main

// binarySearch 二分查找
func binarySearch(array []string, key string) int {
	low, hight := 0, len(array)-1
	for low <= hight {
		m := (low + hight) >> 1
		if array[m] < key {
			low = m + 1
		} else if array[m] > key {
			hight = m - 1
		} else {
			// 相等
			return m
		}
	}
	return -1
}

// sectionSort 选择排序
func sectionSort(values []int) []int {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}
