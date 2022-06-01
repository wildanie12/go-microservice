package product

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wildanie12/go-microservice/product-service/entities/domain"
	_domain "github.com/wildanie12/go-microservice/product-service/entities/domain"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	"github.com/wildanie12/go-microservice/product-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	mc *mongo.Collection
	collectionName string
	mongoCtx context.Context
}

func NewMongo(mongo *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		mc: mongo.Collection("product"),
		mongoCtx: context.TODO(),
	}
}

func (repo MongoProductRepository) FindAll() ([]_domain.Product, error) {
	cursor, err := repo.mc.Find(repo.mongoCtx, bson.D{})
	if err != nil {
		return []domain.Product{}, _web.WebError{
			Code: http.StatusInternalServerError,
			DevMessage: "Mongo FindAll err: " + err.Error(),
			ProdMessage: "Cannot get product data",
		}
	}

	for cursor.Next(context.TODO()) {
		var result bson.D
		err := cursor.Decode(&result)
		if err != nil {
			return []domain.Product{}, _web.WebError{
				Code: http.StatusInternalServerError,
				DevMessage: "Mongo FindAll cussor err: " + err.Error(),
				ProdMessage: "Cannot get product data",
			}
		}
		fmt.Println(utils.JSONEncode(result))
	}
	return []_domain.Product{}, nil
}

func (repo MongoProductRepository) Insert(product domain.Product) (_domain.Product, error) {
	result, err := repo.mc.InsertOne(context.TODO(), product)
	product.ID = result.InsertedID.(primitive.ObjectID).String()
	if err != nil {
		return _domain.Product{}, err
	}
	return _domain.Product{}, nil
}