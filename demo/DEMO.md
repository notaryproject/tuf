# Demo

## Install
[branch of go-tuf](https://github.com/mnm678/go-tuf/tree/tuf-notary-demo)
[ORAS client](https://github.com/oras-project/oras/releases)

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
cat repository/targets.json | jq
```

### Upload delegated targets to net-monitor
```
oras push localhost:5000/net-monitor \
      --artifact-type tuf/example \
      --subject localhost:5000/net-monitor:v1 \
      repository/targets.json:application/json
```

### Upload top-level targets and root to tuf-metadata with root as an artifact that targets refers to
```
oras push localhost:5000/tuf-repository:root repository/root.json:application/json
oras push localhost:5000/tuf-repository:targets \
			--artifact-type tuf/targets \
			--subject localhost:5000/tuf-repository:root \
			repository/targets.json:application/json

```

## Pull from registry
### Find artifact references, fetch delegated targets
```
set DIGEST  (oras discover localhost:5000/net-monitor:v1 -o json | jq -r .digest)
curl localhost:5000/oras/artifacts/v1/net-monitor/manifests/$DIGEST/referrers | jq
oras pull -a \
      localhost:5000/net-monitor@( \
      oras discover \
      -o json \
      --artifact-type=tuf/example \
      localhost:5000/net-monitor:v1 | jq -r .references[0].digest)
```

### Fetch root and top-level targets
```
oras pull -a localhost:5000/tuf-repository:root
Set DIGEST  (oras discover localhost:5000/tuf-repository:root -o json | jq -r .digest)
curl localhost:5000/oras/artifacts/v1/tuf-repository/manifests/$DIGEST/referrers | jq
oras pull -a localhost:5000/tuf-repository@( \
         oras discover \
         -o json \
         --artifact-type=tuf/targets \
         localhost:5000/tuf-repository:root | jq -r .references[0].digest)
```

### Verify with go-tuf
TODO
