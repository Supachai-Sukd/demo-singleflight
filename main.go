package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

var (
	group singleflight.Group
	data  string
	lock  sync.Mutex
)

func main() {
	r := gin.Default()
	// ab -n 100 -c 10 http://localhost:8080/singleflight

	// Route ที่ใช้ singleflight
	r.GET("/singleflight", handleSingleflight)

	// Route ที่ไม่ใช้ singleflight
	r.GET("/normal", handleNormal)

	_ = r.Run(":8080")
}

func handleSingleflight(c *gin.Context) {
	result, err, _ := group.Do("data", func() (interface{}, error) {
		// สำหรับตัวอย่างนี้เราจะสร้างข้อมูลโดยใช้เวลาสุ่ม 1 วินาที
		time.Sleep(1 * time.Second)

		// ในตัวอย่างนี้เราจะสร้างข้อมูลเป็น timestamp ปัจจุบัน
		return time.Now().String(), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result.(string),
	})
}

func handleNormal(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()

	// ถ้ามีการเรียก handleNormal ภายในเวลา 1 วินาที จะเรียกใช้เวลาในการสร้างข้อมูลเสมอ
	time.Sleep(1 * time.Second)

	// ในตัวอย่างนี้เราจะสร้างข้อมูลเป็น timestamp ปัจจุบัน
	data = time.Now().String()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
