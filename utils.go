package goja

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/go-spirit/spirit/mail"
	"io/ioutil"
	"os"

	"github.com/go-spirit/spirit/worker/fbp/protocol"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
)

var (
	jsUtils = new(Utils)
)

type Utils struct {
}

func (p *Utils) CWD() string {
	dir, _ := os.Getwd()
	return dir
}

func (p *Utils) Base64Encode(x string) string {
	return base64.StdEncoding.EncodeToString([]byte(x))
}

func (p *Utils) Base64Decode(x string) string {
	data, _ := base64.StdEncoding.DecodeString(x)
	return string(data)
}

func (p *Utils) ReadBinaryFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithField("file", filename).Errorln(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

func (p *Utils) WriteBinaryFile(filename, base64Txt string) bool {
	data, err := base64.StdEncoding.DecodeString(base64Txt)
	if err != nil {
		logrus.WithField("file", filename).Errorln(err)
		return false
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		logrus.WithField("file", filename).Errorln(err)
		return false
	}
	return true
}

func (p *Utils) ReadTextFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithField("file", filename).Errorln(err)
		return ""
	}
	return string(data)
}

func (p *Utils) WriteTextFile(filename, txt string) bool {
	err := ioutil.WriteFile(filename, []byte(txt), 0644)
	if err != nil {
		logrus.WithField("file", filename).Errorln(err)
		return false
	}
	return true
}

func (p *Utils) MD5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}

func (p *Utils) IsNil(v interface{}) bool {
	return v == nil
}

func (p *Utils) NewUUID() string {
	return uuid.New()
}

func (p *Utils) SetFBPError(session mail.Session, namespace string, code int, message string) {
	session.Payload().Content().SetError(
		&protocol.Error{
			Code:        int64(code),
			Namespace:   namespace,
			Description: message,
		},
	)
}

func (p *Utils) NewFBPError(namespace string, code int, message string) error {
	return &protocol.Error{
		Code:        int64(code),
		Namespace:   namespace,
		Description: message,
	}
}

func (p *Utils) SetBody(session mail.Session, body interface{}) error {
	return session.Payload().Content().SetBody(body)
}
