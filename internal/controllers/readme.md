To deploy on minikube run the following command:

```
minikube docker-env
```

This will give you the command to run to point to the minikube docker environment.

Then build the docker image and deploy to minikube, which will then pull from the local docker registry.