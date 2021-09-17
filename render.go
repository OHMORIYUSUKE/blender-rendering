package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func toInt64(strVal string) int64 {
	rex := regexp.MustCompile("[0-9]+")
	strVal = rex.FindString(strVal)
	intVal, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intVal
}

func Int64ToInt(i int64) int {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	} else {
		return int(i)
	}
}

type Config struct {
	Name   string
	Number int
}

func main() {
	lines := fromFile("renderConfig.txt")

	var start Config
	var end Config
	var video Config
	var file Config

	for _, line := range lines {
		number64 := toInt64(line)
		var intNumber int = Int64ToInt(number64)
		sNumber := strconv.Itoa(intNumber)
		sNumber = "=" + sNumber
		LineString := strings.Replace(line, sNumber, "", 1)

		intNumber = Int64ToInt(number64)

		//--
		r := regexp.MustCompile(`FILENAME`)
		//fmt.Println(r.MatchString(LineString)) // => true
		//--
		if r.MatchString(LineString) {
			fileName := strings.Replace(LineString, "FILENAME=", "", 1)
			//fmt.Println(fileName)
			file = Config{Name: fileName, Number: 1}
		} else if LineString == "START" {
			start = Config{Name: LineString, Number: intNumber}
			// fmt.Println(start.Name , start.Number)
		} else if LineString == "END" {
			end = Config{Name: LineString, Number: intNumber}
			// fmt.Println(end.Name , end.Number)
		} else if LineString == "VIDEO" {
			video = Config{Name: LineString, Number: intNumber}
			// fmt.Println(video.Name , video.Number)
		}
	}
	fmt.Println(file.Name)
	fmt.Println(start.Name, start.Number)
	fmt.Println(end.Name, end.Number)
	fmt.Println(video.Name, video.Number)

	var StrStartNum string
	StrStartNum = strconv.Itoa(start.Number)
	var StrEndNum string
	StrEndNum = strconv.Itoa(start.Number)

	//--コマンド実行--
	// sudo blender --background -noaudio blend/Miraikomachi.blend --threads 0 -E CYCLES --render-output img/anim -s 2330 -e 2331 -a

	// cmd := exec.Command("sudo blender --background -noaudio blend/" + file.Name + " --threads 0 -E CYCLES --render-output img/anim" + "-s " + StrStartNum + " -e " + StrEndNum + " -a")
	// cmd.Start()
	// cmd.Wait()
	var out []byte
	out, err := exec.Command("sudo","blender --background -noaudio blend/" + file.Name + " --threads 0 -E CYCLES --render-output img/anim" + "-s " + StrStartNum + " -e " + StrEndNum + " -a").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result:%s\n", string(out))
	if video.Number == 1 {
		cmd := exec.Command("sudo", "apt-get install -y -q python3 ")
		cmd.Start()
		cmd.Wait()
		cmd1 := exec.Command("sudo", "pip install opencv-python-headless==4.4.0.44")
		cmd1.Start()
		cmd1.Wait()
		cmd2 := exec.Command("sudo", "python3 pngtomp4.py")
		cmd2.Start()
		cmd2.Wait()
	}

}

func fromFile(filePath string) []string {
	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
		os.Exit(1)
	}

	// 関数return時に閉じる
	defer f.Close()

	// Scannerで読み込む
	// lines := []string{}
	lines := make([]string, 0, 100) // ある程度行数が事前に見積もれるようであれば、makeで初期capacityを指定して予めメモリを確保しておくことが望ましい
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// appendで追加
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}

	return lines
}
