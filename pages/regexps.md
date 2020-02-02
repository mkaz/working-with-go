---
title: Regexp
layout: page
order: 21
---

# Regular Expressions

Use the [`regexp`](http://golang.org/pkg/regexp/) standard library to work with regular expressions.


## Basic Regexp

Create regular expression pattern to match 1 or more numbers

```go
pattern := "[0-9]+"

re, err := regexp.Compile(pattern)
if err != nil {
	fmt.Println("Error compiling regex", err)
}
```

### Test if Pattern Matches

Use `MatchString` to test if pattern matches

```go
str := "The 12 monkeys ate 48 bananas"
if re.MatchString(str) {
	fmt.Println("Yes, matched a number")
} else {
	fmt.Println("No, no match")
}
```

### Return Match

Use `FindString` to return match

```go
result := re.FindString(str)
fmt.Println("Number matched:", result)
```

### Return Multiple Matches

Use `FindAllString` to return multiple matches, specify `-1` to return all.

```go
results := re.FindAllString(str, 2)
for i, v := range results {
	fmt.Printf("Match %d: %s\n", i, v)
}
```

### Replace Match

Use `ReplaceAllString` to replace matches

```go
results := re.ReplaceAllString(str, "xx")
fmt.Println("Result:", results)
```

## Case Insensitive

The format for Go regular expression flags is different than typical. In Golang, the flag precedes the pattern, the syntax is, let's just say, not great. Here is an example using the case insensitive flag `i`

```go
ptn := `(?i)^t.`
str := "To be or not to be"

re, err := regexp.Compile(ptn)
if err != nil {
	fmt.Println("Error compiling regex", err)
}

// match string
result := re.FindString(str)
fmt.Println(result)
```

## Submatches

Submatches are what Go calls the matches that are grabbed by `(.*)` inside of a regular expression.

```go
str1 := "Hello @world@ Match"
sub_re, _ := regexp.Compile("@(.*)@")

m := sub_re.FindStringSubmatch(str1)
// FindStringSubmatch returns [@world@ world]
// so to just get the submatch you would use
if len(m) > 1 {
	fmt.Println(m[1]) // submatch
}
```

### Escaping Special Characters

If you wanted to match brackets or other special characters and try to just escape them like so `\[(.*)\]` you will get an error for unknown escape sequence `[`

You need to double up the slashes or a nicer solution is to use string literals and wrap in ticks instead of quotes

```go
str2 := "A [word] here and [there]"
esc_pattern := `\[(.*?)\]`
esc_re, _ := regexp.Compile(esc_pattern)

// this will only find the first
fmt.Println(esc_re.FindStringSubmatch(str2))

// use FindAll with second parameter for # of matches -1 = all
fmt.Println(esc_re.FindAllStringSubmatch(str2, -1))
```

