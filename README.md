### bookstore-sample-controller


A Controller written in kubernetes sample-controller style which watches a custom resource named Bookstore. A resource create a deployment and a NodePort service. The container image is a simple bookstore api server. The NodePort service listen to request on host port 30000.

### How to Use

<h3>Clone repo and move to directory</h3>

```
git clone git@github.com:shn27/bookstore-sample-controller.git
cd bookstore-sample-controller
```
<h3>Create a cluster using Kind </h3>

` kind create cluster --config=clusterconfig.yaml` 

<h3> Create bookstores.shn.com Custom Resource Definition (CRD)</h3>

` kc create -f menifest/crd-status-subresource.yaml `

<h3>Build and Run the controller</h3>

```
go build .
./sample-controller -kubeconfig=clusterconfig.yaml
```
<h3>Create Custom Resource</h3>

`kc create -f menifest/example-bookstore.yaml `

<h3> Port Forward </h3>

`kc port-forward service/bookstore 3000`

<h3> Test </h3>
Use <i>POSTMAN</i> <br>
<i>USERNAME</i> : admin <br>
<i>PASSWORD</i> : 1234<br>
https://github.com/shn27/BookStoreApi-Go <br>
