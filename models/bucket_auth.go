package models

import (
	"fmt"
)

type BucketAuth struct {
	Bucket       string
	Username     string
	Relationship string
}

var (
	RelationshipBucketOwner     = "Owner"
	RelationshipBucketConductor = "Conductor"
	RelationshipBucketViewer    = "Viewer"
	RelationshipBucketNone      = "None"
)

func BucketBindOwner(bucket *Bucket, user *User) (bool, error) {
	if bucket == nil || user == nil {
		return false, fmt.Errorf("输入格式错误")
	}
	db.Delete(&BucketAuth{Bucket: bucket.Name, Relationship: RelationshipBucketOwner})
	db.Create(&BucketAuth{Bucket: bucket.Name, Username: user.Name, Relationship: RelationshipBucketOwner})
	return true, nil
}

func BucketAddConductor(bucket *Bucket, user *User) (bool, error) {
	if bucket == nil || user == nil {
		return false, fmt.Errorf("输入格式错误")
	}
	db.Delete(&BucketAuth{Bucket: bucket.Name, Username: user.Name})
	db.Create(&BucketAuth{Bucket: bucket.Name, Username: user.Name, Relationship: RelationshipBucketConductor})
	return true, nil
}

func BucketAddViewer(bucket *Bucket, user *User) (bool, error) {
	if bucket == nil || user == nil {
		return false, fmt.Errorf("输入格式错误")
	}
	db.Delete(&BucketAuth{Bucket: bucket.Name, Username: user.Name})
	db.Create(&BucketAuth{Bucket: bucket.Name, Username: user.Name, Relationship: RelationshipBucketViewer})
	return true, nil
}

func GetUserBucketAuth(bucket *Bucket, user *User) string {
	if bucket == nil || user == nil {
		return RelationshipBucketNone
	}

	var bucketAuth BucketAuth
	err := db.Where(&BucketAuth{Bucket: bucket.Name, Username: user.Name}).First(&bucketAuth).Error

	if err != nil {
		return RelationshipBucketNone
	}

	return bucketAuth.Relationship
}
