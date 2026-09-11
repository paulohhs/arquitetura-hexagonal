package main

/*
Base main para testar CLI e WEBSERVER adapters
*/
import "github.com/paulohhs/arquitetura-hexagonal/cmd"

func main() {
	cmd.Execute()
}

/*
Base main para testar DB adapter
*/
// import (
// 	"database/sql"
// 	db2 "github.com/paulohhs/arquitetura-hexagonal/adapters/db"
// 	"github.com/paulohhs/arquitetura-hexagonal/application"
// )

// func main() {
// 	db, _ := sql.Open("sqlite3", "db.sqlite")
// 	productDbAdapter := db2.NewProductDb(db)
// 	productService := application.NewProductService(productDbAdapter)

// 	product, _ := productService.Create("Produto Exemplo", 30)
// 	productService.Enable(product)
// 	productService.Disable(product)
// }
