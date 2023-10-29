package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func Stream(c *gin.Context) {
	fileName := "Sandra N - Lion Heart (Official Video) ( 360 X 640 ).mp4"
	filePath := "public/" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening the file")
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting file info")
		return
	}

	http.ServeContent(c.Writer, c.Request, fileName, fileInfo.ModTime(), file)
}

func Stream2(c *gin.Context) {

	fileName := "Sandra N - Lion Heart (Official Video) ( 360 X 640 ).mp4"
	filePath := "public/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening the file")
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}

	fmt.Println(stat)

	bufferSize := 10000 * 1024 // 100KB
	buffer := make([]byte, bufferSize)

	c.Header("Content-Type", "video/mp4")

	for {
		n, err := io.ReadFull(file, buffer)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			c.String(http.StatusInternalServerError, "Error reading the file")
			return
		}
		if n == 0 {
			break
		}
		c.Writer.Write(buffer[:n])
		c.Writer.Flush()
	}

}
