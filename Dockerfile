# Use the official image as a parent image.
FROM ubuntu

# Set the working directory.
WORKDIR /usr/local/bin

# Copy the file from your host to your current location.
COPY golang-boilerplate .

# Inform Docker that the container is listening on the specified port at runtime.
EXPOSE 1323

# Run the specified command within the container.
CMD [ "./golang-boilerplate" ]
