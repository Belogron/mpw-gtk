package main

import (
	"github.com/conformal/gotk3/gtk"
	"log"
)

type MpwLoginGui struct {
	identiconLabel *gtk.Label
	userNameField *gtk.Entry
	passwordField *gtk.Entry
	box *gtk.Box
}


func CreateLoginGui() {

	gui := new(MpwLoginGui)
	gui.createGui()

}

func (gui MpwLoginGui) createGui() {

	loginWindow, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}
	loginWindow.SetTitle("MasterPasswordApp")
	loginWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})
	loginWindow.SetIconFromFile("lastpass")

	gui.createUserNameField()
	gui.createPasswordField()
	gui.createIdenticonLabel()
	gui.createBox()
	gui.box.Add(gui.identiconLabel)
	gui.box.Add(gui.userNameField)
	gui.box.Add(gui.passwordField)

	loginWindow.Add(gui.box)

	loginWindow.SetDefaultSize(400, 400)

	loginWindow.ShowAll()
}

func (gui MpwLoginGui) createBox() {
	gui.box, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	gui.box.SetMarginStart(15)
	gui.box.SetMarginEnd(15)
	gui.box.SetMarginTop(5)
	gui.box.SetMarginBottom(5)

}

func (gui MpwLoginGui) createUserNameField(){
	entry, _ := gtk.EntryNew()
	entry.SetPlaceholderText("Username")
	gui.userNameField = entry
}

func (gui MpwLoginGui) createPasswordField() {
	entry, _ := gtk.EntryNew()
	entry.SetPlaceholderText("Password")
	entry.SetInvisibleChar('●')
	entry.SetVisibility(false)
	gui.passwordField = entry
}

func (gui MpwLoginGui) createIdenticonLabel() {
	label, _ := gtk.LabelNew("aa")
	label.SetUseMarkup(true)
	label.SetMarkup("<span size=\"x-large\" color=\"#ff0000\">Test</span>")

	gui.identiconLabel = label
}
