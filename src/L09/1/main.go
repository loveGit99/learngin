package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 使用现有的基础请求对象解析查询字符串参数。
	// 示例 URI： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式
		name, _ := c.GetQuery("lastname")

		c.String(http.StatusOK, "Hello %s %s %s", firstname, lastname, name)
	})
	router.Run(":8080")
}
