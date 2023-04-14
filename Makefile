# Variables
appname = qe

#  Usage:
# Use make command to run application by golang (https://go.dev/)
#  use start command to run binary mac os file
#  use win command to run binary windows file qe.exe
# use build command to to build application for your GOENV
# after BUILD command the application will be built at the bin/ folder

# Actions
Default:
	go run main.go
start:
	./bin/$(appname)
test:
	go test
win:
	./bin/$(appname)
build:
	go build -o bin/$(appname) main.go