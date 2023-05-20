package mysql

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"gorm.io/gorm"

	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
)

type itemAttributes struct {
	db *gorm.DB
}

func newItemAttributes(ds *datastore) *itemAttributes {
	return &itemAttributes{ds.db}
}

// Create creates a new item attribute.
func (i *itemAttributes) Create(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.CreateOptions) error {
	return i.db.Create(itemAttribute).Error
}

// Update updates an item attribute.
func (i *itemAttributes) Update(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.UpdateOptions) error {
	return i.db.Save(itemAttribute).Error
}

// Delete deletes an item attribute by the attribute ID.
func (i *itemAttributes) Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		i.db = i.db.Unscoped()
	}

	err := i.db.Where("id = ?", id).Delete(&v1.ItemAttributes{}).Error
	if err != nil {
		return err
	}

	return nil
}

// Get returns an item attribute by the attribute identifier.
func (i *itemAttributes) Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.ItemAttributes, error) {
	itemAttribute := &v1.ItemAttributes{}
	err := i.db.Where("id = ?", id).First(itemAttribute).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUnknown, err.Error()) // code.ErrItemAttributeNotFound
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return itemAttribute, nil
}

// FindByAttributes returns a list of items that match the given item attributes.
func (i *itemAttributes) FindByAttributes(ctx context.Context, attributes *v1.ItemAttributes, opts metav1.ListOptions) (*v1.ItemList, error) {
	// This function should be implemented based on your database and query language.
	// Here is a very basic example:

	items := &v1.ItemList{}
	err := i.db.Where(attributes).Find(items).Error
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return items, nil
}
