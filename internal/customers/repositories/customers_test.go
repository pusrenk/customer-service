package repositories_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pusrenk/customer-service/internal/customers/entitites"
	"github.com/pusrenk/customer-service/internal/customers/repositories"
	"github.com/pusrenk/customer-service/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CustomerRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo repositories.CustomerRepository
}

func (s *CustomerRepositoryTestSuite) SetupSuite() {
	s.db, s.mock = test.NewTestDB()
	s.repo = repositories.NewCustomerRepository(s.db)
}

func (s *CustomerRepositoryTestSuite) TearDownTest() {
	assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *CustomerRepositoryTestSuite) TestCreateCustomer() {
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
				s.mock.ExpectBegin()
				s.mock.ExpectQuery("INSERT INTO .customers.").
					WithArgs(customer.Name, customer.Email, customer.Phone, customer.CreatedBy, customer.UpdatedBy, customer.DeletedAt).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
						AddRow(1, time.Now(), time.Now()))
				s.mock.ExpectCommit()
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
				s.mock.ExpectBegin()
				s.mock.ExpectQuery("INSERT INTO .customers.").
					WithArgs(customer.Name, customer.Email, customer.Phone, customer.CreatedBy, customer.UpdatedBy, customer.DeletedAt).
					WillReturnError(gorm.ErrInvalidDB)
				s.mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.customer)
			err := s.repo.CreateCustomer(tt.customer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func (s *CustomerRepositoryTestSuite) TestGetCustomerByID() {
	now := time.Now()
	customer := &entitites.Customer{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Phone:     "+1234567890",
		CreatedAt: now,
		CreatedBy: "system",
		UpdatedAt: now,
		UpdatedBy: "system",
	}

	tests := []struct {
		name    string
		id      uint
		mock    func()
		want    *entitites.Customer
		wantErr bool
	}{
		{
			name: "success get customer",
			id:   1,
			mock: func() {
				s.mock.ExpectQuery(`SELECT \* FROM "customers" WHERE "customers"."id" = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(customer.ID, customer.Name, customer.Email, customer.Phone, customer.CreatedAt, customer.CreatedBy, customer.UpdatedAt, customer.UpdatedBy))
			},
			want:    customer,
			wantErr: false,
		},
		{
			name: "error get customer",
			id:   1,
			mock: func() {
				s.mock.ExpectQuery(`SELECT \* FROM "customers" WHERE "customers"."id" = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
					WithArgs(1, 1).
					WillReturnError(gorm.ErrInvalidDB)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.repo.GetCustomerByID(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func (s *CustomerRepositoryTestSuite) TestGetAllCustomers() {
	now := time.Now()
	customers := []*entitites.Customer{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Phone:     "+1234567890",
			CreatedAt: now,
			CreatedBy: "system",
			UpdatedAt: now,
			UpdatedBy: "system",
		},
		{
			ID:        2,
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			Phone:     "+0987654321",
			CreatedAt: now,
			CreatedBy: "system",
			UpdatedAt: now,
			UpdatedBy: "system",
		},
	}

	tests := []struct {
		name    string
		mock    func()
		want    []*entitites.Customer
		wantErr bool
	}{
		{
			name: "success get all customers",
			mock: func() {
				s.mock.ExpectQuery(`SELECT \* FROM "customers" WHERE "customers"."deleted_at" IS NULL`).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(customers[0].ID, customers[0].Name, customers[0].Email, customers[0].Phone, customers[0].CreatedAt, customers[0].CreatedBy, customers[0].UpdatedAt, customers[0].UpdatedBy).
						AddRow(customers[1].ID, customers[1].Name, customers[1].Email, customers[1].Phone, customers[1].CreatedAt, customers[1].CreatedBy, customers[1].UpdatedAt, customers[1].UpdatedBy))
			},
			want:    customers,
			wantErr: false,
		},
		{
			name: "error get all customers",
			mock: func() {
				s.mock.ExpectQuery(`SELECT \* FROM "customers" WHERE "customers"."deleted_at" IS NULL`).
					WillReturnError(gorm.ErrInvalidDB)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := s.repo.GetAllCustomers()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func (s *CustomerRepositoryTestSuite) TestUpdateCustomer() {
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
				Name:      "John Doe Updated",
				Email:     "john.updated@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(`UPDATE "customers" SET "name"=\$1,"email"=\$2,"phone"=\$3,"created_at"=\$4,"created_by"=\$5,"updated_at"=\$6,"updated_by"=\$7,"deleted_at"=\$8 WHERE "customers"."deleted_at" IS NULL AND "id" = \$9`).
					WithArgs(customer.Name, customer.Email, customer.Phone, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), customer.UpdatedBy, sqlmock.AnyArg(), customer.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "error update customer",
			customer: &entitites.Customer{
				ID:        1,
				Name:      "John Doe Updated",
				Email:     "john.updated@example.com",
				Phone:     "+1234567890",
				UpdatedBy: "system",
			},
			mock: func(customer *entitites.Customer) {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(`UPDATE "customers" SET "name"=\$1,"email"=\$2,"phone"=\$3,"created_at"=\$4,"created_by"=\$5,"updated_at"=\$6,"updated_by"=\$7,"deleted_at"=\$8 WHERE "customers"."deleted_at" IS NULL AND "id" = \$9`).
					WithArgs(customer.Name, customer.Email, customer.Phone, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), customer.UpdatedBy, sqlmock.AnyArg(), customer.ID).
					WillReturnError(gorm.ErrInvalidDB)
				s.mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock(tt.customer)
			err := s.repo.UpdateCustomer(tt.customer)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func (s *CustomerRepositoryTestSuite) TestDeleteCustomer() {
	tests := []struct {
		name    string
		id      uint
		mock    func()
		wantErr bool
	}{
		{
			name: "success delete customer",
			id:   1,
			mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(`UPDATE "customers" SET "deleted_at"=\$1 WHERE "customers"."id" = \$2 AND "customers"."deleted_at" IS NULL`).
					WithArgs(sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "error delete customer",
			id:   1,
			mock: func() {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(`UPDATE "customers" SET "deleted_at"=\$1 WHERE "customers"."id" = \$2 AND "customers"."deleted_at" IS NULL`).
					WithArgs(sqlmock.AnyArg(), 1).
					WillReturnError(gorm.ErrInvalidDB)
				s.mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := s.repo.DeleteCustomer(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCustomerRepositorySuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}
