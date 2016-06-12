package main

import (
	"github.com/conformal/gotk3/gtk"
)

func main() {

	gtk.Init(nil)

	CreateLoginGui()

	gtk.Main()
}
