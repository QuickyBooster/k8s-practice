Lesson 001: First API with go

expose metrics with this post: 
- [service-monitor](https://medium.com/daemon-engineering/exposing-metrics-to-prometheus-with-service-monitors-326f38b2daf1)
- [pod-monitor + service-monitor](https://medium.com/@helia.barroso/a-guide-to-service-discovery-with-prometheus-operator-how-to-use-pod-monitor-service-monitor-6a7e4e27b303)
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
<- k6 -> pushmetrics <- prometheus