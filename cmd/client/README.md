# PQueue Client

### pqctl 

Pronounce: Peak Cuddle

The main client CLI tool for interfacing with the PQueue service

```bash
pqctl --namespace=nova get <queue>
pqctl --namespace=nova list
pqctl --namespace=nova set <queue> <message>
pqctl --namespace=nova delete <queue>
pqctl --namespace=nova empty <queue>
alias pn="pqctl --namespace=nova"
pn get <queue> 
pn list 
pn set <queue> <message>
```

### Creating namespace queues in Kubernetes

```bash
$ kubectl create namespace nova
$ pqctl create <channel> --namespace nova
$ kubectl get pods --namespace nova
  NAME                     STATUS
  pqueue-<channel>-01      ready
$ pqctl set <channel> --namespace=nova <message>
$ pqctl get <channel> --namespace=nova 
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

