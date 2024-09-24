all: start hello integers iteration arrays end

start:
	@echo "Starting Test"

hello:
	@go test ./01_hello -v 

integers:
	@go test ./02_integers -v 

iteration:
	@go test ./03_iteration/ -v

arrays:
	@go test ./04_arrays_slices/ -v

structs:
	@go test ./05_structs_methods_interfaces/ -v

end:
	@echo "End Test"


