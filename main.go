// Create: 2019/05/05 16:53:00 Change: 2019/06/05 11:52:04
// FileName: main.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"example.com/hello/display"
	pkg "github.com/introclass/go_mod_example_pkg"
	pkgv3 "github.com/introclass/go_mod_example_pkg/v3"
	"github.com/lijiaocn/golib/version"
)

func main() {
	version.Show()
	display.Display("display print\n")
	pkg.Vesrion()
	pkgv3.Vesrion()
}
