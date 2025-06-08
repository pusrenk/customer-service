package services

import (
	"github.com/pusrenk/customer-service/internal/customers/entitites"
	"github.com/pusrenk/customer-service/internal/customers/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *entitites.Customer) error
	GetCustomerByID(id uint) (*entitites.Customer, error)
	GetAllCustomers() ([]*entitites.Customer, error)
	UpdateCustomer(customer *entitites.Customer) error
	DeleteCustomer(id uint) error
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// CreateCustomer creates a new customer
func (s *customerService) CreateCustomer(customer *entitites.Customer) error {
	return s.customerRepository.CreateCustomer(customer)
}

// GetCustomerByID returns a customer by ID
func (s *customerService) GetCustomerByID(id uint) (*entitites.Customer, error) {
	return s.customerRepository.GetCustomerByID(id)
}

// GetAllCustomers returns all customers
func (s *customerService) GetAllCustomers() ([]*entitites.Customer, error) {
	return s.customerRepository.GetAllCustomers()
}

// UpdateCustomer updates a customer
func (s *customerService) UpdateCustomer(customer *entitites.Customer) error {
	return s.customerRepository.UpdateCustomer(customer)
}

// DeleteCustomer deletes a customer by ID
func (s *customerService) DeleteCustomer(id uint) error {
	return s.customerRepository.DeleteCustomer(id)
}
