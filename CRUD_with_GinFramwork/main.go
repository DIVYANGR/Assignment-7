package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	Id      string `json:id`
	Name    string `json:name`
	Price   string `json:price`
	Quality string `json:quality`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "divyang"
	dbPass := "divyangk1998"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {
		var p product
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"name":    p.Name,
				"price":   p.Price,
				"quality": p.Quality,
			})
			db := dbConn()
			insForm, err := db.Prepare("INSERT INTO product(name, price, quality) VALUES(?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(p.Name, p.Price, p.Quality)
			fmt.Printf("name: %s; price: %s; quality: %s", p.Name, p.Price, p.Quality)
		}

	})

	router.PUT("/updatename", func(c *gin.Context) {

		var p product
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"name": p.Name,
				//"price":   p.Price,
				//"quality": p.Quality,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE product SET name=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Name, p.Id)
			//fmt.Printf("name: %s; price: %s; quality: %s",  p.Price, p.Quality)
		}
	})

	router.PUT("/updateprice", func(c *gin.Context) {

		var p product
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"price": p.Price,
				//"price":   p.Price,
				//"quality": p.Quality,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE product SET price=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Price, p.Id)
			//fmt.Printf("name: %s; price: %s; quality: %s", p.Name, p.Price, p.Quality)
		}
	})

	router.PUT("/updatequality", func(c *gin.Context) {

		var p product
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"quality": p.Quality,
				//"price":   p.Price,
				//"quality": p.Quality,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE product SET quality=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Quality, p.Id)
			//fmt.Printf("name: %s; price: %s; quality: %s", p.Name, p.Price, p.Quality)
		}
	})

	router.GET("/GET", func(c *gin.Context) {
		var p product
		if c.BindJSON(&p) == nil {
			db := dbConn()
			selDB, err := db.Query("SELECT * FROM product WHERE id=?", p.Id)
			if err != nil {
				panic(err.Error())
			}

			var id, name, price, quality string
			for selDB.Next() {

				err = selDB.Scan(&id, &name, &price, &quality)
				if err != nil {
					panic(err.Error())
				}
			}
			fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)

			c.JSON(200, gin.H{
				"id":      id,
				"name":    name,
				"price":   price,
				"quality": quality,
			})

		}

	})

	router.GET("/getall", func(c *gin.Context) {
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM product")
		if err != nil {
			panic(err.Error())
		}

		var id, name, price, quality string
		for selDB.Next() {

			err = selDB.Scan(&id, &name, &price, &quality)
			c.JSON(200, gin.H{
				"id":      id,
				"name":    name,
				"price":   price,
				"quality": quality,
			})
			fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)
			if err != nil {
				panic(err.Error())
			}
		}
		//fmt.Printf("name: %s; price: %s; quality: %s", name, price, quality)

	})

	router.DELETE("/delete", func(c *gin.Context) {
		var p product
		if c.BindJSON(&p) == nil {
			db := dbConn()
			delForm, err := db.Prepare("DELETE FROM product WHERE name=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(p.Name)
			log.Println("DELETE")
			defer db.Close()
		}

	})

	router.Run(":8080")
}
