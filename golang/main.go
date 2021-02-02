package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"jaysonhelseth/gopictureframe/events"
	"jaysonhelseth/gopictureframe/shared"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	shared.AppWindow = myApp.NewWindow("Picture Frame")

	image := canvas.NewImageFromFile(getImage())
	image.FillMode = canvas.ImageFillContain
	shared.AppWindow.SetContent(image)

	myApp.Settings().SetTheme(myTheme{})
	shared.AppWindow.SetFullScreen(true)
	//shared.AppWindow.Resize(fyne.NewSize(500, 500))

	go func() {
		for {
			time.Sleep(time.Second * 5)
			image.File = getImage()
			image.Refresh()
		}
	}()

	shared.AppWindow.Canvas().SetOnTypedRune(events.KeyListener)
	shared.AppWindow.ShowAndRun()
}

func getImage() string {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(files))
	name := files[index].Name()

	return fmt.Sprintf("images/%s", name)
}
