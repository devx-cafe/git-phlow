#!/bin/sh


DIR_NAME=build/phlow-test-pkg


#CREATE TEST DIRECTORY
mkdir -p -m 0755 $DIR_NAME
cd $DIR_NAME

#CREATE TEST FILES
touch README.md file1.txt file2.txt


#DECLARE GIT FUNCTIONS
create_inital_commit () {
    git init
    git add .
    git commit -m "initial commit"
}

two_commits_on_master () {
    echo "hello" > README.md
    add_all_and_commit "1 commit - changes to readme"
    echo "file1" > file1.txt
    add_all_and_commit "2 commit - changes to file1"
}

branch_foo_additions (){
    git checkout -b  foo
    echo "hello world to README" > README.md
    echo "on branch foo" > file1.txt
    add_all_and_commit "changes on branch foo"
}


branch_master_additions () {
    git checkout master
    echo "hello world to README from master" > README.md
    add_all_and_commit "changes on branch master"
}


branch_issue_additions () {
    git checkout -b 11-issue-bar
    echo "changes in file 2" > file2.txt
    echo "changes in file 1" > file1.txt
    add_all_and_commit "changes on branch 11-issue"
}

return_to_master () {
    git checkout master
}

add_all_and_commit (){
    git add .
    git commit -m "$1"
}


#ACTUAL SCRIPT
echo "CREATING TEST REPOSITORY"
create_inital_commit
two_commits_on_master
branch_foo_additions
branch_master_additions
branch_issue_additions
echo "WRAPPING UP JOB"
