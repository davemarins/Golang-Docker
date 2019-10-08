# Golang Docker simple REST CRUD Application

This app is an example of blog manager with simple functionalities.

### Install

Make sure to have Go installed in your computer. More on https://golang.org/dl/

Clone this repository into your go home path.

> git clone https://github.com/davidemarino0508/Golang-Docker

If you don't have glide installed, make sure to do that with

> export GOPATH=/path/to/home/folder
> export GOBIN=$GOPATH/bin
> export PATH=$PATH:$GOPATH/bin
> glide install
> curl https://glide.sh/get | sh

Now you're able to install all dependencies and run the app

> cd cmd/main/ && go run main.go

### Config

Serving port is in main.go. <br>
JWT secret and issuer name is in utils.go. <br>
DB configuration is in app.go <br>
All non-authenticated API have '/public/' for the sake of simplicity. See more in auth.go. <br>

If you want to install dependencies, type

> glide get github.com/foo/bar

### Future development

Dockerization is about to come real soon. The Angular app is about to be developed.
