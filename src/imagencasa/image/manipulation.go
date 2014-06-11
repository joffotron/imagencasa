package image

import (
		"image"
		 _ "image/jpeg"
	    "os"
)

func LoadImage(filename string) image.Image {
	file, _ := os.Open(filename)
	defer file.Close()

	image, _, _ := image.Decode(file)
	return image
}


