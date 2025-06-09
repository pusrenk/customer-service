package services_test

import (
	"context"
	"testing"

	"github.com/pusrenk/customer-service/internal/customers/entitites"
	"github.com/pusrenk/customer-service/internal/customers/services"
	testMocks "github.com/pusrenk/customer-service/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerServiceTestSuite struct {
	suite.Suite
	service  services.CustomerService
	mockRepo *testMocks.MockCustomerRepository
}

func (s *CustomerServiceTestSuite) SetupSuite() {
	s.mockRepo = testMocks.NewMockCustomerRepository(s.T())
	s.service = services.NewCustomerService(s.mockRepo)
}

// func (s *CustomerServiceTestSuite) TearDownTest() {
// 	assert.True(s.T(), s.mockRepo.AssertExpectations(s.T()))
// }

func (s *CustomerServiceTestSuite) TestCreateCustomer() {
	tests := []struct {
		name     string
		customer *entitites.Customer
		mock     func(customer *entitites.Customer)
		wantErr  bool
	}{
		{
			name: "success create customer",
			customer: &entitites.Customer{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mockRepo.On("CreateCustomer", customer).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "error create customer",
			customer: &entitites.Customer{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mockRepo.On("CreateCustomer", customer).Return(context.DeadlineExceeded).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.customer)
			err := s.service.CreateCustomer(tt.customer)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func (s *CustomerServiceTestSuite) TestGetCustomerByID() {
	tests := []struct {
		name         string
		id           uint
		mock         func(id uint)
		wantErr      bool
		wantCustomer *entitites.Customer
	}{
		{
			name: "success get customer by id",
			id:   1,
			mock: func(id uint) {
				s.mockRepo.On("GetCustomerByID", id).Return(&entitites.Customer{
					ID:        1,
					Name:      "John Doe",
					Email:     "john@example.com",
					Phone:     "+1234567890",
					CreatedBy: "system",
					UpdatedBy: "system",
				}, nil).Once()
			},
			wantCustomer: &entitites.Customer{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
				UpdatedBy: "system",
			},
		},
		{
			name: "error get customer by id",
			id:   1,
			mock: func(id uint) {
				s.mockRepo.On("GetCustomerByID", id).Return(nil, context.DeadlineExceeded).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.id)
			customer, err := s.service.GetCustomerByID(tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantCustomer, customer)
		})
	}
}

func (s *CustomerServiceTestSuite) TestGetAllCustomers() {
	tests := []struct {
		name          string
		mock          func()
		wantErr       bool
		wantCustomers []*entitites.Customer
	}{
		{
			name: "success get all customers",
			mock: func() {
				s.mockRepo.On("GetAllCustomers").Return([]*entitites.Customer{
					{ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "+1234567890", CreatedBy: "system", UpdatedBy: "system"},
				}, nil).Once()
			},
			wantCustomers: []*entitites.Customer{
				{ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "+1234567890", CreatedBy: "system", UpdatedBy: "system"},
			},
		},
		{
			name: "error get all customers",
			mock: func() {
				s.mockRepo.On("GetAllCustomers").Return(nil, context.DeadlineExceeded).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			customers, err := s.service.GetAllCustomers()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantCustomers, customers)
		})
	}
}

func (s *CustomerServiceTestSuite) TestUpdateCustomer() {
	tests := []struct {
		name     string
		customer *entitites.Customer
		mock     func(customer *entitites.Customer)
		wantErr  bool
	}{
		{
			name: "success update customer",
			customer: &entitites.Customer{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mockRepo.On("UpdateCustomer", customer).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "error update customer",
			customer: &entitites.Customer{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mockRepo.On("UpdateCustomer", customer).Return(context.DeadlineExceeded).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.customer)
			err := s.service.UpdateCustomer(tt.customer)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func (s *CustomerServiceTestSuite) TestDeleteCustomer() {
	tests := []struct {
		name    string
		id      uint
		mock    func(id uint)
		wantErr bool
	}{
		{
			name: "success delete customer",
			id:   1,
			mock: func(id uint) {
				s.mockRepo.On("DeleteCustomer", id).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "error delete customer",
			id:   1,
			mock: func(id uint) {
				s.mockRepo.On("DeleteCustomer", id).Return(context.DeadlineExceeded).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.id)
			err := s.service.DeleteCustomer(tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCustomerServiceSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceTestSuite))
}
