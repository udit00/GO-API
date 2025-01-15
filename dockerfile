FROM golang:1.23

WORKDIR /app

# COPY go.mod go.sum ./

COPY . .
# RUN go mod download && go mod verify

RUN go get
RUN go build -o bin .

EXPOSE 10000

# ENTRYPOINT [ "/app/bin" , "--port=10000"] 
ENTRYPOINT [ "/app/bin"] 