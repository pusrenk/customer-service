package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/pusrenk/customer-service/internal/customers/entitites"
	pb "github.com/pusrenk/customer-service/internal/protobuf"
	"github.com/pusrenk/customer-service/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CustomerServiceServerTestSuite struct {
	suite.Suite
	mockService *mocks.MockCustomerService
	server      *CustomerServiceServer
}

func (s *CustomerServiceServerTestSuite) SetupTest() {
	s.mockService = mocks.NewMockCustomerService(s.T())
	s.server = NewCustomerServiceServer(s.mockService)
}

func (s *CustomerServiceServerTestSuite) TestCreateCustomer() {
	now := time.Now()
	tests := []struct {
		name          string
		request       *pb.CreateCustomerRequest
		mockSetup     func()
		expectedError codes.Code
		validate      func(*testing.T, *pb.CreateCustomerResponse)
	}{
		{
			name: "success create customer",
			request: &pb.CreateCustomerRequest{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().CreateCustomer(mock.AnythingOfType("*entitites.Customer")).
					RunAndReturn(func(customer *entitites.Customer) error {
						customer.ID = 1
						customer.CreatedAt = now
						customer.UpdatedAt = now
						return nil
					}).Once()
			},
			expectedError: codes.OK,
			validate: func(t *testing.T, response *pb.CreateCustomerResponse) {
				assert.NotNil(t, response)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, uint64(1), response.Customer.Id)
				assert.Equal(t, "John Doe", response.Customer.Name)
				assert.Equal(t, "john@example.com", response.Customer.Email)
				assert.Equal(t, "+1234567890", response.Customer.Phone)
				assert.Equal(t, "system", response.Customer.CreatedBy)
				assert.Equal(t, "system", response.Customer.UpdatedBy)
				assert.NotNil(t, response.Customer.CreatedAt)
				assert.NotNil(t, response.Customer.UpdatedAt)
			},
		},
		{
			name: "missing name",
			request: &pb.CreateCustomerRequest{
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.CreateCustomerResponse) {},
		},
		{
			name: "missing email",
			request: &pb.CreateCustomerRequest{
				Name:      "John Doe",
				Phone:     "+1234567890",
				CreatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.CreateCustomerResponse) {},
		},
		{
			name: "missing phone",
			request: &pb.CreateCustomerRequest{
				Name:      "John Doe",
				Email:     "john@example.com",
				CreatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.CreateCustomerResponse) {},
		},
		{
			name: "missing created_by",
			request: &pb.CreateCustomerRequest{
				Name:  "John Doe",
				Email: "john@example.com",
				Phone: "+1234567890",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.CreateCustomerResponse) {},
		},
		{
			name: "service error",
			request: &pb.CreateCustomerRequest{
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				CreatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().CreateCustomer(mock.AnythingOfType("*entitites.Customer")).
					Return(assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.CreateCustomerResponse) {},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.mockSetup()

			// Call the handler
			response, err := s.server.CreateCustomer(context.Background(), tt.request)

			// Check error
			if tt.expectedError != codes.OK {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError, st.Code())
			} else {
				assert.NoError(t, err)
				tt.validate(t, response)
			}
		})
	}
}

func (s *CustomerServiceServerTestSuite) TestGetCustomer() {
	now := time.Now()
	tests := []struct {
		name          string
		request       *pb.GetCustomerRequest
		mockSetup     func()
		expectedError codes.Code
		validate      func(*testing.T, *pb.GetCustomerResponse)
	}{
		{
			name: "success get customer",
			request: &pb.GetCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(&entitites.Customer{
						ID:        1,
						Name:      "John Doe",
						Email:     "john@example.com",
						Phone:     "+1234567890",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: "system",
						UpdatedBy: "system",
					}, nil).Once()
			},
			expectedError: codes.OK,
			validate: func(t *testing.T, response *pb.GetCustomerResponse) {
				assert.NotNil(t, response)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, uint64(1), response.Customer.Id)
				assert.Equal(t, "John Doe", response.Customer.Name)
				assert.Equal(t, "john@example.com", response.Customer.Email)
				assert.Equal(t, "+1234567890", response.Customer.Phone)
				assert.Equal(t, "system", response.Customer.CreatedBy)
				assert.Equal(t, "system", response.Customer.UpdatedBy)
				assert.NotNil(t, response.Customer.CreatedAt)
				assert.NotNil(t, response.Customer.UpdatedAt)
			},
		},
		{
			name: "invalid id",
			request: &pb.GetCustomerRequest{
				Id: 0,
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.GetCustomerResponse) {},
		},
		{
			name: "customer not found",
			request: &pb.GetCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: codes.NotFound,
			validate:      func(t *testing.T, response *pb.GetCustomerResponse) {},
		},
		{
			name: "service error",
			request: &pb.GetCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.GetCustomerResponse) {},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.mockSetup()

			// Call the handler
			response, err := s.server.GetCustomer(context.Background(), tt.request)

			// Check error
			if tt.expectedError != codes.OK {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError, st.Code())
			} else {
				assert.NoError(t, err)
				tt.validate(t, response)
			}
		})
	}
}

func (s *CustomerServiceServerTestSuite) TestListCustomers() {
	now := time.Now()
	tests := []struct {
		name          string
		request       *pb.ListCustomersRequest
		mockSetup     func()
		expectedError codes.Code
		validate      func(*testing.T, *pb.ListCustomersResponse)
	}{
		{
			name:    "success list customers",
			request: &pb.ListCustomersRequest{},
			mockSetup: func() {
				s.mockService.EXPECT().GetAllCustomers().
					Return([]*entitites.Customer{
						{
							ID:        1,
							Name:      "John Doe",
							Email:     "john@example.com",
							Phone:     "+1234567890",
							CreatedAt: now,
							UpdatedAt: now,
							CreatedBy: "system",
							UpdatedBy: "system",
						},
					}, nil).Once()
			},
			expectedError: codes.OK,
			validate: func(t *testing.T, response *pb.ListCustomersResponse) {
				assert.NotNil(t, response)
				assert.NotNil(t, response.Customers)
				assert.Equal(t, 1, len(response.Customers))
				assert.Equal(t, uint64(1), response.Customers[0].Id)
				assert.Equal(t, "John Doe", response.Customers[0].Name)
				assert.Equal(t, "john@example.com", response.Customers[0].Email)
				assert.Equal(t, "+1234567890", response.Customers[0].Phone)
				assert.Equal(t, "system", response.Customers[0].CreatedBy)
				assert.Equal(t, "system", response.Customers[0].UpdatedBy)
				assert.NotNil(t, response.Customers[0].CreatedAt)
				assert.NotNil(t, response.Customers[0].UpdatedAt)
			},
		},
		{
			name:    "service error",
			request: &pb.ListCustomersRequest{},
			mockSetup: func() {
				s.mockService.EXPECT().GetAllCustomers().
					Return(nil, assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.ListCustomersResponse) {},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.mockSetup()

			// Call the handler
			response, err := s.server.ListCustomers(context.Background(), tt.request)

			// Check error
			if tt.expectedError != codes.OK {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError, st.Code())
			} else {
				assert.NoError(t, err)
				tt.validate(t, response)
			}
		})
	}
}

func (s *CustomerServiceServerTestSuite) TestUpdateCustomer() {
	now := time.Now()
	tests := []struct {
		name          string
		request       *pb.UpdateCustomerRequest
		mockSetup     func()
		expectedError codes.Code
		validate      func(*testing.T, *pb.UpdateCustomerResponse)
	}{
		{
			name: "success update customer",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(&entitites.Customer{
						ID:        1,
						Name:      "John Doe",
						Email:     "john@example.com",
						Phone:     "+1234567890",
						CreatedAt: now,
						UpdatedAt: now,
						CreatedBy: "system",
						UpdatedBy: "system",
					}, nil).Once()
				s.mockService.EXPECT().UpdateCustomer(mock.AnythingOfType("*entitites.Customer")).
					Return(nil).Once()
			},
			expectedError: codes.OK,
			validate: func(t *testing.T, response *pb.UpdateCustomerResponse) {
				assert.NotNil(t, response)
				assert.NotNil(t, response.Customer)
				assert.Equal(t, uint64(1), response.Customer.Id)
				assert.Equal(t, "John Doe", response.Customer.Name)
				assert.Equal(t, "john@example.com", response.Customer.Email)
				assert.Equal(t, "+1234567890", response.Customer.Phone)
				assert.Equal(t, "system", response.Customer.CreatedBy)
				assert.Equal(t, "system", response.Customer.UpdatedBy)
				assert.NotNil(t, response.Customer.CreatedAt)
				assert.NotNil(t, response.Customer.UpdatedAt)
			},
		},
		{
			name: "invalid id",
			request: &pb.UpdateCustomerRequest{
				Id: 0,
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "missing name",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Email:     "john@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "missing email",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "missing phone",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				UpdatedBy: "system",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "missing updated_by",
			request: &pb.UpdateCustomerRequest{
				Id:    1,
				Name:  "John Doe",
				Email: "john@example.com",
				Phone: "+1234567890",
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "customer not found",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: codes.NotFound,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "get customer by id error",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
		{
			name: "update customer error",
			request: &pb.UpdateCustomerRequest{
				Id:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(&entitites.Customer{
						ID: 1,
					}, nil).Once()
				s.mockService.EXPECT().UpdateCustomer(mock.AnythingOfType("*entitites.Customer")).
					Return(assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.UpdateCustomerResponse) {},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.mockSetup()

			// Call the handler
			response, err := s.server.UpdateCustomer(context.Background(), tt.request)

			// Check error
			if tt.expectedError != codes.OK {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError, st.Code())
			} else {
				assert.NoError(t, err)
				tt.validate(t, response)
			}
		})
	}
}

func (s *CustomerServiceServerTestSuite) TestDeleteCustomer() {
	tests := []struct {
		name          string
		request       *pb.DeleteCustomerRequest
		mockSetup     func()
		expectedError codes.Code
		validate      func(*testing.T, *pb.DeleteCustomerResponse)
	}{
		{
			name: "success delete customer",
			request: &pb.DeleteCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(&entitites.Customer{
						ID: 1,
					}, nil).Once()
				s.mockService.EXPECT().DeleteCustomer(uint(1)).
					Return(nil).Once()
			},
			expectedError: codes.OK,
			validate:      func(t *testing.T, response *pb.DeleteCustomerResponse) {},
		},
		{
			name: "invalid id",
			request: &pb.DeleteCustomerRequest{
				Id: 0,
			},
			mockSetup:     func() {},
			expectedError: codes.InvalidArgument,
			validate:      func(t *testing.T, response *pb.DeleteCustomerResponse) {},
		},
		{
			name: "customer not found",
			request: &pb.DeleteCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: codes.NotFound,
			validate:      func(t *testing.T, response *pb.DeleteCustomerResponse) {},
		},
		{
			name: "get customer by id error",
			request: &pb.DeleteCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(nil, assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.DeleteCustomerResponse) {},
		},
		{
			name: "service error",
			request: &pb.DeleteCustomerRequest{
				Id: 1,
			},
			mockSetup: func() {
				s.mockService.EXPECT().GetCustomerByID(uint(1)).
					Return(&entitites.Customer{
						ID: 1,
					}, nil).Once()
				s.mockService.EXPECT().DeleteCustomer(uint(1)).
					Return(assert.AnError).Once()
			},
			expectedError: codes.Internal,
			validate:      func(t *testing.T, response *pb.DeleteCustomerResponse) {},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			tt.mockSetup()

			// Call the handler
			response, err := s.server.DeleteCustomer(context.Background(), tt.request)

			// Check error
			if tt.expectedError != codes.OK {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedError, st.Code())
			} else {
				assert.NoError(t, err)
				tt.validate(t, response)
			}
		})
	}
}

func TestCustomerServiceServerSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceServerTestSuite))
}
