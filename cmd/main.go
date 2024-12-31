package main

import (
	"github.com/DamiaoCanndido/document-api/config"
	"github.com/DamiaoCanndido/document-api/routes"

	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := routes.SetupRouter(db)
	router.Run(":5000")
}