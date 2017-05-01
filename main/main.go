//main.go

package main

import (
	"github.com/brentritzema/go-cs214/handler"
	"github.com/brentritzema/go-cs214/server"
	"github.com/mattn/go-gtk/gtk"
)

const (
	connHost = ""
	connPort = "6785"
	connType = "tcp"
)

func main() {

	//Setup GTK
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("textview")
	window.Connect("destroy", gtk.MainQuit)

	//GTKVBox
	vbox := gtk.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.NewFrame("Demo")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	vpaned.Pack1(frame1, false, false)

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk.NewStatusbar()
	contextId := statusbar.GetContextId("go-gtk")
	statusbar.Push(contextId, "Not Running")

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------

	//make a channel to hold state message
	quit := make(chan bool)

	// defer close(quit)

	togglebutton := gtk.NewToggleButtonWithLabel("Click to Start")
	togglebutton.Connect("toggled", func() {
		if togglebutton.GetActive() {

			togglebutton.SetLabel("Stop")
			statusbar.Push(contextId, "Running")
			//start listener
			go server.StartListener(
				connHost,
				connPort,
				connType,
				handler.ProcessConnection,
				true,
				quit)
		} else {
			quit <- true
			togglebutton.SetLabel("Click to Start!")
			statusbar.Push(contextId, "Stopped")
		}
	})
	buttons.Add(togglebutton)

	framebox1.PackStart(buttons, false, false, 0)

	framebox1.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(200, 200)
	window.ShowAll()
	gtk.Main()

}
