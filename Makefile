prof:
	time go run . 2
	go tool pprof -http localhost:8888 mem.pprof

bench:
	go test --bench=.

build:
	go build ./
