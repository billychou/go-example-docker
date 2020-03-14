package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/checkpreload.htm", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	r.GET("/test", func(c *gin.Context){
		max := 1000000
		sum := 0
		for i:=0; i <= max; i++ {
			sum += i
		}
		c.JSON(http.StatusOK, gin.H{
			"sum": sum,
		})
	})

	r.Run()
}