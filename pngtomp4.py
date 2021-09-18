import glob
import cv2

def main():
    img_array = []
    for filename in sorted(glob.glob("img/*.png")):
        img = cv2.imread(filename)
        height, width, layers = img.shape
        size = (width, height)
        img_array.append(img)

    name = 'blender.mp4'
    out = cv2.VideoWriter(name, cv2.VideoWriter_fourcc(*'MP4V'), 5.0, size)

    for i in range(len(img_array)):
        out.write(img_array[i])
    out.release()

if __name__=="__main__":
    main()