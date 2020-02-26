# Containers & Kubernetes
![](images/containers-everywhere.jpg)<!-- .element style="border: 0; background: None; box-shadow: None height="80%" width="80%" text-align: center" -->


# Container Security - The Good
- Microservices 
- Pods need to run on a platform that you can lock down
- Security updates and patches are easier to roll out
- limited open ports

Note:
- microservices mean we have a much smaller attack surface as the pod should contain only the libraries and binaries needed for the service
- Updates are easier as you can update your base image, re-build the image, and do a rolling deploy so outages shouldn't be a thing
- We can stop pods running as root, and create security restrictions for what pods can and can't do


# Container Security - The Bad
- So many Microservices 
- Complexity 
- Visibility

Note:
- Gets complicated with so many pods, those 10vm's might be 100 pods now (what ports running, how many replicas, are they all logging?, how do I know if it gets compromised)
- devs may just grab images from open places or use base image that is unkown