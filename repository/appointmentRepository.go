package repository

import (
	"database/sql"

	"github.com/felipehfs/appointment-app/dto"
)

type AppointmentRepository struct {
	Conn *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{
		Conn: db,
	}
}

func (ar AppointmentRepository) Insert(data *dto.CreateAppointmentDto) error {
	tx, err := ar.Conn.Begin()
	if err != nil {
		return err
	}

	sql := `INSERT INTO appointment.appointment(customer_id, schedule_on) VALUES ($1, $2)`
	_, err = tx.Exec(sql, data.CustomerId, data.ScheduleOn)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ar AppointmentRepository) Update(id int, data *dto.UpdateAppointmentDto) error {
	tx, err := ar.Conn.Begin()
	if err != nil {
		return err
	}

	sql := `UPDATE appointment.appointment SET customer_id=$1, schedule_on=$2 WHERE id=$3`

	_, err = tx.Exec(sql, data.CustomerId, data.ScheduleOn, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ar AppointmentRepository) Delete(id int) error {
	tx, err := ar.Conn.Begin()
	if err != nil {
		return err
	}

	sql := `DELETE FROM appointment.appointment WHERE id=$1`

	_, err = tx.Exec(sql, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (ar AppointmentRepository) FindById(id int) (appointment dto.SelectAppointment, err error) {
	sql := `
		SELECT 
		a.id, a.schedule_on,c.id, c.name, c.created_at, sex."name" as sex, c.birthday,
		address.street, address.number, address.neighbor, address.state, address.city
		FROM appointment.appointment a 
		INNER JOIN appointment.customer c ON a.customer_id = c.id
		INNER JOIN appointment.sex ON sex.id = c.sex_id
		INNER JOIN appointment.address ON address.id = c.address_id
		WHERE a.id=$1
	`
	err = ar.Conn.QueryRow(sql, id).Scan(
		&appointment.Id,
		&appointment.ScheduleOn,
		&appointment.Customer.Id,
		&appointment.Customer.Name,
		&appointment.Customer.CreatedAt,
		&appointment.Customer.Sex,
		&appointment.Customer.BirthDay,
		&appointment.Customer.Address.Street,
		&appointment.Customer.Address.Number,
		&appointment.Customer.Address.Neighbor,
		&appointment.Customer.Address.State,
		&appointment.Customer.Address.City,
	)

	if err != nil {
		return
	}

	return
}

func (ar AppointmentRepository) Select() (appointments []dto.SelectAppointment, err error) {
	sql := `
		SELECT 
		a.id, a.schedule_on,c.id, c.name, c.created_at, sex."name" as sex, c.birthday,
		address.street, address.number, address.neighbor, address.state, address.city
		FROM appointment.appointment a 
		INNER JOIN appointment.customer c ON a.customer_id = c.id
		INNER JOIN appointment.sex ON sex.id = c.sex_id
		INNER JOIN appointment.address ON address.id = c.address_id
	`
	rows, err := ar.Conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var appointment dto.SelectAppointment

		err = rows.Scan(&appointment.Id,
			&appointment.ScheduleOn,
			&appointment.Customer.Id,
			&appointment.Customer.Name,
			&appointment.Customer.CreatedAt,
			&appointment.Customer.Sex,
			&appointment.Customer.BirthDay,
			&appointment.Customer.Address.Street,
			&appointment.Customer.Address.Number,
			&appointment.Customer.Address.Neighbor,
			&appointment.Customer.Address.State,
			&appointment.Customer.Address.City,
		)

		if err != nil {
			return
		}

		appointments = append(appointments, appointment)

	}

	return
}
