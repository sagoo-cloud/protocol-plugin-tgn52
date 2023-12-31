BINARY_NAME=protocol-tgn52

local:
	echo "========local============"
	go build -o ${BINARY_NAME}
	mv ${BINARY_NAME} ./built

linux:
	echo "========linux============"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}
	mv ${BINARY_NAME} ./built

windows:
	echo "========windows============"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}.exe
	mv ${BINARY_NAME} ./built