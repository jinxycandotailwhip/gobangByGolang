# golang实现基于alpha-beta剪枝的五子棋
## 项目目录
/bin： 编译后的可执行文件  
/frontEnd：前端html页面和js  
/negative_max：alpha-beta剪枝算法实现  
/server：golang http server文件存放  
/test：单元测试  
/variable：全局变量及常量  
## 如何将项目跑起来
### 本项目的前端跨域是用nginx代理实现的
1. 在本机下载nginx
2. nginx配置config文件
```
server {
    listen 80;
    server_name localhost;
    location / {
        root /home/jinxiaoyang/workspace/hciCogn/myGobang/frontEnd;
        index gobang.html;
    }
    location /api {
        proxy_pass http://localhost:8090;
    }
}
 ```
### 配置完nginx后将后端跑起来
```
make run
```
### 访问localhost就可以看到前端页面了

### 单元测试
```
make e2etest
```
## 最后
本项目参考了一个大哥用python写的代码和另一位大哥用vue.js写的代码，深表感谢，代码仓库地址如下:  
> https://github.com/colingogogo/gobang_AI
> https://github.com/lihongxun945/gobang

我的js写的很不规范，属于是拿来了个页面能跑就行，请不要参考我的js代码。


