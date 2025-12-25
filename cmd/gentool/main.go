package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Config file path
var configFile = flag.String("f", "configs/code-genesis-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Connect to database
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Initialize Generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/query",
		ModelPkgPath: "internal/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Use the DB connection
	g.UseDB(db)

	// Generate all tables
	g.ApplyBasic(g.GenerateAllTable()...)

	// Execute
	g.Execute()
	fmt.Println("Code Generation Complete!")
}
