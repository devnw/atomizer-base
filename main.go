// Copyright © 2021 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"fmt"
	"os"

	"atomizer.io/base/atoms"
	"atomizer.io/cmd"
	"atomizer.io/engine"
)

func init() {
	engine.Register(&atoms.Plugin{})
}

func main() {
	err := cmd.Initialize("Atomizer Test Agent")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
