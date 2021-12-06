all: build push

IMAGENAME=gc21-pyqa
REGISTRY=dwhitena

build:
	docker build --force-rm=true -t ${REGISTRY}/${IMAGENAME} .

test:
	docker run ${REGISTRY}/${IMAGENAME} pytest -v

push:
	docker push ${REGISTRY}/${IMAGENAME}
