package ipnet

type PNetError interface {
	IsPNetError() bool
}

func NewError(err string) error {
	return pnetErr("privnet: " + err)
}

func IsPNetError(err error) bool {
	v, ok := err.(PNetError)
	return ok && v.IsPNetError()
}

type pnetErr string

var _ PNetError = (PNetError)(pnetErr(""))

func (p pnetErr) Error() string {
	return string(p)
}

func (pnetErr) IsPNetError() bool {
	return true
}
