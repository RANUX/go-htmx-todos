package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/anthdm/slick"
)

func StaticFile(c *slick.Context, status int, file []byte) error {
	// Return image or other binary data
	c.Response.Header().Set("Content-Type", http.DetectContentType(file))
	c.Response.WriteHeader(status)
	_, err := c.Response.Write(file)
	return err
}

// Handle static files from public directory
func HandlePublicStaticFile(c *slick.Context) error {
	// get current project directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	path := c.Request.URL.Path[1:]

	// load file from public directory
	f, err := os.Open(dir + "/" + path)
	if err != nil {
		return err
	}
	defer f.Close()
	// read file
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	return StaticFile(c, http.StatusOK, b)

}
