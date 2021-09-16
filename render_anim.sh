#!/bin/bash

# Blenderをレンダリング CYCLES
# https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html
sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E EEVEE --render-output img/anim -s 1000 -e 1030 -a

# 動画生成
sudo pip install opencv-python-headless==4.4.0.44
sudo python3 pngtomp4.py