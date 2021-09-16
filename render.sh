#!/bin/bash
# 触るな
# Blenderをレンダリング
# https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html
sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E CYCLES --render-output img/anim --render-frame 1510
