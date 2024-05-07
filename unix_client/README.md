# <center>UnixClient</center>

# 前言

自研的unix 客户端

unix 客户端，主要用于unix的测试场景，linux环境上，curl命令也可以实现。

curl --unix-socket /tmp/nginx-status-server.sock http://localhost/nginx_status

当前工具跟这个curl命令功能相似。只是当前初步加了一点自己的东西而已

unixclient编译时，需要根据build中的文件，配置GoLand对应的编译参数。

当前编译指导是按照linux配置的，windows下的sock监听server端没配置过，所以client端也没有配置的必要。

如果要在本地进行功能测试，debug，可以在build配置中的**Program arguments**中配置程序的入参，然后运行调试

# 一、Usage

    usage: ./unixclient [-options][value]... /socket_file /url
      -b string
            msg body
      -debug
            output debug log
      -del
            send a DELETE msg to socket server
      -delay int
            after pre proc, delay times second
      -get
            send a GET msg to socket server
      -h    usage help
      -json
            output response body by json format
      -local
            use http://localhost/xxx
      -patch
            send a PATCH msg to socket server
      -post
            send a POST msg to socket server
      -pre int
            repeat send message before method times
      -print
            output request header (default true)
      -put
            send a PUT msg to socket server
      -test int
            send multi msg to socket server

## 1.1 options

* b
    * 消息的body，这个会直接按照string读取填入请求消息中去，所以这里参数的格式要自己注意
* debug
    * 调试日志打印
* del
    * DELETE 方法
* delay
    * 延时处理，在pre处理之后，正式发消息之前，用处是pre结束后，等待一个时间，然后在发送一条请求消息。用于构造特殊测试场景
* get
    * GET 方法
* h
    * 命令帮助
* json
    * json格式输出消息响应

* patch
    * PATCH 方法
* post
    * POST 方法
* pre
    * 预处理，和其他方法结合使用，在发送消息正常的处理器，重复发送一定次数的请求消息，构造特殊测试场景
* print
    * 打印响应消息体，默认是true，即打印，不过是按照字符串打印的。
    * 如果需要按照json格式打印，请加上-json参数。
    * 如果都选择了，优先json输出。
* put
    * PUT 方法
* test
    * 测试功能，支持连续重复发送多条相同消息

## 1.2 socket_file

socket文件绝对路径

如：/opt/local/testsock.sock

## 1.3 url

url路径

如：

/routes

/plugins

# 二、example

## 2.1 json输出

    ./unixclient -get -json /opt/mdisk/apigwuds/apigwlocalhost8001.sock /plugins/f05d896a-695b-4185-a9c3-4b3f3cb0a210


## 2.2 普通输出

    ./unixclient -get /opt/mdisk/apigwuds/apigwlocalhost8001.sock /plugins/f05d896a-695b-4185-a9c3-4b3f3cb0a210


## 2.3 debug

    ./unixclient -get -debug /opt/mdisk/apigwuds/apigwlocalhost8001.sock /plugins/f05d896a-695b-4185-a9c3-4b3f3cb0a210