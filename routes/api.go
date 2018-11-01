package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gupiao/bootstrap"
	"net/http"
	"time"
)

func Run() {
	var router gin.IRouter = bootstrap.GetBooterInstance().Router

	router.GET("/", func(c *gin.Context) {
		format := "2006-01-02 15:04:05"
		//go func() {
		//	time.Sleep(3*time.Second)
		//	log.Println(time.Now().Format(format))
		//}()
		c.String(http.StatusOK, "api works "+time.Now().Format(format))
	})


	// 加载admin路由
	AdminRun()
}
