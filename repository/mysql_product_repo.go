package repository

import (
	"fmt"

	"github.com/danurwijayanto/api-go-ecomm/config"
	"github.com/danurwijayanto/api-go-ecomm/entity"
)

type mysqlRepository struct{}

// newMysqlProductRepository : constructor
func NewMysqlProductRepository() ProductRepository {
	return &mysqlRepository{}
}

func (*mysqlRepository) Add(product *entity.Product) (int, error) {
	db, err := config.GetMysqlConnect()

	if err != nil {
		fmt.Println("Mysql error", err.Error())
	}

	query := "INSERT INTO product (product_name)" +
		"VALUES (?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Prepare error", err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(product.ProductName)
	if err != nil {
		fmt.Println("Exec error", err.Error())
	}

	ID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId error", err.Error())
	}

	return int(ID), nil
}

func (*mysqlRepository) GetAll() ([]*entity.Product, error) {
	var list []*entity.Product
	db, err := config.GetMysqlConnect()

	if err != nil {
		fmt.Println("Mysql error", err.Error())
	}

	query := "SELECT id, product_name FROM product"
	row, err := db.Query(query)

	if err != nil {
		fmt.Println("Query error", err.Error())
	}

	for row.Next() {
		e := new(entity.Product)

		row.Scan(&e.ID, &e.ProductName)

		list = append(list, e)
	}

	return list, err
}
