package v1

import (
	"context"
	"sync"

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
	users, err := u.store.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	var m sync.Map

	for _, user := range users.Items {
		wg.Add(1)

		go func(user *v1.User) {
			defer wg.Done()

			policies, err := u.store.Policies().List(ctx, user.Name, metav1.ListOptions{}) // metav1.ListOptions{}
			if err != nil {
				errChan <- errors.WithCode(code.ErrDatabase, err.Error())

				return
			}

			m.Store(
				user.ID,
				&v1.User{
					ObjectMeta: metav1.ObjectMeta{
						ID:         user.ID,
						InstanceID: user.InstanceID,
						Name:       user.Name,
						Extend:     user.Extend,
						CreatedAt:  user.CreatedAt,
						UpdatedAt:  user.UpdatedAt,
					}, Nickname: user.Nickname,
					Email:       user.Email,
					Phone:       user.Phone,
					TotalPolicy: policies.TotalCount,
					LoginedAt:   user.LoginedAt,
				})
		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*v1.User, 0, len(users.Items))
	for _, user := range users.Items {
		info, _ := m.Load(user.ID)
		infos = append(infos, info.(*v1.User))
	}

	log.L(ctx).Debugf("get %d users from backend storage.", len(infos))

	return &v1.UserList{ListMeta: users.ListMeta, Items: infos}, nil
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
