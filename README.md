#备份并做了部分优化修改与汉化工作

###已经编译好的客户端及服务端下载地址：</br>https://github.com/chengangwin/ngrok/releases

客户端为：ngrok</br>
启动方式1：ngrok 80</br>
启动方式2：ngrok  start-all</br></br>
客户端配置文件：</br>
1、新建一个txt文本文件，输入：server_addr: "www.gdmpmiu.tk:4443" 后保存，并将文件名改为：ngrok.cfg</br>
2、ngrok.cfg 与ngrok客户端 放在同一目录下。</br></br></br>

服务端为：ngrokd</br>
启动方式：ngrokd -domain="gdmpmiu.tk"</br></br>
以上启动命令均是在cmd模式下运行。</br></br>

此版本在原版1.7的基础上进行了部分汉化以及优化工作，让客户端可以连接其它版本的服务端。
配置文件使用与1.7版一致！


###原源码地址：https://github.com/inconshreveable/ngrok
