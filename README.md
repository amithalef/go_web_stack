# go_web_stack


To run application use:

        docker-compose build
        docker-compose up

Mocks:
   Sample code to generate mocks
   
        mockgen -destination=/Users/a.nair/myprojects/my_workspace/src/github.com/amithnair91/go_web_stack/go_web_starter/app/commands/mocks/mock_storage.go -source=/Users/a.nair/myprojects/my_workspace/src/github.com/amithnair91/go_web_stack/go_web_starter/app/commands/storage/storage.go
        

1. Services in GO (Hexagonal/Usecase Driven/TDD)
2. RPC
3. Kafka (message broker)
4. Load Testing/Perf Test (vegeta/gatling)
5. Kubernetes Yaml
6. Helm Chart
7. Minikube setup
8. Istio
9. Knative
10. Distributed Tracing
11. Circuit breakers/ Fault Tolerance (at infrastructure layer)
12. Debezium
13. Propogate some additional data in the network layer(Feature toggle information)


References:
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--data-aggregation-pipeline
