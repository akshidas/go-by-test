all: start hello integers end

start:
	@echo "Starting Test"

hello:
	@go test ./01_hello -v 

integers:
	@go test ./02_integers -v 

end:
	@echo "End Test"


