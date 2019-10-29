package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				fmt.Println(time.Now(), ": event:", event)
				log.Println("event:", event)

				if event.Op == fsnotify.Create {
					fmt.Println(event.Name, "size is", getFileSize(event.Name))
					log.Println(event.Name, "size is", getFileSize(event.Name))
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println(time.Now(), ": modified file:", event.Name)
					log.Println("modified file:", event.Name)
					fmt.Println(event.Name, "size is", getFileSize(event.Name))
					log.Println(event.Name, "size is", getFileSize(event.Name))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				fmt.Println("error:", err)
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/Users/yang/file")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	<-done
}

func init() {
	file := "./" + "logFile" + ".txt"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile) // 将文件设置为log输出的文件
}

func getFileSize(filename string) int64 {
	var result int64

	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})

	return result
}
