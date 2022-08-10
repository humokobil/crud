FROM golang:1.19rc1-bullseye AS build

COPY . /go/src/crud/.
WORKDIR /go/src/crud/migration

# Just for corporate network
ARG cert_location=/usr/local/share/info-watch
ADD ./certs/info-watch.crt /etc/ssl/certs/
RUN update-ca-certificates

ENV HTTP_PROXY=http://192.168.0.19:4480
ENV HTTPS_PROXY=http://192.168.0.19:4480
# End

ENV CGO_ENABLED=0


RUN go mod download \
 && go build -a -installsuffix cgo -o crud_migration .

FROM scratch AS runtime
COPY --from=build /go/src/crud/migration/crud_migration ./
COPY --from=build /go/src/crud/conf ./conf/
ENTRYPOINT ["./crud_migration", "--conf-file=./conf/config.yaml"]
