# 编译指导

## unixclient_x86
Run Kind: 选择Package

Package Path: codehub.huawei.com/hemyzhao/tools/efficiency/unixclient

<font color=#FF0000>**去勾选“Run After Build”**</font>

Output directory: 当前工程的release目录。可以设置为`release\`

Working directory: 工程目录

Environment: <font color=#FF0000>**GOOS=linux;GOARCH=amd64**</font>

Go tool arguments: <font color=#FF0000>**-o release\unixclient_x86**</font>

Module：选择unixclient

如图：

![unixclient_x86.png](unixclient_x86.png)

## unixclient_arm
Run Kind: 选择Package

Package Path: codehub.huawei.com/hemyzhao/tools/efficiency/unixclient

<font color=#FF0000>**去勾选“Run After Build”**</font>

Output directory: 当前工程的release目录。可以设置为`release\`

Working directory: 工程目录

Environment: <font color=#FF0000>**GOOS=linux;GOARCH=arm64**</font>

Go tool arguments: <font color=#FF0000>**-o release\unixclient_arm**</font>

Module：选择unixclient

如图：

![unixclient_arm.png](unixclient_arm.png)

## unixclient_windows