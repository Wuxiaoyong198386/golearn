/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-26
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

//本节主要学习nsq
/**
 1、分布式消息队列中间件，分布式和去中心化拓扑结构
 2、无单点故障，故障容错，高可用
 3、布置简单，容易使用，支持延时消息


 核心点：
 一、三个必要的组建
 	1、nsqd
 		a、nsqd 负责接收消息，存储队列和将消息发送给客户端的守护进程
 		b、nsqd可以在多台机器上布置，当你使用客户端给一个topic(话题)发送消息时，可以配多个nsqd地址
 			消息会随机的分布到各个nsqd上。
 		c、nsqd会把消息优先存在内存channel(通道)上，当内存channel满了后，会把消息写在磁盘文件。
 		d、nsqd会侦听两个tcp端口,一个用来服务客户端(默认：4150)，一个用来提供http的接口（默认：4151），

 		nsqd --lookupd-tcp-address=127.0.0.1:4160

 		参数说明：
 		nsqd --lookupd-tcp-address=127.0.0.1:4160 --broadcast-address=127.0.0.1

 		也可以指定端口 与数据目录，其他配置项可详见官网
 		nsqd --lookupd-tcp-address=127.0.0.1:4160    解析 TCP 地址名字 (可能会给多次)
 				--broadcast-address=127.0.0.1		 通过 lookupd  注册的地址（默认名是 OS）
 				-tcp-address=127.0.0.1:4150			 TCP 客户端 监听的 <addr>:<port>
 				-http-address="0.0.0.0:4151"         为 HTTP 客户端监听 <addr>:<port>
 				--data-path=/data/nsqdata			缓存消息的磁盘路径

 	2、nsqlookupd

 		nsqlookupd --http-address="0.0.0.0:4161" --tcp-address="0.0.0.0:4160"
 		a、负责服务发现、负责nsqd的心跳、状态监测、给客户端
 		参数说明：
 		-http-address="0.0.0.0:4161": <addr>:<port> 监听 HTTP 客户端（管理nsqadmin服务）
 		-tcp-address="0.0.0.0:4160": TCP 客户端监听的 <addr>:<port>（管理nsqd服务）
		-inactive-producer-timeout=5m0s: 从上次 ping 之后，生产者驻留在活跃列表中的时长
		-broadcast-address: 这个 lookupd 节点的外部地址, (默认是 OS 主机名)
		-tombstone-lifetime=45s: 生产者保持 tombstoned  的时长
		-verbose=false: 允许输出日志
		-version=false: 打印版本信息

 	3、nsqadmin
		a、是一个web管理界面 默认的访问地址：http://127.0.0.1:4171
		nsqadmin --lookupd-http-address=127.0.0.1:4161（与nsqlookupd进行通信）


 		-graphite-url="": URL to graphite HTTP 地址
		-http-address="0.0.0.0:4171": <addr>:<port> HTTP clients 监听的地址和端口
		-lookupd-http-address=[]: lookupd HTTP 地址 (可能会提供多次)
		-notification-http-endpoint="": HTTP 端点 (完全限定) ，管理动作将会发送到
		-nsqd-http-address=[]: nsqd HTTP 地址 (可能会提供多次)
		-proxy-graphite=false: Proxy HTTP requests to graphite
		-template-dir="": 临时目录路径
		-use-statsd-prefixes=true: expect statsd prefixed keys in graphite (ie: 'stats_counts.')
		-version=false: 打印版本信息
 	其中nsqd和nsqlookupd是必须要的组建



	三者之间的关系及启动服务顺序：

 	第一步：
 	侦听本机两个端口
 	nsqlookupd --tcp-address="0.0.0.0:4160" --http-address="0.0.0.0:4161"
 	TCP: listening on [::]:4160
	HTTP: listening on [::]:4161

 	第二步：
 	通过4160的端口和第一个服务通信
 	nsqd --lookupd-tcp-address=127.0.0.1:4160
 	LOOKUP connecting to 127.0.0.1:4160
	TCP: listening on [::]:4150   				//队列生产者通过此端口和nsqd建立对话
    HTTP: listening on [::] 4151

 	第三步：
 	通过4161的端口和第一个服务通信
 	nsqadmin --lookupd-http-address=127.0.0.1:4161
 	HTTP: listening on [::]:4171
 	默认侦听4171端口,如果想改端口可以加上--http-address="127.0.0.1:9001"
 	HTTP: listening on 127.0.0.1:9001

	消息通过api处理

 	nsq_to_http --topic=sms --lookupd-http-address=127.0.0.1:4161
                --post=http://127.0.0.1:8001/callBack/logConsumer/logConsumer

	post如何接数据：

		buf:=make([]byte,1024)
		n,err:=req.Body.Read(buf)
		if err!=nil{
			fmt.Println(err.Error())
		}
		fmt.Println(string(buf[0:n]))

	nsq_to_http --topic=sms --lookupd-http-address=127.0.0.1:4161
                --get=http://127.0.0.1:8001/callBack/logConsumer/logConsumer?msg=%s

	get如保接数据

		query:=req.URL.Query()
		get_act:=query["msg"][0]
		w.Write([]byte(get_act))
		fmt.Println(get_act)



	消息log文本文件存储，默认./tmp文件夹
	nsq_to_file --topic=Jsondata --lookupd-http-address=127.0.0.1:4161 --output-dir=./tmp



*/

package main

func main() {

}
