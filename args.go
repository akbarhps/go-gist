package main

import (
	"log"
	"os"
	"strings"
)

var (
	stateDelete               = "DELETE"
	stateCreateFromDir        = "CREATE_FROM_DIR"
	stateCreateFromFile       = "CREATE_FROM_FILE"
	stateCreateFromFileMerged = "CREATE_FROM_FILE_MERGED"
	stateCreateFromClipboard  = "CREATE_FROM_CLIPBOARD"
)

type Args struct {
	ExecDir         string
	SaveToClipboard bool
	StateData       map[string][]string
}

func parseArgs() *Args {
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	args := &Args{
		ExecDir:   execDir,
		StateData: make(map[string][]string),
	}

	currentState := ""
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		content := arg
		switch arg {
		case "-cf":
			currentState = stateCreateFromFile
			continue
		case "-cfm":
			currentState = stateCreateFromFileMerged
			continue
		case "-cd":
			currentState = stateCreateFromDir
			continue
		case "-cc":
			currentState = stateCreateFromClipboard
		case "-d":
			currentState = stateDelete
			continue
		case "-c":
			args.SaveToClipboard = true
			continue
		}

		switch {
		case strings.HasPrefix(arg, "-f"):
			content = strings.Replace(arg, "-f=", "", 1)
		}

		args.StateData[currentState] = append(args.StateData[currentState], content)
	}

	if currentState == "" {
		log.Fatal(`HELP:
'-cd' create gist from files in specified directory from executed app directory
example:

	go-gist -cd [dirName]

'-cf' create gist from files in executed app directory
example:

	go-gist -cf [file1] [file2]

'-cfm' create gist from files in executed app directory, merged into 1 gist
example:

	go-gist -cfm [file1] [file2]

'-cc' create gist from clipboard
example:

	go-gist -f=[fileName (optional)] -cc

'-d' delete remote gist
example:

	go-gist -d [gistId]

'-c' save created gist id and url to clipboard
example:
	
	go-gist -c -cf [file1]

`)
	}

	return args
}
