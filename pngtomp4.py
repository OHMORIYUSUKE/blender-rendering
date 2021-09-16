import sys
import cv2

def main():
    # encoder(for mp4)
    fourcc = cv2.VideoWriter_fourcc('m', 'p', '4', 'v')
    # output file name, encoder, fps, size(fit to image size)
    video = cv2.VideoWriter('video.mp4',fourcc, 20.0, (1240, 1360))

    if not video.isOpened():
        print("can't be opened")
        sys.exit()

    for i in range(0, 10+1):
        # hoge0000.png, hoge0001.png,..., hoge0090.png
        img = cv2.imread('./img/anim%04d.png' % i)

        # can't read image, escape
        if img is None:
            print("can't read")
            break

        # add
        video.write(img)
        print(i)

    video.release()
    print('written')

if __name__=="__main__":
    main()