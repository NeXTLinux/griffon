FROM docker.io/golang:1.20@sha256:690e4135bf2a4571a572bfd5ddfa806b1cb9c3dea0446ebadaf32bc2ea09d4f9
WORKDIR /go/src/github.com/anchore/test/
COPY golang/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-app .

FROM scratch
COPY --from=0 /go/src/github.com/anchore/test/go-app ./
COPY . .
