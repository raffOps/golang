package main

import "fmt"

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

func printArrayInt32(arr []int32) {
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	println("\n")
}

func main() {
	arr := []int32{5, 5, 5, 5, 5}
	printArrayInt32(arr)

	orderedArr := quicksort(arr)
	printArrayInt32(orderedArr)

}
