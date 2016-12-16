#备份并做了部分修改与汉化工作

###已经编译好的客户端及服务端下载地址：https://github.com/chengangwin/ngrok/releases

客户端为：ngrok</br>
启动方式1：ngrok.exe 80</br>
启动方式2：ngrok.exe -log-level=error -config ngrok.cfg start-all</br></br>
服务端为：ngrokd</br>
启动方式：ngrokd.exe -domain="gdmpmiu.tk" -httpAddr=":80" -httpsAddr=":443" -tunnelAddr=":4443"</br></br>
以上启动命令均是在cmd模式下运行。</br></br>

此版本在原版1.7的基础上进行了部分汉化以及优化工作，让客户端可以连接其它版本的服务端。
配置文件使用与1.7版一致！


###原源码地址：https://github.com/inconshreveable/ngrok
