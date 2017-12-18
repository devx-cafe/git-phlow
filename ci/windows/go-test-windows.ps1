#SET GOPATH TO TMP DIRECTORY, BECAUSE WE WANT THE DEPENDENCIES LOCALLY

# this line is important. The worker have been configured through a build session, so the powershell
# have not been reloaded to update the path. Therefore we must set it each time. 
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine")

$env:GOPATH = "$PWD\"
$env:GOBIN = "$PWD\bin"

##NAVIGATE TO FOLDER
mkdir -Path $env:GOPATH/src/github.com/praqma
cp -R git-phlow/ $env:GOPATH/src/github.com/praqma

cd $env:GOPATH/src/github.com/praqma/git-phlow
ls

##GET DEPENDENCIES
go get -d -t -v ./...

##RUN TESTS
cd plugins
go test -v ./...

## USES THE LATEST EXITCODE, WHICH IS FROM GO TEST, AND EXITS WITH THAT. THIS FIXIS 
# ISSUE #152
exit $lastexitcode