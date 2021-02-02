package events

import (
	"log"
	"os/exec"
)

var (
	mplayer *exec.Cmd
	isPlaying bool = false
)

func TogglePlayer() {
	if isPlaying {
		mplayer.Process.Kill()
		mplayer.Wait()
		isPlaying = false
		return
	}

	// This goroutine will run until it finishes or is killed
	go func() {
		// Run the audio file 9 times, which is around 40ish minutes of music.
		mplayer = exec.Command("mplayer", "-volume", "100", "-loop", "9", "/home/jayson/Music/FixYou.m4a")
		isPlaying = true
		err := mplayer.Run()
		if err != nil {
			if killedError, otherError := err.(*exec.ExitError); otherError {
				log.Println("Killed the mplayer")
			} else {
				log.Fatal(killedError)
			}
		}
		// When done playing mark this as done.
		isPlaying = false
	}()
}
