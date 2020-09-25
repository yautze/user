package usecase

import (
	"context"
	"time"
	"user/app/domin/model"
	"user/app/domin/repository"
	"user/app/domin/service"

	"github.com/yautze/tools/st"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecase -
type UserUsecase interface {
	Create(in *model.User) error
	GetByID(id primitive.ObjectID) (*model.User, error)
	GetByFBIDAndPassword(fbID, password string) (*model.User, error)
	UpdateInfo(id primitive.ObjectID, info *model.Info) error
	UpdatePassword(id primitive.ObjectID, fbID, password string) error
	DeleteByID(id primitive.ObjectID) error
}

type userUsecase struct {
	userRepository repository.UserRepository
	userService    *service.UserService
}

// NewUserUsecase -
func NewUserUsecase(userRepository repository.UserRepository, userService *service.UserService) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
		userService:    userService,
	}
}

// Create -
func (u *userUsecase) Create(in *model.User) error {
	// check fbID duplicate
	duplicate, err := u.userService.DuplicateByFBID(in.FBID)
	if err != nil {
		return err
	}

	if duplicate {
		return st.ErrorDataIsExists
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	in.Key = u.userService.HashKey(in.FBID, in.Password)
	// create at time
	in.CreateAt = time.Now().Unix()

	// call userRepository - create
	if err := u.userRepository.Create(ctx, in); err != nil {
		return err
	}

	return nil
}

// GetByID -
func (u *userUsecase) GetByID(id primitive.ObjectID) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// call  userRepository - GetByID
	res, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetByFBIDAndPassword -
func (u *userUsecase) GetByFBIDAndPassword(fbID, password string) (*model.User, error) {
	key := u.userService.HashKey(fbID, password)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// call  userRepository - GetByKey
	res, err := u.userRepository.GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateInfo -
func (u *userUsecase) UpdateInfo(id primitive.ObjectID, info *model.Info) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"info.name":  info.Name,
			"info.phone": info.Phone,
			"udpateAt":   time.Now().Unix(),
		},
	}

	// call  userRepository - Update
	if err := u.userRepository.Update(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

// UpdatePassword -
func (u *userUsecase) UpdatePassword(id primitive.ObjectID, fbID, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// call userRepository - GetByID for check fbID
	user, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	// check user.fbID equal fbID
	if user.FBID != fbID {
		return st.ErrorDatabaseUpdateFailed
	}

	key := u.userService.HashKey(fbID, password)

	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"key":      key,
			"udpateAt": time.Now().Unix(),
		},
	}

	// call  userRepository - DeleteByID
	if err := u.userRepository.Update(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

// DeleteByID -
func (u *userUsecase) DeleteByID(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// call  userRepository - DeleteByID
	if err := u.userRepository.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
