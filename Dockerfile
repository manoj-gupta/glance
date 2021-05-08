FROM golang:1.16-alpine AS builder

# Set necessary environment variables
ENV GO111MODULE=on

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o glance .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/.env . && cp /build/glance .

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /dist/glance /
COPY /.env /

# Command to run when starting the container
CMD ["/glance"]

