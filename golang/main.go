package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"io/ioutil"
	"jaysonhelseth/gopictureframe/driver"
	"log"
	"math/rand"
	"time"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Picture Frame")

	image := canvas.NewImageFromFile(getImage())
	image.FillMode = canvas.ImageFillContain
	w.SetContent(image)

	myApp.Settings().SetTheme(myTheme{})
	//w.SetFullScreen(true)
	w.Resize(fyne.NewSize(200, 200))

	go func() {
		for {
			time.Sleep(time.Second * 5)
			image.File = getImage()
			image.Refresh()
		}
	}()

	go driver.Start()

	w.ShowAndRun()
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
