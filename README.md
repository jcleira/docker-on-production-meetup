#Docker on production - sysadmin@galicia meetup

This repository contains the assets used on the ["Docker in production" sysadmin@galicia](https://www.meetup.com/es-ES/Sysadmin-Galicia/events/235823265/) meetup on 26/01/2017.

The live demo on the meetup consists on a serie of compile, build & operate examples of GOlang web services deployed using Docker images on a Kubernetes Raspberry PI cluster.
Here is the command list used during the demo:

```bash
# Build the service to be shown on a browser.
go build -o service . 

# Build the service deployable on the Kubernetes PI cluster by using the arm ARCH.
GOARCH=arm go build -o service .

# Create the deployable Docker image.
docker build -t jcorral/sample-web-service:v1 .

# Get the Docker image ready to download.
docker push jcorral/sample-web-service:v1
```
