# fileWatcher

A cross-platform Go CLI tool that watches files and directories and automatically creates compressed backups whenever files change.

## Features

- Watch a single file or an entire directory tree recursively
- Detect changes using SHA-256 content hashing
- Automatically back up files the first time they are seen and every time they change
- Compress backups with Zstandard (`.zst`) for smaller storage
- Configurable polling interval, persisted in `config.ini`
- Cross-platform: builds and runs on Windows and Linux
- Zero-config auto-setup: creates a default `config.ini` on first run

---

# User Guide

This guide is for people who just want to use the tool to keep automatic backups of their files. No programming knowledge needed.

## 1. Get the program

Download the zip archive that matches your computer:

- **Windows** on an Intel/AMD CPU → `fileWatcher-windows-amd64.zip`
- **Windows** on an ARM CPU → `fileWatcher-windows-arm64.zip`
- **Linux** on an Intel/AMD CPU → `fileWatcher-linux-amd64.zip`
- **Linux** on an ARM CPU → `fileWatcher-linux-arm64.zip`

Archives are attached to every GitHub release. Each zip contains one executable named `fileWatcher`. Extract it to a folder of your choice and run it from that folder.

On Linux you may need to make the file executable first:

```bash
chmod +x fileWatcher
```

## 2. Start watching a folder

Run the program from a command line (terminal). On Windows open `cmd` or `PowerShell`, on Linux open your terminal.

```bash
# Windows
fileWatcher watch C:\Users\you\Documents\my-project
# or using shortcut:
fileWatcher -w C:\Users\you\Documents\my-project

# Linux
./fileWatcher watch /home/you/Documents/my-project
# or using shortcut:
./fileWatcher -w /home/you/Documents/my-project
```

The program will:

1. Create a `config.ini` file the first time it runs (if it doesn't already exist).
2. Scan the folder *and all subfolders*.
3. Back up every file it sees for the first time.
4. Keep watching, and back up a file again whenever its content changes.

The watcher runs until you stop it with `Ctrl + C`. On restart it
remembers what it already saw (from a state file), so it does **not**
re-back-up files that haven't changed since the last run.

## 3. Where are the backups?

Backups are saved inside a folder named `.backup/`, created **wherever the program is run from** (its working directory), not next to the watched files.

Example backup filename:

```
my-document_20260914_103000.docx.zst
```

- `my-document.docx` – the original file name
- `20260914_103000` – the date and time the backup was created
- `.zst` – back up is compressed with Zstandard

If two backups would get the same filename (for example within the same
second), the program automatically adds a suffix: `..._103000_1.docx.zst`.

Besides the backups, `.backup/` also holds a `state.json` file that records
which files and contents have already been backed up (the timestamp in the
filename uses the backup time — not the file's modification time).

## 4. Restore a file

A `.zst` file must be decompressed before use. Use any Zstandard-compatible tool:

```bash
# zstd CLI
zstd -d my-document_20260914_103000.docx.zst

# 7-Zip also opens .zst archives
```

## 5. Change how often it checks for changes

The default check is every 10 seconds. To change it, use the `change-time` (or `-ct`) command (the program does **not** need to be running to change the config):

```bash
# Windows
fileWatcher change-time 30s
# or using shortcut:
fileWatcher -ct 30s

# Linux
./fileWatcher change-time 5m
# or using shortcut:
./fileWatcher -ct 5m
```

| Suffix | Meaning            | Example            |
| ------ | ------------------ | ------------------ |
| `s`    | seconds            | `change-time 30s`  |
| `m`    | minutes            | `change-time 5m`   |
| `h`    | hours              | `change-time 1h`   |

## 6. Check version

```bash
fileWatcher version
# or using shortcut:
fileWatcher -v
```

## 7. Get help

```bash
fileWatcher help
# or using shortcut:
fileWatcher -h
```

---

# Developer Guide

This guide is for developers who want to build, modify, or contribute to the project.

## Requirements

- Go 1.26 or later
- Git

## Set up

```bash
git clone https://github.com/mahmoudplay/fileWatcher.git
cd fileWatcher
```

There are no third-party setup steps beyond downloading dependencies:

```bash
go mod download
```

## Build

```bash
go build -o fileWatcher .
```

Build for another platform (cross-compile):

```bash
# Windows amd64
env GOOS=windows GOARCH=amd64 go build -o fileWatcher .

# Linux arm64
env GOOS=linux GOARCH=arm64 go build -o fileWatcher .
```

**Note for Windows shells:** use `$env:GOOS="windows"` / `$env:GOARCH="amd64"` (PowerShell) or `set GOOS=windows` (cmd) instead of `env`.

The project is pure Go (no CGo), so building statically with `CGO_ENABLED=0` works everywhere.

## Run locally

```bash
go run . watch ./your-files     # watch the sample directory
go run . change-time 10s        # change interval
go run . help                   # show help
```

## Check your code

```bash
gofmt -l .        # formatting
go vet ./...      # static analysis
go test ./...    # tests (none exist yet)
```

## Project structure

```
fileWatcher/
├── main.go                 # CLI entry point; subcommand dispatch
├── config.ini              # Runtime configuration (auto-generated, gitignored)
├── .backup/                # Backup output directory (gitignored)
├── gg/                     # Test/watch directory (gitignored)
├── utils/
│   ├── backup.go           # Reads a file and writes a compressed backup
│   ├── compress.go         # Zstandard compression wrapper
│   ├── file_name.go        # Generates timestamped backup filenames
│   ├── fileHash.go         # SHA-256 content hashing
│   ├── help.go             # Help text
│   ├── iniFile.go          # config.ini read/write + unit parsing
│   └── registery.go        # Recursive watcher, change detection
├── go.mod
└── go.sum
```

## How it works

1. `watch` loads the interval from `config.ini` (creating it with a default of 10s if missing).
2. Every interval, it walks the path from `main.go` (`RegisterFiles()` in `utils/registery.go`):
   - **File:** computes the SHA-256 hash. If the hash is new or different from the last seen hash, it creates a backup (`makeBackup()` in `utils/backup.go`) and updates the stored hash.
   - **Directory:** recurses into each entry.
3. Seen hashes are persisted in `.backup/state.json` and re-loaded on startup, so files are only backed up again once their content changes.
4. Backups are compressed with Zstandard (`CompressFile()` in `utils/compress.go`) and written to `.backup/` with a name like `<name>_<YYYYMMDD_HHMMSS>.<ext>.zst`. Compression streams the file through the encoder, so even very large files use little memory.

## Extending the tool

- **New command:** add a `case` in the `switch` in `main.go`, then document it in `utils/help.go` and the README.
- **Config keys:** read/write them in `utils/iniFile.go`.
- **Backup format:** the filename pattern lives in `utils/file_name.go`; the compression is isolated in `utils/compress.go`.

### The interval unit parser

`change-time` values must end in `s`, `m`, or `h`, parsed by `timeTranslater()` in `utils/iniFile.go`. The stored `config.ini` value is always in **seconds**:

```
[backup]
interval = 10
```

## Contributing

- Keep code formatted (`gofmt`) and run `go vet ./...` before committing.
- Follow the existing package style: all functionality lives in `utils/`, `main.go` only dispatches commands.
- Backups and config files (`config.ini`, `.backup/`) are gitignored — don't commit them.

---

## License

This project is licensed under the [MIT License](LICENSE).

Copyright (c) 2026 [mahmoudplay](https://github.com/mahmoudplay)