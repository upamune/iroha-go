flatbuf:
	flatc --gen-mutable --go --grpc -o . schema/*.fbs
deps: 
	glide install	
test:
	go test -v
