# fsnotify
This is the experimental fsnotify code combined with kubernetes configmap. The goal is to test whether the update of the configmap
could be watched by fsnotify.

# steps
- create configmap1.yaml
```shell
kubectl create -f configmap1.yaml
```
- create fsnofity rc
```shell
kubectl create -f test.yaml
```
- update configmap
```shell
kubectl apply -f configmap2.yaml
```
- check logs

# caveat
the actual directory structure with configmap is as follows:
```shell
root@fsnotify-gihzt:/tmp# ls -arl
total 4
lrwxrwxrwx  1 root root   10 Apr 28 03:26 foo -> ..data/foo
lrwxrwxrwx  1 root root   31 Apr 28 03:27 ..data -> ..4984_28_04_11_27_36.107821517
drwx------  2 root root   60 Apr 28 03:27 ..4984_28_04_11_27_36.107821517
drwxr-xr-x 32 root root 4096 Apr 28 03:28 ..
drwxrwxrwt  3 root root  100 Apr 28 03:27 .
```
when configmap is updated, the old directory `..4984_28_04_11_27_36.107821517` will be removed and a new one will be created.
So we cannot simply watch the file `/tmp/foo`. Instead we should watch the whole directory.
