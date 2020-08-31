FROM golang:1.15
EXPOSE 8081
WORKDIR /go/src/
COPY . . 
RUN GOOS=linux 
CMD ["./catalog"]