package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/controller/v1/item"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/controller/v1/policy"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/controller/v1/secret"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/controller/v1/user"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/options"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store/mysql"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/middleware"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/middleware/auth"
	"github.com/spf13/viper"

	// custom gin validators.
	_ "github.com/marmotedu/iam/pkg/validator"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/log"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	// Refresh time can be longer than token timeout
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})

	storeIns, _ := mysql.GetMySQLFactoryOr(nil)

	// _, err := storage.GetFileStorageFactoryOr(opts.FileStorageOptions)
	// if err != nil {
	// 	log.Fatalf("Failed to create file storage instance: %v", err)
	// }

	v1 := g.Group("/v1")
	{
		// user RESTful resource
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)

			userv1.POST("", userController.Create)
			userv1.Use(auto.AuthFunc(), middleware.Validation())
			// v1.PUT("/find_password", userController.FindPassword)
			userv1.DELETE("", userController.DeleteCollection) // admin api
			userv1.DELETE(":name", userController.Delete)      // admin api
			userv1.PUT(":name/change-password", userController.ChangePassword)
			userv1.PUT(":name", userController.Update)
			userv1.GET("", userController.List)
			userv1.GET(":name", userController.Get) // admin api
		}

		v1.Use(auto.AuthFunc())

		// policy RESTful resource
		policyv1 := v1.Group("/policies", middleware.Publish())
		{
			policyController := policy.NewPolicyController(storeIns)

			policyv1.POST("", policyController.Create)
			policyv1.DELETE("", policyController.DeleteCollection)
			policyv1.DELETE(":name", policyController.Delete)
			policyv1.PUT(":name", policyController.Update)
			policyv1.GET("", policyController.List)
			policyv1.GET(":name", policyController.Get)
		}

		// secret RESTful resource
		secretv1 := v1.Group("/secrets", middleware.Publish())
		{
			secretController := secret.NewSecretController(storeIns)

			secretv1.POST("", secretController.Create)
			secretv1.DELETE(":name", secretController.Delete)
			secretv1.PUT(":name", secretController.Update)
			secretv1.GET("", secretController.List)
			secretv1.GET(":name", secretController.Get)
		}
	}

	v2 := g.Group("/v2")
	{
		// item RESTful resource
		itemv2 := v2.Group("/items", middleware.Publish())
		{
			itemController := item.NewItemController(storeIns)
			itemv2.Use(auto.AuthFunc())
			itemv2.POST("", itemController.Create)
			itemv2.DELETE(":itemID", itemController.Delete)
			itemv2.PUT(":itemID", itemController.Update)
			itemv2.GET(":itemID", itemController.Get)
			itemv2.GET("", itemController.List)

			// Adding new route for item images
			fileStorageOptions := options.NewOptions().FileStorageOptions
			viper.UnmarshalKey("fileStorage", &fileStorageOptions)
			itemImageController, _ := item.NewItemImageController(storeIns, fileStorageOptions)
			itemv2.GET(":itemID/images", itemImageController.List)
		}

		itemAtrriV2 := v2.Group("/itemAttris", middleware.Publish())
		{
			itemAttriController := item.NewItemAttributesController(storeIns)
			itemAtrriV2.Use(auto.AuthFunc())
			itemAtrriV2.POST("", itemAttriController.Create)
			itemAtrriV2.PUT(":attributeID", itemAttriController.Update)
			itemAtrriV2.GET(":attributeID", itemAttriController.Get)
			itemAtrriV2.DELETE(":attributeID", itemAttriController.Delete)
		}

		// itemImage RESTful resource
		itemImageV2 := v2.Group("/itemImages", middleware.Publish())
		{

			fileStorageOptions := options.NewOptions().FileStorageOptions
			err := viper.UnmarshalKey("fileStorage", &fileStorageOptions)
			if err != nil {
				log.Errorf("Unable to decode into struct, %v", err)
			}

			itemImageController, err := item.NewItemImageController(storeIns, fileStorageOptions)
			if err != nil {
				log.Errorf("Failed to create itemImageController: %v", err) // Handle this error according to your project's logging strategy
			}

			itemImageV2.Use(auto.AuthFunc())
			itemImageV2.POST("", itemImageController.Create)
			itemImageV2.PUT(":id", itemImageController.Update)
			itemImageV2.GET(":id", itemImageController.Get)
			itemImageV2.DELETE(":id", itemImageController.Delete)
			// itemImageV2.GET("item/:item_id", itemImageController.List)
			log.Info("item images initialized")
		}

	}

	return g
}
