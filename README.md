1. curl -sSL https://github.com/operator-framework/operator-sdk/releases/download/v1.16.0/operator-sdk_v1.16.0_darwin_amd64.tar.gz | tar xz -C /usr/local/bin
2. operator-sdk init --domain brbarmex --repo github.com/brbarmex/rollout-lab
3. operator-sdk create api --group rollout --version v1 --kind RolloutApp --resource --controller
4. make bundle IMG="brbarmex/operator:v0.0.1"
5. make bundle-build bundle-push BUNDLE_IMG="brbarmex/operator:v0.0.1"
6. export WATCH_NAMESPACE=default
7. operator-sdk run bundle docker.io/brbarmex/operator:v0.0.1


debug

1. go get -u github.com/go-delve/delve/cmd/dlv



