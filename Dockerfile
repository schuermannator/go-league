FROM golang:1.14 as builder

# Copy the code from the host and compile it
WORKDIR BUILD
COPY . ./
RUN go get -u github.com/gin-gonic/gin
RUN CGO_ENABLED=0 go build -a -installsuffix nocgo -o /app .
COPY ./static /static
COPY ./templates /templates

FROM scratch
COPY --from=builder /app ./
COPY --from=builder /static ./static
COPY --from=builder /templates ./templates
EXPOSE 8080
ENTRYPOINT ["./app"]
