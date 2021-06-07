package lumincore

import "github.com/eensymachines-in/scheduling"

/*Auth microservice will use this as payload to send messages over on the socket, On the receiver side this message then can be unmarshalled to read its properties */
type ISockMessage interface {
	Pass() bool     // lets the client know if reg and auth have passed
	Serial() string // gets the device serial
}
type IAuthSockMsg interface {
	IsRegPass() bool
	IsAuthPass() bool
	SetAuth(v bool)
}
type ISchedSockMsg interface {
	JRStates() interface{} // json relay states
}

// Its this message that this microservice will shuttle thru the socket
type SockMessage struct {
	Auth bool   `json:"auth"`
	Reg  bool   `json:"reg"`
	SID  string `json:"serial"`
}

// SetAuth : sets the auth status for the message, this is a setter property for SockMessage
func (m *SockMessage) SetAuth(v bool) {
	m.Auth = v
}
func (m *SockMessage) Pass() bool {
	return m.Auth && m.Reg // pass is ok only when both reg and auth are true
}
func (m *SockMessage) Serial() string {
	return m.SID
}
func (m *SockMessage) IsRegPass() bool {
	return m.Reg == true
}
func (m *SockMessage) IsAuthPass() bool {
	return m.Auth == true
}

// SchedSockMessage : just message with extra field for schedules, extension on the SockMessage
type SchedSockMessage struct {
	*SockMessage
	Scheds []scheduling.JSONRelayState `json:"scheds"`
}

func (sm *SchedSockMessage) JRStates() interface{} {
	return sm.Scheds
}
