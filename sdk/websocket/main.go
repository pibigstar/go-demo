package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"time"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout: 3 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var basePath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath = filepath.Dir(currentFile)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, filepath.Join(basePath, "ws.html"))
	})

	r.GET("/ws", func(c *gin.Context) {
		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade", err)
			return
		}
		defer func() {
			if err := ws.Close(); err != nil {
				log.Println("failed close ws", err)
			}
		}()
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("failed read message", err)
				return
			}
			log.Printf("recv message: %s", message)
			err = ws.WriteMessage(mt, []byte(fmt.Sprintf("你说的是: %s", message)))
			if err != nil {
				log.Println("failed write message", err)
				return
			}
		}
	})

	if err := r.Run(":8088"); err != nil {
		panic(err)
	}
}
