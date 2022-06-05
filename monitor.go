package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var fileChan chan *Message

type Message struct {
	custom   string
	filename string
}

func initMonitor() {
	fileChan = make(chan *Message, 10000)
	go monitorFile()
}

func monitorFile() {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}
	defer w.Close()
	// 遍历当前文件夹下的目录，将所有的目录添加但监听列表
	filepath.Walk("./upload", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.Add(path)
			if err != nil {
				return err
			}
			log.Println("开始监控:", path)
		}
		return nil
	})
	for {
		select {
		case ev := <-w.Events:
			if ev.Op&fsnotify.Create == fsnotify.Create {
				info, err := os.Stat(ev.Name)
				if err != nil {
					log.Println(err)
				}
				// 判断是否是文件
				if !info.IsDir() {
					log.Println(ev.Name, "created!!!")
					dir := filepath.Dir(ev.Name)
					custom := filepath.Base(dir)
					message := &Message{
						custom:   custom,
						filename: ev.Name,
					}
					fileChan <- message
				}
			}
		case err := <-w.Errors:
			log.Println(err)
			return
		}
	}
}
