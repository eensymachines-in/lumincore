package lumincore

/*Auth microservice will use this as payload to send messages over on the socket, On the receiver side this message then can be unmarshalled to read its properties */
type ISockMessage interface {
	Pass() bool     // lets the client know if reg and auth have passed
	Serial() string // gets the device serial
}
type IAuthSockMsg interface {
	IsRegPass() bool
	IsAuthPass() bool
}
type ISchedSockMsg interface {
	JRStates() interface{} // json relay states
}

// Its this message that this microservice will shuttle thru the socket
type Message struct {
	Auth bool   `json:"auth"`
	Reg  bool   `json:"reg"`
	SID  string `json:"serial"`
}

func (m *Message) Pass() bool {
	return m.Auth && m.Reg // pass is ok only when both reg and auth are true
}
func (m *Message) Serial() string {
	return m.SID
}
func (m *Message) IsRegPass() bool {
	return m.Reg == true
}
func (m *Message) IsAuthPass() bool {
	return m.Auth == true
}

// SchedsMessage : just message with extra field for schedules, extension on the Message
type SchedsMessage struct {
	*Message
	Scheds interface{} `json:"scheds"`
}

func (sm *SchedsMessage) JRStates() interface{} {
	return sm.Scheds
}
