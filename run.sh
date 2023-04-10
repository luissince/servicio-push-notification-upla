docker stop servicio-push && docker rm servicio-push

docker image rm servicio-push

docker build -t servicio-push .

docker run -d --name servicio-push --net=upla -p 8890:80 servicio-push