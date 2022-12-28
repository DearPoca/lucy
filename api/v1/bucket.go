package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"lucy/middleware/jwt"
	"lucy/models"
	"lucy/pkg/respond"
	"lucy/pkg/setting"

	"github.com/gin-gonic/gin"
)

const bucketDefaultCapacity = 512 * 1024 * 1024 * 1024

func CreateBucket(c *gin.Context) {
	tmp, ok := c.Get(jwt.KeyOfUsername)
	if !ok {
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	username := tmp.(string)

	var body struct {
		Bucket string `json:"bucket"`
	}

	if err := c.BindJSON(&body); err != nil {
		log.Printf("An error occurred while CreateBucket parse param: %s", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeBucketNameInvalidError))
		return
	}

	bucket, _ := models.GetBucket(body.Bucket)

	if bucket != nil {
		log.Printf("bucket [%s] existed", bucket.Name)
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeBucketNameInvalidError, "bucket existed"))
		return
	}

	ok, err := models.CreateBucket(body.Bucket, bucketDefaultCapacity)

	if !ok || err != nil {
		log.Printf("An error occurred while create bucket: %s", err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeBucketNameInvalidError))
		return
	}

	models.BucketBindOwner(body.Bucket, username)

	path := fmt.Sprintf("%s/%s", setting.AppSetting.AppRoot, body.Bucket)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Printf("Create directory [%s] failed: %s", path, err.Error())
		c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeUnknownError))
		return
	}
	c.JSON(http.StatusOK, respond.CreateRespond(respond.CodeSuccess))
}
