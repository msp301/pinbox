.PHONY: build-dev start stop

build-dev:
	docker build --rm -f Dockerfile.dev -t pinbox .

build-prod:
	docker build --rm -f Dockerfile -t pinbox-prod .

attach:
	docker run -it --name pinbox-dev -p 4200:4200 -w /pinbox --mount type=bind,source="$(shell pwd)",target=/pinbox pinbox:latest \
		/bin/bash

start-prod:
	docker run -d -p 80:80 --name pinbox pinbox-prod

start:
	docker run -d -it --name pinbox-dev -p 4200:4200 -w /pinbox --mount type=bind,source="$(shell pwd)",target=/pinbox pinbox:latest \
		/bin/bash -c "yarn && ng serve --host 0.0.0.0 --proxy-config proxy.conf.json"

stop-prod:
	docker stop pinbox
	docker rm pinbox

stop:
	docker stop pinbox-dev
	docker rm pinbox-dev
