#!/bin/bash
# The script does automatic checking on a Go package and its sub-packages, including:
# 1. gofmt         (http://golang.org/cmd/gofmt/)
# 2. goimports     (https://github.com/bradfitz/goimports)
# 3. golint        (https://github.com/golang/lint)
# 4. go vet        (http://golang.org/cmd/vet)
# 5. ineffassign   (https://github.com/gordonklaus/ineffassign)
# 6. race detector (http://blog.golang.org/race-detector)
# 7. test coverage (http://blog.golang.org/cover)

set -e

# Automatic checks
test -z "$(gofmt -l -w .     | tee /dev/stderr)"
test -z "$(golint .          | tee /dev/stderr)"

CI_SERVICE="travis-ci"
DIR_SOURCE="$(find . -maxdepth 10 -type f -not -path '*/Godeps*' -name '*.go' | xargs -I {} dirname {} | sort | uniq)"

go vet ${DIR_SOURCE}

# Run test coverage on each subdirectories and merge the coverage profile.

echo "mode: count" > profile.cov

for dir in ${DIR_SOURCE};
do
    go test -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> profile.cov
        rm $dir/profile.tmp
    fi
done

go tool cover -func profile.cov

# To submit the test coverage result to coveralls.io,
# use goveralls (https://github.com/mattn/goveralls)
if [ -n "${CI_SERVICE+1}" ]; then
    echo "goveralls with" $CI_SERVICE
    if [ -n "${COVERALLS_TOKEN+1}" ]; then
        goveralls -coverprofile=profile.cov -service=$CI_SERVICE -repotoken $COVERALLS_TOKEN
    else
        goveralls -coverprofile=profile.cov -service=$CI_SERVICE
    fi
fi