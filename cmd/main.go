package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/karlaugust1/code_pix_go/application/grpc"
	"github.com/karlaugust1/code_pix_go/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
