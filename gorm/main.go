package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// gorm model
type Record struct {
	RID   int `gorm:"primaryKey"`
	Name  string
	Price float64
}

type MyStrategy struct {
	schema.NamingStrategy
}

func (s MyStrategy) ColumnName(table, column string) string {
	return strings.ToLower(column)
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm1.db"), &gorm.Config{
		NamingStrategy: MyStrategy{},
	})
	if err != nil {
		log.Fatalln(err)
	}

	if !db.Migrator().HasTable(&Record{}) {
		fmt.Println("no table 'records' present: creating...")
		if err := db.Migrator().CreateTable(&Record{}); err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("table 'records' created")
			fmt.Println("preparing records to be inserted in the 'records' table...")
			records := make([]Record, 5)
			for i := range records {
				records[i] = Record{RID: i + 1, Name: fmt.Sprintf("name-%v", i), Price: float64(i) + 1.2}
			}
			//fmt.Println(records)
			db.Create(&records)
			fmt.Printf("records inserted in the 'records' table:\n %+v\n", records)
		}
	} else { // has table, do query and then update
		fmt.Println("'records' table found. Querying...")

		fmt.Println("SELECT * FROM records WHERE price > 100;")
		records := []Record{}
		result := db.Where("price > ?", 100).Find(&records)
		fmt.Println("rows =", result.RowsAffected, result.Error)

		fmt.Println("SELECT * FROM records WHERE rid = 10;")
		record := Record{}
		result = db.First(&record, 10)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				fmt.Println("record not found")
			} else {
				fmt.Println(result.Error)
			}
		} else {
			fmt.Println(record)
		}
	}
}
