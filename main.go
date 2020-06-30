package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	gorm.Model
	Content string
}

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("データベースに接続できませんでした。")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
	db.Create(&Todo{Content: "明日やるべきこと"})
	var item Todo
	db.First(&item, 1)
	fmt.Println(item)
}
