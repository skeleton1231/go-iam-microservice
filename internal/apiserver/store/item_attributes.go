package store

import (
	"context"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	v1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
)

// ItemAttributeSrv defines functions used to handle item attribute requests.
type ItemAttributesStore interface {
	Create(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.CreateOptions) error
	Update(ctx context.Context, itemAttribute *v1.ItemAttributes, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, id int, opts metav1.DeleteOptions) error
	Get(ctx context.Context, id int, opts metav1.GetOptions) (*v1.ItemAttributes, error)
}
