set -e

coveragedir=cover

run_coverage(){
rm -r cover
mkdir $coveragedir

#Iterate through all packages and generate cover profile
for pkg in "$@"; do
    
        file="$coveragedir/$(echo $pkg | tr / -).out"    
        echo $file
        go test -covermode=count -coverprofile=$file "$pkg"
done

#Write the mode in a coverage.out file
echo "mode: count" > $coveragedir/coverage.out 

#Apppends all the coverage reports into one file
grep -h -v "^mode" $coveragedir/*.out >> $coveragedir/coverage.out
}

push_to_coveralls() {
    goveralls -service concourse-ci -coverprofile $coveragedir/coverage.out -repotoken $TOKEN
}

test_percentage() {
    
    #Runs go tool cover and calculates the total coverage in percentage
    #Filters the output, get the percentage number and removes the percentage sign
    percentage=$(go tool cover -func=$coveragedir/coverage.out | awk 'NR==0; END{print}' | awk '{print $3}' | sed 's/.$//')
    echo $percentage
}

test_percentage

#!/usr/bin/env bash

#set -e
#echo "" > coverage.txt

#for d in $(go list ./... | grep -v vendor); do
#    go test -race -coverprofile=profile.out -covermode=atomic $d
#    if [ -f profile.out ]; then
#        cat profile.out >> coverage.txt
#        rm profile.out
#    fi
#done