# Soccer League Ranking
[![Go Report Card](https://goreportcard.com/badge/github.com/dcaponi/soccer-league)](https://goreportcard.com/report/github.com/dcaponi/soccer-league)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-100%25-brightgreen.svg?longCache=true&style=flat)</a>

## The problem

A command-line application that calculates the ranking table for a soccer league.

### Input/output

The input is a text file (.txt) provided by a file passed by name on the command line. 

`soccer-league-2020 input.txt`

The input contains results of games, one per line. The output should be ordered from most to least points.


#### Supplied Input
```
Robots 3, Spammers 3
Thieves 1, FC Fraudsters 0
Robots 1, FC Fraudsters 1
Thieves 3, Spammers 1
Robots 4, Grandparents 0
```

#### Supplied Expected Output
```
1. Thieves, 6 pts
2. Robots, 5 pts
3. FC Fraudsters, 1 pt
3. Spammers, 1 pt
4. Grandparents, 0 pts
```

Assumes the input is well-formed. There is no special handling for malformed input files.

Assumes one file input at a time and the file will always be .txt type

### The rules

In this league, a draw (tie) is worth 1 point and a win is worth 3 points. A
loss is worth 0 points. If two or more teams have the same number of points,
they should have the same rank and be printed in alphabetical order (as in the
tie for 3rd place in the sample data).

### Build / Test / Run

To run the tests & report coverage (requires golang) `make test`

To build/install as a cli utility with golang (only tested on mac - windows may vary) `make install`

To add the pre-built binary to path (macos) `make install-macos`

#### ⚠️ A Note About MacOS Binaries & Security ⚠️
You'll likely get hit with a security warning when you try running the pre-built binary for the first time.

To fix, go to System Preferences > Security & Privacy

you'll be presented with the warning about the binary not being from an identified developer. Allow this app to run.

Try running the command again and click Open from the popup and you should be good to go.

Run without adding to bin or path (requires golang) `make run-supplied` or `make run-n-ties` or `make run-all-ties`

Run smoke test on sample-inputs and compare to expected outputs (requires golang) `make check`

To run one of the builds without golang - involke the binary whos name matches your system architecture as you would when running compiled code from the command line. e.g. `./build/soccer-league-2020-darwin-amd64` or use one of the makefile commands (not tested on windows) 

### Todo

Add script for running convenience checks on windows machines

Dockerize

Add Github action for CI/CD

Install script beyond the `make install` command that is only tested on macos
