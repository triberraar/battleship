# battle ship game

Proof of concept battleship game.
Ui written with phaser 3 and typescript.
Backend written in golang with gorilla

## demo

on heroku: https://whispering-caverns-71033.herokuapp.com/

## How to build

First build the ui by running `npm install` `npm run build` in the root folder.

Next build the backend by running `go build -o bin/go-battleship ./cmd/go-battleship/main.go` in the root folder

Run the whole thing by executing `./bin/go-battleship`

## How to develop

To develop you need to run two servers. One runs and watches the frontend code, the other one runs the go server

Run the ui by running `npm run watch` in the root folder. This will run a server on port 100001.
Run the backend by running `go run ./cmd/go-battleship/` in the root folder. This will run a server on port 100002.
You can use either the 100001 or 10002 url as they both serve the same files, the watch command ll recompile.

## How to deploy

Make sure the frontend is built, commit to master, and run `git push heroku master`. You can then open the webpage by running `heroku open` and view the logs with `heroku logs --tail`.
