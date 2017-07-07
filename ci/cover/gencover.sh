set -e

coveragedir=coverfiles

#rm -r $coveragedir
mkdir $coveragedir

#Iterate through all packages and generate cover profile
for pkg in "$@"; do
        echo $pkg
        file="$coveragedir/$(echo $pkg | tr / -).out"    
        echo $file
        go test -p 1 -covermode=count -coverprofile=$file "$pkg"
done

#Write the mode in a coverage.out file
echo "mode: count" > $coveragedir/coverage.cov

#Apppends all the coverage reports into one file
grep -h -v "^mode" $coveragedir/*.out >> $coveragedir/coverage.cov


push_to_coveralls() {
    goveralls -service concourse-ci -coverprofile $coveragedir/coverage.out -repotoken $TOKEN
}

test_percentage() {
    
    #Runs go tool cover and calculates the total coverage in percentage
    #Filters the output, get the percentage number and removes the percentage sign
    percentage=$(go tool cover -func=$coveragedir/coverage.cov | awk 'NR==0; END{print}' | awk '{print $3}' | sed 's/.$//')
    echo $percentage > $coveragedir/percentage
}

test_percentage