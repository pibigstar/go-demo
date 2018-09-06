package main

import (
	"os"
	"log"
	"image"
	)

func main() {

	width, height := getImageSize("image.jpg")

	log.Printf("width:%d, height:%d",width,height)

}
// 得到图片的宽和高
func getImageSize(imagePath string)(width,height int)  {
	file, err := os.Open(imagePath)
	if err!=nil {
		log.Println("open file failed:",err)
	}

	image, _, err := image.DecodeConfig(file)
	if err!=nil {
		log.Println("decode image config failed:",err)
	}
	return image.Width,image.Height
}
