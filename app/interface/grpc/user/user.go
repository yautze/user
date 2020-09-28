package user

import (
	"context"
	pb "user/app/interface/grpc/user/protobuf"
	"user/app/usecase"

	"github.com/yautze/tools/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

// Apply -
func Apply(server *grpc.Server, service *service) {
	pb.RegisterUserServiceServer(server, service)
}

type service struct {
	userUsecase usecase.UserUsecase
}

// New -
func New(userUsecase usecase.UserUsecase) *service {
	return &service{
		userUsecase: userUsecase,
	}
}

// Create -
func (s service) Create(ctx context.Context, in *pb.User) (*pb.Empty, error) {
	// check fbID & password not empty
	if in.GetFbid() == "" || in.GetPassword() == "" {
		return nil, st.ErrorInvalidParameter
	}

	// convert
	data, err := userToModel(in)
	if err != nil {
		return nil, err
	}

	// call userUsecase Create
	if err := s.userUsecase.Create(data); err != nil {
		return nil, st.ErrorDatabaseCreateFailed
	}

	return &pb.Empty{}, nil
}

// GetByID -
func (s *service) GetByID(ctx context.Context, in *pb.GetByIDRequest) (*pb.User, error) {
	// check id to objectID
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, st.ErrorInvalidParameter
	}

	// call userUsecase GetByID
	user, err := s.userUsecase.GetByID(objectID)
	if err != nil {
		return nil, st.ErrorDataNotFound
	}

	return userToPb(user), nil
}

// GetByFBIDAndPassword -
func (s *service) GetByFBIDAndPassword(ctx context.Context, in *pb.GetByFBIDAndPasswordRequest) (*pb.User, error) {
	// check fbID & password not empty
	if in.GetFbid() == "" || in.GetPassword() == "" {
		return nil, st.ErrorInvalidParameter
	}

	// call userUsecase GetByFBIDAndPassword
	user, err := s.userUsecase.GetByFBIDAndPassword(in.GetFbid(), in.GetPassword())
	if err != nil {
		return nil, st.ErrorDataNotFound
	}

	return userToPb(user), nil
}

// UpdateInfo -
func (s *service) UpdateInfo(ctx context.Context, in *pb.UpdateInfoRequest) (*pb.Empty, error) {
	// check id to objectID
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, st.ErrorInvalidParameter
	}

	// call userUsecase UpdateInfo
	if err := s.userUsecase.UpdateInfo(objectID, infoToModel(in.GetInfo())); err != nil {
		return nil, st.ErrorDataNotFound
	}

	return &pb.Empty{}, nil
}

// UpdatePassword -
func (s *service) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*pb.Empty, error) {
	// check id to objectID
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, st.ErrorInvalidParameter
	}

	// call userUsecase UpdatePassword
	if err := s.userUsecase.UpdatePassword(objectID, in.GetFbid(), in.GetPassword()); err != nil {
		return nil, st.ErrorDataNotFound
	}

	return &pb.Empty{}, nil
}

// DeleteByID -
func (s *service) DeleteByID(ctx context.Context, in *pb.DeleteByIDRequest) (*pb.Empty, error) {
	// check id to objectID
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, st.ErrorInvalidParameter
	}

	// call userUsecase UpdateInfo
	if err := s.userUsecase.DeleteByID(objectID); err != nil {
		return nil, st.ErrorDataNotFound
	}

	return &pb.Empty{}, nil
}
