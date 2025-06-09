package repositories

import (
	"github.com/pusrenk/customer-service/internal/customers/entitites"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *entitites.Customer) error
	GetCustomerByID(id uint) (*entitites.Customer, error)
	GetAllCustomers() ([]*entitites.Customer, error)
	UpdateCustomer(customer *entitites.Customer) error
	DeleteCustomer(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// CreateCustomer create customer
func (r *customerRepository) CreateCustomer(customer *entitites.Customer) error {
	return r.db.Create(customer).Error
}

// GetCustomerByID get customer by id
func (r *customerRepository) GetCustomerByID(id uint) (*entitites.Customer, error) {
	var customer entitites.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetAllCustomers get all customers
func (r *customerRepository) GetAllCustomers() ([]*entitites.Customer, error) {
	var customers []*entitites.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

// UpdateCustomer update customer
func (r *customerRepository) UpdateCustomer(customer *entitites.Customer) error {
	return r.db.Save(customer).Error
}

// DeleteCustomer delete customer by id
func (r *customerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&entitites.Customer{}, id).Error
}
