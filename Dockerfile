# Use an official Go base image (I use go version '1.22.1')
FROM golang:1.22.1

# Set working directory
WORKDIR /app

# Copy project source files
COPY . /app

# Compile the Go application (if needed)
RUN go build -o ./app/bin/tiny-service

# Define the final image
# Use the /app directory as the base for your app
# Copy the binary in the /app directory to be used later
COPY --from=build-context ./app/bin/tiny-service /app/tiny-service 

# Copy static assets (and directory structure) from /public to /app/public
COPY --from=build-context ./public/ ./app/public/

# Expose port for web application service
EXPOSE 8080

ENTRYPOINT [ "./tiny-service" ]
