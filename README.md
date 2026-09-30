# Wordle

A terminal Wordle game written in Go.

## Get the binary

Install Git and Go 1.27.1 or newer (the version required by `go.mod`), then clone and build:

```sh
git clone https://github.com/nshern/wordle.git
cd wordle
go build -o wordle .
./wordle
```

This creates the `wordle` binary in the repository directory. On Windows, build with `go build -o wordle.exe .` and run `.\wordle.exe`.

Keep the cloned repository at its original location: the game reads `words.txt` and `valid-words.txt` from the source directory recorded during the build. The binary is not standalone; copying those files next to a relocated binary does not change where it looks for them.
