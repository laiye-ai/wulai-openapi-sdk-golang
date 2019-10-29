package wulai

/*****************用户属性类型*****************/

//AttributeValueType 用户属性值类型
type AttributeType string

const (
	//USER_ATTRIBUTE_TYPE_DEFAULT :默认类型
	USER_ATTRIBUTE_TYPE_DEFAULT AttributeType = "USER_ATTRIBUTE_TYPE_DEFAULT"

	//USER_ATTRIBUTE_VALUE_TYPE_STRING :系统属性
	USER_ATTRIBUTE_TYPE_SYSTEM AttributeType = "USER_ATTRIBUTE_TYPE_SYSTEM"

	//USER_ATTRIBUTE_TYPE_TEMP :临时属性
	USER_ATTRIBUTE_TYPE_TEMP AttributeType = "USER_ATTRIBUTE_TYPE_TEMP"
)

/*****************用户属性值类型*****************/

//AttributeValueType 用户属性值类型
type AttributeValueType string

const (
	//USER_ATTRIBUTE_VALUE_TYPE_DEFAULT :默认类型
	USER_ATTRIBUTE_VALUE_TYPE_DEFAULT AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_DEFAULT"

	//USER_ATTRIBUTE_VALUE_TYPE_STRING :字符串类型
	USER_ATTRIBUTE_VALUE_TYPE_STRING AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_STRING"

	//USER_ATTRIBUTE_VALUE_TYPE_ENUM :枚举类型
	USER_ATTRIBUTE_VALUE_TYPE_ENUM AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_ENUM"

	//USER_ATTRIBUTE_VALUE_TYPE_TIME :时间类型
	USER_ATTRIBUTE_VALUE_TYPE_TIME AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_TIME"

	//USER_ATTRIBUTE_VALUE_TYPE_CITY :城市类型
	USER_ATTRIBUTE_VALUE_TYPE_CITY AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_CITY"

	//USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING :多行文本类型
	USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_MULTILINE_STRING"

	//USER_ATTRIBUTE_VALUE_TYPE_INT :整型类型
	USER_ATTRIBUTE_VALUE_TYPE_INT AttributeValueType = "USER_ATTRIBUTE_VALUE_TYPE_INT"
)
