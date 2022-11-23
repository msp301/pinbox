FROM golang:1.19.3-bullseye
RUN apt-get update && apt-get -y install notmuch libnotmuch-dev
RUN mkdir -p /pinbox-server /mail
WORKDIR /pinbox-server
COPY . /pinbox-server/
RUN make build
CMD [ "./pinbox-server", "configs/docker.toml" ]