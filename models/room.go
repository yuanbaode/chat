package models

import (
	"sync"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"math/rand"
	"time"
)

var TEXTARRAY []string
var AmountArray []int64
var UserArray []string

func init() {
	TEXTARRAY = make([]string, 0, 100)
	for i := 0; i < 28; i++ {
		TEXTARRAY = append(TEXTARRAY, strconv.Itoa(i))
	}
	array := []string{
		"大", "双", "单", "极小", "极大", "大双", "大单", "小双", "小单",
	}
	TEXTARRAY = append(TEXTARRAY, array ...)
	AmountArray = make([]int64, 0, 100)
	for i := 1; i < 11; i++ {
		AmountArray = append(AmountArray, int64(i*5))
	}
	UserArray = []string{
		"酒与心事", "淡墨", "酒伴久身", "无力自拔", "时间路人", "無心", "零度°", "凉薄之人", "空心", "社会不惯人", "垂涎自由", "执酒共酌", "寄与心", "挽歌少年", "暮年", "孤独酒馆", "痴骨ら", "同归", "不抽烟C", "曾天真现成熟", "我怕的是人心", "淡然一笑", "岁月之沉淀", "单身泽", "孤寂芳心", "笑魇醉看", "喜喜", "不羁", "余生没有北方", "暴雨柴舟", "白衣决绝提剑", "耳边情话",
	}
}

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
	go func() {

		for {
			sleepTime := getRandNum(40)
			time.Sleep(time.Duration(sleepTime) * time.Second)
			RandomData(room)
		}

	}()
	return room
}

func RandomData(room *Room) {
	apiUrl := `http://test.lxh.wiki/api/lottery/latest/period`
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	type respData struct {
		Data int
	}
	data := &respData{}
	if err = json.Unmarshal(body, data); err != nil {
		return
	}
	user := User{Name: getNickName()}
	room.In <- &Message{GROUPCHAT, user, "第" + strconv.Itoa(data.Data) + `期：
` + getText() + "/" + strconv.Itoa(getAmount())}
}
func getRandNum(range_ int) int {
	return rand.Intn(range_)
}
func getNickName() string {
	l := len(UserArray)
	index := getRandNum(l - 1)
	return UserArray[index]
}
func getAmount() int {
	l := len(AmountArray)
	index := getRandNum(l - 1)
	return int(AmountArray[index])
}
func getText() string {
	l := len(TEXTARRAY)
	index := getRandNum(l - 1)
	return TEXTARRAY[index]
}
