package storage

import "rent-car/models"


type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
Create(models.Car)(string,error)
GetByID(id string)(models.Car,error)
GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
Update(models.Car) (string, error)
Delete(string) error
}

type ICustomerStorage interface {
Create(models.Customer)(string,error)
GetByID(id string)(models.Customer,error)
GetAll(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
Update(models.Customer) (string, error)
Delete(string) error
}
type IOrderStorage interface {
	Create(models.CreateOrder)(string,error)
	// GetByID(id string)(models.,error)
	// GetAll(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	Update(models.UpdateOrder) (string, error)
	Delete(string) error
	}