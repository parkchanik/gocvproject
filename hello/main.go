package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		fmt.Println("test")
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
		////테스트트트
	}
}
