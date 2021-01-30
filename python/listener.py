import keyboard
import subprocess
import threading
import time

is_playing = False
process = None


def fire(key_type):
    global is_playing
    global process

    print(key_type)
    if key_type != "equals":
        return

    if is_playing:
        process.kill()
        is_playing = False
        return

    is_playing = True
    # process = subprocess.Popen(["mplayer", "/home/pi/Music/FixYou.m4a"])
    process = subprocess.Popen(["./count.sh"])


def check_process():
    global process
    global is_playing

    while True:
        if process is not None and process.poll() is not None:
            is_playing = False
        time.sleep(3)


# keyboard.on_press_key("=", lambda _:print("pressed"))
# keyboard.on_press_key("q", lambda _:exit())
keyboard.add_hotkey("=", fire, args=["equals"])
keyboard.add_hotkey("windows,r", fire, args=["button"])

thread = threading.Thread(target=check_process)
thread.start()

keyboard.wait()

