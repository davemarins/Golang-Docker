docker:
	docker build -t golang_web_server .
	docker run -p 8080:8080 golang_web_server

build:
	go build -a -o main
	chmod +x main

run:
	./main

angular:
	cd frontend && ng build --prod --aot && cd ..