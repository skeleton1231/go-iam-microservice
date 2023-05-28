// internal/pkg/store/mysql/itemimage.go

package mysql

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"gorm.io/gorm"
)

type itemImages struct {
	db *gorm.DB
}

func newItemImages(ds *datastore) *itemImages {
	return &itemImages{ds.db}
}

// Create creates a new item image.
func (i *itemImages) Create(ctx context.Context, itemImage *v1.ItemImage, opts metav1.CreateOptions) error {
	return i.db.Create(itemImage).Error
}

// Update updates an item image.
func (i *itemImages) Update(ctx context.Context, itemImage *v1.ItemImage, opts metav1.UpdateOptions) error {
	return i.db.Save(itemImage).Error
}

// Delete deletes an item image by the image ID.
func (i *itemImages) Delete(ctx context.Context, id uint64, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		i.db = i.db.Unscoped()
	}

	err := i.db.Where("id = ?", id).Delete(&v1.ItemImage{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Get returns an item image by the image ID.
func (i *itemImages) Get(ctx context.Context, id uint64, opts metav1.GetOptions) (*v1.ItemImage, error) {
	itemImage := &v1.ItemImage{}
	err := i.db.Where("id = ?", id).First(itemImage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUnknown, err.Error()) // code.ErrItemImageNotFound
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return itemImage, nil
}

// store/mysql/itemimage.go

func (s *itemImages) List(ctx context.Context, itemID uint64, opts metav1.ListOptions) ([]*v1.ItemImage, error) {
	var itemImages []*v1.ItemImage
	err := s.db.Where("item_id = ?", itemID).Find(&itemImages).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUnknown, err.Error()) // code.ErrItemImageNotFound
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return itemImages, nil
}
