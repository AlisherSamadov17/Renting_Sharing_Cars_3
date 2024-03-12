package storage

import (
	"database/sql"
	"rent-car/models"

	"github.com/google/uuid"
)



type orderRepo struct {
	db *sql.DB
}

func NewOrder(db *sql.DB) orderRepo {
	return orderRepo{
		db: db,
	}
}
func (o *orderRepo) Create(or models.CreateOrder) (string,error) {
	id :=uuid.New()

	query :=`insert into orders(
		id,
		car_id,
		customer_id,
		from_date,
		to_date,
		status,
		paid,
		updated_at=CURRENT_TIMESTAMP
	) values($1,$2,$3,$4,$5,$6,$7)`

	_,err:=o.db.Exec(query,id.String(),or.CarId,or.CustomerId,or.FromDate,or.ToDate,or.Status,or.Paid,or.CreatedAt)
	if err != nil {
		return "",err
	}
	return id.String(),nil
}

func (o *orderRepo) Update(or models.UpdateOrder) (string,error) {
	query:=`update orders set
	   from_date=$1,
	   to_date=$2,
	   status=$3,
	   paid=$4,
       updated_at=CURRENT_TIMESTAMP
	   WHERE id=$5 and car_id=$6 and customer_id=$7
	`
	_,err:=o.db.Exec(query,or.FromDate,or.ToDate,or.Status,or.Paid,or.Id,or.CarId,or.CustomerId)
	if err != nil {
		return "",err
	}
	return or.Id,nil
}

// Get All

// Get by id

func (o *orderRepo) Delete(id string) error {
	query :=`delete from orders where id = &1`
	_,err := o.db.Exec(query,id)
	if err != nil {
		return err
	}
	return nil
}