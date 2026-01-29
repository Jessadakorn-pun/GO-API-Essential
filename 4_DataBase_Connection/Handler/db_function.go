package handler

import (
	database "sql-connection/Database"
	domain "sql-connection/Domain"
)

func CreateProduct(product *domain.Product) error {
	_, err := database.DB.Exec(
		`INSERT INTO public.products(name, price) VALUES ($1, $2);`, product.Name, product.Price,
	)

	return err
}

func GetProduct(id int) (domain.Product, error) {
	var p domain.Product
	row := database.DB.QueryRow("SELECT id, name, price FROM products WHERE id=$1", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return domain.Product{}, err
	}

	return p, err
}

func GetProducts() ([]domain.Product, error) {
	rows, err := database.DB.Query("SELECT id, name, price FROM products;")

	if err != nil {
		return nil, err
	}

	var product []domain.Product

	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		product = append(product, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProduct(id int, product *domain.Product) (domain.Product, error) {
	// _, err := database.DB.Exec(
	// 	`UPDATE public.products SET name=$1, price=$2 WHERE id = $3;`,
	// 	product.Name, product.Price, id,
	// )

	// return err

	var p domain.Product
	row := database.DB.QueryRow(
		`UPDATE public.products SET name=$1, price=$2 WHERE id = $3 RETURNING id, name, price;`,
		product.Name, product.Price, id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, err

}

func DeleteProduct(id int) error {
	// var p domain.Product
	// row := database.DB.QueryRow(
	// 	`DELETE FROM public.products WHERE id = $1 RETURNING name`, id,
	// )

	// err := row.Scan(&p.Name)
	// if err != nil {
	// 	return domain.Product{}, err
	// }
	// return p, err

	_, err := database.DB.Exec(
		`DELETE FROM public.products WHERE id = $1;`, id,
	)

	return err

}
