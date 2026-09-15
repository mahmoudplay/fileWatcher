package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fileWatcher/utils"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Wrong command, type fileWatcher help")
		return
	}

	switch os.Args[1] {
	case "watch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: fileWatcher watch <path>")
			return
		}

		if err := utils.CreateConfig(0); err != nil {
			fmt.Println(err)
			return
		}

		interval, err := utils.GetConfigTime()
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		count, err := utils.CountFiles(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Watching %d file(s): %s\n", count, os.Args[2])

		for {
			if err := utils.RegisterFiles(os.Args[2]); err != nil {
				fmt.Print(err)
				break
			}

			select {
			case <-ctx.Done():
				fmt.Println("\nStopped watching.")
				return
			case <-time.After(time.Duration(interval) * time.Second):
			}
		}

	case "change-time":
		if len(os.Args) < 3 {
			fmt.Println("Usage: fileWatcher change-time <number>")
			return
		}

		err := utils.SetConfigTime(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

	case "help":
		utils.HelpCli()

	default:
		fmt.Println("Wrong command, type fileWatcher help")
	}
}
