# Go Picture Frame

make sure you create an images folder and place your images in there. 

The project uses Fyne.io and fyne-cross to build for the raspberry pi.
```
# Make sure to be in the picture frame directory
fyne-cross linux -arch=arm
```

## Desktop Shortcut

* If you are using a touchscreen enable single click desktop applications.
* You may also need to disable the menus that popup when selecting a desktop shortcut.
* Create a `.desktop` file in your ~/Desktop folder.
* You may need to do `chmod +x` on your new desktop file.
Here is an example of a Desktop Shortcut
```
[Desktop Entry]
Version=1.0
Type=Application
Encoding=UTF-8
Name=Picture Frame
Comment=Picture Frame
# Icon=/home/pi/Downloads/icon.png
Exec=/home/pi/projects/pictureframe/slideshow.sh
Terminal=false
Categories=Graphics
```