PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
OUTPUT = PACKAGE/app/core/structs

build: generate
	   go build .

generate:
	protoc -Iproto --go_opt=module=${PACKAGE} --go_out=.${OUTPUT}*.proto

clean:
	rm ${OUTPUT}/*.pb.go