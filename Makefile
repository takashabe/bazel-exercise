test: deps
	bazel test //...

deps:
	bazel run //:gazelle -- update
