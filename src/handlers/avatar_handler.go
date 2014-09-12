package handlers

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/martini-contrib/sessions"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func UploadAvatar(session sessions.Session, w http.ResponseWriter, r *http.Request) string {
	if session.Get("auth_id") != "" {
		id := session.Get("auth_id").(string)
		file, header, err := r.FormFile("file") // the FormFile function takes in the POST input id file
		defer file.Close()

		if err != nil {
			fmt.Fprintln(w, err)
			return "Error"
		}

		if os.MkdirAll("avatar/" + id + "/", 0777) != nil {
			panic("Unable to create directory for tagfile!" + err.Error())
		}
		out, err := os.Create("avatar/" + id + "/" + header.Filename)
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return "Error"
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintf(w, "File uploaded successfully : ")
		fmt.Fprintf(w, header.Filename)
		width, height := getImageDimension("avatar/" + id + "/" + header.Filename)
		if width != height {
			return "Image not square"
		}
		return "OK"
	}
	return "no auth"
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return -1, -2
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
		return -1, -2
	}
	return img.Width, img.Height
}
