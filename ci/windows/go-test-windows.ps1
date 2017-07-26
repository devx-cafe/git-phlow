#SET GOPATH TO TMP DIRECTORY, BECAUSE WE WANT THE DEPENDENCIES LOCALLY
$env:GOPATH = "$PWD\"
$env:GOBIN = "$PWD\bin"

#NAVIGATE TO FOLDER
mkdir -p $env:GOPATH/src/github.com/praqma
cp -R git-phlow/ $env:GOPATH/src/github.com/praqma

cd $env:GOPATH/src/github.com/praqma/git-phlow

#GET DEPENDENCIES
go get -d -t -v ./...

#RUN TESTS
cd executor
go test -p 1 -v

# USES THE LATEST EXITCODE, WHICH IS FROM GO TEST, AND EXITS WITH THAT. THIS FIXIS 
# ISSUE #152
exit $lastexitcode