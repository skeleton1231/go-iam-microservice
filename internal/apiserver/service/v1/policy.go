package v1

import (
	"context"

	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-gin-restful-api-boilerplate/internal/apiserver/store"
)

// PolicySrv defines functions used to handle policy request.
type PolicySrv interface {
	Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error
	Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts metav1.GetOptions) (*v1.Policy, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error)
}

type policyService struct {
	store store.Factory
}

var _ PolicySrv = (*policyService)(nil)

func newPolicies(srv *service) *policyService {
	return &policyService{store: srv.store}
}

func (s *policyService) Create(ctx context.Context, policy *v1.Policy, opts metav1.CreateOptions) error {
	return nil
}

func (s *policyService) Update(ctx context.Context, policy *v1.Policy, opts metav1.UpdateOptions) error {
	return nil
}

func (s *policyService) Delete(ctx context.Context, username, name string, opts metav1.DeleteOptions) error {
	return nil
}

func (s *policyService) DeleteCollection(
	ctx context.Context,
	username string,
	names []string,
	opts metav1.DeleteOptions,
) error {

	return nil
}

func (s *policyService) Get(ctx context.Context, username, name string, opts metav1.GetOptions) (*v1.Policy, error) {
	return nil, nil
}

func (s *policyService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PolicyList, error) {
	return nil, nil
}
