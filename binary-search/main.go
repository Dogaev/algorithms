package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func main() {
	x := 7
	arr := []int{1,2,3,4,5,6,7,8,9,10,11}
	pos, err := BinarySearch(x, arr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("the position of val %d in array %v is %d\n", x, arr, pos)
}

// BinarySearch - бинарный поиск по отсортированному массиву O(logN)
func BinarySearch(x int, arr []int) (pos int, err error) {
	if !sort.IntsAreSorted(arr) {
		return 0, errors.New("invalid parameters: array doesn't sorted")
	}

	var low = 0
	var end = len(arr)-1
	for low = 0; low <= end; {
		mid := (low+end)/2 // при делении округляется в меньшую сторону
		switch {
		case x == arr[mid]:
			return mid+1, nil
		case x > arr[mid]:
			low = mid
		default:
			end = mid
		}
		continue
	}

	return 0, fmt.Errorf("can not found %d in array", x)
}