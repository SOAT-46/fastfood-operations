package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/SOAT-46/fastfood-operations/internal/orders/adapters/repositories/models"
	"github.com/SOAT-46/fastfood-operations/internal/orders/domain/exceptions"
	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GormOrdersRepository struct {
	client *mongo.Collection
}

func NewGormOrdersRepository(client *mongo.Database) *GormOrdersRepository {
	return &GormOrdersRepository{
		client: client.Collection("orders"),
	}
}

func (itself *GormOrdersRepository) ListAll(
	ctx context.Context, pagination entities.Pagination) (entities.PaginatedEntity[models.MongoOrder], error) {
	var orders []models.MongoOrder
	filter := bson.M{"status": bson.M{"$in": []string{"RECEIVED", "PREPARATION", "READY"}}}

	totalElements, err := itself.client.CountDocuments(ctx, filter)
	if err != nil {
		return entities.PaginatedEntity[models.MongoOrder]{},
			fmt.Errorf("can't count documents in orders repository: %w", err)
	}
	pagination.TotalElements = totalElements

	// Define sorting options: Sort by status (descending) and received_at (ascending)
	findOptions := options.Find().
		SetSort(bson.D{
			{Key: "status", Value: -1},     // DESCENDING order
			{Key: "received_at", Value: 1}, // ASCENDING order
		}).
		SetSkip(int64(pagination.Offset())). // Skip the appropriate number of records
		SetLimit(int64(pagination.Size))     // Limit results per page

	cursor, err := itself.client.Find(ctx, filter, findOptions)
	if err != nil {
		return entities.NewEmptyPaginatedEntity[models.MongoOrder](pagination),
			fmt.Errorf("can't fetch paginated orders: %w", err)
	}
	defer cursor.Close(ctx)

	if cursorErr := cursor.All(ctx, &orders); cursorErr != nil {
		return entities.NewEmptyPaginatedEntity[models.MongoOrder](pagination),
			fmt.Errorf("failed to decode MongoDB results: %w", err)
	}

	return entities.NewPaginatedEntity(orders, pagination), nil
}

func (itself *GormOrdersRepository) GetByID(ctx context.Context, id string) (*models.MongoOrder, error) {
	order := &models.MongoOrder{}
	filter := bson.M{"number": id}

	err := itself.client.FindOne(ctx, filter).Decode(order)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, exceptions.ErrOrderNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("can't GetByID in orders repository. Reason: %w", err)
	}
	return order, nil
}

func (itself *GormOrdersRepository) Save(ctx context.Context, order models.MongoOrder) (*models.MongoOrder, error) {
	_, err := itself.client.InsertOne(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("can't Save in orders repository. Reason: %w", err)
	}
	return &order, nil
}

func (itself *GormOrdersRepository) Update(ctx context.Context, order models.MongoOrder) (*models.MongoOrder, error) {
	filter := bson.M{"_id": order.ID}
	update := bson.M{"$set": order}

	result, err := itself.client.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("can't update order in repository. Reason: %w", err)
	}
	if result.MatchedCount == 0 {
		return nil, exceptions.ErrOrderNotFound
	}
	return &order, nil
}
