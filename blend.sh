#!/bin/bash
#python
sudo apt-get install -y -q python3 
# aptアップデート,snapインストール,Blenderインストール
sudo apt -y update
sudo apt list --upgradable
sudo apt install snapd
sudo snap install blender --classic
