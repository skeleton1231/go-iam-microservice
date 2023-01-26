package v1

import (
	"context"
	"fmt"

	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/apiserver/store"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/pkg/code"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ChangePassword(ctx context.Context, user *v1.User) error
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

// List returns user list in the storage. This function has a good performance.
func (u *userService) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	user, err := u.store.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	fmt.Println(user)

	return nil, nil
}

// ListWithBadPerformance returns user list in the storage. This function has a bad performance.
func (u *userService) ListWithBadPerformance(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {

	return nil, nil
}

func (u *userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {

	return nil
}

func (u *userService) DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error {

	return nil
}

func (u *userService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {

	return nil
}

func (u *userService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {

	return nil, nil
}

func (u *userService) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return nil
}

func (u *userService) ChangePassword(ctx context.Context, user *v1.User) error {
	return nil
}
