# eBPF_Visualization

提供更加易用的 eBPF 组件管理系统，并提供 eBPF组件 工作过程和数据的可视化展示。项目目标如下：

1. 管理 eBPF 组件；
2. 快速便捷地实时展示 Linux 系统性能数据；
3. 可视化展示**应用/进程**从用户态到内核态执行链路；

项目采用前后端分离开发方式，整体由 Web 应用和服务端应用组成，主要功能为eBPF插件数据可视化和内核场景化分析可视化。

##### web 端

该项目的目标是收集梳理基于 eBPF 的可观测性组件，从各种实现方式，各个角度对内核进行可观测。

Web 应用目标如下：

1. 用户信息管理；
2. eBPF 组件信息管理；
3. eBPF 数据可视化；
4. eBPF 场景化分析可视化；

##### server 端

服务端应用目标如下：

1. 提供 `RESTful` 风格 `API` 和接口说明文档；
2. 基于 `JWT` 的用户校验；
3. 进行 eBPF 组件信息和用户信息的持久化存储；



#### eBPF插件数据可视化说明

目标是快速实现eBPF数据的可视化观测。

欢迎大家体验使用这里的组件，更欢迎大家贡献更多的基于 eBPF 的可观测性程序。

插件类型目录如下：

| 类型 | 介绍                               |
| ---- | ---------------------------------- |
| BCC  | BCC 的 eBPF 可观测性程序           |
| C    | 使用 C 开发的 eBPF 可观测性程序    |
| GO   | 使用 GO 开发的 eBPF 可观测性程序   |
| RUST | 使用 RUST 开发的 eBPF 可观测性程序 |



#### eBPF内核场景化分析说明

旨在利用 eBPF 实现 应用/内核事件 的场景化分析，即实现某一 应用/进程 从用户态到内核态执行链路的关键事件分析和可视化。目的如下：

1. 探索 eBPF 能力
2. 为初学者、内核爱好者学习Linux内核提供一种新的可视化方式；
3. 熟悉应用/内核事件执行流程，为使用 eBPF 技术积累知识储备；
4. 体验同一内核事件的不同执行链路，例如不同用户态操作会使用相同的内核事件，但内核事件的执行方式不同；
5. 解决生产问题时，如已定位到某个内核事件，可快速使用仓库中的eBPF代码脚手架；

场景化分析目录如下：

| 项目名称   | 场景介绍                 | 技术栈      | 维护者 |
| ---------- | ------------------------ | ----------- | ------ |
| write      | 将10个字符写入一个空文件 | cilium/ebpf | 孙张品 |
| tr_package | 本地一个数据包的收发流程 | cilium/ebpf | 张翔哲 |



# Linux microscope

LMP是一个基于BCC(BPF Compiler Collection)的Linux系统性能数据实时展示的web工具，它使用BPF(Berkeley Packet Filters)，也叫eBPF，目前LMP在ubuntu18.04上测试通过，内核版本4.15.0。

## 代码结构
```
├── config  # 配置
├── docs    # 文档
├── pkg     # golang 服务代码
├── plugins # python ebpf 代码
├── static  # 网页代码
├── test    # 测试数据
└── vendor  # golang vendor 
```
## 项目架构

![](../eBPF_docs/static/imgs/LMP-arch4.png)
![](../eBPF_docs/static/imgs/LMP-arch5.png)

## 界面截图

![homepage](../eBPF_docs/static/imgs/homepage.png)

![homepage](../eBPF_docs/static/imgs/grafana1.png)

![grafana4](../eBPF_docs/static/imgs/grafana4.png)


##  运行lmp

###  Ubuntu-source

#### 从源码构建lmp，需要的基本环境：

- golang：go1.12及以上；
- docker：influxdb、grafana；
- bcc环境
- mysql：5.7.29测试通过

###  安装依赖docker镜像

```
# For grafana
sudo docker pull grafana/grafana
# For Influxdb
sudo docker pull influxdb
```


### 编译并安装

```
 git clone https://github.com/linuxkerneltravel/lmp
 cd lmp
 sudo make db
 > 输入您的 mysql root用户密码
 make
```

## 单机节点，本地运行

```
# 项目的所有配置均位于config.yaml中，grafana的默认端口为3000端口，influxdb的默认端口为8086，修改配置信息的方式如下：
 vim lmp/config/config.yaml

#run grafana
 sudo docker run -d \
   -p 3000:3000 \
   --name=grafana \
   grafana/grafana

#run influxdb，按照如下命令启动influxdb之后，会自动带有database lmp，influxdb的用户名和密码位于config.yaml中
 sudo docker run -d \
    -p 8083:8083 \
    -p 8086:8086 \
    --name influxdb \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/default.conf:/etc/influxdb/influxdb.conf \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/data:/var/lib/influxdb/data \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/meta:/var/lib/influxdb/meta \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/wal:/var/lib/influxdb/wal influxdb1.8

#run lmp
 cd lmp/
 make
 sudo ./lmp -h
```

### 观测-步骤

在shell执行`sudo ./lmp`，即启动观测功能

在单机节点上部署完成lmp并启动之后，通过浏览器访问8080端口即可。如果是本地查看，则访问localhost:8080，如果远程访问，则访问remoteip:8080即可。

8080端口返回页面如下，该页面仅用于观测指标下发，输入栏输入的是观测时间，单位是分钟。可一次下发多个指标，但是注意实际环境使用bcc的开销问题，建议单个指标下发对比开销之后，再组合多个指标观测。

![homepage](../eBPF_docs/static/imgs/homepage.png)

另外开启grafana页面，通过浏览器访问3000端口即可。如果是本地查看，则访问localhost:3000，如果远程访问，则访问remoteip:3000即可。

grafana用于指标数据观测，进入grafana之后，首先需要登录进入grafana，初始用户名和密码均为admin，之后需要配置grafana连接influxdb：

![grafana1](../eBPF_docs/static/imgs/grafana1.png)

按照自己的ip地址配置完成以后，点击save&test按钮，测试influxdb是否连接成功，出现如下提示说明连接成功：

![grafana2](../eBPF_docs/static/imgs/grafana2.png)

接下来导入/lmp/test/grafana-JSON下的lmp.json文件，即可自动创建grafana的dashboard：

![grafana3](../eBPF_docs/static/imgs/grafana3.png)

在Dashboards中Manage中，点击Import，上传lmp.json文件即可观测：

![grafana4](../eBPF_docs/static/imgs/grafana4.png)

统计时间到时后，lmp会自动关闭后台bcc插件，之后继续在8080端口页面下发指标即可。

### 扩展功能：例如机器学习
LMP通过命令行参数引入机器学习模型，你可以自己实现模型后对接到项目中，我们的想法是利用观测功能提取到数据后，利用这些数据来训练模型，目前仅可引入机器学习模型，使用示例如下：
```shell
⇒  ./lmp -h 
NAME:
   LMP - LMP is a web tool for real-time display of Linux system performance data based on BCC (BPF Compiler Collection). 
To get more info of how to use lmp:
  # lmp help


USAGE:
   lmp [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   cluster  Density peak clustering
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```
命令介绍中介绍了cluster命令，在命令行执行：
```shell
⇒  ./lmp cluster -h
NAME:
   lmp cluster - Density peak clustering

USAGE:
   lmp cluster [command options] [APP_NAME]

DESCRIPTION:
   Density peak clustering, Can be used for anomaly detection.
       example: ./lmp cluster --data /YOUR_PATH
       example: ./lmp cluster -d /YOUR_PATH

OPTIONS:
   --data value, -d value  specified the the dataset to run
   --help, -h              show help (default: false)

```
即可看到使用方法，关于如何增加一个模型见→[增加一个模型](http://lmp.kerneltravel.net/start/addmodel/)

### 如何增加插件
LMP目前支持BCC类型的插件程序，增加的方法见→[增加一个插件](http://lmp.kerneltravel.net/start/addplugin/)

### 如何调试项目
推荐使用 `postman` 下发web请求，具体调试方法见→[调试项目](http://lmp.kerneltravel.net/start/getstat/)

### 总结出的监控方案
[基于 eBPF 的 prometheus 监控方案](http://lmp.kerneltravel.net/study/ebpf_prometheus/)

[快速实现Linux系统性能数据提取、存储和可视化展示](http://lmp.kerneltravel.net/study/exporter_prometheuse_grafana/)

#

![](../eBPF_docs/static/imgs/LMP-logo.png)

# Linux microscope

LMP is a web tool for real-time display of Linux system performance data based on BCC (BPF Compiler Collection), which uses BPF (Berkeley Packet Filter), also known as eBPF. Currently, LMP is tested on ubuntu18.04 and the kernel version is 4.15.0.

## Project architecture

![](../eBPF_docs/static/imgs/LMP-arch4.png)

## Interface screenshot

![homepage2](../eBPF_docs/static/imgs/homepage2.png)

![homepage](../eBPF_docs/static/imgs/grafana.png)

![homepage](../eBPF_docs/static/imgs/data.png)


## Project structure overview  

<details>
<summary>Expand to view</summary>
<pre><code>.
.
├── LICENSE
├── README.md
├── config  # 配置
├── docs    # 文档
├── pkg     # golang 服务代码
├── plugins # python ebpf 代码
├── static  # 网页代码
├── test    # 测试数据
└── vendor  # golang vendor 
</code></pre>
</details>


##  install lmp

###  Ubuntu-source

#### Build lmp from source，The basic environment required is as follows：

- golang
- docker
- bcc

###  Install dependent docker image

```
# For prometheus 
sudo docker pull prom/prometheus
# For grafana
sudo docker pull grafana/grafana
# For Influxdb
sudo docker pull influxdb
```

### Compile and install

```
 git clone https://github.com/linuxkerneltravel/lmp
 cd lmp
 make
 sudo make install
```

##  Single machine node, Run locally

```
# Modify configuration file
 vim lmp/config/config.yaml

#run grafana
 sudo docker run -d \
   -p 3000:3000 \
   --name=grafana \
   -v /opt/grafana-storage:/var/lib/grafana \
   grafana/grafana
   
#run influxdb
 sudo docker run -d \
    -p 8083:8083 \
    -p 8086:8086 \
    --name influxdb \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/default.conf:/etc/influxdb/influxdb.conf \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/data:/var/lib/influxdb/data \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/meta:/var/lib/influxdb/meta \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/wal:/var/lib/influxdb/wal influxdb
    
#run lmp
 cd lmp/
 make
 sudo ./lmp
```

### observation

http://localhost:8080/  After logging in to grafana, view it.

### Uninstall

```
make clean
```

## Thanks for the support of the following open source projects

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [bcc] - [https://github.com/iovisor/bcc](https://github.com/iovisor/bcc)

# 