ship:
	go build ./... && sh ship.sh github.com/dcaponi/soccer-league-2020

test:
	go get github.com/jpoles1/gopherbadger
	gopherbadger -md="readme.md" -png=false

install-macos:
	cp ./build/soccer-league-2020-darwin-amd64 /usr/local/bin/soccer-league-2020

run-supplied:
	go build ./... && go run . ./sample-inputs/supplied-input.txt

run-n-ties:
	go build ./... && go run . ./sample-inputs/n-ties-input.txt

run-all-ties:
	go build ./... && go run . ./sample-inputs/all-ties-input.txt

check:
	go build ./... && sh check.sh
