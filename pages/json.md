---
title: JSON
layout: page
order: 16
---

# Working with JSON

Use the [`encoding/json`](http://golang.org/pkg/encoding/json/) standard library for encoding to and decoding from JSON.


## Encoding to JSON

Encoding a Go struct to JSON format uses the [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) method, this returns a `[]byte` of the struct.

```go
// Create a struct to match the format of JSON
type Person struct {
	Name string
	City string
}

p := Person{ Name: "Marcus", City: "San Francisco" }

json, err := json.Marshal(p)
if err != nil {
	fmt.Println("JSON Encoding Error", err)
}

fmt.Println(string(json))
```

## Decode from JSON

To decode a JSON string to a Go struct, use the [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) method. Unmarshal accepts the `[]byte` of the JSON string to decode, and a pointer to the struct that matches the data structure.

```go
json_str := `{ "Name": "Marcus", "City": "San Jose"}`
var p Person

if err := json.Unmarshal([]byte(json_str), &p); err != nil {
	fmt.Println("Error parsing JSON: ", err)
}

// output result
fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)
```

### Partial Data

If a variable is already defined with values, the Unmarshaling will "fill in" the variable with the additional values from JSON.

### Read JSON from a File

Since `Unmarshal` expects a `[]byte` to be read, reading from a file is straight-forward.

```go
var people []Person

file, err := ioutil.ReadFile("names.json")
if err != nil {
	fmt.Println("Error reading file")
}

// the names.json file has an array of person objects, so read into people
if err := json.Unmarshal(file, &people); err != nil {
	fmt.Println("Error parsing JSON", err)
}

// output result
fmt.Println(people)
```

## Mapping Struct and JSON Field

The JSON field names may not always match the struct names. Go will automatically manage uppercase and lowercase fields, but if they are completely different you can use struct field tags to label.

```go
type Person {
	Name string `json:"username"`
	City string `json:"municipality"`
}
```

By defining the struct with the field tags, the Marshal and Unmarshaling will take the Go struct and write them to those JSON fields.

```go
p := Person{ Name: "Marcus", City: "San Francisco" }
json, _ := json.Marshal(p)
fmt.Println(string(json))

>> { "username": "Marcus", "municipality": "San Francisco" }
```

### Ignore Field

You can ignore a field from being encoded or decoded to JSON, using `-` definition.

```go
type Person {
	Name string
	City string
	Phone string `json:"-"`
}
```

The `Phone` field will be omitted from JSON operations.

### Omit Empty Fields

Set a field to be ignored when empty, use `omitempty` so JSON encoding will not include the field.

```go
type Person {
	Name string
	City string
	Phone string `json:",omitempty"`
}
```

