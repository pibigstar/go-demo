package fsnotify_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

func TestFileChangeListen(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()
	// 监听此目录
	err = watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Printf("%s %s\n", event.Name, event.Op)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		case <-ctx.Done():
			return
		}
	}
}
