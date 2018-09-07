package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

const imgPath = "utils/images/image.jpg"

func main() {

	width, height := getImageSize(imgPath)

	log.Printf("width:%d, height:%d", width, height)

	Resize(imgPath, 100, 100)

}

// 得到图片的宽和高
func getImageSize(imagePath string) (width, height int) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Println("open file failed:", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Println("decode image config failed:", err)
	}
	return image.Width, image.Height
}

// 修改图片尺寸
func Resize(imagePath string, width, height uint) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Println("open file failed:", err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Println("decode file failed:", err)
	}
	file.Close()

	newImg := resize.Resize(width, height, img, resize.Lanczos3)

	newFile, err := os.Create("utils/images/new.jpg")
	if err != nil {
		log.Println("create file failed:", err)
	}
	defer newFile.Close()

	jpeg.Encode(newFile, newImg, nil)

}
