


TEMP="$(dirname "$0")"
BASE_DIR=$(echo "$(basename ${TEMP})")
CURRENT_DIR=${PWD}/${BASE_DIR}


kubectl delete -f "${CURRENT_DIR}"/auth-service.yaml
kubectl delete -f "${CURRENT_DIR}"/s3-service.yaml
#kubectl delete -f "${CURRENT_DIR}"/mysql.yaml
kubectl delete -f "${CURRENT_DIR}"/configmap.yaml
kubectl delete -f "${CURRENT_DIR}"/secret.yaml
kubectl delete -f "${CURRENT_DIR}"/s3-secret.yaml
#kubectl delete -f "${CURRENT_DIR}"/mysql-persistent-storage-claim.yaml
#kubectl delete -f "${CURRENT_DIR}"/mysql-persistent-storage.yaml
kubectl get pod
