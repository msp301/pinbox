.PHONY: build-dev start stop

build-dev:
	docker build --rm -f Dockerfile.dev -t pinbox .

start:
	docker run -d -it --name pinbox-dev -p 4200:4200 -w /pinbox --mount type=bind,source="$(shell pwd)",target=/pinbox pinbox:latest \
		ng serve --host 0.0.0.0 --proxy-config proxy.conf.json

stop:
	docker stop pinbox-dev
	docker rm pinbox-dev
