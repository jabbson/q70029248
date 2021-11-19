# STAGE 1: Build
FROM golang:latest AS build
WORKDIR /go/src/app
COPY [ "./src/go-app.go", "./src/go.mod", "/go/src/app/" ]
RUN go build -o go-app

# -------------------------------------------------
# STAGE 2: Deployment
FROM scratch
COPY --from=build /go/src/app/go-app /go-app
ENTRYPOINT [ "/go-app" ]
