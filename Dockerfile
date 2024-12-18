FROM golang:1.23.4-alpine AS build
RUN mkdir /src
WORKDIR /src
ADD ./go.mod .
ADD ./*.go ./

RUN go build -o archTest
RUN chmod +x archTest

FROM scratch
COPY --from=build /src/archTest /usr/local/bin/archTest
CMD ["archTest"]
