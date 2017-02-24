FROM golang:1.7
MAINTAINER <groenborg> 

#Working directory for the phlow project
WORKDIR $GOPATH/src/github.com/praqma/git-phlow

COPY . .

# -d retreievs the dependencies without installing them
# -v is for verbse mode
# ./... runs through all packages and downloads test deps and prod deps
RUN go get -d -t -v ./...

#Run tests in sequence and not parallel
#the tests setup does not support parallel 
#testing at the moment
ENTRYPOINT go test -v -p 1 ./...