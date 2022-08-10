FROM golang:latest AS build

COPY . /go/src/crud/.
WORKDIR /go/src/crud/cmd/app

# Just for corporate network
ARG cert_location=/usr/local/share/info-watch
ADD /certs/info-watch.crt /etc/ssl/certs/
RUN update-ca-certificates

ENV HTTP_PROXY=http://192.168.0.19:4480
ENV HTTPS_PROXY=http://192.168.0.19:4480
# End

ENV CGO_ENABLED=0
 
RUN go mod download \
    && go build -a -installsuffix cgo -o crud .

FROM scratch AS runtime
COPY --from=build /go/src/crud/cmd/app/crud ./
COPY --from=build /go/src/crud/conf ./
ENTRYPOINT ["./crud"]
