# Go-Gist

Simple CLI application to create and delete github gist.

## Installation

- Create your [Github Gist Token](https://github.com/settings/tokens)

- Open `http.go` file, and add your token in `Authorization` header

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
➜ go run .
2021/12/09 11:07:31 HELP:
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

        go-gist -cc [fileName (optional)]

'-d' delete remote gist
example:

        go-gist -d [gistId]

'-c' save created gist id and url to clipboard
example:

        go-gist -c -cf [file1]

```

## Example

```bash
go-gist -c -cf test.go
```

Output:

```bash
➜ go run . -c -cf test.go
2021/12/08 22:40:18 Id: 4d67ace8ac028a515ced11e397681aed | Url: https://gist.github.com/4d67ace8ac028a515ced11e397681aed
```

## Using Vim Command

`.vimrc`:
```
vnoremap <space>0 "+y :!go-gist -c -cc %<CR>
```

- make sure you set `go-gist` folder in environment path  
- output automatically write to clipboard
