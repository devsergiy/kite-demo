GOPATH := $(shell pwd):$(shell pwd)
ENV = GOPATH=${GOPATH}

PORT = PORT=6000
UNAME = USERNAME="kontrol"
STORE = STORAGE="postgres"
URL = KONTROLURL="http://127.0.0.1:6000/kite"
PUB_KEY = PUBLICKEYFILE="certs/key_pub.pem"
PRIVATE_KEY = PRIVATEKEYFILE="certs/key.pem"
PG_NAME = POSTGRES_DBNAME="kontrol"
PG_UNAME = POSTGRES_USERNAME="postgres"
PG_PASS = POSTGRES_PASSWORD="postgres"

KONTROL_ENV = $(addprefix KONTROL_, ${PORT} ${UMANE} ${STORE} ${URL} ${PUB_KEY} ${PRIVATE_KEY} ${PG_NAME} ${PG_UNAME} ${PG_PASS})

kontrol_build:
	cd ../src/github.com/koding/kite/kontrol && ${ENV} go build

setup:
	${KONTROL_ENV} ../bin/kontrol --initial

run:
	${KONTROL_ENV} ../bin/kontrol

get:
	cd .. && ${ENV} go get github.com/koding/kite/kontrol/kontrol

db_schema:
	psql postgres -f db/001-schema.sql -U postgres

db_create:
	psql -c 'CREATE DATABASE kontrol owner kontrol;' -U postgres

db_structure:
	psql kontrol -f db/002-table.sql -U postgres

db_build: db_schema db_create db_structure
