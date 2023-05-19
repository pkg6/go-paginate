package adapter

import (
	"github.com/pkg6/go-paginate"
	"xorm.io/xorm"
)

type XORM struct {
	session *xorm.Session
}

func XORMAdapter(session *xorm.Session) paginate.IAdapter {
	return &XORM{session: session}
}

func (x XORM) Length() (int64, error) {
	return x.session.Count()
}

func (x XORM) Slice(offset, length int64, data any) error {
	return x.session.Limit(int(length), int(offset)).Find(data)
}
