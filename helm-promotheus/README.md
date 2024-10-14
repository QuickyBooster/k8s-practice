Install prometheus using helm

### 2. Add repo and modify values
```
# adding repo
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

# updating repo
helm repo update

# searching chart
helm search repo kube-prometheus-stack

# showing version to yaml file (add --version if needed)
helm show values prometheus-community/kube-prometheus-stack --version 65.2.0 > ./helm-promotheus/prom-values.yaml

# edit the prom-values.yaml file

# commonLabels:                           #line 27
#   prometheus: devops
# defaultRules.rules.etcd: false          #line 51
# adminPassword: admin                    #line 984
# kubeControllerManager:                  #line 1469
#   enabled: false
# kubeEtcd:                               #line 1744
#   enabled: false
# kubeScheduler:                          #line 1849
#   enabled: false
# serviceMonitorSelector:                 #line 3637
#   matchLabels:
#     prometheus: devops
```
### 3. Install chart
```
helm install monitoring \
prometheus-community/kube-prometheus-stack \
--values ./helm-promotheus/prom-values.yaml \
--version 65.2.0 \
--namespace monitoring \
--create-namespace
```
