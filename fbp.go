package goja

import (
	"github.com/go-spirit/spirit/mail"

	"github.com/go-spirit/spirit/worker/fbp/protocol"
)

type fbp struct {
	session mail.Session
}

func (p *fbp) NewError(namespace string, code int, message string) error {
	return &protocol.Error{
		Code:        int64(code),
		Namespace:   namespace,
		Description: message,
	}
}

func (p *fbp) SetError(namespace string, code int, message string) {
	p.session.Payload().Content().SetError(
		&protocol.Error{
			Code:        int64(code),
			Namespace:   namespace,
			Description: message,
		},
	)
}

func (p *fbp) SetBody(body interface{}) error {
	return p.session.Payload().Content().SetBody(body)
}

func (p *fbp) Object() map[string]interface{} {
	v := make(map[string]interface{})
	err := p.session.Payload().Content().ToObject(&v)
	if err != nil {
		return nil
	}

	return v
}
