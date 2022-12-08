FROM golang:1.19 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN GO111MODULE=on go mod download

COPY ./cmd ./cmd
COPY ./main.go ./main.go
RUN CGO_ENABLED=0 go build -v -o nomnemonic-cli

# Create a "nobody" non-root user for the next image by crafting an /etc/passwd
# file that the next image can copy in. This is necessary since the next image
# is based on scratch, which doesn't have adduser, cat, echo, or even sh.
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# No need extra files
FROM scratch

COPY --from=build /app/nomnemonic-cli /
COPY --from=build /etc_passwd /etc/passwd

USER nobody

ENTRYPOINT ["/nomnemonic-cli"]
CMD ["help"]
