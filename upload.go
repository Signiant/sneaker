package sneaker

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	fpath "path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload encrypts the given secret with a KMS data key and uploads it to S3.
func (m *Manager) Upload(path string, r io.Reader) error {
	plaintext, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	ciphertext, err := m.Envelope.Seal(m.KeyId, m.context(path), plaintext)
	if err != nil {
		return err
	}

	if _, err := m.Objects.PutObject(context.TODO(),
		&s3.PutObjectInput{
			ContentLength: int64(len(ciphertext)),
			ContentType:   aws.String(contentType),
			Bucket:        aws.String(m.Bucket),
			Key:           aws.String(fpath.Join(m.Prefix, path)),
			Body:          bytes.NewReader(ciphertext),
		},
	); err != nil {
		return err
	}
	return nil
}

const (
	contentType = "application/octet-stream"
)
