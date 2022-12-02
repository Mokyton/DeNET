all: build

build:
	go build

rebuild: clear build

run:
	go run main.go helpers.go

dvi:
	open https://github.com/Mokyton/DeNET/blob/main/README.md

clear:
	rm -rf DeNET