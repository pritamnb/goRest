package repository
import (
	"context"
	"log"
	"testing"
	"projectx/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	link := "mongodb+srv://pritambhalerao9284:LGhVtk8zFG40naGu@cluster0.xb7nxzz.mongodb.net/?retryWrites=true&w=majority&appName=cluster0"

	mongoTestClient, err := mongo.Connect(context.Background(),options.Client().ApplyURI(link))

	if err !=nil {
		log.Fatal("error while connecting mongodb", err)
	}

	log.Println("mongodb successfully connected. ")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Ping failed", err)
	}

	log.Println("Ping successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data	
	emp1 := uuid.New().String()
	// emp2 = uuid.New().String()

	// Connect to the collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}
	
	// Insert Employee 1 data

	t.Run("Insert Employee 1", func(t *testing.T){
		emp :=model.Employee{
			Name: "Tony Start",
			Department: "Physics",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}

		t.Log("Insert 1 successful", result)
	})

}




























