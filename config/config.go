package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		Stage string
		Name  string
		Url   string
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
		ApiDuration      int64
	}

	Kafka struct {
		Url       string
		ApiKey    string
		ApiSecret string
	}

	Grpc struct {
		AuthUrl      string
		ItemUrl      string
		PlayerUrl    string
		InventoryUrl string
		PaymentUrl   string
	}

	Paginate struct {
		ItemNextPageBaseUrl  string
		InventoryPageBaseUrl string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env fail")
	}

	return Config{
		App: App{
			Stage: os.Getenv("APP_STAGE"),
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading access duration failed!")
				}
				return result
			}(),
			RefreshDuration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading refresh duration failed!")
				}
				return result
			}(),
		},
		Kafka: Kafka{
			Url:       os.Getenv("KAFKA_URL"),
			ApiKey:    os.Getenv("KAFKA_API_KEY"),
			ApiSecret: os.Getenv("KAFKA_API_SECRET"),
		},

		Grpc: Grpc{
			AuthUrl:      os.Getenv("GRPC_AUTH_URL"),
			ItemUrl:      os.Getenv("GRPC_ITEM_URL"),
			PlayerUrl:    os.Getenv("GRPC_PLAYER_URL"),
			InventoryUrl: os.Getenv("GRPC_INVENTORY_URL"),
			PaymentUrl:   os.Getenv("GRPC_PAYMENT_URL"),
		},

		Paginate: Paginate{
			ItemNextPageBaseUrl:  os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			InventoryPageBaseUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
	}
}
