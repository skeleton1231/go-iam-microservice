// internal/pkg/service/v1/itemimage.go

package v1

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
)

type ItemImageSrv interface {
	Create(ctx context.Context, itemImage *model.ItemImage, opts metav1.CreateOptions) error
	Update(ctx context.Context, itemImage *model.ItemImage, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error
	Get(ctx context.Context, id int, opts metav1.GetOptions) (*model.ItemImage, error)
	List(ctx context.Context, itemID int, opts metav1.ListOptions) ([]*model.ItemImage, error) // 添加这行

}

type itemImageService struct {
	store store.Factory
}

var _ ItemImageSrv = (*itemImageService)(nil)

func newItemImages(srv *service) *itemImageService {
	return &itemImageService{store: srv.store}
}

func (s *itemImageService) Create(ctx context.Context, itemImage *model.ItemImage, opts metav1.CreateOptions) error {
	return s.store.ItemImage().Create(ctx, itemImage, opts)
}

func (s *itemImageService) Update(ctx context.Context, itemImage *model.ItemImage, opts metav1.UpdateOptions) error {
	return s.store.ItemImage().Update(ctx, itemImage, opts)
}

func (s *itemImageService) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	return s.store.ItemImage().Delete(ctx, id, opts)
}

func (s *itemImageService) Get(ctx context.Context, id int, opts metav1.GetOptions) (*model.ItemImage, error) {
	return s.store.ItemImage().Get(ctx, id, opts)
}

func (s *itemImageService) List(ctx context.Context, itemID int, opts metav1.ListOptions) ([]*model.ItemImage, error) {
	return s.store.ItemImage().List(ctx, itemID, opts)
}
