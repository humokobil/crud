# CRUD
Otus02 CRUD service

# Run
> go run .\cmd\app\main.go --conf-file=conf/config.yaml
> go run .\cmd\migration\main.go --conf-file=conf/config.yaml

# Image build
> docker build --progress plain . -t humokobil/crud:v1 -f .\cmd\app\Dockerfile 
> docker build --progress plain . -t humokobil/crud_migration:v1 -f .\cmd\migration\Dockerfile 

# Image pull
> docker pull humokobil/crud:v1
> docker pull humokobil/crud_migration:v1

# Запуск через kubectl
 kubectl apply -f ./k8s/namespace.yaml -f ./k8s/postgres/pv.yaml -f ./k8s/postgres/pvc.yaml -f ./k8s/postgres/configmap.yaml  -f ./k8s/postgres/secrets.yaml -f ./k8s/postgres/deployment.yaml  -f ./k8s/postgres/service.yaml -f ./k8s/migration/job.yaml -f ./k8s/api/configmap.yaml  -f ./k8s/api/deployment.yaml -f ./k8s/api/service.yaml -f ./k8s/api/ingress.yaml -f ./k8s/prometheus/servicemonitor.yaml


 # Helm chart
 > helm install app helm/crud
 
 # Prometheus
 > helm install -n monitoring prom prometheus-community/kube-prometheus-stack -f .\helm\prometheus\prometheus.yaml --atomic --set prometheusOperator.admissionWebhooks.enabled=false --set prometheusOperator.admissionWebhooks.patch.enabled=false --set prometheusOperator.tlsProxy.enabled=false
 
 > kubectl -n monitoring port-forward service/prom-kube-prometheus-stack-prometheus 9090
 > kubectl -n monitoring port-forward service/prom-grafana 9000:80


 # Ingress
 ## Nginx-ingress-controller with prometheus metrics
 > helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
 > helm install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --set controller.metrics.enabled=true --set-string controller.podAnnotations."prometheus\.io/scrape"="true" --set-string controller.podAnnotations."prometheus\.io/port"="10254"

# Postgres exporter
> url : https://artifacthub.io/packages/helm/prometheus-community/prometheus-postgres-exporter

> helm install prom-pg-exporter prometheus-community/prometheus-postgres-exporter -n monitoring

 # Load testing
 k6 run --vus 10 --duration 10m .\postman_collection\k6-script.js --rps 10