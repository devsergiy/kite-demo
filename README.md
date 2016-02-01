# Micro-demo
Micro-demo is a microservices sample application powered by [koding/kite](https://github.com/koding/kite) library

### Setup

You need installed golang version 1.5 and postgres.
Get project dependencies by running in project root next command:
```sh
make get
```
Application is using kite kontrol service for interaction between services("kites"). To setup it you have to run sequence of commands.
Note: Probably you will need to modify postgres credentials in the makefile and maybe modify db commands to use proper user and password. After complete that you could continue with commands
```sh
$ cd kontrol
$ make db_build
$ make setup
```
After that you will get ~/.kite/kite.key file signed by kontrol for use on local machine and properly setuped db schema.

### Run
To be able to interact between services you will need running instance of kontrol
```sh
$ cd kontrol
$ make run
```
To run services use next command.
Note: key -j4 is required to be able to run all services in parallel
```sh
$ make -j4 run
```

### Interaction with app

##### Login query

|  |  |
|---|---|
| Request type: | **post** |
| Url |  **http://localhost:3000/login** |
| Params | username: "username", password: "password" |
| Curl query| curl -X POST -F "username=username" -F "password=password" "http://localhost:3000/login" |

Response will be
```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.2mR6Oa_ZjoYZBK4Mayd6jYgSpe_z0HZQS_cBEEdkSjU"}
```
This token should be used in all next queries

##### Profile query

|  |  |
|---|---|
| Request type: | **get** |
| Url |  **http://localhost:3000/profile** |
| Headers | should contain token |
| Curl query| curl -X GET -H "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.2mR6Oa_ZjoYZBK4Mayd6jYgSpe_z0HZQS_cBEEdkSjU" -H "Cache-Control: no-cache" -H "Postman-Token: 9e3742cc-38ca-2f95-b478-d175acb512fc" "http://localhost:3000/profile" |

Response will be
```json
{"age":"20","id":"1","name":"Sammy The Bug"}
```

##### Todos query

|  |  |
|---|---|
| Request type: | **get** |
| Url |  **http://localhost:3000/todos** |
| Headers | should contain token |
| Curl query| curl -X GET -H "token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.2mR6Oa_ZjoYZBK4Mayd6jYgSpe_z0HZQS_cBEEdkSjU" -H "Cache-Control: no-cache" -H "Postman-Token: 4227ade7-4fde-a3a0-7654-dd1c58fa59d7" "http://localhost:3000/todos" |

Response will be
```json
["cleanup things","buy a new phone","sort pencils","change the habits"]
```

### Implementation drawbacks
- You could need some time to understand logic of sending params between services
- Not all error cases are covered
- Skipped test implementation excepting authorizer
- Not all handlers extracted to separated files or objects
- Logging done via fmt to the app default output

### Reason of choose Kite
- Built-in support for routing via gorilla/mux
- Kontrol - tool for registering services to each other
- Rather clean api
- Built-in support for different auth types for services

### Implementation alternatives

- https://github.com/micro/go-micro - uses protobufs and has rather clean api
- https://github.com/go-kit - api is not so clear
- Custom solution build with use of https://github.com/docker/libchan for talk between services as with simple channels and https://github.com/gorilla/mux for the routing
- https://github.com/goadesign/goa - tool for generation services with built in tools for api docs
