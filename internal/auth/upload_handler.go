package auth

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Ошибка получения файла:", err)
		c.JSON(400, gin.H{"error": "Не удалось прочитать файл"})
		return
	}
	if file.Size > 10*1024*1024 {
		log.Println("Файл больше разрешенного размера")
		c.JSON(400, gin.H{"error": "Файл больше разрешенного размера"})
		return
	}
	extention := filepath.Ext(file.Filename)
	newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), extention)
	err = c.SaveUploadedFile(file, "./uploads/"+newName)
	if err != nil {
		log.Println("Ошибка загрузки файла на сервер:", err)
		c.JSON(500, gin.H{"error": "Ошибка загрузки файла на сервер"})
		return
	} else {
		c.JSON(200, gin.H{"file_url": "/uploads/" + newName})
	}
}
