# Container Platform Security


# Platform Security - DO
- Choose a k8s distribution that is secure like Openshift & OKD
- Control access to the nodes, SSH should be off
- Limit access to resources to stop DOS
- On-boarding process to get services into the cluster
- Have a tooling cluster 
- Invest in some security tooling, monitor and alert on security violations

Note:
- Most cloud hosted K8s make it hard to access the shell. On-prem you shouldn't need too. OCP4 story
- Set requests and limits so pods or namespaces can't use all the resources 
- should be an ops team or automated process to set up the users, namespaces, and quotas so the cluster isn't a free for all, dis-allow authenticated users from creating projects
- Have a tooling cluster where containers that require elevated permissions can live, lock this down to a few ops members for security and ops  
- Openshift is more locked down out of the box and "enterprise ready" good option


# Platform Security - DON'T
- Presume the platform includes all security
- Permit privileged access instead of educating users
- Allow unverified images to run
- Run without a PodSecurityPolicy

Note:
- Kubernetes doesn't have a ton of tooling out of the box, look for some add on security tooling to help enhance the cluster. Cloud native, kubernetes ecosystem is huge and growing
- Will need to spend time educating users on the new platform and it may be harder for them to re-config there app to not run as root
- Probably shouldn't allow images from dockerhub to run on your prod cluster
- If your allowing containers to run as root or map volumes back to the host your really asking for trouble. Set admistion controller for and use pod security policies
