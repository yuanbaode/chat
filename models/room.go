package models

import "sync"

type Room struct {
	Id      int64          `orm:"pk" json:"id"`
	Name    string         `orm:"size(20)" json:"name"`
	Clients *SyncClientMap `orm:"-" json:"members"`
	In      chan *Message  `orm:"-" json:"-"`
}
type RoomSet struct {
	Data map[int64]*Room
	Lock sync.RWMutex
}

func NewRoom(name string) *Room {
	room := new(Room)
	room.Name = name
	room.Clients = NewSyncClientMap()
	room.In = make(chan *Message, 100)
	go func() {
		for {
			select {
			case msg := <-room.In:
				room.Clients.Lock()
				for _, v := range room.Clients.Data {
					v.Out <- msg
				}
				room.Clients.UnLock()
			}
		}

	}()
	return room
}
