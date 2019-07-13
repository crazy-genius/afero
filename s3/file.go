package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"sync"
)

type uploaderProvider func () (*s3manager.Uploader, error)
type downloaderProvider func() (*s3manager.Downloader, error)
type fileData []byte

type File struct {
	uploaderProvider uploaderProvider
	downloaderProvider downloaderProvider
	data fileData
	input *s3manager.UploadInput
}

func (fd fileData) WriteAt(p []byte, off int64) (n int, err error) {
	mut := sync.Mutex{}
	mut.Lock()
	defer mut.Unlock()
	buf := bytes.NewBuffer(fd)

	if buf.Len() <= int(off) + len(p) {
		buf.Grow(int(off) + len(p))
	}

	return buf.Write(p)
}

func CreateFile(uploaderProvider uploaderProvider, downloaderProvider downloaderProvider, input *s3manager.UploadInput) File  {
	return File{
		uploaderProvider: uploaderProvider,
		downloaderProvider: downloaderProvider,
		input:    input,
	}
}

func (f File) Close() error {
	 uploader, err := f.uploaderProvider()

	 if err != nil {
	 	return err
	 }

	 if _, err := uploader.Upload(f.input); err != nil {
	 	return err
	 }

	 return nil
}

func (f File) Read(p []byte) (n int, err error) {
	downloader, err := f.downloaderProvider()

	if err != nil {
		return 0, err
	}

	readLen, err := downloader.Download(f.data, '')

	if err != nil {
		return 0, err
	}

	return int(readLen), nil
}

func (f File) ReadAt(p []byte, off int64) (n int, err error) {
	panic("implement me")
}

func (f File) Seek(offset int64, whence int) (int64, error) {
	panic("implement me")
}

func (f File) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (f File) WriteAt(p []byte, off int64) (n int, err error) {
	panic("implement me")
}

func (f File) Name() string {
	panic("implement me")
}

func (f File) Readdir(count int) ([]os.FileInfo, error) {
	panic("implement me")
}

func (f File) Readdirnames(n int) ([]string, error) {
	panic("implement me")
}

func (f File) Stat() (os.FileInfo, error) {
	panic("implement me")
}

func (f File) Sync() error {
	panic("implement me")
}

func (f File) Truncate(size int64) error {
	panic("implement me")
}

func (f File) WriteString(s string) (ret int, err error) {
	panic("implement me")
}


