#!/usr/bin/env bash
NAME="csdn博客导出"
mkdir -p output
cp conf.yml output/

name=$1
echo $name "release is buliding"
if [[ $name = "" ]];then
    go build -o output/${NAME}.out
    chmod +x output/${NAME}.out
elif [[ $name =~ "mac" ]];then
    GOOS=darwin GOARCH=amd64 go build -o output/${NAME}_osx.out
    chmod +x output/${NAME}_osx.out
elif [[ $name =~ "linux" ]];then
    GOOS=linux GOARCH=amd64 go build -o output/${NAME}_linux.out
    chmod +x output/${NAME}_linux.out
else
    GOOS=windows GOARCH=amd64 go build -o output/${NAME}.exe
fi