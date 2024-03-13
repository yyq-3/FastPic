package pic

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

const (
	WIDTH1600 uint = 1600
	WIDTH800  uint = 800
	HEIGHT600 uint = 600
	HEIGHT900 uint = 900
)

func CutTo1600x900(c *gin.Context) {
	buf := cutPic(c, WIDTH1600, HEIGHT900)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

func CutTo800x600(c *gin.Context) {
	buf := cutPic(c, WIDTH800, HEIGHT600)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

func CustomerCut(c *gin.Context) {
	width := c.PostForm("width")
	height := c.PostForm("height")
	iWidth, err := strconv.Atoi(width)
	if err != nil {
		c.JSON(200, gin.H{"message": "输入宽度不合法", "success": false})
		return
	}
	iHeight, err := strconv.Atoi(height)
	if err != nil {
		c.JSON(200, gin.H{"message": "输入高度不合法", "success": false})
		return
	}
	buf := cutPic(c, uint(iWidth), uint(iHeight))
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(buf.Bytes())
}

func cutPic(c *gin.Context, width, height uint) *bytes.Buffer {
	file, _, err := c.Request.FormFile("file")
	if nil == file {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("file is nil"),
		})
		return nil
	}
	defer file.Close()

	img, err := jpeg.Decode(file)

	if err != nil {
		c.JSON(200, gin.H{"message": "解码失败，异常原因: " + err.Error()})
		return nil
	}
	newImg := resize.Resize(width, height, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	png.Encode(buf, newImg)
	return buf
}
