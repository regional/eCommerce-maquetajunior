FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o main .

# Use a minimal base image to reduce the size of the final image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the built Go binary from the build stage
COPY --from=build /app/main .

# Copy wait-for-it.sh script
# COPY wait-for-it.sh .

# Ensure wait-for-it.sh has execution permissions
# RUN chmod +x wait-for-it.sh

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
# CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]
CMD ["sh", "-c", "sleep 30 && ./main"]