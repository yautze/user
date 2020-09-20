package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
	"user/app/domin/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

// HASHKEY -
const HASHKEY = "YauTz"

// UserService -
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService -
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// DuplicateByFBID - if duplicate return true
func (u *UserService) DuplicateByFBID(fbID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := u.userRepo.GetByFBID(ctx, fbID)
	if err != nil {
		if mongo.ErrNoDocuments == err { // no data in db
			return false, nil
		}
		return false, err
	}

	if res == nil { // Not in db
		return false, nil
	}

	return true, nil
}

// HashKey -
func (u *UserService) HashKey(fbID, password string) string {
	// bind hash
	h := hmac.New(sha256.New, []byte(HASHKEY))

	// hash fbid & password
	data := strings.TrimSpace(fbID) + strings.TrimSpace(password)

	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
