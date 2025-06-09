package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/pusrenk/customer-service/internal/customers/entitites"
	pb "github.com/pusrenk/customer-service/internal/protobuf"
	"github.com/pusrenk/customer-service/internal/customers/services"
	"github.com/pusrenk/customer-service/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type CustomerServiceServer struct {
	pb.UnimplementedCustomerServiceServer
	customerService services.CustomerService
}

func NewCustomerServiceServer(customerService services.CustomerService) *CustomerServiceServer {
	return &CustomerServiceServer{
		customerService: customerService,
	}
}

func (s *CustomerServiceServer) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	// Validate request
	if req.Name == "" {
		log.Errorf("name is required")
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Email == "" {
		log.Errorf("email is required")
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Phone == "" {
		log.Errorf("phone is required")
		return nil, status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.CreatedBy == "" {
		log.Errorf("created_by is required")
		return nil, status.Error(codes.InvalidArgument, "created_by is required")
	}

	// Convert request to entitites
	customer := &entitites.Customer{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UpdatedBy: req.CreatedBy,
	}

	// Call service layer
	if err := s.customerService.CreateCustomer(customer); err != nil {
		log.Errorf("failed to create customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to create customer")
	}

	// Convert entities to response
	response := &pb.CreateCustomerResponse{
		Customer: &pb.Customer{
			Id:        uint64(customer.ID),
			Name:      customer.Name,
			Email:     customer.Email,
			Phone:     customer.Phone,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			CreatedBy: customer.CreatedBy,
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
			UpdatedBy: customer.UpdatedBy,
		},
	}

	return response, nil
}

func (s *CustomerServiceServer) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	if req.Id == 0 {
		log.Errorf("id is required")
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	customer, err := s.customerService.GetCustomerByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("customer not found")
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		log.Errorf("failed to get customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to get customer")
	}

	response := &pb.GetCustomerResponse{
		Customer: &pb.Customer{
			Id:        uint64(customer.ID),
			Name:      customer.Name,
			Email:     customer.Email,
			Phone:     customer.Phone,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			CreatedBy: customer.CreatedBy,
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
			UpdatedBy: customer.UpdatedBy,
		},
	}

	return response, nil
}

func (s *CustomerServiceServer) ListCustomers(ctx context.Context, req *pb.ListCustomersRequest) (*pb.ListCustomersResponse, error) {
	customers, err := s.customerService.GetAllCustomers()
	if err != nil {
		log.Errorf("failed to list customers: %v", err)
		return nil, status.Error(codes.Internal, "failed to list customers")
	}

	response := &pb.ListCustomersResponse{
		Customers: make([]*pb.Customer, len(customers)),
	}

	for i, customer := range customers {
		response.Customers[i] = &pb.Customer{
			Id:        uint64(customer.ID),
			Name:      customer.Name,
			Email:     customer.Email,
			Phone:     customer.Phone,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			CreatedBy: customer.CreatedBy,
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
			UpdatedBy: customer.UpdatedBy,
		}
	}

	return response, nil
}

func (s *CustomerServiceServer) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	if req.Id == 0 {
		log.Errorf("id is required")
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.Name == "" {
		log.Errorf("name is required")
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Email == "" {
		log.Errorf("email is required")
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Phone == "" {
		log.Errorf("phone is required")
		return nil, status.Error(codes.InvalidArgument, "phone is required")
	}
	if req.UpdatedBy == "" {
		log.Errorf("updated_by is required")
		return nil, status.Error(codes.InvalidArgument, "updated_by is required")
	}

	// Get existing customer
	existingCustomer, err := s.customerService.GetCustomerByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("customer not found")
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		log.Errorf("failed to get customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to get customer")
	}

	// Update customer fields
	existingCustomer.Name = req.Name
	existingCustomer.Email = req.Email
	existingCustomer.Phone = req.Phone
	existingCustomer.UpdatedBy = req.UpdatedBy
	existingCustomer.UpdatedAt = time.Now()

	// Save changes
	if err := s.customerService.UpdateCustomer(existingCustomer); err != nil {
		log.Errorf("failed to update customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to update customer")
	}

	response := &pb.UpdateCustomerResponse{
		Customer: &pb.Customer{
			Id:        uint64(existingCustomer.ID),
			Name:      existingCustomer.Name,
			Email:     existingCustomer.Email,
			Phone:     existingCustomer.Phone,
			CreatedAt: timestamppb.New(existingCustomer.CreatedAt),
			CreatedBy: existingCustomer.CreatedBy,
			UpdatedAt: timestamppb.New(existingCustomer.UpdatedAt),
			UpdatedBy: existingCustomer.UpdatedBy,
		},
	}

	return response, nil
}

func (s *CustomerServiceServer) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	if req.Id == 0 {
		log.Errorf("id is required")
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	// Check if customer exists
	_, err := s.customerService.GetCustomerByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("customer not found")
			return nil, status.Error(codes.NotFound, "customer not found")
		}
		log.Errorf("failed to get customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to get customer")
	}

	// Delete customer
	if err := s.customerService.DeleteCustomer(uint(req.Id)); err != nil {
		log.Errorf("failed to delete customer: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete customer")
	}

	return &pb.DeleteCustomerResponse{}, nil
}
