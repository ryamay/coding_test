#!/bin/sh

# This script creates a new module in the new directory
mkdir $1
cd $1
go mod init github.com/ryamay/coding_test/$1
touch main.go

# add go.work
cd ../
go work use ./$1
