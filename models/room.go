package models

import (
	"sync"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"math/rand"
	"time"
	"net"
	"crypto/tls"
)

var TEXTARRAY []string
var AmountArray []int64
var UserArray []User

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
	for i := 1; i < 6; i++ {
		AmountArray = append(AmountArray, int64(i*10+50))
	}
	for i := 1; i < 5; i++ {
		AmountArray = append(AmountArray, int64(i*50+100))
	}
	for i := 1; i < 3; i++ {
		AmountArray = append(AmountArray, int64(i*100+100))
	}

	UserNameArray := []string{
		"刘思恬4月飞日本",
		"蔡莉",
		"余生",
		"李萱",
		" sunru ",
		"类类",
		"蕾蕾",
		"夏磊瓒",
		"乐乐",
		"楽",
		"Zora",
		"997_🙃",
		"冷雪纯",
		"🔆Catherine",
		"Macro",
		"恩戴米恩",
		"章鱼妹妹",
		"Kate HK",
		"LeunGyH",
		"无恙",
		"ZjhZp",
		"🦍🦍久",
		"Fanfan",
		"空白",
		"热爱",
		"SEUNGRI",
		"思瑶",
		"J",
		"Better me",
		"梁小朵儿",
		"Elio",
		"kaill",

	}
	AvatarUrl := map[string]string{
		"刘思恬4月飞日本": "http://wx.qlogo.cn/mmhead/ver_1/nUIE5cFNcUjbZ2psXz8s9YbE8wWRkb5WtWmiariawqA6HzSbAia9s08ITJx9iccm2Wvgke3oXQ6lV9Kgs9f0gAc42naThfib97U7KAmIlic8FHDP8/0",
		"蔡莉":       "http://wx.qlogo.cn/mmhead/ver_1/AyznXDIVWWSlb5PwscBtIXgAYX2aicxJNicz8lAmxeGJItAYtnrCMNiciaNSL3ia3ZTC1kHbP9qGnQGFokvAQfRZlzictCHYGuZAZzicnfDzxwMibt4/0",
		"余生":       "http://wx.qlogo.cn/mmhead/ver_1/WicYeqZmJqnzWc8dXL1mv9b3Qge1fFQ58h0e38YmVVW1WkJwpokm9yssDbRyBIDiaJr7msppGnqOeLXyFyvcbcyFq8CsrqB2pkUKUUl8KJSr0/0",
		"李萱":       "http://wx.qlogo.cn/mmhead/ver_1/cD4KiafGAQWzv1fRCRMFdVpkGvlcEQJf1nxLBKIuQsdoeNza4dz24ibC0suHjaicBI1aHD3vB2eG3DcgzAItwHfJA/0",
		" sunru ":     "http://wx.qlogo.cn/mmhead/ver_1/LP8joYcjEVXxJibadJ3SUwneMhNmyfxU4knlgDLgd62AA4Vqd07PGe98aOiaostaG0ntYyNxbZtfc9kGZc4w1malcpwdIicprYCTyngER7h0qk/0",
		"类类":       "http://wx.qlogo.cn/mmhead/ver_1/VGLl0k5AL3VlwTvBHrk3BTh29JJS9Nq0qqVXRibCRURURWK8QFZf8oQfMUoH0JeSreZSkKzFYuvl4n99lVXIt5H6uPALCLEa0cqYlGODTh5g/0",
		"蕾蕾":      "http://wx.qlogo.cn/mmhead/ver_1/e0RAMQyJmX5bRM19TqK0UHn8kEBUOH9XxDoBBEGviaYRIXab8OiaEsHP0yTWglM1eaLpJIyOJcJlxzRa69hGkCDMyElCC8ScN3uOKibDdCTIT4/0",
		"夏磊瓒":     "http://wx.qlogo.cn/mmhead/ver_1/lm61K5jVHQGV4AmTM3j1ibO3YUMnMFxr4t4MKxs6C9gvQG6dCOBTiaQKaWaM0V1UG6Bob9tuicGIQ1M5vTia8ejC59xQr03ReDK5NV5F3gPyibNs/0",
		"乐乐":       "http://wx.qlogo.cn/mmhead/ver_1/FmjAjfmdAt13P9HTB5OuKgNqDIFeAc3WpNaSficzQFlQ2ZxTMCa3qtHs67aJr6NLibGDzndzSQEv9IGQ5BaahLd9RG1WiaQggNiaYhxK6BicnhK4/0",
		"楽":    "http://wx.qlogo.cn/mmhead/ver_1/sic79jWOBBPickrqLOAJNS1nkjnNtSnRcxVE8DsOJNibvTuibBCNBZqf4fEooSa8xYCgQjSaaGkG6o9qzg3GZeZmyKwhx6c1W9FfkRCZzibje8Vs/0",
		"Zora":     "http://wx.qlogo.cn/mmhead/ver_1/G79uqtUmOSl8keZzC7QKI5HyX1XMaGleKnWXGajYDs7LP5vXGTjJfJqy4Bc3458qrsHSrRbEo7hFH6pDiauC6X1wibCWiaJ0icHehOY961bVSu0/0",
		"997_🙃":     "http://wx.qlogo.cn/mmhead/ver_1/tz0jZy069QmgJRxOtWaA70E3l7FbdayJ8jlic6VU5SXrq2NrJpKqU83Ra4JjvAIXVoTXGHupia0XXtPRoibcxPlBGvGsMJFTZtjJfw0YBX0kDo/0",
		"冷雪纯":      "http://wx.qlogo.cn/mmhead/ver_1/7KKibrXnW7CuNQ12EEUvUqzzP7kWFaSD0SzuQkibP50nVgyJbG2H34YabNDyicBYMyAq68p3kq7pBW3rI2YJMcJmnsZl871uzjLUib7GxsKiciavs/0",
		"🔆Catherine":     "http://wx.qlogo.cn/mmhead/ver_1/xjSVHwgWYf0Jre8JEILqyfF7p8mpsnCkyRCic9bt7icuInGBaDLGIFiaL0m8mwia8BIL6xqdM2bFwyn6AYGX7Ek5AbD6RicKjMibMYuz81ibqPticiag/0",
		"Macro":       "http://wx.qlogo.cn/mmhead/ver_1/ziaAPickWibrvQUs4icHZrUJe5bH0byxejTnuoCQ0bKxu22MibSQtEDbsIgtEnHgBiaAiahcMWvA7X5iaXN2GTNVn0h2GQ/0",
		"恩戴米恩":     "http://wx.qlogo.cn/mmhead/ver_1/RYvhphP7ibxoH5YpTsEwkiczRDTppb4qXYiaKKhp27hugS4doI27oxvmK9HFdljRwk5Btj38FicvCKOBJU2U1ygQ31pUqnGZmAsqzDSADdf6rsU/0",
		"章鱼妹妹":      "http://wx.qlogo.cn/mmhead/ver_1/2gGxR9e05fC2X64YKk5gKT6tavvo4U6Od3hHibShY9xjM9c9ibcrBbbMHDflnX5Fw8fcjpDFFric57uzdqTd8S9J7B9BgwhQicDkmxGsM7WPw08/0",
		"Kate HK":       "http://wx.qlogo.cn/mmhead/ver_1/jzDWywOuK2MWerqNbcKRIicUj1xQib1RtcHu02lzfsUcVIvPotDST1MHGBWx0Cc7XbeLec3DABViaOa8gicvHorPaaTxt9zFpTSu3rNjHZXvmibk/0",
		"LeunGyH":     "http://wx.qlogo.cn/mmhead/ver_1/wxedibNGZxQyNS1Bzx9c1z1JOgCpOKX0ZXFltkiaeZQxiaY0EtxuRc86dhxJicQ1sL2qvToh5820PhUDohdTVJrO8D8KJCTcE8VMCV0S7yiaW7G0/0",
		"无恙":   "http://wx.qlogo.cn/mmhead/ver_1/ABn7gMSOFMibmV5ZU7MoztTARGHoVEYbNcicucCxKGrsmYRv81kCVL5ial54PRSTNhBqNpU3SIm1jRhZrW4dibkOgNiaEORzdVtQ3iaoHSBv2k9lQ/0",
		"ZjhZp":   "http://wx.qlogo.cn/mmhead/ver_1/VLz5IPlxgpSDmMgx3rwicwib90yjSlR5IkSPiciaiaU1ibONBia9e5dGlTnJGpV9WqM0SNDz00z12CZEVJoJqAg464UB5ibEUG09cm3QNImOlGwyu1E/0",
		"🦍🦍久":     "http://wx.qlogo.cn/mmhead/ver_1/u9vRj55tZbsBibCIvpOKV8toIaLeWGGBEZgeZoML5ktboXlJU7opSZEs7N2v1B03tL1dwvHunafsD6FUFcUJaH1bnOvIpfESSmBtiab2mcbwM/0",
		"Fanfan":    "http://wx.qlogo.cn/mmhead/ver_1/XO5gh6yZyUbicpGEoh6b0PrVRvlyJxjV3IiavboEicBhjqwxJQIGzQgTPUJwvcUMGONbnx3EstM6L2z8kHLBVtlhZ6I12Y7ia1pgm0ibZTFQf3aQ/0",
		"空白":      "http://wx.qlogo.cn/mmhead/ver_1/gRjkE42XemnoUQlnicpIAjaowfic3kofIuKMzdI1zWGvw6ujbjHVGRESo6E7kSxVsciaHY0tiavFNYVb3On1UIgcicQ/0",
		"热爱":     "http://wx.qlogo.cn/mmhead/ver_1/kH5ke8U2iaerHc2m2ibsicGZc6BWtk10rTYp2HHiafJdhAF7iapp2TyymvbgY1HdcD5RrEBDjtOm3OkQGNPEibTFp4Ns5lKGYFJaEOO1bW8M0yDFA/0",
		"SEUNGRI":     "http://wx.qlogo.cn/mmhead/cypR72jV8BHjDwNh3Nc1YcsgzmiaZacpR1dgiaibt4QuMs/0",
		"思瑶":       "http://wx.qlogo.cn/mmhead/ver_1/8ic8gRyChQEhN89e4joeLjRm0jOoQaTXcLHKNv2AxaNxannuEL9U4TbJwp03yZ1uy5gN8lmrVZZkcOZjfDf1VKzpcjiaH7GqwtlRGR6ZUKEmI/0",
		"J":       "http://wx.qlogo.cn/mmhead/ver_1/H8IwapZfbbr1pWxQFwyZvuDbLzK7vHnHKkRDyzbibd0JhYiaaTXdaFzdPayjJIAD3ORfiaibSThtwU4ABfegVsypFS6tMGziaju2WdUOic6HVRUcs/0",
		"Better me":   "http://wx.qlogo.cn/mmhead/ver_1/u7Prpo2OHmapV9u6U7vaqy4ouP778aY4GI4icutyd82CjGtzG73OMSuwOtrhibibIWSCjSmxtuyCrRnVrKXqecodA/0",
		"梁小朵儿":     "http://wx.qlogo.cn/mmhead/ver_1/lqkl3A5CrcKHcaJNWbCBa5pmNGKW9dXg1Ju9RfQyPubibwbUt6gkYfrRg1bAy3paxSsHNiaUicaXZoicl3BIIgMiaOeBlCSfDaFPMF5AYS5hSTZg/0",
		"Elio":   "http://wx.qlogo.cn/mmhead/ver_1/ethF5twU40vQKyEcR9XNnUrOY9X6YiaKfEvTTRuBZ4ciaJ6He1KrG1oP03GL78O1N5bIhhIOkxT5NjHN1iamqc1T8cm1KsV7AHbHrtN3hIyYpU/0",
		"kaill":     "http://wx.qlogo.cn/mmhead/ver_1/vCiclpLjonLUzmInas7fYw85bEWe9ibvRbg3ygXavtrz1o73tJWANNLbIrw1pNtnYAnhFXkmv29RAT67e4V4N5O9Fec0WV8wN6iah21RaZ4jgE/0",
	}

	UserArray = make([]User, 0, 100)
	for i := 0; i < len(UserNameArray); i++ {
		UserArray = append(UserArray, User{Name: UserNameArray[i], AvatarUrl: AvatarUrl[UserNameArray[i]]})
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
		switch room.Id {
		case 1:
			randomMessage(room,"https://api.erong28.com/api/lottery/latest/period/short",2,210)
		case 2:
			randomMessage(room,"https://api.erong28.com/api/lottery/latest/period/short",20,210)
		case 3:
			randomMessage(room,"https://api.erong28.com/api/lottery/latest/period/long",2,300)
		case  4:
			randomMessage(room,"https://api.erong28.com/api/lottery/latest/period/long",20,300)
		default:

		}

	}()
	return room
}

func randomMessage(room *Room,url string ,sleepRange int, interval int64) {
	for {
		sleepTime := getRandNum(sleepRange)
		difference := time.Now().Unix() - 1553834100
		if difference%interval < 5 || difference%interval > (interval-10) {
			continue
		}
		time.Sleep(time.Duration(sleepTime) * time.Second)
		RandomData(room, url)
	}
}
func RandomData(room *Room, apiUrl string) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * 5,
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
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
	room.In <- &Message{GROUPCHAT, getRandonUser(), "第" + strconv.Itoa(data.Data) + `期：
` + getText() + "/" + strconv.Itoa(getAmount())}
}
func RandomData12(room *Room, apiUrl string) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
	}
	client := &http.Client{
		Timeout:   time.Second * 5,
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
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
	count := getRandNum(4)
	for i := 0; i < count; i++ {
		room.In <- &Message{GROUPCHAT, getRandonUser(), "第" + strconv.Itoa(data.Data) + `期：
` + getText() + "/" + strconv.Itoa(getAmount())}
	}

}
func getRandNum(range_ int) int {
	return rand.Intn(range_)
}
func getRandonUser() User {
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
