package item

import (
	"path/filepath"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/item/v1/model"
	srvv1 "github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/service/v1"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/apiserver/store"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/code"
	"github.com/skeleton1231/go-iam-ecommerce-microservice/internal/pkg/options"
	storage "github.com/skeleton1231/go-iam-ecommerce-microservice/pkg/file_storage"
)

type itemImageController struct {
	srv     srvv1.Service
	storage storage.FileStorage
}

func NewItemImageController(store store.Factory, storageOpts *options.FileStorageOptions) (*itemImageController, error) {

	fs, err := storage.GetFileStorageFactoryOr(storageOpts)
	if err != nil {
		return nil, err
	}

	return &itemImageController{
		srv:     srvv1.NewService(store),
		storage: fs,
	}, nil
}

func (ctrl *itemImageController) Create(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil)
		return
	}

	files := form.File["upload[]"]

	for _, file := range files {
		// Save file to local storage
		localPath := filepath.Join("/tmp", file.Filename)
		err := c.SaveUploadedFile(file, localPath)
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
			return
		}

		// Upload file to S3 and get file URL
		url, err := ctrl.storage.Upload(localPath)
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil) // code.ErrStorage
			return
		}

		node, err := snowflake.NewNode(1)
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrEncrypt, err.Error()), nil)

			return
		}

		imageID := uint64(node.Generate().Int64())

		// Create a new item image
		image := &model.ItemImage{
			ID:       int(imageID),
			ItemID:   itemID,
			ImageURL: url,
		}
		err = ctrl.srv.ItemImage().Create(c, image, metav1.CreateOptions{})
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
			return
		}
	}

	core.WriteResponse(c, nil, gin.H{
		"message": "Files uploaded successfully",
	})
}

// Implement other methods (Update, Delete, Get, List) with similar structure
func (ctrl *itemImageController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil)
		return
	}

	// Handle file update logic here
	file, err := c.FormFile("upload")
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	// Save file to local storage
	localPath := filepath.Join("/tmp", file.Filename)
	err = c.SaveUploadedFile(file, localPath)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	// Upload file to S3 and get file URL
	url, err := ctrl.storage.Upload(localPath)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil) // code.ErrStorage
		return
	}

	image := &model.ItemImage{
		ID:       id,
		ImageURL: url,
	}

	err = ctrl.srv.ItemImage().Update(c, image, metav1.UpdateOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"message": "File updated successfully",
	})
}

func (ctrl *itemImageController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil)
		return
	}

	image, err := ctrl.srv.ItemImage().Get(c, id, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	err = ctrl.storage.Delete(image.ImageURL)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil) // code.ErrStorage
		return
	}

	err = ctrl.srv.ItemImage().Delete(c, id, metav1.DeleteOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, gin.H{
		"message": "File deleted successfully",
	})
}

func (ctrl *itemImageController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil)
		return
	}

	image, err := ctrl.srv.ItemImage().Get(c, id, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, image)
}

func (ctrl *itemImageController) List(c *gin.Context) {
	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, err.Error()), nil)
		return
	}

	images, err := ctrl.srv.ItemImage().List(c, itemID, metav1.ListOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, images)
}
