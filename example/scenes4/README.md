### 场景四：异步定制对话

通过异步方式实现更复杂的定制化对话。

实现效果：调用方将用户发送的消息传给吾来，吾来通过消息路由将机器人的回复发给调用方，调用方对回复内容进行二次加工（如：修改机器人直接回复的内容、将机器人回复内容作为客服的候选回复等），加工完成后再将最终要发送的回复内容传回消息路由，吾来通过消息投递接口将最终回复发送给用户。

<img src="../../docs/static/scenes4-1.png"  width="800" height="400">

对接前提:

    * 调用方提供消息收发的渠道。
    * 渠道具有 WebSocket 能力。
    * 已经在吾来平台-渠道设置-开放平台(新版)设置了消息路由接口。
    * 已经在吾来平台-渠道设置-开放平台(新版)设置了消息投递接口。
    * 已经在吾来机器人平台搭建并训练了机器人。

对接步骤:

    1. 调用 user/create 接口，传入 user_id，创建用户；
    2. 如果需要个性化回复，在此处调用 user/user-attribute/create 接口，传入 user_id 和该用户的用户属性；
    3. 调用 msg/receive 接口，将用户的问题发给吾来；
    4. 在消息路由中接受机器人返回的结果，并根据业务需要修改相应内容；
        1）如果需修改机器人回复的内容，修改 msg_body 的内容。
        2）如果不希望这条消息发给用户，修改 is_send 为 false。
        3）如果希望将机器人回复作为候选回复展示给客服，可提取 suggested_response/response 中的内容。
        该方法需要调用方的客服工作台支持前端展示页面，展示形式可参考吾来的实现方式（下图）：


<img src="../../docs/static/scenes4-2.png">

        4）"是否派发"参数 is_dispatch 表示机器人判断是否应该转人工（true 代表应该转人工）。不管 is_dispatch 是 true 或 false ，调用方都可以将一条消息转人工处理。如果决定要做“转人工”处理，需要完成以下步骤：
        - 在消息路由的消息体中，修改 `is_send` 的值为 `false` ，这样这条消息就不会被发送给用户。  
        - 同时将这条消息推送到调用方人工服务平台，并根据调用方规则安排客服。  
        - 在人工接入状态下，接收的每一条用户消息，都需要分别做”转人工“处理。
        5）将消息路由的消息体处理完成之后，作为 Response 返回给吾来平台。

    5. 吾来平台会回调消息投递接口，将最终回复发给用户。
  

链接: [异步对话文档](../../docs/callback.md)
      