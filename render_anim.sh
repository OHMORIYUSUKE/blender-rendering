#!/bin/bash

# Blenderをレンダリング
# https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html
sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E CYCLES --render-output img/anim -s 0 -e 500 -a

# 動画生成
sudo pip install opencv-python-headless==4.4.0.44
sudo python3 pngtomp4.py