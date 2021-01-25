package customer

import (
	"api/db"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CustomerRepository struct
type CustomerRepository struct {
	resource   *db.Resource
	collection *mongo.Collection
}

//Repository interface
type Repository interface {
	Create(CustomerRegistorBody) (Customer, error)
	FindOne(query bson.M) (*Customer, error)
	Update(query bson.M, update bson.M) error
}

//NewCustomerRepository repo
func newCustomerRepository(resource *db.Resource) Repository {
	collection := resource.DB.Collection("customer")
	repository := &CustomerRepository{resource: resource, collection: collection}
	return repository
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

//Create customer
func (customer *CustomerRepository) Create(request CustomerRegistorBody) (Customer, error) {
	cus := Customer{
		Id:        primitive.NewObjectID(),
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		Lastname:  request.Lastname,
		Age:       request.Age,
	}
	ctx, cancel := initContext()
	defer cancel()
	_, err := customer.collection.InsertOne(ctx, cus)

	if err != nil {
		return Customer{}, err
	}

	return cus, nil
}

//FindOne get customer detail
func (customer *CustomerRepository) FindOne(query bson.M) (*Customer, error) {
	var cus Customer

	ctx, cancel := initContext()
	defer cancel()

	if err := customer.collection.FindOne(ctx, query).Decode(&cus); err != nil {
		return nil, err
	}
	return &cus, nil
}

//Update update customer detail
func (customer *CustomerRepository) Update(query bson.M, update bson.M) error {
	ctx, cancel := initContext()
	defer cancel()

	_, err := customer.collection.UpdateOne(
		ctx,
		query,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}
