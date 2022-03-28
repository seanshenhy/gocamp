#  socket 粘包
数据在TCP层是以字节流传输，发送方把数据拆分成多个数据包发送，接收方在缓冲区组合多个小包，这种组合小包就会涉及某种规则解析数据。

## 解析数据方式

定长（fix length）
 * 发送方以某种固定字节长度发送，接收方以固定长度解析数据。

特殊限制符（delimiter based）
 * 发送方以\r\n作为定界符发送数据，接收方获取后以\r\n作为分隔符解析数据。
 
长度编码（length field based frame decoder）
 * 发送方在每个数据小包加上包长度，接收方获取后按照包长度读取解析数据。

采用包头固定长度

代码作业见ch9/pkg/proto下
```shell script
cd server && go run .
```

```shell script
cd client && go run .
```