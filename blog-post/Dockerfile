FROM golang:1.13

WORKDIR /go/src/blog-post

COPY . .

RUN go get -d -v ./...

#RUN go build -o posts .

#EXPOSE 8000

#ENTRYPOINT [ "/posts" ]