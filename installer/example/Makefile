build:
	go build -v -o ./build/hw ./src
	cp ../../LICENSE ./build/LICENSE
	cp -r ./static ./build
pkg:
	go generate
dist:
	go build -v -ldflags "-s -w" -o ./dist/hwi .
upx:
	upx ./dist/hwi