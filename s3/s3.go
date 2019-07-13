package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/afero"
	"os"
	"time"
)

type Fs struct {
	bucket string
}

func (f Fs) makeUploader() *s3manager.Uploader {
	sess := session.Must(session.NewSession())
	return s3manager.NewUploader(sess)
}

func (f Fs) makeDownloader() *s3manager.Downloader {
	sess := session.Must(session.NewSession())
	return s3manager.NewDownloader(sess)
}

func (f Fs) Create(name string) (afero.File, error) {
	uploadInput := &s3manager.UploadInput{Bucket: aws.String(f.bucket), Key: aws.String(name)}

	return CreateFile(f.makeUploader(), uploadInput), nil
}

func (f Fs) Mkdir(name string, perm os.FileMode) error {
	panic("implement me")
}

func (f Fs) MkdirAll(path string, perm os.FileMode) error {
	panic("implement me")
}

func (f Fs) Open(name string) (afero.File, error) {
	panic("implement me")
}

func (f Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	panic("implement me")
}

func (f Fs) Remove(name string) error {
	panic("implement me")
}

func (f Fs) RemoveAll(path string) error {
	panic("implement me")
}

func (f Fs) Rename(oldname, newname string) error {
	panic("implement me")
}

func (f Fs) Stat(name string) (os.FileInfo, error) {
	panic("implement me")
}

func (f Fs) Name() string {
	panic("implement me")
}

func (f Fs) Chmod(name string, mode os.FileMode) error {
	panic("implement me")
}

func (f Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("implement me")
}

func New(bucket string) afero.Fs {
	return &Fs{bucket: bucket}
}

