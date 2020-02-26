# Demo

Note:
- Look at Aqua set up and the security features it brings
- Show podsecurity policy 
  - kubectl get psp default-psp -o yaml
  - kubectl create -f root-pod.yaml --kubeconfig=k8s-bob-user/bob-k8s-config
  - kubectl describe pod nginx-privileged -n preso

# Aquasec Architecture 
![](images/aqua-architecture.jpg)<!-- .element style="border: 0; background: None; box-shadow: None height="100%" width="100%" text-align: center" -->


# Free Resources 
- Kube-Scan - Scan your workloads and get a risk score :: https://github.com/octarinesec/kube-scan
- Kube Bench - Checks nodes against CIS k8s benchmark :: https://github.com/aquasecurity/kube-bench 
- Kubeaudit - Audit various parts of k8s cluster, provides some fixes :: https://github.com/Shopify
- Kube-hunter - Looks for known security issues :: https://github.com/aquasecurity/kube-hunter