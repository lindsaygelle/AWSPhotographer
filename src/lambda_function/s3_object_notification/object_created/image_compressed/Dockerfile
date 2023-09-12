# Stage 1: Build the Go application
FROM golang:1.21 as build
WORKDIR /function

# Copy all Go source files
COPY . .

# Build the Go application
RUN go build -o main

# Stage 2: Create a clean image for the Lambda function
FROM public.ecr.aws/lambda/provided:al2

# Copy the built executable from the previous stage
COPY --from=build /function/main ./main

# Set the entry point
ENTRYPOINT [ "./main" ]
