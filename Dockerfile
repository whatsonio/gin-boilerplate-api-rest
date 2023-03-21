FROM golang:1.19 AS build

RUN useradd -u 10001 appuser

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y tzdata && \
    apt-get clean

COPY docker/scripts/install_certificates.sh .
RUN bash ./install_certificates.sh

WORKDIR /usr/src/app

COPY ./app/go.mod  ./
RUN go mod download && go mod verify

COPY ./app ./
RUN go build -v -o /usr/src/app/cmd/web ./...

#
# Second stage
# ------------
FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /usr/src/app/cmd/web /web
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo

USER appuser

ENV PORT ":8080"
EXPOSE 8080

ENTRYPOINT [ "/web" ]