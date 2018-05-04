package archive

type Archive interface {
	RemoveFile(filePath string, bucketName string) error
	BackupFile(bucketName string, destinationFolder string, filePath string) error
	UploadFile(bucketName string, destinationFolder string, filePath string) error
	CheckFileIntergrity(bucketName string, destinationFolder string, filePath string) (bool, error)
	GetAuthDataPath() string
	GetReserveDataBucketName() string
	GetStatDataBucketName() string
	GetPriceAnalyticPath() string
}
