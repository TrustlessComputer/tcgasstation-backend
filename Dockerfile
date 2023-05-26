FROM golang:1.18-bullseye as deps

RUN apt-get -y update && apt-get -y upgrade && \
    apt-get -y install git && \
    apt-get -y install make

ARG ENV=dev

ENV ENV=${ENV} \
    CGO_ENABLED=1

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN make init

RUN go mod download

FROM deps as builder
COPY  . .

RUN echo "✅ Build for Linux"; make build


FROM debian:bullseye as runner

RUN apt-get -y update \
    && apt-get install -y ca-certificates

WORKDIR /app

COPY --from=builder /app/tc-gasstation-backend /app/tc-gasstation-backend

CMD ["/app/tc-gasstation-backend"]
