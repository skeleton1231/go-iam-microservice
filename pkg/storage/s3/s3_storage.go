/*
这些数据通常保存在 AWS CLI 的共享凭证文件中，其默认位置是在您的用户主目录下的`.aws/credentials`文件。这个文件通常对应的路径如下：

- 在 Unix 或类 Unix 系统（包括 Linux 和 Mac OS X）上：`~/.aws/credentials`
- 在 Windows 上：`C:\Users\USERNAME\.aws\credentials`

其中 `USERNAME` 是你的用户名。这个文件不会自动被创建，所以如果你还没有创建过，你需要手动创建这个文件。

在这个文件中，你可以存储多个凭证集，每个集用方括号里的名称来标记，`default` 是默认的凭证集。每个凭证集需要包括 `aws_access_key_id` 和 `aws_secret_access_key`，如果你使用的是临时的安全凭证，还需要 `aws_session_token`。每个属性占据一行，属性名和值之间用等号分隔。例如：

```ini
[default]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY
aws_session_token = YOUR_SESSION_TOKEN

[anotherProfile]
aws_access_key_id = ANOTHER_ACCESS_KEY
aws_secret_access_key = ANOTHER_SECRET_KEY
```

请注意，在这个文件里的数据是敏感的，因为它们可以用来访问和修改你的 AWS 资源，所以你需要保证这个文件的安全，不要把这个文件的内容泄露出去，也不要把它放在不安全的地方。
*/
package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	Region  string
	Bucket  string
	Session *session.Session
}

func NewS3Storage(region, bucket string) (*S3Storage, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}

	return &S3Storage{
		Region:  region,
		Bucket:  bucket,
		Session: sess,
	}, nil
}

func (s *S3Storage) Upload(file string) (string, error) {
	uploader := s3manager.NewUploader(s.Session)

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
		Body:   f,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func (s *S3Storage) Delete(file string) error {
	svc := s3.New(s.Session)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			fmt.Fprintf(os.Stderr, "file does not exist\n")
		}
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(file),
	})

	return err
}
