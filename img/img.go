package img

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/cantevenread/cantevengag/internal"
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

type FindResult struct {
	Coord internal.Coordinate
	Err   error
	Completed bool
}

func CaptureScreenToMat() (mat gocv.Mat, err error) {
	img, err := robotgo.CaptureImg()
	if err != nil {
		return gocv.Mat{}, err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return gocv.Mat{}, err
	}

	mat, err = gocv.IMDecode(buf.Bytes(), gocv.IMReadColor)
	if err != nil {
		return gocv.Mat{}, err
	}

	return mat, nil
}

// FindTemplateOnScreen searches for the template and signals completion via channel if not nil
func FindTemplateOnScreen(templatePath string, threshold float32, completed chan bool) (internal.Coordinate, error) {
	tmpl := gocv.IMRead(templatePath, gocv.IMReadColor)
	defer tmpl.Close()
	if tmpl.Empty() {
		if completed != nil {
			completed <- false
		}
		return internal.Coordinate{}, fmt.Errorf("template image failed to load")
	}

	screenMat, err := CaptureScreenToMat()
	if err != nil {
		if completed != nil {
			completed <- false
		}
		return internal.Coordinate{}, fmt.Errorf("screenshot failed: %v", err)
	}
	defer screenMat.Close()

	result := gocv.NewMat()
	defer result.Close()

	gocv.MatchTemplate(screenMat, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	if maxVal >= threshold {
		screenWidth, screenHeight := robotgo.GetScreenSize()
		scaleWidth, scaleHeight := robotgo.GetScaleSize()

		scaleX := float64(scaleWidth) / float64(screenWidth)
		scaleY := float64(scaleHeight) / float64(screenHeight)

		// Calculate center coordinate of the matched template, adjusted for scale
		adjustedX := int(float64(maxLoc.X)/scaleX + float64(tmpl.Cols())/(2*scaleX))
		adjustedY := int(float64(maxLoc.Y)/scaleY + float64(tmpl.Rows())/(2*scaleY))

		if completed != nil {
			completed <- true
		}
		return internal.Coordinate{X: adjustedX, Y: adjustedY}, nil
	}

	if completed != nil {
		completed <- false
	}
	return internal.Coordinate{}, fmt.Errorf("no match found (maxVal=%.2f)", maxVal)
}

// FindTemplateOnScreenAsync searches for the template asynchronously,
// returning a read-only channel that will deliver the FindResult.
func FindTemplateOnScreenAsync(templatePath string, threshold float32) <-chan FindResult {
	resultChan := make(chan FindResult, 1)

	go func() {
		defer close(resultChan)

		tmpl := gocv.IMRead(templatePath, gocv.IMReadColor)
		defer tmpl.Close()
		if tmpl.Empty() {
			resultChan <- FindResult{Err: fmt.Errorf("template image failed to load")}
			return
		}

		screenMat, err := CaptureScreenToMat()
		if err != nil {
			resultChan <- FindResult{Err: fmt.Errorf("screenshot failed: %v", err)}
			return
		}
		defer screenMat.Close()

		result := gocv.NewMat()
		defer result.Close()

		gocv.MatchTemplate(screenMat, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat())

		_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
		if maxVal >= threshold {
			screenWidth, screenHeight := robotgo.GetScreenSize()
			scaleWidth, scaleHeight := robotgo.GetScaleSize()

			scaleX := float64(scaleWidth) / float64(screenWidth)
			scaleY := float64(scaleHeight) / float64(screenHeight)

			adjustedX := int(float64(maxLoc.X)/scaleX + float64(tmpl.Cols())/(2*scaleX))
			adjustedY := int(float64(maxLoc.Y)/scaleY + float64(tmpl.Rows())/(2*scaleY))

			resultChan <- FindResult{Coord: internal.Coordinate{X: adjustedX, Y: adjustedY}, Err: nil, Completed: true}
			return
		}

		resultChan <- FindResult{Err: fmt.Errorf("no match found (maxVal=%.2f)", maxVal), Completed: false}
	}()

	return resultChan
}
