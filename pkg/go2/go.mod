module github.com/takashabe/bazel-exercise/src/go2

go 1.14

require (
	github.com/go-sql-driver/mysql v1.4.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/takashabe/bazel-exercise/lib/go v0.0.0-20200609103620-5a28c182695b
	google.golang.org/appengine v1.6.6 // indirect
)

replace github.com/takashabe/bazel-exercise/lib/go => ../../lib/go
