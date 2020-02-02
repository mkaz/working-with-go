---
title: Maps
layout: page
order: 8
---

# Working with Maps

A map is an associative array data type. You can define a map that uses any data type for the key or value.

## Make a Map

Map uses the built-in `make` function to initialize, this creates an empty map with no keys.

```go
m := make(map[string]string)
m["c"] = "Cyan"
m["y"] = "Yellow"
m["m"] = "Magenta"
m["k"] = "Black"
```

You can also create the values on initialization using.

```go
var m = map[string]string{
	"c": "Cyan",
	"y": "Yellow",
	"m": "Magenta",
	"k": "Black",
}
```

## Iterate over a Map

Use the same `range` function for iterating over a map, it returns the `key, value` pair.

```go
for k, v := range m {
	fmt.Printf("Key: %s, Value: %s", k, v)
}
```

## Retrieve a Map Item

Get a single map item using brackets

```go
c = m["c"]
```


## Delete a Map Item

Use the `delete` built-in function, passing map and key.

```go
delete(m, "c")
```

## Test Item Exists

Test if a map contains an item, by checking the second value returned when fetching an item. If the item does not exists, it will be false.

If the item does not exist the value returned will be the zero value, however the item may exist as the zero value in the map, so it is best practice to check the second value returned.

```go
c, ok = m["p"]
```


