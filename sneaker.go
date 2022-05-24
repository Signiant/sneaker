// Package sneaker provides an integrated system for securely storing sensitive
// information using Amazon's Simple Storage Service (S3) and Key Management
// Service (KMS).
package sneaker

import (
	"context"
	"fmt"
	fpath "path"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ObjectStorage is a sub-set of the capabilities of the S3 client.
type ObjectStorage interface {
	ListObjects(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) (*s3.ListObjectsOutput, error)
	DeleteObject(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

// KeyManagement is a sub-set of the capabilities of the KMS client.
type KeyManagement interface {
	GenerateDataKey(context.Context, *kms.GenerateDataKeyInput, ...func(*kms.Options)) (*kms.GenerateDataKeyOutput, error)
	Decrypt(context.Context, *kms.DecryptInput, ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

// A File is an encrypted secret, stored in S3.
type File struct {
	Path         string
	LastModified time.Time
	Size         int
	ETag         string
}

// A Manager allows you to manage files.
type Manager struct {
	Objects           ObjectStorage
	Envelope          Envelope
	KeyId             string
	EncryptionContext map[string]string
	Bucket, Prefix    string
}

func (m *Manager) context(path string) map[string]string {
	ctxt := make(map[string]string, len(m.EncryptionContext)+1)
	for k, v := range m.EncryptionContext {
		ctxt[k] = v
	}
	ctxt["Path"] = fmt.Sprintf("s3://%s/%s", m.Bucket, fpath.Join(m.Prefix, path))
	return ctxt
}
