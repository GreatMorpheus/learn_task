export tag=v1.0

build:
	echo "building httpserver"
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -o bin/httpserver .
release:build
	echo "building httpserver image"
	docker build -t 569646632/task_3:${tag} .
	#docker build -t 127.0.0.1:5000/httpserver:${tag} .
push:release
	echo "pushing httpserver to Docker hub"
	docker push 569646632/task_3:${tag}
	#docker push 127.0.0.1:5000/httpserver:${tag}
