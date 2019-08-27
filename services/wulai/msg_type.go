package wulai

//Text 消息类型:文本消息
type Text struct {
	Content string `json:"content"` //文本消息体 [ 1 .. 2048 ] characters
}

//Image 消息类型:图片消息
type Image struct {
	ResourceURL string `json:"resource_url"` //图片的资源链接,必填 [ 1 .. 512 ] characters
}

//Custom 消息类型:自定义消息
type Custom struct {
	Content string `json:"content"` //消息内容
}

//Video 消息类型:视频消息
type Video struct {
	ResourceURL string `json:"resource_url"` //视频的资源链接,必填 [ 1 .. 512 ] characters
	Thumb       string `json:"thumb"`        //视频的缩略图 <= 512 characters
	Description string `json:"description"`  //视频描述 <= 512 characters
	Title       string `json:"title"`        //视频标题 <= 128 characters
}

//File 消息类型:文件消息
type File struct {
	FileName    string `json:"file_name"`    //文件名 <= 128 characters
	ResourceURL string `json:"resource_url"` //视频的资源链接,必填 [ 1 .. 512 ] characters
}

//Voice 消息类型:语音消息
type Voice struct {
	UserID string `json:"user_id"` //用户唯一标识 [ 1 .. 128 ] characters
	Extra  string `json:"extra"`   //自定义字段 <= 1024 characters
}

//Event 消息类型:事件消息
type Event struct {
	//TODO:待开发
}

//ShareLink 消息类型:图文消息
type ShareLink struct {
	Description    string `json:"description"`     //视频描述 <= 512 characters
	DestinationURL string `json:"destination_url"` //链接目标地址,必填 [ 1 .. 512 ] characters
	CoverURL       string `json:"cover_url"`       //封面图片地址,必填 [ 1 .. 512 ] characters
	Title          string `json:"title"`           //链接的文字标题,必填 [ 1 .. 128 ] characters
}
