FROM golang:alpine

# Create a group named "app" and an user named "app"
RUN addgroup -S app && adduser -S app -G app; \
    # Install task runner
    wget -O - https://taskfile.dev/install.sh | sh

# Use the user created
USER app

# Specify the volume to mount source files and tests
VOLUME [ "/app/test" ]
WORKDIR /app/test

# Run test through task runner
CMD [ "task", "test" ]
