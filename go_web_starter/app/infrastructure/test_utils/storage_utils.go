package test_utils

import (
	"context"
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const DatabaseName = "testing"

type MongoTestContainer struct {
	IP      string
	Port    string
	mongo   testcontainers.Container
	context context.Context
}

func (mtc *MongoTestContainer) Start(t *testing.T) (mongo testcontainers.Container) {
	port := "27017"
	ctx := context.Background()
	mtc.context = ctx
	req := testcontainers.ContainerRequest{
		Image:        "mongo:4.2.6-bionic",
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", port)},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)),
	}
	mongo, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	mtc.mongo = mongo
	if err != nil {
		t.Error(err)
	}
	ip, err := mongo.Host(ctx)
	if err != nil {
		t.Error(err)
	}
	mappedPort, err := mongo.MappedPort(ctx, nat.Port(port))
	if err != nil {
		t.Error(err)
	}
	mtc.Port = mappedPort.Port()
	mtc.IP = ip
	return
}

func (mtc *MongoTestContainer) Stop() {
	mtc.mongo.Terminate(mtc.context)
}

func StartMongoDbForTest(t *testing.T) (MongoTestContainer, *mongo.Database) {
	container := MongoTestContainer{}
	container.Start(t)
	database, error := mongo_storage.Connect(container.IP, container.Port, DatabaseName)
	if error != nil {
		fmt.Println("failed to Start Mongo Db for Test")
		panic(t)
	}
	return container, database
}

func StopMongoDbContainer(container MongoTestContainer) {
	container.Stop()
}
