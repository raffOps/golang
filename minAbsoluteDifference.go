package main

import (
	"fmt"
	"math"
)

func quicksort(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	} else {
		pivotIndex := int(len(arr) / 2)
		pivot := arr[pivotIndex]
		var leftNumbers, rightNumbers, orderedArr []int32
		for i := 0; i < len(arr); i++ {
			if arr[i] <= pivot && i != pivotIndex {
				leftNumbers = append(leftNumbers, arr[i])
			} else if arr[i] > pivot && i != pivotIndex {
				rightNumbers = append(rightNumbers, arr[i])
			}
		}

		leftNumbers = quicksort(leftNumbers)
		rightNumbers = quicksort(rightNumbers)

		orderedArr = append(leftNumbers, pivot)
		for index, _ := range rightNumbers {
			orderedArr = append(orderedArr, rightNumbers[index])
		}
		return orderedArr
	}
}

func abs(number1 int64, number2 int64) int64 {
	if number1 >= number2 {
		return number1 - number2
	} else {
		return number2 - number1
	}
}

func minimumAbsoluteDifference(arr []int32) int64 {
	// Write your code here
	sortedArr := quicksort(arr)
	var minAbsoluteDifference int64 = math.MaxInt64
	for i := 0; i < len(sortedArr)-1; i++ {
		value := abs(int64(sortedArr[i]), int64(sortedArr[i+1]))
		if value < minAbsoluteDifference {
			minAbsoluteDifference = value
		}
	}
	return minAbsoluteDifference
}

func main() {
	arr := []int32{math.MaxInt32, math.MinInt32, 0}
	fmt.Printf("%d", minimumAbsoluteDifference(arr))
}
