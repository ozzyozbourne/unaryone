PROTO_DIR = protofiles
PROTO_OUT = pbout
BUILD_DIR = build
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

#protoc code gen command
generate:
	find ${PROTO_DIR} -name "*.proto" -exec protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. {} \;

#clean and generated resource folder
rs_create:
	mdkir savedxlsxfiles

rs_clean:
	rm -rf savedxlsxfiles

rs:	rs_clean rs_create


#clean the output directories
clean_proto:
	rm -rf pbout

clean_server:
	rm -f ${BUILD_DIR}/server_linux
	rm -f ${BUILD_DIR}/server_windows
	rm -f ${BUILD_DIR}/server_mac

clean_client:
	rm -f ${BUILD_DIR}/client_linux
	rm -f ${BUILD_DIR}/client_windows
	rm -f ${BUILD_DIR}/client_mac

clean:	clean_proto clean_server clean_client


#Cross platform build commands
build_server:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/server_linux ./server
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/server_mac ./server
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/server_windows.exe ./server

build_client:
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/client_linux ./client
	GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/client_mac ./client
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/client_windows.exe ./client

build:	build_server build_client

#build from scratch
all:	rs clean generate build



