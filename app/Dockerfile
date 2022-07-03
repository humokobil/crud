FROM golang:latest AS build

COPY . /go/src/crud/.
WORKDIR /go/src/crud/app

# Just for corporate network
ARG cert_location=/usr/local/share/info-watch
ADD /certs/info-watch.crt /etc/ssl/certs/
RUN update-ca-certificates

ENV CGO_ENABLED=0

RUN go mod download \
 && go build -a -installsuffix cgo -o crud .

FROM scratch AS runtime
COPY --from=build /go/src/crud/app/crud ./
COPY --from=build /go/src/crud/conf ./
ENTRYPOINT ["./crud"]
