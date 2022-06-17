# SimpleChatWindow

## 这是一个基于redis的订阅与发布模型做的一个聊天室的小程序，实现了简单的通话模型。

** 怎样去使用?

基于Windows10操作系统
1. 修改DataSource文件夹的yaml配置文件，将redis以及mysql的hostname配置为服务器的hostname，注意一定要在同一个网络中才能够使用。
2. 使用 go build chatOutPut 将源代码编译成可执行文件exe(chatOutPut用来收集对话框里面的内容)
3. 使用 go build chatServer 将聊天的客户端编译成可执行的exe文件
4.  执行chatServer文件即可

**注意:因为作者自身的原因，导致客户端可以直接创建到mysql的链接，此方式可能会导致一定的不安全性，所以以上代码仅作参考**
