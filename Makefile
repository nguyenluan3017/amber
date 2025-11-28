amber: main.go
	mkdir -p bin
	go build -o bin/amber

unittest: 
	go test ./... 

clean:
	rm -rfv bin