.PHONY: all tidy vendor zip

all:
	GOOS=linux go build -mod=vendor ./...
	mv $$(ls ./cmd) ./bin

tidy:
	go mod tidy

vendor:
	go mod vendor

zip:
	cd ./bin; for file in $$(ls); do zip "$${file}" "$${file}"; done
	cd ./bin; mv *.zip ../zips
