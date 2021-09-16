#!/bin/bash
# 触るな
# Blenderをレンダリング
# https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html CYCLES BLENDER_EEVEE
sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E BLENDER_EEVEE --render-output img/anim --render-frame 2310
