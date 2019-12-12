### 场景六：知识库搭建

知识库搭建在大部分情况下可在吾来平台人工完成。如果项目需要不登录吾来平台的情况下就搭建和管理知识库，可以按照以下流程接入。

对接步骤:

```
	1. 创建知识点
		a. 调用 QaKnowledgeTagList 方法，查询可用的知识点分类.
		b. 调用 QaKnowledgeCreate 方法，创建知识点.

	2. 创建属性组
		a.调用 QaUserAttributeGroupItemList 方法，查询可用于定义属性组的用户属性.
		b.调用 QaUserAttributeGroupItemCreate 接口，创建属性组.

	3. 创建属性组回复
		a.调用 QaUserAttributeGroupItemList 方法,查询属性组.
		b.调用 QaKnowledgeItemList 方法,查询知识点.
		c.调用 QaUserAttributeGroupAnswerCreate 方法,创建属性组回复.
```
        

