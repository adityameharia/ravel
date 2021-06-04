# parent image
FROM golang:1.15.6-alpine3.12

# workspace directory
WORKDIR /home/adi/golang/src/github.com/adityameharia/ravel

# copy `go.mod` and `go.sum`
ADD ./go.mod ./go.sum ./

# install dependencies
RUN go mod download

RUN apk add build-base

# copy source code
COPY . .

# build executable
RUN go build ./cmd/ravel_cluster_admin 

# expose ports
EXPOSE 42000

# set entrypoint
ENTRYPOINT [ "./ravel_cluster_admin" ]