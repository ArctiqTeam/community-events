while read p; do kubectl --kubeconfig=.kube/config rollout status deployment/$p ; done < $GITHUB_WORKSPACE/microservices.txt
