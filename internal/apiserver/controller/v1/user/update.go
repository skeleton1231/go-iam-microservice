package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/marmotedu/api/apiserver/v1"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/log"
)

// Update update a user info by the user identifier.
func (u *UserController) Update(c *gin.Context) {
	log.L(c).Info("update user function called.")

	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	user, err := u.srv.Users().Get(c, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.Extend = r.Extend

	if errs := user.ValidateUpdate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	// Save changed fields.
	if err := u.srv.Users().Update(c, user, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
