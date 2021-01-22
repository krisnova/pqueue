# PQueue Client

### pqctl 

Pronounce: Peak Cuddle

The main client CLI tool for interfacing with the PQueue service

```bash
pqctl --namespace=<namespace> get <queue>
pqctl --namespace=<namespace> list
pqctl --namespace=<namespace> set <queue> <message>
pqctl --namespace=<namespace> delete <queue>
pqctl --namespace=<namespace> empty <queue>
alias pn="pqctl --namespace=<namespace>"
pn get <queue> 
pn list 
pn set <queue> <message>
```

### Creating namespace queues in Kubernetes

```bash
$ kubectl create namespace <namespace>
$ pqctl create <channel> --namespace=<namespace>
$ kubectl get pods --namespace=<namespace>
  NAME                     STATUS
  pqueue-<channel>-01      ready
$ pqctl set <channel> --namespace=<namespace> <message>
$ pqctl get <channel> --namespace=<namespace>
  '
  {
      "message": "great message"
  }
  '
$ alias pn="pqctl --namespace=<namespace>"
$ pn list
  <channel1>
  <channel2>
  <channel3>
  ...
  <channelN> 
```

