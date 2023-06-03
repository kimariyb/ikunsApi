package controllers

import (
	"encoding/json"
	"fmt"
	"ikunsApi/models"
	"ikunsApi/services/es"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// 频道页 - 获取顶部广告
// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		this.ServeJSON()
	}
}

// 频道页-获取正在热播
// @router /channel/hot [*]
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelHotList(channelId)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// 频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (this *VideoController) ChannelRecommendRegionList() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	if regionId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道地区")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// 频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecomendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道类型")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}

}

// 根据传入参数获取视频列表
// @router /channel/video [*]
func (this *VideoController) ChannelVideo() {
	//获取频道ID
	channelId, _ := this.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := this.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := this.GetInt("typeId")
	//获取状态
	end := this.GetString("end")
	//获取排序
	sort := this.GetString("sort")
	//获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// 我的视频管理
// @router /user/video [*]
func (this *VideoController) UserVideo() {
	uid, _ := this.GetInt("uid")
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定用户")
		this.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// 获取视频详情
// @router /video/info [*]
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}
	video, err := models.GetVideoInfo(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", video, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		this.ServeJSON()
	}
}

// 获取视频剧集列表
// @router /video/episodes/list [*]
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		this.ServeJSON()
	}
}

// 保存用户上传视频信息
// @router /video/save [*]
func (this *VideoController) VideoSave() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")
	aliyunVideoId := this.GetString("aliyunVideoId")
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "请先登录")
		this.ServeJSON()
	}
	if playUrl == "" {
		this.Data["json"] = ReturnError(4002, "视频地址不能为空")
		this.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}
}

// 搜索接口
// @router /video/search [*]
func (this *VideoController) Search() {
	//获取搜索关键字
	keyword := this.GetString("keyword")
	//获取翻页信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if keyword == "" {
		this.Data["json"] = ReturnError(4001, "关键字不能为空")
		this.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}

	sort := []map[string]string{map[string]string{"id": "desc"}}
	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"title": keyword,
				},
			},
		},
	}

	res := es.EsSearch("fyouku_video", query, offset, limit, sort)
	total := res.Total.Value
	var data []models.Video

	for _, v := range res.Hits {
		var itemData models.Video
		err := json.Unmarshal([]byte(v.Source), &itemData)
		if err == nil {
			data = append(data, itemData)
		}
	}
	if total > 0 {
		this.Data["json"] = ReturnSuccess(0, "success", data, int64(total))
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// 生成测试视频数据
// @router /video/save/all [*]
func (this *VideoController) SaveAll() {
	var i = 1
	for {
		i++
		rand.Seed(time.Now().UnixNano())
		uidi := rand.Intn(10)
		uid := uidi + 10

		models.SaveVideo(strconv.Itoa(i)+"鸣人柯南一护路飞由诺阿斯塔"+strconv.Itoa(i), "蜡笔小新樱桃小丸子", 1, 2, 2, "/static/video/coverr-sparks-of-bonfire-1573980240958.mp4", uid, "")
		i++
		fmt.Println(i)
	}
}

// 导入ES脚本
// @router /video/send/es [*]
func (this *VideoController) SendEs() {
	_, data, _ := models.GetAllList()
	for _, v := range data {
		body := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"is_end":               v.IsEnd,
			"channel_id":           v.ChannelId,
			"status":               v.Status,
			"region_id":            v.RegionId,
			"type_id":              v.TypeId,
			"episodes_update_time": v.EpisodesUpdateTime,
			"comment":              v.Comment,
			"user_id":              v.UserId,
			"is_recommend":         v.IsRecommend,
		}
		es.EsAdd("fyouku_video", "video-"+strconv.Itoa(v.Id), body)
	}
}
