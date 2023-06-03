package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"ikunsApi/models"
	"math/rand"
	"strconv"
	"time"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

// 获取评论列表
// @router /comment/list [*]
func (this *CommentController) List() {
	//获取剧集数
	episodesId, _ := this.GetInt("episodesId")
	//获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if episodesId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		this.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err == nil {
		var data []CommentInfo
		var commentInfo CommentInfo

		//获取uid channel
		uidChan := make(chan int, 12)
		closeChan := make(chan bool, 5)
		resChan := make(chan models.UserInfo, 12)
		//把获取到的uid放到channel中
		go func() {
			for _, v := range comments {
				uidChan <- v.UserId
			}
			close(uidChan)
		}()
		//处理uidChannel中的信息
		for i := 0; i < 5; i++ {
			go chanGetUserInfo(uidChan, resChan, closeChan)
		}
		//判断是否执行完成，信息聚合
		go func() {
			for i := 0; i < 5; i++ {
				<-closeChan
			}
			close(resChan)
			close(closeChan)
		}()

		userInfoMap := make(map[int]models.UserInfo)
		for r := range resChan {
			userInfoMap[r.Id] = r
		}
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			//获取用户信息
			commentInfo.UserInfo, _ = userInfoMap[v.UserId]
			data = append(data, commentInfo)
		}

		this.Data["json"] = ReturnSuccess(0, "success", data, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}
func chanGetUserInfo(uidChan chan int, resChan chan models.UserInfo, closeChan chan bool) {
	for uid := range uidChan {
		res, err := models.GetUserInfo(uid)
		fmt.Println(res)
		if err == nil {
			resChan <- res
		}
	}
	closeChan <- true
}

// 保存评论
// @router /comment/save [*]
func (this *CommentController) Save() {
	content := this.GetString("content")
	uid, _ := this.GetInt("uid")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")

	if content == "" {
		this.Data["json"] = ReturnError(4001, "内容不能为空")
		this.ServeJSON()
	}
	if uid == 0 {
		this.Data["json"] = ReturnError(4002, "请先登录")
		this.ServeJSON()
	}
	if episodesId == 0 {
		this.Data["json"] = ReturnError(4003, "必须指定评论剧集ID")
		this.ServeJSON()
	}
	if videoId == 0 {
		this.Data["json"] = ReturnError(4005, "必须指定视频ID")
		this.ServeJSON()
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "succes", "", 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}
}

// 生成评论数据
// @router /comment/save/all [*]
func (this *CommentController) SaveAll() {
	var i = 1
	for {
		i++
		rand.Seed(time.Now().UnixNano())
		uidi := rand.Intn(10)
		uid := uidi + 10

		models.SaveComment(strconv.Itoa(i)+"博人智商很高，在忍者学校成绩都满分。悟性也都是很高螺旋丸很快就能掌握，但不知道为什么博人在战斗的时候总分不清情况。什么时候该打，什么时候该跑总是表现得一股脑。无论敌人什么实力总是要向前冲，一点也表现不出博人的战术分析。按理说博人比鸣人更有理智，但表现得总是差强人意。我偶尔就看的十分尴尬。"+strconv.Itoa(i), uid, 1, 1)
		i++
		fmt.Println(i)
	}
}
