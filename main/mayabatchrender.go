package main

import (
//"github.com/axgle/mahonia"
	"os/exec"
	"fmt"
//"bytes"
	"time"
	"os"
//"io/ioutil"
	"strconv"
)


//Common flags for the command line renderer
//https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-0280AB86-8ABE-4F75-B1B9-D5B7DBB7E25A-htm.html

//Render from the command line using mental ray
//https://knowledge.autodesk.com/support/maya/learn-explore/caas/CloudHelp/cloudhelp/2016/ENU/Maya/files/GUID-FCB01DEF-10DF-4222-A047-AF0122782397-htm.html

//convert "E:\My\Desktop\Pal2\art\maya\Songzhou Outskirts\images\Outskirts.png" -gamma 2.2 "E:\My\Desktop\Pal2\art\maya\Songzhou Outskirts\images\Outskirtsgammad.png"

//"C:\Program Files\Autodesk\Maya2016\bin\Render.exe" -r mr -proj "E:\My\Desktop\Pal2\art\maya\Songzhou Outskirts" -cam "light_AD:camera45gate" -of "png" "E:\My\Desktop\Pal2\art\maya\Songzhou Outskirts\scenes\Outskirts.mb"
type RenderSettings struct {
	ProjectPath string
	SceneName   string
	CameraName  string
	ImageName   string
	PercentResolution int // 50 means 50%
}

func main() {
	rs := RenderSettings{}
	rs.ProjectPath = "E:\\My\\Desktop\\Pal2\\art\\maya\\Songzhou Outskirts"
	rs.SceneName = "Outskirts"
	rs.CameraName = "light_AD:camera45gate"
	rs.ImageName = rs.SceneName + " " + time.Now().Format("2006-01-02 150405")
	rs.PercentResolution = 100

	RenderCommand(rs)
	ConvertCommand(rs.ImageName)
	var enterToExit string
	fmt.Println("press Enter to exit")
	fmt.Scanf(enterToExit)
	//fmt.Println(string(out))
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
func ConvertCommand(imageName string) {
	dir, _ := os.Getwd()
	idir := dir + "/images"
	imageFullName := idir + "/" + imageName + ".png"
	cmd := exec.Command("convert",
		imageFullName,
		"-gamma", "2.2",
		imageFullName);
	CommandRun(cmd)
}

func RenderCommand(rs RenderSettings) {
	dir, _ := os.Getwd()
	render := "C:/Program Files/Autodesk/Maya2016/bin/Render.exe"
	//imageName := rs.SceneName + " " + time.Now().Format("2006-01-02 150405")

	cmd := exec.Command(render,
		"-r", "mr",
		"-rd", dir + "/images",
		"-im", rs.ImageName,
		"-of", "png",
		"-percentRes",strconv.Itoa(rs.PercentResolution),
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
