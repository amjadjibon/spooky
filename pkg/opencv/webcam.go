package opencv

import (
	"github.com/urfave/cli/v2"
	"gocv.io/x/gocv"
)

func OpenWebcam(c *cli.Context) error {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
