package filestore

import "github.com/gin-gonic/gin"

func UploadHandler(storage Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, _ := c.Request.FormFile("file")
		defer file.Close()

		info, err := storage.Upload(
			c.Request.Context(),
			file,
			header.Size,
			header.Filename,
			UploadOptions{
				ContentType: header.Header.Get("Content-Type"),
			},
		)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, info)
	}
}
