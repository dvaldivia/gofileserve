FROM golang as golayer


ADD go.mod /go/src/github.com/dvaldivia/gofileserve/go.mod
#ADD go.sum /go/src/github.com/dvaldivia/gofileserve/go.sum
WORKDIR /go/src/github.com/dvaldivia/gofileserve/

RUN go mod download


COPY . /go/src/github.com/dvaldivia/gofileserve/

RUN go build --tags=kqueue -ldflags "-w -s" -a -o gofileserve .

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.7
EXPOSE 4000

COPY --from=golayer /go/src/github.com/dvaldivia/gofileserve/gofileserve .

ENTRYPOINT ["./gofileserve"]