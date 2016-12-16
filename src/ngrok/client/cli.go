package client

import (
	"flag"
	"fmt"
	"ngrok/version"
	"os"
)

const usage1 string = `Usage: %s [OPTIONS] <本地端口或地址>
Options:
`

const usage2 string = `
示例:
	ngrok 80
	ngrok -subdomain=example 8080
	ngrok -proto=tcp 22
	ngrok -hostname="example.com" -httpauth="user:password" 10.0.0.1


Advanced usage: ngrok [OPTIONS] <command> [command args] [...]
命令:
	ngrok start [tunnel] [...]    通过配置文件中的名称启动隧道
	ngork start-all               启动配置文件中定义的所有隧道
	ngrok list                    从配置文件列出隧道名称
	ngrok help                    显示帮助
	ngrok version                 显示ngrok版本

示例:
	ngrok start www api blog pubsub
	ngrok -log=stdout -config=ngrok.cfg start ssh
	ngrok start-all
	ngrok version

`

type Options struct {
	config    string
	logto     string
	loglevel  string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}


	config := flag.String(
		"config",
		"ngrok.cfg",
		"ngrok配置文件的路径. (default: $HOME/.ngrok)")

	logto := flag.String(
		"log",
		"none",
		"将日志消息写入此文件. 'stdout' 和 'none' 有特殊意义")

	loglevel := flag.String(
		"log-level",
		"ERROR",
		"要记录的消息级别. 其中之一: DEBUG, INFO, WARNING, ERROR")

	authtoken := flag.String(
		"authtoken",
		"",
		"用于标识帐户的身份验证令牌")

	httpauth := flag.String(
		"httpauth",
		"",
		"username:password 保护公共隧道端点的HTTP基本认证信任")

	subdomain := flag.String(
		"subdomain",
		"",
		"从ngrok服务器请求自定义子域. （仅限HTTP和HTTS模式）")

	hostname := flag.String(
		"hostname",
		"",
		"从ngrok服务器请求自定义域名. （仅限HTTP和HTTPS）（需要自定义域名CNAME解析至ngrok服务器）")

	protocol := flag.String(
		"proto",
		"http+https",
		"通过隧道的流量的协议 {'http', 'https', 'tcp'} (default: 'http+https')")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		loglevel:  *loglevel,
		httpauth:  *httpauth,
		subdomain: *subdomain,
		protocol:  *protocol,
		authtoken: *authtoken,
		hostname:  *hostname,
		command:   flag.Arg(0),
	}

	switch opts.command {
	case "list":
		opts.args = flag.Args()[1:]
	case "start":
		opts.args = flag.Args()[1:]
	case "start-all":
		opts.args = flag.Args()[1:]
	case "version":
		fmt.Println(version.MajorMinor())
		os.Exit(0)
	case "help":
		flag.Usage()
		os.Exit(0)
	case "":

		err = fmt.Errorf("错误: 指定一个要隧道连接到本地的端口，或" +
			"一条ngrok命令.\n\n例: \n要映射端口80，请运行 " +
			"'ngrok 80'\n要启动配置文件ngrok.cfg里所有隧道，请运行'ngrok start-all'")
		return
       

	default:
		if len(flag.Args()) > 1 {
			err = fmt.Errorf("您可以在命令行上指定一个端口以便隧道到达 %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
		
	}

	return
}
