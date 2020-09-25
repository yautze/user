package user

import (
	"user/app/domin/model"
	pb "user/app/interface/grpc/user/protobuf"

	"github.com/yautze/tools/st"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func userToModel(in *pb.User) (*model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, st.ErrorInvalidParameter
	}

	return &model.User{
		ID:       objectID,
		FBID:     in.GetFbid(),
		Password: in.GetPassword(),
		Key:      in.GetKey(),
		Info:     *infoToModel(in.GetInfo()),
		CreateAt: in.GetCreateAt(),
		UpdateAt: in.GetUpdateAt(),
		LoginAt:  in.GetLoginAt(),
	}, nil
}

func infoToModel(in *pb.Info) *model.Info {
	return &model.Info{
		Name:  in.GetName(),
		Phone: in.GetPhone(),
	}
}

func userToPb(in *model.User) *pb.User {
	return &pb.User{
		Id:       in.ID.Hex(),
		Fbid:     in.FBID,
		Info:     infoToPb(&in.Info),
		CreateAt: in.CreateAt,
		UpdateAt: in.UpdateAt,
		LoginAt:  in.LoginAt,
	}
}

func infoToPb(in *model.Info) *pb.Info {
	return &pb.Info{
		Name:  in.Name,
		Phone: in.Phone,
	}
}
