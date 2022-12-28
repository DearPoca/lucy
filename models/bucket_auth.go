package models

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

func BucketBindOwner(bucket string, user string) (bool, error) {
	db.Delete(&BucketAuth{Bucket: bucket, Relationship: RelationshipBucketOwner})
	db.Create(&BucketAuth{Bucket: bucket, Username: user, Relationship: RelationshipBucketOwner})
	return true, nil
}

func BucketAddConductor(bucket string, user string) (bool, error) {
	db.Delete(&BucketAuth{Bucket: bucket, Username: user})
	db.Create(&BucketAuth{Bucket: bucket, Username: user, Relationship: RelationshipBucketConductor})
	return true, nil
}

func BucketAddViewer(bucket string, user string) (bool, error) {
	db.Delete(&BucketAuth{Bucket: bucket, Username: user})
	db.Create(&BucketAuth{Bucket: bucket, Username: user, Relationship: RelationshipBucketViewer})
	return true, nil
}

func GetUserBucketAuth(bucket string, user string) string {
	var bucketAuth BucketAuth
	err := db.Where(&BucketAuth{Bucket: bucket, Username: user}).First(&bucketAuth).Error

	if err != nil {
		return RelationshipBucketNone
	}

	return bucketAuth.Relationship
}

func GetBucketAuthsRelatedUser(username string) []*BucketAuth {
	bucketAuth := make([]*BucketAuth, 0)
	db.Where(&BucketAuth{Username: username}).Find(&bucketAuth)
	return bucketAuth
}
