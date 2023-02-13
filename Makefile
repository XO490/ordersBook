.PHONY: all build clean
.SILENT:

filename = "ordersBook"
binpath = "bin"
maingo = "cmd/apiserver/main.go"

mkdir:
	mkdir -p $(binpath)

copyConfig:
	cp configs/config.json bin

build: mkdir
#	Linux
	GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -gcflags "-m" -ldflags "-s -H linux" -o $(binpath)/$(filename)-l64 $(maingo)
	#GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -gcflags "-m" -ldflags "-s -H linux" -o $(binpath)/$(filename)-l32 $(maingo)
	#GOOS=linux GOARCH=386 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H linux" -o $(binpath)/$(filename)-l32 $(maingo)

#	Linux ARM
	#GOOS=linux GOARCH=arm64 CC="aarch64-linux-gnu-gcc" CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -gcflags "-m" -ldflags "-s -H linux" -o $(binpath)/$(filename)-arm64 $(maingo)
	#GOOS=linux GOARCH=arm CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -gcflags "-m" -ldflags "-s -H linux" -o $(binpath)/$(filename)-arm32 $(maingo)

#	Windows
	#GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H windows" -o $(binpath)/$(filename)-w64.exe $(maingo)

run: build copyConfig
	$(binpath)/$(filename)-l64 -conf="bin/config.json"
