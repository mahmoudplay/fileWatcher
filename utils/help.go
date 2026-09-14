package utils

import "fmt"

func HelpCli() {
	fmt.Println("fileWatcher - watches files and directories and creates backups on change")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  fileWatcher <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  watch <path>        Watch a file or directory and back up files when they change")
	fmt.Println("  change-time <value> Change the polling interval (e.g. 10s, 5m, 1h)")
	fmt.Println("  help                Show this help message")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  fileWatcher watch ./gg")
	fmt.Println("  fileWatcher change-time 5s")
	fmt.Println("  fileWatcher help")
}
