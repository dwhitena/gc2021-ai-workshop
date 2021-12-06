all: compile build push clean

IMAGENAME=gc21-goqa
REGISTRY=dwhitena

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o gcqa

build:
	docker build --no-cache=true --force-rm=true -t ${REGISTRY}/${IMAGENAME} .

push:
	docker push ${REGISTRY}/${IMAGENAME}

clean:
	rm gcqa

test:
	go test -v

