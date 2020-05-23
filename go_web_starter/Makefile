fixfmt:
	@echo "Fixing format"
	@gofmt -w -l .

test:
	@echo "Running Tests"
	@go test -v ./...

mocks:
	@mockgen -destination=$(pwd)/usecase/mocks/mock_storage.go -source=$(pwd)//usecase/storage/storage.go

compile:
	@echo "Building Binaries"
	@GOOS=darwin GOARCH=amd64 go build -o=build/godownload-darwin app/main.go
	@GOOS=linux GOARCH=amd64 go build -o=build/godownload-linux app/main.go
	@GOOS=windows GOARCH=amd64 go build -o=build/godownload-windows app/main.go

clean:
	@echo "Removing existing builds"
	@-rm -rf build

docker:
	@docker build -t go-web-starter .

# install-hyperkit:
# 	@brew install docker-machine-driver-hyperkit
# 	@sudo chown root:wheel /usr/local/opt/docker-machine-driver-hyperkit/bin/docker-machine-driver-hyperkit
# 	@sudo chmod u+s /usr/local/opt/docker-machine-driver-hyperkit/bin/docker-machine-driver-hyperkit

minikube-start:
	@minikube start --memory=8192 --cpus=4   --kubernetes-version=v1.13.0   --vm-driver=hyperkit   --disk-size=10g   --extra-config=apiserver.enable-admission-plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"

install-istio:
	@kubectl apply --filename https://github.com/knative/serving/releases/download/v0.4.0/istio-crds.yaml && curl -L https://github.com/knative/serving/releases/download/v0.4.0/istio.yaml \
     | sed 's/LoadBalancer/NodePort/' \
     | kubectl apply --filename -

install-knative-serving:
	@curl -L https://github.com/knative/serving/releases/download/v0.3.0/serving.yaml \
     | sed 's/LoadBalancer/NodePort/' \
     | kubectl apply --filename -

install-knative-build:
	@kubectl apply --filename https://github.com/knative/build/releases/download/v0.3.0/release.yaml

# list-knative-services:
# 	@ INGRESSGATEWAY=istio-ingressgateway
# 	@echo $(minikube ip):$(kubectl get svc $INGRESSGATEWAY --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')

# actual command := echo $(minikube ip):$(kubectl get svc istio-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')

# knative := https://knative.dev/docs/install/any-kubernetes-cluster/
# https://knative.dev/v0.12-docs/install/knative-with-minikube/