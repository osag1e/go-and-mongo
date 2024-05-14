package db

import (
	"context"
	"os"

	"github.com/osag1e/go-and-mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MongoDBNameEnvName = "MONGO_DB_NAME"
)

type DataStore struct {
	Movie MovieDataStore
}

const movieCollection = "movies"

type MovieDataStore interface {
	InsertMovie(context.Context, *model.MovieTicket) (*model.MovieTicket, error)
	GetMovies(context.Context) ([]*model.MovieTicket, error)
	DeleteMovie(context.Context, string) error
}

type MongoMovieDataStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoMovieDataStore(client *mongo.Client) *MongoMovieDataStore {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoMovieDataStore{
		client: client,
		coll:   client.Database(dbname).Collection(movieCollection),
	}
}

func (s *MongoMovieDataStore) InsertMovie(ctx context.Context, ticket *model.MovieTicket) (*model.MovieTicket, error) {
	res, err := s.coll.InsertOne(ctx, ticket)
	if err != nil {
		return nil, err
	}
	ticket.ID = res.InsertedID.(primitive.ObjectID)
	return ticket, nil
}

func (s *MongoMovieDataStore) GetMovies(ctx context.Context) ([]*model.MovieTicket, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var tickets []*model.MovieTicket
	if err := cur.All(ctx, &tickets); err != nil {
		return []*model.MovieTicket{}, nil
	}
	return tickets, nil
}

func (s *MongoMovieDataStore) DeleteMovie(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
