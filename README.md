# k8s-scheduling-framework-small-sample
small sample implementation of kubernetes scheduling-framework

# build
```
go build sample-scheduler.go
```

# run
* execute the following command on k8s master
```
sudo ./sample-scheduler --authentication-kubeconfig=/etc/kubernetes/scheduler.conf --authorization-kubeconfig=/etc/kubernetes/scheduler.conf --config=scheduler-config.yaml --secure-port=10260
```

# test
* run this command
```
kubectl create -f sample-pod.yml
```

* then you will see the following stdout from sample-scheduler
```
I1225 07:09:30.512923   10010 sample-scheduler.go:27] pre filter called for pod sample-scheduler-pod
I1225 07:09:30.516019   10010 sample-scheduler.go:34] pod sample-scheduler-pod is binded to node-hoge-hoge
```
