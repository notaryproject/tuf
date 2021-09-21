# Demo

## Install
[branch of go-tuf](https://github.com/mnm678/go-tuf/tree/tuf-notary-demo)

## Push to registry
### create net-monitory image, upload to localhost
```
docker run -d -p 5000:5000 ghcr.io/oras-project/registry:v0.0.3-alpha
docker build -t localhost:5000/net-monitor:v1 https://github.com/wabbit-networks/net-monitor.git#main
docker push localhost:5000/net-monitor:v1
```

### sign with go-tuf library
```
tuf init
tuf add <digest>
tuf snapshot
tuf timestamp
tuf commit
```

### Upload delegated targets to net-monitor
```
oras push localhost:5000/net-monitor \
      --artifact-type tuf/example \
      --subject localhost:5000/net-monitor:v1 \
      repository/targets.json:application/json
```

### Upload top-level targets and root to tuf-metadata
TODO

## Pull from registry
### Find artifact references, fetch delegated targets
```
Set DIGEST  (oras discover localhost:5000/net-monitor:v1 -o json | jq -r .digest)
curl localhost:5000/oras/artifacts/v1/net-monitor/manifests/$DIGEST/referrers | jq
oras pull -a \
      localhost:5000/net-monitor@( \
      oras discover \
      -o json \
      --artifact-type=tuf/example \
      localhost:5000/net-monitor:v1 | jq -r .references[0].digest)
```

### Fetch root and top-level targets
TODO

### Verify with go-tuf
TODO
