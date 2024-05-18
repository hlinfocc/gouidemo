package web

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int
	Msg  string
	Data string
}

func checkPortStatus(port int) bool {
	// 监听 端口
	log.Println("check listener Port", port)
	listenerPort := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", listenerPort)
	if err != nil {
		// 如果监听失败，则说明端口已被占用
		return false
	}
	// 关闭监听器
	defer listener.Close()

	// 如果监听成功，则说明端口未被占用
	return true
}

func writePort(port int) {
	filePath := "/tmp/hlinfo-cyssh-server.port"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(port))
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
}

/**
* 启动web服务
 */
func StartWebServer() {
	port := 1088
	for !checkPortStatus(port) {
		port = port + 1
	}
	writePort(port)
	log.Println("监听的端口:", port)
	httpPort := fmt.Sprintf(":%d", port)
	// 创建一个默认的Gin引擎
	router := gin.Default()

	// 使用嵌入的静态资源
	// router.StaticFS("/static", assets.FileSystem)

	// 使用cookie存储session信息
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("PHPSESSIONID", store))

	// 定义一个GET请求的路由，当访问根路径"/"时，返回"Hello, Gin!"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
		// 在这里提供Vue.js的入口HTML文件
		// indexHTML, err := assets.GetIndexHtml() //embededFiles.ReadFile("assets/index.html")
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, "Failed to read index.html: %v", err)
		// 	return
		// }
		// c.Data(http.StatusOK, "text/html", indexHTML)
	})

	// 启动Gin服务器，监听端口
	router.Run(httpPort)

}
