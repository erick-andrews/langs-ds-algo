package main

import (
	"fmt"
)

var a1 = []string{"a", "b", "c", "d"}
var a2 = []string{"z", "y", "i"}

var a3 = []string{"a", "b", "c", "d", "x"}
var a4 = []string{"z", "y", "x"}

func main() {
	containsOverlap := containsCommonItem(a3, a4)
	fmt.Printf("These arrays contain an overlap: %v\n", containsOverlap)

	genericOverlapTest := containsCommonItemGeneric(a1, a2)
	fmt.Printf("These arrays contain an overlap: %v\n", genericOverlapTest)

}

// func containsCommonItem(array1 []string, array2 []string) bool {
// 	// brute force O^2 nested for loop - not ideal, but start with brute force.
// 	// Naive, bad approach, but baseline nonetheless.
// 	// declare boolean return
// 	overlap := false
// 	for _, str1 := range array1 {
// 		for _, str2 := range array2 {
// 			if str1 == str2 {
// 				overlap := true
// 				return overlap
// 			}
// 		}
// 	}
// 	return overlap
// }

func containsCommonItem(array1 []string, array2 []string) bool {
	// this solution has O(a + b) time complexity.
	// Faster than above.
	// declare boolean return
	overlap := false
	// thinking... array to map. O(n)
	//array1 -> object{
	// a:true
	// b:true
	//c: true
	//x: true
	//}
	// array2[index] == obj.properties

	var array1Map = make(map[string]bool, len(array1))
	// loop through array and create map from array w/ bools
	for _, key := range array1 {
		array1Map[key] = true
	}
	// loop through second array and check if second array exists on created object
	for _, element := range array2 {
		if array1Map[element] {
			overlap := true
			return overlap
		}
	}
	return overlap
}

// Creating a generic to deal with int comparisons, or string comparisons, etc.
func containsCommonItemGeneric[T comparable](array1 []T, array2 []T) bool {
	// Take in our array1, make map
	array1Map := make(map[T]bool, len(array1))
	for _, key := range array1 {
		array1Map[key] = true
	}
	// compare array2 to map.
	for _, element := range array2 {
		if array1Map[element] {
			return true
		}
	}

	return false

}
