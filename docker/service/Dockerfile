FROM golang:1.22-alpine as build

WORKDIR /build

COPY ./src/ .
RUN go build -o ./output/high-score


FROM alpine:3.19 as final
COPY --from=build /build/output .
COPY --from=build /build/config /etc/highscore
ENTRYPOINT ["high-score"]