package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Prayag2003/connecting_go_and_mongo/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the env file", err)
	}
}

const dbName = "netflix"

// NOTE: Collection
const colName = "watchlist"

var connectionString = os.Getenv("MONGO_URL")

// IMPORTANT:
var collection *mongo.Collection

// connect with MongoDB
// init method runs at the very first time only
func init() {

	// Client options -> responsible for getting a connection
	clientOption := options.Client().ApplyURI(connectionString)

	// connect req
	// Context has to passed always whenever one is doing database operations
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection is successfully established with MongoDB")

	// need a reference of the variables
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection reference is ready")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// IMPORTANT:
// 2 set of methods, one from MongoDB
// -- MongoDB Helper Methods: Receive the data into DB
// -- Basic Methods:  Bring up all data from URL Params or req body, will do auth and they will call mongoDB helpers.

// MONGODB Helpers
// helper methods to be written in lowercase

// Insert 1 movie into DB
func insertOneMovie(movie model.Netflix) {

	data, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with ID :", data.InsertedID)
}

// Update 1 movie
func updateOneMovie(movieId string) {
	// converting to _id from string which is acceptible by MongoDB
	_id, err := primitive.ObjectIDFromHex(movieId)
	handleErr(err)

	// everything inside MongoDB is BSON
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"watched": true}}

	// updating
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count :", res.ModifiedCount)

}

// Delete 1 movie
func deleteOneMovie(movieId string) {
	_id, err := primitive.ObjectIDFromHex(movieId)
	handleErr(err)

	filter := bson.M{"_id": _id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	handleErr(err)

	fmt.Println("Movie deleted with delete count:", deleteCount)
}

// Delete many movie
func deleteAllMovies() int64 {
	// filter := bson.D{{}}// empty key value pairs
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	handleErr(err)

	fmt.Println("Totals movie deleted:", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}

// get all movies from DB
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	handleErr(err)

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie)
		handleErr(err)

		// pushing into array
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// IMPORTANT:
// Actual Controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteCount := deleteAllMovies()
	json.NewEncoder(w).Encode(deleteCount)
}
