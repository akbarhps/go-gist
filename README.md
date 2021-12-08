# Go-Gist

Simple CLI application to create and delete github gist.

## Installation

- Open `http.go` file, and add your token in `Authorization` header

- Run ```bash go build . ```

## Usage

```
~\..\test
➜ go-gist
2021/12/08 22:41:40 HELP:

to create gist, use:
        gist -c file1 file2 file3

to delete gist, user:
        gist -d gistId1 gistId2

to do both, use:
        gist -c file1 file2 -d gistId1 gistId2
```

## Example

```bash
go-gist -c test.go
```

Output:

```bash
~\..\test
➜ go run . -c test.go
2021/12/08 22:40:18 Id: 4d67ace8ac028a515ced11e397681aed | Url: https://gist.github.com/4d67ace8ac028a515ced11e397681aed
```

## 