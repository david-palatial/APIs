FROM alpine:3.14.1

# Create a directory for the binary
RUN mkdir instance-customisation-plugin

# Changing our working directory to this new directory
WORKDIR /instance-customisation-plugin

# Copy the binary over. This requires that the binary has already been built
COPY bin/instance-customisation-plugin ./server

# Copy kubeconfig over so that the server can access the cluster
COPY config .

# Set an entry point that runs when this container is launched
ENTRYPOINT [ "/instance-customisation-plugin/server" ]
