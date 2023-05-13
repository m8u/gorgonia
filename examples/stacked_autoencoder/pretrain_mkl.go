//go:build !native
// +build !native

package main

import (
	. "github.com/m8u/gorgonia"
	"github.com/m8u/gorgonia/blase"
)

func init() {
	Use(blase.Implementation())
}
