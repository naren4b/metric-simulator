.PHONY: all

all: clean build

build: 
  go build main.go

clean: 
  go clean  

run: 
  go run main.go

dockerbuild:
 docker build -t naren4b/metric-simulator .

docckerpush:
 docker push naren4b/metric-simulator

 


