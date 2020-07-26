# A simple golang web application which can be deployed in the Docker container. 

In order to deploy it, please follow the steps:

1. wget https://github.com/Const-antine/golang_docker/archive/master.zip
2. unzip master.zip && cd golang_docker-master/golan/
3. docker build .
4. docker run -e "API_KEY=|your API key|" -p 8080:8000 (instead of 8080, you may specify you own port via '-e "PORT=|your port|" parameter') |name of built image|

Do not forget to create your own .env file where the variables (for connection to DB) for your app will be stored. Here is an example

```
db_name = mongo-api
db_pass = ****
db_user = postgres
db_type = postgres
//db_host = localhost
// if you use Mac and need to deploy the app on localhost:
db_host = host.docker.internal 
db_port = 5434
token_password =  *******************

//these dynamic variables are pasted to docker-compose.yaml 
POSTGRES_PASSWORD=****
POSTGRES_DB=mongo-api
DBPORT=5434
```


