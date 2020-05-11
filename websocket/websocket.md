# websocket详解与实践

## 概述

WebSocket是一种网络传输协议，可在单个TCP连接上进行全双工通信，位于OSI模型的应用层。

下图就是一个websocket与http之间的关系
![avatar](https://pic1.zhimg.com/80/6651f2f811ec133b0e6d7e6d0e194b4c_720w.jpg)

在websocket.go会使用websocket协议实现一个群聊天的websocket会话

## websocket的创建

在客户端想与服务器建立Websocket链接就会经历以下流程

- 客户端发出握手请求：在客户端与服务端建立TCP连接之后，客户端以HTTP报文形式发送握手请求到服务器

注意： 

HTTP报文必须合法，且请求的方式为GET；

Http报文必须包含以下请求消息头：

```go
GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: x3JJHMbDL1EzLkh9GBhXDw==
Sec-WebSocket-Protocol: chat, superchat
Sec-WebSocket-Version: 13
Origin: http://example.com

```
Sec-WebSocket-Key：  是一个Base64 encode的值，这个是浏览器随机生成的，告诉服务器：我要求验证你是不是能处理websocket请求的能力。提供基本的防护，比如恶意的连接，或者无意的连接。

### Sec-WebSocket-Accept的计算

Sec-WebSocket-Accept根据客户端请求首部的Sec-WebSocket-Key计算出来。

计算公式为：

- 将Sec-WebSocket-Key跟258EAFA5-E914-47DA-95CA-C5AB0DC85B11拼接。
- 再通过SHA1计算出摘要，并转成base64字符串。

客户端会对返回的Sec-WebSocket-Accept进行校验，以检查服务端是否能够正确的处理WebSocket协议。

## Sec-WebSocket-Extensions字段

在建立完websocket之后服务端会先加载Sec-WebSocket-Extensions字段下所需要的插件

## Sec-WebSocket-Protocol字段

Sec_WebSocket-Protocol 是一个用户定义的字符串，用来区分同URL下，不同的服务所需要的协议。简单理解：今晚我要服务A，不要弄错了

## Sec-WebSocket-Version 

Sec-WebSocket-Version 表示 WebSocket 的版本，最初 WebSocket 协议太多，不同厂商都有自己的协议版本，不过现在已经定下来了。如果服务端不支持该版本，需要返回一个 Sec-WebSocket-Versionheader，里面包含服务端支持的版本号。