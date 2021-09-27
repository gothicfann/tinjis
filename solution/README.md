# Bonus Questions

This solution is for demo purposes. I have not used any advanced topics here.

1. **How would a new deployment look like for these services? What kind of tools would you use?**  
    For production deployment I would use some CI/CD solution it to make it automate.  
    Also would be better to use [Helm](https://helm.sh/) instead of raw k8s manifests or write custom operator in [operatorframework](https://operatorframework.io/.  
    Its wise also to define set of resource requests/limits for each microservice, and configure [HPA's](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/).  
    Except that I would use [NetworkPolicies](https://kubernetes.io/docs/concepts/services-networking/network-policies/), and enhance [PodSecurityContext](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) to improve security of running environment.  
    I would not use NodePort service for production deployment. Better to setup some kind of Ingress Controller ([Nginx](https://kubernetes.github.io/ingress-nginx/), [Traefik](https://traefik.io/), [Istio](https://istio.io/)) and configure Ingress resources. Ideally I would use **Istio** as Service Mesh with **Nginx** as Ingress in front of Istio [Gateway](https://istio.io/latest/docs/reference/config/networking/gateway/).
1. **If a developers needs to push updates to just one of the services, how can we grant that permission without allowing the same developer to deploy any other services running in K8s?**  
    To solve this we may need to generate client certificates for developers (or use any other [supported](https://kubernetes.io/docs/reference/access-authn-authz/authentication/) auth mechanisms) to authenticate against kube-apiserver and setup Roles, ClusterRoles, RoleBindings, ClusterRoleBindings to grant required permissions in Kubernetes cluster.
1. **How do we prevent other services running in the cluster to talk to your service. Only Antaeus should be able to do it.**  
    For this we can use NetworkPolicies. For this we need CNI plugin which supports NetworkPolicies ([calico](https://www.tigera.io/project-calico/), [weave](https://www.weave.works/docs/net/latest/overview/), [cilium](https://cilium.io/), etc.)  
    Example manifest for payment service:
    ```yaml
    apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
    name: payment
    spec:
      podSelector:
        matchLabels:
          app: payment
      policyTypes:
      - Ingress
      ingress:
      - from:
          - podSelector:
              matchLabels:
                app: antaeus
          ports:
          - protocol: TCP
            port: 8080
    ```
    In this case only antaeus pods will be able to access payment on port 8080.