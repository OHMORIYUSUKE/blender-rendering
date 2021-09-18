#!/bin/sh

sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E CYCLES --render-output img/anim -s 1980 -e 1999 -a

sudo apt-get install -y -q python3

sudo pip install opencv-python-headless==4.4.0.44

sudo python3 pngtomp4.py