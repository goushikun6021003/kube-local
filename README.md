# kube-local
  - 本项目为发送信息的RPC服务器端；
  - 服务端口为8888(也可配置文件中自行更改）
  - 调用方法为Sender.SendMessage,发送内容为RecvData格式，返回信息为string格式，
  - 具体结构体可查看pkg.model.rule里内容。
  - RecvData结构体中DstObj字段有"slack","mail","dingding","lark","wechat"五种类型。
