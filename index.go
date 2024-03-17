package main
import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)
func main(){
	myApp := app.New()
	myWindow := myApp.NewWindow("hello boy")
	label := myApp.NewWindow("Hello, World!")
	myWindow.SetContent(label)
	myWindow.ShowAndRun()
}
