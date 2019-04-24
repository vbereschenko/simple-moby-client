#!/bin/bash

go get && \
go build simple-moby-client.go
echo "Container list check 1"
./runner list

containerId=`./simple-moby-client run alpine /bin/sleep 60`
echo "Created container: ${containerId}"

echo "Container list check 2"
./simple-moby-client list

./simple-moby-client stop "${containerId}"

echo "Container list check 3"
./simple-moby-client list
rm ./simple-moby-client