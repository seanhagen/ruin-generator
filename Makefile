#!make

GO_BUILD_ENV:=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

OUTPUT="./cmd"
DEBUG_OUTPUT="./debug"

LDFLAGS=$(LDFLAGSBASE) -X main.Repo=github.com/seanhagen/ruin-generator
LDFLAGBUILD=-ldflags "$(LDFLAGS) -s -w"
LDFLAGDEBUG=-ldflags "$(LDFLAGS)"

TEST_FILES=$(shell find -name '*_test.go')
RESULTS_DIR=./results

TEST_OUTPUT=$(RESULTS_DIR)/test-output.xml
COV_OUTPUT=$(RESULTS_DIR)/coverage.out
COV_TXT_OUTPUT=$(RESULTS_DIR)/coverage.txt
COV_HTML_OUTPUT=$(RESULTS_DIR)/coverage.html
COB_OUTPUT=$(RESULTS_DIR)/coverage.xml
CHECK_OUTPUT=$(RESULTS_DIR)/checkstyle-result.xml


$(OUTPUT): generate  $(SOURCES)
	$(GO_BUILD_ENV) go build -a ${LDFLAGBUILD} -o ${OUTPUT} -installsuffix cgo .

$(DEBUG_OUTPUT): generate  $(SOURCES)
	go build -a ${LDFLAGDEBUG} -o ${DEBUG_OUTPUT} -gcflags="-N -l" -installsuffix cgo -race .

$(RESULTS_DIR):
	mkdir -p $(RESULTS_DIR)

$(CHECK_OUTPUT):  $(RESULTS_DIR)
	gometalinter.v2 -j 2 --deadline=60s --checkstyle \
		--enable=nakedret --enable=unparam --enable=megacheck \
		--skip=client --vendor `go list ./... | grep -v vendor` > $(CHECK_OUTPUT) || true

vet: $(CHECK_OUTPUT)

$(TEST_OUTPUT) $(COV_OUTPUT): $(RESULTS_DIR)
	go test -v -coverprofile=$(COV_OUTPUT) -covermode count -short \
		`go list ./... | grep -v vendor ` \
		2>&1 | go-junit-report > $(TEST_OUTPUT)

junit-test: $(TEST_OUTPUT)

$(COV_HTML_OUTPUT): $(COV_OUTPUT)
	go tool cover -html=$(COV_OUTPUT) -o $(COV_HTML_OUTPUT)

$(COB_OUTPUT): $(COV_OUTPUT)
	gocover-cobertura < $(COV_OUTPUT) > $(COB_OUTPUT)

coverage: $(COV_HTML_OUTPUT) $(COB_OUTPUT)

test:
	go test -v `go list ./... | grep -v vendor` -covermode count -short

citest: clean vendor junit-test vet coverage

build: $(OUTPUT)

debug: $(DEBUG_OUTPUT)

clean:
	if [ -f ${OUTPUT} ] ; then rm ${OUTPUT} ; fi
	if [ -f ${DEBUG_OUTPUT} ] ; then rm ${DEBUG_OUTPUT} ; fi
	if [ -d results ] ; then rm -rf results ; fi


.DEFAULT_GOAL: build-clean
.PHONY: clean generate test vet deps build build-container build-clean run deploy grpc cideploy
