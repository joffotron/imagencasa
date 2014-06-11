package imagencasa

import (
	"net/http"
	"image/png"
	manip "imagencasa/image"
)


func Start() {
	http.HandleFunc("/totez", totezTotes)
	http.ListenAndServe(":8080", nil)
}


func totezTotes(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "image/png")
//	func DetectContentType(data []byte) string

	artFile := manip.LoadImage("data/base.jpg")
	png.Encode(response, artFile)
}


