package handler

import (
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

//nolint:golint,unused
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	extra.RegisterFuzzyDecoders()
}
