# Wordle

A terminal Wordle game written in Go. Six guesses. Five letters. Your move.

The colored WORDLE logo stays above the centered guess rows throughout the game.

## Get the binary

Install Git and Go 1.27.1 or newer (the version required by `go.mod`), then clone and build:

```sh
git clone https://github.com/nshern/wordle.git
cd wordle
go build -o wordle .
./wordle random
```

This creates the `wordle` binary in the repository directory. On Windows, build with `go build -o wordle.exe .` and run `.\wordle.exe random`.

Keep the cloned repository at its original location: the game reads `words.txt` and `valid-words.txt` from the source directory recorded during the build. The binary is not standalone; copying those files next to a relocated binary does not change where it looks for them.

## Run from anywhere

From the cloned repository, install the command:

```sh
go install .
```

By default, Go installs the binary into `$(go env GOPATH)/bin`. Add that directory to your PATH once using the instructions for your shell.

For **zsh**:

```sh
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

For **fish**:

```fish
fish_add_path (go env GOPATH)/bin
```

If you have set a custom `GOBIN`, add that directory to your PATH instead.

You can then play from any directory:

```sh
wordle random
wordle daily
```

Keep the cloned repository in its original location so the installed command can find the word lists. Run `go install .` from the repository again after updating the code.

## Play

Choose a mode:

```sh
./wordle random  # Pick a random word from the local word list
./wordle daily   # Fetch a puzzle answer from the New York Times
```

Running `./wordle` without a mode, or with an unknown mode, shows the help screen.

You can also run directly from the repository with `go run . random` or `go run . daily`.

Enter a five-letter word and press Enter. Guesses are case-insensitive and must appear in the accepted word list. Invalid guesses do not use an attempt.

- **Green text:** the letter is in the correct position.
- **Yellow text:** the letter appears in the answer in another position.
- **Plain text:** the letter does not appear in the answer.

Solve the word within six guesses. If you run out of guesses, the game reveals the answer. Use a terminal with ANSI color support to see the colored logo and letter feedback.

## Current limitations

- Daily mode requires internet access and currently fetches the puzzle for **October 5, 2026**; the date is hardcoded rather than following today's date.
- Repeated-letter feedback checks whether a letter appears anywhere in the answer; it does not yet account for how many times that letter occurs.
