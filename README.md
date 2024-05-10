### How to Use
```
git clone git@github.com:shn27/bookstore-sample-controller.git
cd bookstore-sample-controller
go mod tidy & go mod vendor
./hack/update-codegen.sh
kc create -f menifest/crd-status-subresource.yaml
go build .
./sample-controller
kc create -f menifest/example-bookstore.yaml 
```