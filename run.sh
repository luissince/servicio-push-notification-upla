docker stop servicio-bbva && docker rm servicio-bbva

docker image rm servicio-bbva

docker build -t servicio-bbva .

docker run -d --name servicio-bbva --net=upla -p 8890:80 servicio-bbva