# Golang API app with MySQL for persistence
RESTful API based on gin-gonic with MySQL database which can be easily containerized and deployed on K8s

The goal of this project is to create an API on Golang for self-development and practice. The app supports such methods as:
* add user;
* get all users;
* get info about a particular user.

The app can be deployed in several ways:
* docker-compose stack
* build and run golang app on PC and make it connect to already running DB
* use helm to deploy to K8s (e.g KinD or Minikube, for home environment)


Also, Swagger documentation was integrated, which simplifies the understanding of the API methods and objects:


![Image](https://raw.githubusercontent.com/Const-antine/golangapi/main/Screenshot%202021-02-04%20at%2007.56.11.png)

## Quickstart
***Before any further steps, make sure that the .env file is created and filled with the necessary values.***
Here is an example (a default .env is also stored in the repository):
```
DBNAME=test
DBUSER=user
DBPASS=password
DBPORT=3306
MYSQL_CONTAINER_NAME_PREFIX=mysql_container
DBHOST=host.docker.internal
DBTABLE=user
MYSQL_ROOT_PASSWORD=cksackjbnadkjvnckajdbnvckjadbkvcb
DB_ROOT_PASSWORD=ojclasclascla
```

To build a binary:
``` 
$ make build 
```

To build a Docker image for the app:
``` 
$ make docker-image 
```

To regenerate Swagger documentation:
``` 
$ make docs 
```

To make an integration test using a local docker-compose stack:
``` 
$ make test 
```

To deploy on Kubernetes using Helm:
``` 
$ make helm 
```


## Deployment steps for KinD

```
# creating KinD cluster (it's also possible to use custom config file here)
$ kind create cluster

# Fetching the repo
$ git clone https://github.com/Const-antine/golangapi.git && cd golangapi

# Fill up the .env file or use the pre-configured file 
$ vim .env

# Make sure that helm cli is installed & kubectl is configured, then
$ make helm

# Make the service accessible locally in the browser by port-forwarding
$ kubectl port-forward service/golangapi-service  8080:8080

# Feel free to visit http://localhost:8080/v1/users/ or http://localhost:8080/swagger/index.html

```

