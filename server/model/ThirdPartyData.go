package model

/*
Text
返回企微的数据结构
*/
type Text struct {
	Content string `json:"content"`
}

/*
QyData
返回企微的数据结构
*/
type QyData struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}

/*
---------------------------------------------------------
Markdown格式
*/

type MdData struct {
	Msgtype  string `json:"msgtype"`
	Markdown Text   `json:"markdown"`
}

/*
---------------------------------------------------------
卡片模板
*/

/*
TemplateCardContent
卡片模板内容
*/
type TemplateCardContent struct {
	CardType string `json:"card_type"`
}

/*
MainTitle
模版卡片的主要内容，包括一级标题和标题辅助信息
*/
type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

/*
TemplateCardType
卡片类型
*/
type TemplateCardType struct {
	Msgtype      string              `json:"msgtype"`
	TemplateCard TemplateCardContent `json:"template_card"`
	MainTitle    MainTitle           `json:"main_title"`
}
