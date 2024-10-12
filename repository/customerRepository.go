package repository

import (
	"database/sql"

	"github.com/felipehfs/appointment-app/dto"
)

type CustomerRepository struct {
	Conn *sql.DB
}

func NewCustomerRepository(connection *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		Conn: connection,
	}
}

func (cr CustomerRepository) Insert(data dto.CreateCustomer) error {
	tx, err := cr.Conn.Begin()
	if err != nil {
		return err
	}

	var lastAddressId int

	sql := `
		INSERT INTO appointment.address(street, number,neighbor, state, city)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id 
	`

	err = tx.QueryRow(sql, data.Address.Street, data.Address.Number, data.Address.Neighbor,
		data.Address.State, data.Address.City).
		Scan(&lastAddressId)
	if err != nil {
		return err
	}

	sql = `
		INSERT INTO appointment.customer(name, sex_id, address_id, birthday) VALUES ($1, $2, $3, $4)
	`

	_, err = tx.Exec(sql, data.Name, data.SexId, lastAddressId, data.BirthDay)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (cr CustomerRepository) Select() ([]dto.SelectCustomerDTO, error) {
	var result []dto.SelectCustomerDTO

	sql := `
		SELECT 
			c.id, c.name, c.created_at, s.name as sex, c.birthday,
			a.street, a.number, a.neighbor, a.state, a.city
		FROM appointment.customer c 
		INNER JOIN appointment.sex s ON c.sex_id = s.id
		INNER JOIN appointment.address a ON a.id = c.address_id
	`

	rows, err := cr.Conn.Query(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var customer dto.SelectCustomerDTO

		err = rows.Scan(&customer.Id, &customer.Name, &customer.CreatedAt, &customer.Sex, &customer.BirthDay, &customer.Address.Street,
			&customer.Address.Number, &customer.Address.Neighbor, &customer.Address.State, &customer.Address.City)

		if err != nil {
			return nil, err
		}

		result = append(result, customer)
	}

	return result, nil
}

func (cr CustomerRepository) FindById(id int) (*dto.SelectCustomerDTO, error) {
	sql := `
	SELECT 
		c.id, c.name, c.created_at, s.name as sex, c.birthday, 
		a.street, a.number, a.neighbor, a.state, a.city
	FROM appointment.customer c 
	INNER JOIN appointment.sex s ON c.sex_id = s.id
	INNER JOIN appointment.address a ON a.id = c.address_id
	WHERE c.id=$1
	`
	var customer dto.SelectCustomerDTO

	err := cr.Conn.QueryRow(sql, id).Scan(&customer.Id, &customer.Name, &customer.CreatedAt, &customer.Sex, &customer.BirthDay, &customer.Address.Street,
		&customer.Address.Number, &customer.Address.Neighbor, &customer.Address.State, &customer.Address.City)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (cr CustomerRepository) getAddressId(customerId int) (int, error) {
	sql := `
	SELECT address_id FROM appointment.customer WHERE id=$1
`
	var addressId int

	err := cr.Conn.QueryRow(sql, customerId).Scan(&addressId)
	if err != nil {
		return -1, err
	}

	return addressId, nil
}

func (cr CustomerRepository) Remove(id int) error {

	tx, err := cr.Conn.Begin()
	if err != nil {
		return err
	}

	addressId, err := cr.getAddressId(id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM appointment.customer WHERE id=$1", id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM appointment.address WHERE id=$1", addressId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (cr CustomerRepository) Update(id int, data dto.UpdateCustomer) error {
	tx, err := cr.Conn.Begin()
	if err != nil {
		return err
	}

	addressId, err := cr.getAddressId(id)

	if err != nil {
		return err
	}

	sql := `	
		UPDATE appointment.address 
		SET street=$1, number=$2, neighbor=$3, state=$4, city=$5
		WHERE id=$6
	`

	_, err = tx.Exec(sql, data.Address.Street, data.Address.Number, data.Address.Neighbor, data.Address.State, data.Address.City, addressId)
	if err != nil {
		tx.Rollback()
		return err
	}

	sql = `
		UPDATE appointment.customer
		SET name=$1, sex_id=$2, birthday=$3
		WHERE id=$4
	`

	_, err = tx.Exec(sql, data.Name, data.SexId, data.Birthday, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
