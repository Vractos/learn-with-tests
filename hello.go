package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

const englishHelloPrefix = "Hello, "

func HelloTDD(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
}
