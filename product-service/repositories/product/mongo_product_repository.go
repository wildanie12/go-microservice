package product

import (
	"context"
	"net/http"

	"github.com/wildanie12/go-microservice/product-service/entities/domain"
	_domain "github.com/wildanie12/go-microservice/product-service/entities/domain"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoProductRepository main struct
type MongoProductRepository struct {
	mc *mongo.Collection
	collectionName string
	mongoCtx context.Context
}

// NewMongo act as a method constructor to make 
// MongoProductRepository struct object
func NewMongo(mongo *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		mc: mongo.Collection("product"),
		mongoCtx: context.TODO(),
	}
}

// FindAll returns product resources in an array of entity struct
// based on a given filter and sort query parameter 
func (repo MongoProductRepository) FindAll(filters []map[string]string) ([]_domain.Product, error) {
	cursor, err := repo.mc.Find(repo.mongoCtx, bson.D{})
	if err != nil {
		return []domain.Product{}, _web.Error{
			Code: http.StatusInternalServerError,
			DevMessage: "Mongo FindAll err: " + err.Error(),
			ProdMessage: "Cannot get product data",
		}
	}
	products := []_domain.Product{}
	cursor.All(repo.mongoCtx, &products)
	return products, nil
}

// Insert will insert the given product parameter to data source
// an error is returned when there was an error
func (repo MongoProductRepository) Insert(product domain.Product) (_domain.Product, error) {
	result, err := repo.mc.InsertOne(context.TODO(), product)
	product.ID = result.InsertedID.(primitive.ObjectID).String()
	if err != nil {
		return _domain.Product{}, err
	}
	return _domain.Product{}, nil
}