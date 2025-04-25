FROM alpine:3.18

WORKDIR /root/

# Copy the pre-built Go binary into the container
COPY main .

# Grant execute permission to the Go binary
RUN chmod +x ./main

# Expose the port the app will listen on
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
