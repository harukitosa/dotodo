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
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("データベースに接続できませんでした。")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{})
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	switch args[0] {
	case "add":
		err := add(args[1])
		if err != nil {
			panic("DATABASE NOT FOUND")
		}
		break
	case "all":
		err := all()
		if err != nil {
			panic("ALL ERROR")
		}
		break
	default:
		break
	}
	// db.Create(&Todo{Content: "明日やるべきこと"})
	// var item Todo
	// db.First(&item, 1)
	// fmt.Println(item)
}

func add(content string) (err error) {
	fmt.Println("ADD COMMAND")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(&Todo{Content: content})
	return nil
}

func all() (err error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	var todos []Todo
	db.Find(&todos)
	for i := 0; i < len(todos); i++ {
		fmt.Println(todos[i].Content)
	}
	return nil
}
