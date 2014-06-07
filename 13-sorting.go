/**
 * sorting.go
 *
 * Example program on how to sort various arrays and maps
 * See: http://golang.org/pkg/sort/
 */

package main

import (
	"fmt"
	"sort"
)

func main() {

	// sort alphabetical
	abc := []string{"jkl", "ghi", "abc", "def"}
	nums := []int{4, 2, 12, 5, 1, 3}

	// show default slices out of order
	fmt.Println("Raw ABC : ", abc)
	fmt.Println("Raw Nums: ", nums)

	// sort using Strings() method
	sort.Strings(abc)
	fmt.Println("Sorted ABC:", abc)

	// sort using Ints() method
	sort.Ints(nums)
	fmt.Println("Sorted Nums:", nums)

	// reverse sort, need to cast abc as a StringSlice reverse and sort
	sort.Sort(sort.Reverse(sort.StringSlice(abc)))
	fmt.Println("Reverse ABC:", abc)

	//-------------------------------------------------
	// Sort Maps
	//-------------------------------------------------
	// sort map by key
	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	// view map to see its out of order
	for k, v := range hash {
		fmt.Printf("%s => %v\n", k, v)
	}

	// to sort, create array of keys and sort array
	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// display hash by looping ordered keys
	for i := range keys {
		fmt.Printf("%s => %v\n", keys[i], hash[keys[i]])
	}

}
