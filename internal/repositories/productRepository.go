package repositories

import (
	"database/sql"
	"fmt"

	"github.com/vitormoschetta/go-rest-api/internal/models"
)

type IProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id string) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) error
	Delete(id string) error
}

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	rows, err := r.Db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id string) (models.Product, error) {
	var product models.Product
	err := r.Db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, nil
		}
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) Create(product models.Product) (models.Product, error) {
	res, err := r.Db.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
	if err != nil {
		return product, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return product, err
	}
	if rowsAffected != 1 {
		return product, fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return product, nil
}

func (r *ProductRepository) Update(product models.Product) error {
	res, err := r.Db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	res, err := r.Db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", rowsAffected)
	}
	return nil
}
