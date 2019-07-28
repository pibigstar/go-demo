package images

import (
	"testing"
)

const (
	imgPath = "image.jpg"
	newPath = "new.jpg"
)

func TestResizeImage(t *testing.T) {
	width, height := GetImageSize(imgPath)

	t.Logf("width:%d, height:%d", width, height)

	Resize(imgPath, newPath, 100, 100)
}
