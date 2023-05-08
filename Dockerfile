FROM golang:1-alpine as builder 

RUN apk --no-cache --no-progress add ca-certificates tzdata make \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

#COPY . /projet-go
WORKDIR /api/projet-go

# Download go modules
COPY go.mod .
COPY go.sum .

#RUN go get .
#RUN go install
RUN go mod download -x

RUN go get github.com/swaggo/swag/cmd/swag

COPY . .

COPY --from=itinance/swag /root/swag /usr/local/bin
RUN make godoc
RUN make build

# Create a minimal container to run a Golang static binary
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /api/projet-go/.env .
COPY --from=builder /api/projet-go/users-api .

ENTRYPOINT ["/users-api"]
EXPOSE 8080
