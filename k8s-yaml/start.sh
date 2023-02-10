

TEMP="$(dirname "$0")"
BASE_DIR=$(echo "$(basename ${TEMP})")
CURRENT_DIR=${PWD}/${BASE_DIR}

kubectl apply -f "${CURRENT_DIR}/"configmap.yaml
kubectl apply -f "${CURRENT_DIR}/"secret.yaml
kubectl apply -f "${CURRENT_DIR}/"s3-secret.yaml
#kubectl apply -f "${CURRENT_DIR}/"mysql-persistent-storage.yaml
#kubectl apply -f "${CURRENT_DIR}/"mysql-persistent-storage-claim.yaml
#kubectl apply -f "${CURRENT_DIR}/"mysql.yaml
kubectl apply -f "${CURRENT_DIR}/"auth-service.yaml
kubectl apply -f "${CURRENT_DIR}/"s3-service.yaml
kubectl apply -f "${CURRENT_DIR}/nginx-ingress/"ingress-nginx-controller.yaml
kubectl apply -f "${CURRENT_DIR}/nginx-ingress/"ingress.yaml

