# Go-Gist

Simple CLI application to create and delete github gist.

## Installation

- Create your [Github Gist Token](https://github.com/settings/tokens)

- Open `config.json` file, and add your token to `token` field.

- Build  
```bash
go mod tidy
```

then

```bash
go build .
```

## Usage

```
2021/12/09 16:12:19 HELP:
'-dir' create gist from files in specified directory from executed app directory
example:

        go-gist -dir [dirName]

'-f' create gist from files in executed app directory
example:

        go-gist -f [file1] [file2]

'-fm' create gist from files in executed app directory, merged into 1 gist
example:

        go-gist -fm [file1] [file2]

'-c' create gist from clipboard
example:

        go-gist -c [fileName (optional)]

'-d' delete remote gist
example:

        go-gist -d [gistId]

'-s' save created gist id and url to clipboard
example:

        go-gist -s -f [file1]

```

## Example

```bash
go-gist -s -f test.go
```

Output:

```bash
go-gist -s -f test.go

2021/12/08 22:40:18 Id: 4d67ace8ac028a515ced11e397681aed | Url: https://gist.github.com/4d67ace8ac028a515ced11e397681aed
```

# Automate Go-Gist

- make sure you set `go-gist` folder in environment path

## Using VSCode

- create `.vscode` directory in your project
- create `tasks.json` file in `.vscode` directory
- add current config

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "gist last clipboard",
      "type":"shell",
      "command": "gist -s -c ${fileBasename}"
    }
  ]
}
```

- copy code
- press `CTRL + SHIFT + P`, type Tasks: Run Task
- Choose `gist last clipboard` then press enter

## Using Vim

- in your `.vimrc` file, add :
 
```
vnoremap <space>0 "+y :!go-gist -s -c %<CR>
```
