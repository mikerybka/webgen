package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	buildNeeded := true
	go watch("ui/src", func() {
		buildNeeded = true
	})
	cmd := run()
	for {
		if buildNeeded {
			buildNeeded = false
			kill(cmd)
			cmd = run()
		}
		waitPeriod := 3 * time.Second
		time.Sleep(waitPeriod)
	}
}

func kill(cmd *exec.Cmd) {
	err := cmd.Process.Kill()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}

func run() *exec.Cmd {
	cmd := exec.Command("bunx", "tailwindcss", "-i", "./src/input.css", "-o", "./public/index.css")
	cmd.Dir = "ui"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("bun", "run", "dev.tsx")
	cmd.Dir = "ui"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd
}

func watch(dir string, callback func()) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					callback()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(dir)
	if err != nil {
		panic(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
