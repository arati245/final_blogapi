FROM golang:1.24.5

WORKDIR /Final_Blogapi-master

COPY . .

RUN go mod tidy
RUN go build -o blog-api main.go

EXPOSE 8080

CMD ["./blog-api"]