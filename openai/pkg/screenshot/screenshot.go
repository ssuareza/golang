package screenshot

import (
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// Take takes a screenshot of the specified display and saves it to the given file path.
func Take(image string) error {
	display := 0
	bounds := screenshot.GetDisplayBounds(display)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return err
	}

	file, _ := os.Create(image)
	defer file.Close()
	png.Encode(file, img)

	return nil
}
