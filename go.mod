apiVersion: v1
kind: Pod
metadata:
  name: lneu-pod
spec:
  containers:
  - name: lneu-app
    image: lneu-img
    volumeMounts:
    - name: vol
      mountPath: /app/config

  - name: cnf-sidecar
    image: sidecar-img
    volumeMounts:
    - name: cnf-vol
      mountPath: /config
    - name: kubeconfig
      mountPath: /kubeconfig

  volumes:
  - name: cnf-vol
    emptyDir: {}

  - name: kubeconfig
    configMap:
      name: kubeconfig
