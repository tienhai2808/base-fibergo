package database

import (
	"be-fiber/config"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	instance *MongoDB
	once     sync.Once
)

type Config struct {
	URI      string
	Database string
	Timeout  time.Duration
	PoolSize uint64
}

func ConnectMongoDB(cfg *config.MongoDBConfig) *MongoDB {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().
			ApplyURI(cfg.URI).
			SetMaxPoolSize(100)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Không thể kết nối MongoDB: %v", err)
		}

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			log.Fatalf("Không thể ping đến MongoDB: %v", err)
		}

		instance = &MongoDB{
			Client:   client,
			Database: client.Database(cfg.Database),
		}

		log.Println("Kết nối MongoDB thành công!")
	})

	return instance
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}

func (m *MongoDB) Close() {
	if m.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		if err := m.Client.Disconnect(ctx); err != nil {
			log.Printf("Lỗi khi đóng kết nối MongoDB: %v", err)
		}
	}
}