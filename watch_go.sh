#! /usr/bin/env sh

GOPATH=`pwd`

./bin/srv &

inotifywait -m -r -e close_write src/ | while read line
do 
	go build -o bin/srv app && pkill srv && ./bin/srv &
done

