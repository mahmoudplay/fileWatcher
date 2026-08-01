package main

import (
	"fmt"
	"os"
	"fileWatcher/utils"
	"time"
)


func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Wrong command, type fileWatcher help")
		return
	}

	switch(os.Args[1]){
		case "watch":
			if err := utils.CreateConfig(); err != nil {
				fmt.Println(err)
				return
			}

			interval, err := utils.GetConfigTime()

			if err != nil {
				fmt.Println(err)
				return
			}

			for true {
				if err := utils.RegisterFiles(os.Args[2]); err != nil {
					fmt.Print(err)
					break;
				}
				
				time.Sleep(time.Duration(interval) * time.Second)
			}

		case "change-time":
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