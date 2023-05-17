all:
	go build ./cmd/vol
	go build ./cmd/back

back:
	go run ./cmd/back

vol:
	go run ./cmd/vol