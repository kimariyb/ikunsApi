# ikuns视频播放网

ikuns 视频播放网站是一个使用 Go 语言和 Beego 框架搭建的在线视频播放平台，其中应用了 Redis、RabbitMQ 和 Elasticsearch 技术栈，允许用户上传和观看视频。ikuns 视频播放网站，为江西理工大学 kimariyb 的毕业设计。

## 功能

- 用户可以注册和登录账号。
- 用户可以上传视频，并设置视频标题和描述。
- 用户可以浏览视频列表和观看视频。
- 用户可以搜索视频和用户。
- 平台支持视频评论和点赞功能。
- 用户可以发送弹幕

## 技术栈

- 编程语言：Go
- Web 框架：Beego
- 数据库：MySql
- 缓存数据库：Redis
- 消息中间件：RabbitMQ
- 搜索引擎：Elasticsearch

## 部署方式

1. 克隆代码仓库：`git clone https://github.com/kimariyb/ikuns-video.git`
2. 安装依赖包：`go get`
3. 启动 MySql 数据库。
4. 启动 Redis 缓存数据库。
5. 启动 RabbitMQ 消息中间件。
6. 在 `conf/app.conf` 文件中配置相关参数。
7. 启动应用程序：`bee run`

应用程序默认监听 8098 端口，可以在 `conf/app.conf` 文件中进行修改。

## 开发者

- 开发者：*@kimariyb*
- 电子邮件：kimariyb@163.com

## 版本历史

- 版本 1.0.0（2023 年 5 月 27 日）：首次发布。

## 贡献者

- kimariyb

## 许可证

本项目使用 Apache 开源许可证。