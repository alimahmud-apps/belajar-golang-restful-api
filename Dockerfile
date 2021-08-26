#lightweiight container for go
FROM golang:alpine

#update container and packages and istall git
RUN apk update && apk add git --no-cache git

#set /app to be each active directory
WORKDIR /app

#copy all files from outside container , into the container
COPY . .

#download dependency
RUN go mod tidy
#build binary
RUN go build -o binary

#set the entry point of the binary
ENTRYPOINT [ "/app/binary" ]