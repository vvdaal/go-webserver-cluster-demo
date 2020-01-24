# go-webserver-cluster-demo
A simple go webserver that prints some basic information about the environment it's running on and can have a environment variable configured. 

This could be used with kubernetes deployments for example in a cluster or multiple clusters. I used this to test a multi-region anycasted solution in Kubernetes to ensure the routing was working as intended.

You can use the kubernetes examples for an example rollout. Please notice that I didn't provide an ingress on purpose, that's up to you.