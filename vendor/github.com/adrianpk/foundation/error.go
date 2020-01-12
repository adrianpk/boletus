package kabestan

type (
	Err struct {
		msgID string
		Err   error
	}
)

func NewErr(msgID string, err error) Err {
	return Err{
		msgID: msgID,
		Err:   err,
	}
}

func (e Err) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.msgID
}

func (e Err) MsgID() string {
	return e.msgID
}
