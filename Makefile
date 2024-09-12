test:
	go build .
	@time ./palindromes --limit 1_000_000_000_0 v3
	@echo "--------------------------------------------------"
	@time ./palindromes --limit 1_000_000_000_0 v4

clearrun:
	clear;
	go run . v4
