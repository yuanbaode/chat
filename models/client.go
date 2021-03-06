package models

import (
	"github.com/gorilla/websocket"
	"sync"
	"encoding/json"
	"mychatroom/log"
)

type Client struct {
	User User
	Out  chan *Message
	Conn *websocket.Conn
}

func NewClient(user User, conn *websocket.Conn) *Client {
	client := new(Client)
	client.User = user
	client.Out = make(chan *Message, 100)
	client.Conn = conn
	go func() {
		for {
			select {
			case msg := <-client.Out:
				data, err := json.Marshal(msg)
				if err != nil {
					log.Error("WriteMessage  marshal err:%s", err.Error())
				}
				client.Conn.WriteMessage(websocket.TextMessage, data)
			}
		}
	}()
	return client
}

type SyncClientMap struct {
	mux  *sync.RWMutex
	Data map[string]*Client
}

func NewSyncClientMap() *SyncClientMap {
	data := make(map[string]*Client, 100)
	return &SyncClientMap{
		mux:  &sync.RWMutex{},
		Data: data,
	}
}
func (m *SyncClientMap) SET(key string, value *Client) {
	m.mux.Lock()
	if old, ok := m.Data[key]; ok && old != nil {
		old.Conn.Close()
	}
	if value == nil {
		delete(m.Data, key)
	} else {

		m.Data[key] = value
	}
	//m.Data[key] = value
	m.mux.Unlock()
}

func (m *SyncClientMap) GET(key string) (value *Client, ok bool) {
	m.mux.Lock()
	value, ok = m.Data[key]
	m.mux.Unlock()
	return
}

func (m *SyncClientMap) Lock() {
	m.mux.Lock()
}
func (m *SyncClientMap) UnLock() {
	m.mux.Unlock()
}
