FROM golang:1.16
RUN mkdir /api
WORKDIR /api
ADD . /api/
RUN go mod download && go get -u github.com/cosmtrek/air
