FROM golang:1.11 as builder

# Copy the code from the host and compile it
WORKDIR ~/Desktop/Dev/go-league
COPY . ./
RUN go get -u github.com/gin-gonic/gin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .
COPY ./static /static
COPY ./templates /templates

FROM scratch
COPY --from=builder /app ./
COPY --from=builder /static ./static
COPY --from=builder /templates ./templates
ENTRYPOINT ["./app"]