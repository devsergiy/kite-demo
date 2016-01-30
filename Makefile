PROJECT_PATH = src/demo
GOPATH := $(shell pwd):$(shell pwd)/$(PROJECT_PATH)
ENV = GOPATH=${GOPATH}

router:
	cd $(PROJECT_PATH) && ${ENV} go run routerkite.go

app:
	cd $(PROJECT_PATH) && ${ENV} go run appkite.go

get:
	${ENV} go get github.com/koding/kite/kontrol/kontrol
	${ENV} go get github.com/koding/kite/
	# ${ENV} go get github.com/smartystreets/goconvey
	${ENV} go get gopkg.in/yaml.v2
	${ENV} go get -u golang.org/x/tools/cmd/goimports

# test:
# 	cd $(PROJECT_PATH) && ${ENV} go test -v -race

# webtest:
# 	cd $(PROJECT_PATH) && ${ENV} goconvey
