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

You should also use a .env file in order to load all configurations which you don't want to commit and push into a
public repository. In this code provided, your .env file should look something like this:

```
MYSQL_USERNAME=YOUR-USERNAME
MYSQL_PASSWORD=YOUR-PASSWORD
MYSQL_DATABASE=YOUR-DATABASE-NAME
JWT_SECRET=YOUR-SECRET
ISSUER_NAME=YOUR-ISSUER-NAME
```

```.env``` and ```.env.prod``` should be present in order to change environment.
CI/CD integration has been made through GCP Cloud Build and Container Registry. Once you type ```make deploy```
if configured in CGP the image is built, published into the Container Registry and the same image is substituted
in all running VM instances of the group (so that downtime is not present).
