kind: DaemonSet
spec:
  template:
    spec:
      serviceAccountName: kubeip-service-account
      tolerations:
        - key: startup-taint.cluster-autoscaler.kubernetes.io/kubeip-not-ready
          operator: Exists
          effect: NoSchedule
      securityContext:
        runAsNonRoot: true
        runAsUser: 1001
        runAsGroup: 1001
        fsGroup: 1001
      containers:
        - name: kubeip
          image: doitintl/kubeip-agent
          env:
            - name: TAINT_KEY
              value: startup-taint.cluster-autoscaler.kubernetes.io/kubeip-not-ready
          securityContext:
            privileged: false
            allowPrivilegeEscalation: false
            capabilities:
              drop:
                - ALL
            readOnlyRootFilesystem: true