def main():
    dict_cmd = {}

    fileobj = open("renderConfig.txt", "r", encoding="utf_8")
    while True:
        line = fileobj.readline()
        if line:
            dict_cmd[line.split('=')[0]] = line.split('=')[1].replace('\n', '')
        else:
            break

    cmdString = "#!/bin/sh"
    # Pythonを使う場合
    if dict_cmd['PYTHONFILENAME'] != '0':
        cmdString = cmdString + "\n\n"
        cmdString = cmdString + "sudo blender --background -noaudio blend/" + dict_cmd['FILENAME'] + " --threads 0 -E CYCLES --python python/" + dict_cmd['PYTHONFILENAME'] + " --render-output img/anim" + " -s " + dict_cmd['START'] + " -e " + dict_cmd['END'] + " -a"
    else:
        cmdString = cmdString + "\n\n"
        cmdString = cmdString + "sudo blender --background -noaudio blend/" + dict_cmd['FILENAME'] + " --threads 0 -E CYCLES --render-output img/anim" + " -s " + dict_cmd['START'] + " -e " + dict_cmd['END'] + " -a"
    # 動画を出力する場合
    if dict_cmd['VIDEO'] == '1':
        cmdString = cmdString + "\n\n"
        cmdString = cmdString + "sudo pip install opencv-python-headless==4.4.0.44"
        cmdString = cmdString + "\n\n"
        cmdString = cmdString + "sudo python3 pngtomp4.py"


    text_file = open("make.sh", "wt")
    text_file.write(cmdString)

if __name__=="__main__":
    main()