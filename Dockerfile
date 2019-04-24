FROM golang:1.8

WORKDIR /go/src/mtgrestservice
COPY . .

RUN go get -u -v github.com/gorilla/mux
RUN go get mtgrestservice
#RUN go install
RUN go build

CMD ["./mtgrestservice"]


#RUN go get -d -v ./...
#RUN go install -v ./...
