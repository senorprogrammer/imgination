install:
	go get
	go clean
	go install

run:
	go run imgination.go

clean:
	go clean
	rm -f imgination
