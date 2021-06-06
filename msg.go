package lumincore

/*Reg authentication with luminapi server will emit messages of this datashape. Since we need this in the luminapi as well as autolumin we capture this as common package here
Data objects here are accessible over interfaces */
type IMessage interface {
	Pass() bool     // lets the client know if reg and auth have passed
	Serial() string // gets the device serial
}
type IAuthMessage interface {
	IsRegPass() bool
	IsAuthPass() bool
}
type ISchedMessage interface {
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
