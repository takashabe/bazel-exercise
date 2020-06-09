package main

import (
	logger "github.com/takashabe/bazel-exercise/lib/go"
)

func main() {}

func A() string {
	return "A"
}

func B() {
	logger.Info("hoge")
}
