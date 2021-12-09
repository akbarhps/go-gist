package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	stateDelete               = "DELETE"
	stateCreateFromDir        = "CREATE_FROM_DIR"
	stateCreateFromFile       = "CREATE_FROM_FILE"
	stateCreateFromFileMerged = "CREATE_FROM_FILE_MERGED"
	stateCreateFromClipboard  = "CREATE_FROM_CLIPBOARD"

	argCreateFromDir        = "-dir"
	argCreateFromFile       = "-f"
	argCreateFromFileMerged = "-fm"
	argCreateFromClipboard  = "-c"
	argSaveToClipboard      = "-s"
	argDelete               = "-d"
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
		case argCreateFromDir:
			currentState = stateCreateFromDir
			continue
		case argCreateFromFile:
			currentState = stateCreateFromFile
			continue
		case argCreateFromFileMerged:
			currentState = stateCreateFromFileMerged
			continue
		case argCreateFromClipboard:
			currentState = stateCreateFromClipboard
			continue
		case argDelete:
			currentState = stateDelete
			continue
		case argSaveToClipboard:
			args.SaveToClipboard = true
			continue
		}

		switch currentState {
		case stateCreateFromClipboard:
			content = filepath.Base(arg)
		}

		args.StateData[currentState] = append(args.StateData[currentState], content)
	}

	if currentState == "" {
		helpString := fmt.Sprintf(`HELP:
'%s' create gist from files in specified directory from executed app directory
example:

	go-gist %s [dirName]

'%s' create gist from files in executed app directory
example:

	go-gist %s [file1] [file2]

'%s' create gist from files in executed app directory, merged into 1 gist
example:

	go-gist %s [file1] [file2]

'%s' create gist from clipboard
example:

	go-gist %s [fileName (optional)]

'%s' delete remote gist
example:

	go-gist %s [gistId]

'%s' save created gist id and url to clipboard
example:
	
	go-gist %s -f [file1]

`, argCreateFromDir, argCreateFromDir,
			argCreateFromFile, argCreateFromFile,
			argCreateFromFileMerged, argCreateFromFileMerged,
			argCreateFromClipboard, argCreateFromClipboard,
			argDelete, argDelete,
			argSaveToClipboard, argSaveToClipboard)
		log.Fatal(helpString)
	}

	return args
}
