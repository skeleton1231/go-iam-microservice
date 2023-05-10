package item

import (
	v1svc "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/stretchr/testify/mock"
)

type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) Item() v1svc.ItemSrv {
	args := m.Called()
	return args.Get(0).(v1svc.ItemSrv)
}
