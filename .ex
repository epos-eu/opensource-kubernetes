apiVersion: rabbitmq.com/v1beta1
kind: RabbitmqCluster
metadata:
  labels:
    app: rabbitmq
    version: 
  name: netrabbit
spec:
  replicas: 1
  resources:
    requests:
      cpu: 500m
      memory: 1Gi
    limits:
      cpu: 2
      memory: 2Gi
  persistence:
    storage: 20Gi
  override:
    statefulSet:
      spec:
        template:
          spec:
            containers:
            - name: rabbitmq
              readinessProbe:
                tcpSocket:
                  port: amqp
                initialDelaySeconds: 10
                periodSeconds: 10
                timeoutSeconds: 5
                successThreshold: 1
                failureThreshold: 3
  rabbitmq:
    additionalConfig: |
        log.console.level = info
        channel_max = 1700
        default_user =  
        default_pass = 
        default_vhost = 
        default_user_tags.administrator = true
    additionalPlugins:
      - rabbitmq_top
      - rabbitmq_shovel
      - rabbitmq_peer_discovery_k8s
      - rabbitmq_management
      - rabbitmq_prometheus
      - rabbitmq_tracing
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: rabbitmq-ingress
  labels:
    app: rabbitmq
    version: 
  annotations:
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/rewrite-target: /$1
spec:
  rules:
  - http:
      paths:
      - path: /deploy/pathrabbitmq/?(.*)
        pathType: Prefix
        backend:
          service:
            name: netrabbit
            port:
              number: 15672