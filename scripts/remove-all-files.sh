#!/usr/bin/env bash

set -eu

if [ -z "${1-}" ]; then 
    echo "Use: $0 [pvc-name]"
    exit 1
fi

kubectl apply -f - <<EOF
apiVersion: batch/v1
kind: Job
metadata:
  name: cleanup-files
spec:
  template:
    metadata:
      creationTimestamp: null
    spec:
      volumes:
      - name: backups
        persistentVolumeClaim:
          claimName: $1
      containers:
      - image: alpine
        name: cleanup-files
        command:
        - sh
        - -c
        - "rm -rf /backups/*"
        volumeMounts:
        - name: backups
          mountPath: /backups
        resources: {}
      restartPolicy: Never
EOF

kubectl wait --for=condition=complete job/cleanup-files
kubectl delete job cleanup-files
