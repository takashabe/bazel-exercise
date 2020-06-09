package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_A(t *testing.T) {
	assert.Equal(t, "", "")
	logrus.Infof("%s", "hoge")
}
