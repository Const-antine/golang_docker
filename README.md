# A simple golang web application which can be deployed in the Docker container. 

In order to deploy it, please follow the steps:

1. wget https://github.com/Const-antine/golang_docker/archive/master.zip
2. unzip master.zip && cd golang_docker-master/golan/
3. docker build .
4. docker run -e "API_KEY=|your API key|" -p 8080:8000 (instead of 8080, you may specify you own port via '-e "PORT=|your port|" parameter') |name of built image|




