#Docker on production - sysadmin@galicia meetup
![Meetup logos](http://i.imgur.com/2MeSFWS.png)

This repository contains the assets used on the ["Docker in production" sysadmin@galicia](https://www.meetup.com/es-ES/Sysadmin-Galicia/events/235823265/) meetup on 26/01/2017.

During the meetup we performed a live demo about how-to compile, build, deploy and operate a simple GO web service on a Raspberry PI Kubernetes cluster, the repo root contains the service code `service.go` the `Dockerfile` and the Kubernetes deployment and service configuration files `k8s-deployment.yaml` and `k8s-service.yaml` that we used on the live demo.

The non live demo slides can be found [here](https://docs.google.com/presentation/d/1zFVJMEyPL05P9EGfQ53MMPbePEMSU34JUCFp8929ZMY/pub?start=false&loop=false&delayms=10000&slide=id.g35f391192_00).

Commands used to create the GO arm executable.

```bash
# Build the service to be shown on a browser.
go build -o service . 

# Build the service deployable on the Kubernetes PI cluster by using the arm ARCH.
GOARCH=arm go build -o service .
```

Commands used to build and publish the docker image.

```bash
# Create the deployable Docker image.
docker build -t jcorral/sample-web-service:v1 .

# Get the Docker image ready to download.
docker push jcorral/sample-web-service:v1
```

Commands used to deploy and operate the service on the Kubernetes Cluster. 

```bash
# Create a Kubernetes Deployment with the pushed image.
kubectl create -f k8s-deployment.yaml

# Create a Kubernetes Service for the created Deployment in order to accept requests.
kubectl create -f k8s-service.yaml

# Update the Docker image used on a Deployment.
kubectl set image deployment/sample-web-service sample-web-service=jcorral/sample-web-service:v2

# Update the Kubernetes deployment
kubectl edit deployment sample-web-service

# Scale the Kubernetes deployment
kubectl scale --replicas=10 sample-web-service

# Status commands to be used on the demo

# Get Nodes & Status
watch kubectl get nodes

# Infinite curl loop to get the service version
while true; do curl <service ip:port>/version; printf "\n"; sleep 1; done;

# Get Kubernetes Deployment rollout status
watch kubectl rollout status deployment/sample-web-service

# Get Kubernetes Pods Status
watch kubectl get po

# Get Kubernetes Pod Nodes
watch -n1 "kubectl describe po | grep Node:"
```
