---
title: Sorting
layout: page
order: 15
---

# Sorting Arrays and Maps

The [`sort`](https://golang.org/pkg/sort/) standard library is used for sorting slices, maps, and user define collections.

## Sort Alphabetical

Use `sort.Strings()` to sort alphabetically

```go
abc := []string{"jkl", "ghi", "abc", "def"}
sort.Strings(abc)
fmt.Println("Sorted ABC:", abc)
```

## Sort Numeric

Use `sort.Ints()` to sort numerically

```go
nums := []int{4, 2, 12, 5, 1, 3}
sort.Ints(nums)
fmt.Println("Sorted Nums:", nums)
```

## Reverse Sort

Reverse sort, need to cast `abc` as a StringSlice to reverse and sort

```go
sort.Sort(sort.Reverse(sort.StringSlice(abc)))
fmt.Println("Reverse ABC:", abc)
```

## Sort Maps

To sort a map by keys, pull an array of keys out of the map, sort the keys and then iterate over the sorted keys. A map by itself is not sortable.

```go
hash := map[string]int{
	"c": 3,
	"a": 1,
	"b": 2,
	"e": 5,
	"d": 4,
}

// Create array of kys
var keys []string
for k := range hash {
	keys = append(keys, k)
}

// Sort keys
sort.Strings(keys)

// Use ordered keys to loop through hash
for i := range keys {
	fmt.Printf("%s => %v\n", keys[i], hash[keys[i]])
}
```
