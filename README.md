# CRUD service
Otus02 CRUD service

# postman collection
 /postman_collection

# image build
docker build --progress plain . -t humokobil/crud:v1 -f .\app\Dockerfile
 
docker build --progress plain . -t humokobil/crud_migration:v1 -f .\migration\Dockerfile


# kubectl 
 kubectl apply \
 -f ./k8s/namespace.yaml \
 -f ./k8s/postgres/pv.yaml \
 -f ./k8s/postgres/pvc.yaml \
 -f ./k8s/postgres/configmap.yaml \
 -f ./k8s/postgres/secrets.yaml \
 -f ./k8s/postgres/deployment.yaml \
 -f ./k8s/postgres/service.yaml \
 -f ./k8s/migration/job.yaml \
 -f ./k8s/api/configmap.yaml \
 -f ./k8s/api/deployment.yaml \
 -f ./k8s/api/service.yaml \
 -f ./k8s/api/ingress.yaml


 # helm
 helm install app helm/crud
