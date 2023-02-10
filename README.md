# Go File Serve

This repo has an example just to reproduce a problem where serving a video file from `Go` through `kubectl port-forward` causes the port forward itself to fail.

# To repdroduce

(optionally: build your own docker image or use the one I'm sharing at `miniodev/gofileserve`)

## 1. Deploy this app as a deplyment

```shell
kubectl apply -f kubernetes/deployment.yaml
```

## 2. Port Forward to this service
```shell
kubectl port-forward svc/gofileserve 4001:4000
```

## 3. On another terminal, try to get `Spinning-Cat.mp4`

```shell
curl http://localhost:4001/assets/media/Spinning-Cat.mp4
```

You'll see how the port-forward breaks

```shell
➜ kubectl port-forward svc/gofileserve 4001:4000
Forwarding from 127.0.0.1:4001 -> 4000
Forwarding from [::1]:4001 -> 4000
Handling connection for 4001
E0210 13:22:27.776615   61159 portforward.go:379] error copying from remote stream to local connection: readfrom tcp4 127.0.0.1:4001->127.0.0.1:52390: write tcp4 127.0.0.1:4001->127.0.0.1:52390: write: broken pipe
```