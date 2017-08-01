#!/bin/sh

set -e -u -x

export VERSION=$(cat gp-version/version)

#Generate the scoop manifest
python git-phlow/ci/scoop/scoop.py 

#Print the manifest to screen 
cat scoop-bucket/git-phlow.json

#Commit everything in the scoop bucket
cd scoop-bucket
git add --all
git commit -m "$VERSION released"
