package main

import (
//"github.com/axgle/mahonia"
	"fmt"
	"os/exec"
//"bytes"
	"os"
	"time"
//"io/ioutil"
	"strconv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Common flags for the command line renderer
//https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-0280AB86-8ABE-4F75-B1B9-D5B7DBB7E25A-htm.html

//Render from the command line using mental ray
//https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-FCB01DEF-10DF-4222-A047-AF0122782397-htm.html

type RenderSettings struct {
	ProjectPath       string `yaml:"ProjectPath"`
	SceneName         string `yaml:"SceneName"`
	CameraName        string `yaml:"CameraName"`
	ImageName         string
	PercentResolution int `yaml:"PercentResolution"` // 50 means 50%
}

func main() {
	rs := RenderSettings{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &rs)
	if err != nil {
		fmt.Println("error: %v", err)
	}

	rs.ImageName = rs.SceneName + " " + time.Now().Format("2006-01-02 150405")
	fmt.Printf("%+v\n", rs)

	RenderCommand(rs)
	ConvertCommand(rs.ImageName)
	ConvertCommand2(rs.ImageName)

	var enterToExit string
	fmt.Println("press Enter to exit")
	fmt.Scanf(enterToExit)
}

//var printedOut bytes.Buffer
//var cmdOut bytes.Buffer
//
//func PrintOutRealtime() {
//	//for {
//		//time.Sleep(time.Millisecond * 100)
//		//if 0 != cmdOut.Len() {
//			//fmt.Print(cmdOut.String())
//			//cmdOut.Reset()
//		//}
//	//}
//}

// gamma 2.2
func ConvertCommand(imageName string) {
	program := "ImageMagick-6.9.3-0-portable-Q16-x64/convert.exe"
	dir, _ := os.Getwd()
	idir := dir + "/images"
	imageFullName := idir + "/" + imageName + ".png"
	cmd := exec.Command(program,
		imageFullName,
		"-gamma", "2.2",
		imageFullName)
	CommandRun(cmd)
}

// png to jpg
func ConvertCommand2(imageName string) {
	program := "ImageMagick-6.9.3-0-portable-Q16-x64/convert.exe"
	dir, _ := os.Getwd()
	idir := dir + "/images"
	imageFullNameWithoutSuffix := idir + "/" + imageName
	cmd := exec.Command(program,
		imageFullNameWithoutSuffix + ".png",
		"-quality", "92",
		imageFullNameWithoutSuffix + ".jpg")
	CommandRun(cmd)
}

func RenderCommand(rs RenderSettings) {
	dir, _ := os.Getwd()
	program := "C:/Program Files/Autodesk/Maya2016/bin/Render.exe"
	//imageName := rs.SceneName + " " + time.Now().Format("2006-01-02 150405")

	cmd := exec.Command(program,
		"-r", "mr",
		"-rd", dir + "/images",
		"-im", rs.ImageName,
		"-of", "png",
		"-percentRes", strconv.Itoa(rs.PercentResolution),
		"-proj", rs.ProjectPath,
		"-cam", rs.CameraName,
		rs.ProjectPath + "/scenes/" + rs.SceneName + ".mb")

	//cmd := exec.Command("D:\\projects\\tickers.exe")
	startTime := time.Now()
	CommandRun(cmd)
	endTime := time.Now()
	dur := endTime.Sub(startTime)
	fmt.Print(dur.Minutes())
	fmt.Println(" minitues")
	//if err != nil {
	//	fmt.Println(err)
	//}
}
func CommandRun(cmd *exec.Cmd) {
	fmt.Println("cmd.Path = " + cmd.Path)
	fmt.Println("cmd.Args = ")
	fmt.Println(cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	//go PrintOutRealtime()
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)

	//fmt.Println(cmdOut.String())

	//write to file to read
	//f, err := os.Create("out.txt")
	//defer f.Close()
	//f.WriteString(os.Stdout. cmdOut.String())

}

func StringAddDoubleQuotes(str string) string {
	return string('"') + str + string('"')
}
