//go:build integration

package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories"
	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GormOrdersRepositorySuite struct {
	suite.Suite
	client  *mongo.Client
	repo    *repositories.GormOrdersRepository
	cleanup func()
	order   models.MongoOrder
	orderID string
}

func setupMongoDBContainer(ctx context.Context) (*mongo.Client, func(), error) {
	mongoContainer, err := mongodb.Run(ctx, "mongo:latest")
	if err != nil {
		return nil, nil, err
	}

	// Get the connection string
	connStr, err := mongoContainer.ConnectionString(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Connect to the MongoDB container
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, nil, err
	}

	// Return the client and a cleanup function
	cleanup := func() {
		client.Disconnect(ctx)
		mongoContainer.Terminate(ctx)
	}

	return client, cleanup, nil
}

func (suite *GormOrdersRepositorySuite) SetupSuite() {
	ctx := context.Background()

	client, cleanup, err := setupMongoDBContainer(ctx)
	assert.NoError(suite.T(), err)

	suite.client = client
	suite.cleanup = cleanup

	db := client.Database("testdb")
	suite.repo = repositories.NewGormOrdersRepository(db)

	suite.orderID = "1234"
	now := time.Now()
	suite.order = models.MongoOrder{
		ID:         primitive.NewObjectID(),
		Status:     "RECEIVED",
		ReceivedAt: &now,
		Number:     suite.orderID,
	}
}

func (suite *GormOrdersRepositorySuite) TearDownSuite() {
	suite.cleanup()
}

func (suite *GormOrdersRepositorySuite) SetupTest() {
	_, err := suite.repo.Save(context.Background(), suite.order)
	suite.Require().NoError(err)
}

func (suite *GormOrdersRepositorySuite) TearDownTest() {
	filter := bson.M{"number": suite.orderID}
	_, err := suite.client.Database("testdb").
		Collection("orders").
		DeleteOne(context.Background(), filter)
	suite.Require().NoError(err)
}

func (suite *GormOrdersRepositorySuite) TestGetByID() {
	// given
	ctx := context.Background()

	// when
	retrievedOrder, err := suite.repo.GetByID(ctx, suite.orderID)

	// then
	suite.Require().NoError(err)
	suite.Equal(suite.orderID, retrievedOrder.Number)
}

// TestUpdate tests the Update method
func (suite *GormOrdersRepositorySuite) TestUpdate() {
	ctx := context.Background()

	// Update the order
	suite.order.Status = "PREPARATION"
	updatedOrder, err := suite.repo.Update(ctx, suite.order)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "PREPARATION", updatedOrder.Status)

	// Retrieve the updated order
	retrievedOrder, err := suite.repo.GetByID(ctx, suite.orderID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "PREPARATION", retrievedOrder.Status)
}

// TestListAll tests the ListAll method
func (suite *GormOrdersRepositorySuite) TestListAll() {
	ctx := context.Background()

	// Create a second order
	now := time.Now()
	order2 := models.MongoOrder{
		ID:         primitive.NewObjectID(),
		Status:     "READY",
		ReceivedAt: &now,
	}
	_, err := suite.repo.Save(ctx, order2)
	assert.NoError(suite.T(), err)

	pagination := entities.Pagination{Page: 1, Size: 10}
	paginatedOrders, err := suite.repo.ListAll(ctx, pagination)

	suite.Require().NoError(err)
	suite.Len(paginatedOrders.Content, 2)
	suite.Equal(int64(2), paginatedOrders.Pagination.TotalElements)
}

func TestGormOrdersRepositorySuite(t *testing.T) {
	suite.Run(t, new(GormOrdersRepositorySuite))
}
