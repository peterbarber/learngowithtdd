package main

const englishHelloPrefix = "Hello, "

func Hello(name string) (returnString string) {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
