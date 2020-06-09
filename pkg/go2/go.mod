module github.com/takashabe/bazel-exercise/src/go2

go 1.14

require (
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/takashabe/bazel-exercise/lib/go v0.0.0-20200609103620-5a28c182695b
)

replace github.com/takashabe/bazel-exercise/lib/go => ../../lib/go
