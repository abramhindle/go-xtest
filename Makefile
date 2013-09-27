
run:
	rm $(GOPATH)/src/xtest || echo "Ok"
	ln -sf `pwd`/xtest $(GOPATH)/src/xtest
	go install xtest
	go run driver.go
