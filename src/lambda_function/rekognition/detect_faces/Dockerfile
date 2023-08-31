# Stage 1: Build stage
FROM golang:latest as build
WORKDIR /function

# Copy dependencies list
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application code
COPY main.go .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -tags lambda.norpc -o main main.go

# Stage 2: Final image
FROM public.ecr.aws/lambda/provided:al2

# Copy the binary from the build stage
COPY --from=build /function/main ./main

# Set the entrypoint
ENTRYPOINT [ "./main" ]
