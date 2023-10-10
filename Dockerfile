FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .
RUN make clean build test-unit xbuild-all test-integration
RUN make install

#CMD ["/app/main"]