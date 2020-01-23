test:
	go test ./...

cleanup:
	rm -rf btcd/data
	rm -rf btcwallet/simnet