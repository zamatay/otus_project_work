package models

type NetInfo struct {
	State        string
	RecvQ        int
	SendQ        int
	LocalAddress string
	PeerAddress  interface{}
}
