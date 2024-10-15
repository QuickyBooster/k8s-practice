Lesson 001: First API with go

### 1. Build docker image:
```
docker image build -t quickybooster/go-app:2.0.0 -f go-app/Dockerfile go-app/
```

### 2. Push to dockerhub
```
docker push quickybooster/go-app:2.0.0
```

### 3. Deploy deployment file
```
kubectl apply -f k8s/go-app/deployment.yaml
kubectl get pod -n default
```
