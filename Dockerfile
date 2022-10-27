FROM registry.imroc.cc/library/build-tools:latest as builder
COPY . /build
RUN cd /build && go build -o /server

FROM ubuntu:22.04
COPY --from=builder /server /server
CMD ["/server"]