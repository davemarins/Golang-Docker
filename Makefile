docker:
	docker build -t golang_web_server .
	docker run -p 8080:8080 golang_web_server

clean:
	docker image prune -a
	docker container prune
	docker volume prune
	docker image prune -a

build:
	go build -a -o main
	chmod +x main

run:
	./main

install:
	glide install
	cd Angular/frontend/
	npm install

deploy:
	gcloud builds submit --tag eu.gcr.io/<PROJECT-NAME>/<CONTAINER-IMAGE-NAME>
	gcloud compute instances update-container <VM-INSTANCE-NAME>/<VM-INSTANCE-GROUP-NAME>

frontend:
	cd Angular/frontend && ng build --prod --aot && cd ..
