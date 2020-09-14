package domain

import (
	"context"
	"reflect"
	"testing"

	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/patrickmn/go-cache"
)

func TestUserService_GetByID(t *testing.T) {
	type fields struct {
		repo  UserRepository
		cache *cache.Cache
	}
	type args struct {
		ctx context.Context
		r   *pb.RequestUser
	}

	repository := NewInMemoryRepository()
	service := NewUserService(repository)

	f := fields{
		repo:  service.repo,
		cache: service.cache,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "WhenGetByID_1",
			fields: f,
			args: args{
				ctx: context.Background(),
				r:   &pb.RequestUser{Id: "1"},
			},
			want:    "1",
			wantErr: false,
		},
		{
			name:   "WhenGetByID_3",
			fields: f,
			args: args{
				ctx: context.Background(),
				r:   &pb.RequestUser{Id: "3"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repo:  tt.fields.repo,
				cache: tt.fields.cache,
			}
			got, err := s.GetByID(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetId(), tt.want) {
				t.Errorf("UserService.GetByID() = %v, want %v", got.GetId(), tt.want)
			}
		})
	}
}
