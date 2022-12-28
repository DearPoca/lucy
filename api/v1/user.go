package v1

import (
	"net/http"

	"lucy/middleware/jwt"
	"lucy/models"
	"lucy/pkg/respond"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	tmp, ok := c.Get(jwt.KeyOfUsername)
	if !ok {
		c.JSON(http.StatusOK, respond.ResUnknownError())
		return
	}
	username := tmp.(string)

	type bucket struct {
		Name string `json:"name"`
		Auth string `json:"auth"`
	}
	u := struct {
		Name    string   `json:"name"`
		Buckets []bucket `json:"buckets"`
	}{Name: username}

	bucketAuths := models.GetBucketAuthsRelatedUser(username)
	for i, _ := range bucketAuths {
		u.Buckets = append(u.Buckets, bucket{Name: bucketAuths[i].Bucket, Auth: bucketAuths[i].Relationship})
	}

	c.JSON(http.StatusOK, respond.ResSuccess(u))
}
