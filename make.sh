#!/bin/sh

sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E CYCLES --render-output img/anim -s 1976 -e 1976 -a

sudo pip install opencv-python-headless==4.4.0.44

sudo python3 pngtomp4.py