package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Importa o driver MySQL (é usado com o _ para importar sem usar explicitamente)
	"github.com/google/uuid"           // Importa a biblioteca uuid para gerar IDs únicos
)

// Definindo a estrutura de dados do produto
type Product struct {
	ID    string
	Name  string
	Price float64
}

// Função para criar um novo produto com um ID único gerado automaticamente
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(), // Gera um novo ID único
		Name:  name,
		Price: price,
	}
}

func main() {
	// Conectar ao banco de dados MySQL
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Fecha a conexão com o banco de dados quando a função main() terminar

	// Criar um novo produto
	product := NewProduct("Gatilho de Dedo", 75)

	// Inserir o produto no banco de dados
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Atualizar o preço do produto
	product.Price = 100
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// Selecionar o produto do banco de dados
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de R$ %.2f", p.Name, p.Price)

	// Selecionar todos os produtos do banco de dados
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	// Deletar o produto do banco de dados
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

// Função para inserir um produto no banco de dados
func insertProduct(db *sql.DB, product *Product) error {
	// Prepara uma declaração SQL para inserção
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa a declaração SQL com os valores do produto
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

// Função para atualizar um produto no banco de dados
func updateProduct(db *sql.DB, product *Product) error {
	// Prepara uma declaração SQL para atualização
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa a declaração SQL com os valores atualizados do produto
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// Função para selecionar um produto do banco de dados por ID
func selectProduct(db *sql.DB, id string) (*Product, error) {
	// Prepara uma declaração SQL para seleção
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Executa a declaração SQL com o ID como parâmetro e lê o resultado
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Função para selecionar todos os produtos do banco de dados
func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Executa uma consulta SQL para selecionar todos os produtos
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Lê as linhas de resultados e cria uma lista de produtos
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Função para deletar um produto do banco de dados por ID
func deleteProduct(db *sql.DB, id string) error {
	// Prepara uma declaração SQL para exclusão
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa a declaração SQL com o ID como parâmetro
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
