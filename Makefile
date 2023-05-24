BINARY_NAME=lwe

build:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ${BINARY_NAME}_mac main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ${BINARY_NAME}_linux main.go
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ${BINARY_NAME}_win64.exe main.go
run:
	./${BINARY_NAME}

release:
	# Clean
	go clean
	rm -rf ./*.gz

	# Build for mac
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ${BINARY_NAME}_mac main.go
	tar czvf ${BINARY_NAME}_mac_${VERSION}.tar.gz ./${BINARY_NAME}_mac
	rm -rf ./${BINARY_NAME}_mac

	# Build for linux
	go clean
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ${BINARY_NAME}_linux main.go
	tar czvf ${BINARY_NAME}_linux_${VERSION}.tar.gz ./${BINARY_NAME}_linux
	rm -rf ./${BINARY_NAME}_linux

	# Build for win
	go clean
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ${BINARY_NAME}_win64.exe main.go
	tar czvf ${BINARY_NAME}_win64_${VERSION}.tar.gz ./${BINARY_NAME}_win64.exe
	rm -rf ./${BINARY_NAME}_win64.exe
	go clean

clean:
	go clean
	rm -rf ./${BINARY_NAME}_mac
	rm -rf ./${BINARY_NAME}_linux
	rm -rf ./${BINARY_NAME}_win64.exe
	rm -rf ./*.gz


# 发版命令
# make release VERSION=1.0.0