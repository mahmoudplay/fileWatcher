package utils

import "fmt"

func HelpCli() {
	fmt.Println("fileWatcher - watches files and directories and creates backups on change")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  fileWatcher <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  watch, -w <path>        Watch a file or directory and back up files when they change")
	fmt.Println("  change-time, -ct <val>  Change the polling interval (e.g. 10s, 5m, 1h)")
	fmt.Println("  version, -v             Show CLI version")
	fmt.Println("  help, -h                Show this help message")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  fileWatcher watch ./gg")
	fmt.Println("  fileWatcher -w ./gg")
	fmt.Println("  fileWatcher change-time 5s")
	fmt.Println("  fileWatcher -ct 5s")
	fmt.Println("  fileWatcher version")
	fmt.Println("  fileWatcher -v")
	fmt.Println("  fileWatcher help")
	fmt.Println("  fileWatcher -h")
}
