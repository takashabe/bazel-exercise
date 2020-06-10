test: deps
	bazel test --test_output=errors //...

deps:
	bazel run //:gazelle -- update
