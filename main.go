package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

func monitorClipboard(ch chan string) {
	cache := make(map[string]bool)
	for {
		t, err := clipboard.ReadAll()
		if err != nil {
			log.Print("clipboard.ReadAll error", err)
		}

		if t == "" {
			log.Printf("listenyt: It's possible that you need to install either xclip or xsel " +
				"for clipboard management to work properly")
		}

		if strings.Contains(t, "youtube.com") {
			_, ok := cache[t]
			if !ok {
				log.Printf("monitorClipboard: Found youtube link %v", t)
				ch <- t
				cache[t] = true
			}
		}

		time.Sleep(time.Second * 1)
	}
}

const youtubeDL = "youtube-dl"

var youtubeDLArgs = []string{"-f", "140", "-o", "~/youtube-dl/%(title)s.%(ext)s"}

func linkDownloader(ch chan string) {
	for {
		lnk := <-ch
		log.Printf("linkDownloader: Downloading link: %v", lnk)
		cmd := exec.Command(youtubeDL, append(youtubeDLArgs, lnk)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Printf("linkDownloader: cmd.Run error: %v", err)
		}
	}
}

func main() {
	_, err := exec.LookPath("youtube-dl")
	if err != nil {
		log.Panic("youtube-dl is not installed!")
	}

	ch := make(chan string)
	go monitorClipboard(ch)
	go linkDownloader(ch)

	log.Println("listenyt started. Copy some links!!")
	waitForShutdown()
}

func waitForShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	s := <-ch
	log.Printf("app: Received OS signal %v", s)
}
