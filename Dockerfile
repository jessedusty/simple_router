# build stage
FROM golang:alpine AS build-env
ADD . /src/
RUN cd /src && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server
# final stage
FROM scratch
COPY --from=build-env /src/server /
EXPOSE 80
CMD [ "/server" ]
