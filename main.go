package main
import "github.com/gin-gonic/gin"
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8080  #因8080端口被占用，修改默认访问端口8080，改为8888
}
