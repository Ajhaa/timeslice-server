FROM golang:1.16 AS build-stage

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/timeslice/main.go


FROM scratch

# ARG COMMIT_SHA
# ENV COMMIT_SHA=${COMMIT_SHA}

COPY --from=build-stage /app/server /app/server

ENTRYPOINT [ "/app/server" ]