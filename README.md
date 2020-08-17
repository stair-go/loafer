### 起因: 

迫于最近行情不错,公司又明令禁止上班进行与工作无关活动,且不得使用QQ,微信等通讯工具。

奈何人力终有穷,身为码农,搬砖累了想摸鱼风险又太高,手上有台闲置的服务器(当然本地启动也行),便在周末撸了个本项目--较全面+安全摸鱼+社交(拉上基友们在一个群,交流频率提升50%)的理财工具。

(***摸鱼是不好的***,一定要先完成本职工作!)

---

### 项目说明: 
均使用原生Go库,无任何依赖组件,拉取即可启动,拉取基友一起摸鱼吧。
实现很粗糙,欢迎大家提出建议,一起学习,欢迎star~

##### 项目地址 
[https://github.com/stair-go/loafer](https://github.com/stair-go/loafer)

##### 功能
- 工作日9:00 -- 15:00 每半个小时(默认),抓取配置的基金、大盘及股票数据至钉钉机器人
- 通过Http请求实时请求爬虫和推送
- 通过Http请求,运行中添加基金或股票代码

#####  预览
![image.png](https://upload-images.jianshu.io/upload_images/18017519-4ea1b2858f514267.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![image.png](https://upload-images.jianshu.io/upload_images/18017519-3537c7cd1af3331e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


##### 安装：
***环境需求***:
- 钉钉群,并创建钉钉机器人
[https://help.aliyun.com/document_detail/112831.html](https://help.aliyun.com/document_detail/112831.html)
注:若要安全设置,可添加"加"、"加特么的"关键字
![image.png](https://upload-images.jianshu.io/upload_images/18017519-e50ccf104a5b84b4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- go 环境
  - linux 安装:
[https://www.jianshu.com/p/c43ebab25484](https://www.jianshu.com/p/c43ebab25484)
  - windows安装
[https://www.jianshu.com/writer#/notebooks/37276991/notes/75888711/preview](https://www.jianshu.com/writer#/notebooks/37276991/notes/75888711/preview)

***拉取本项目代码***:
```
git clone https://github.com/stair-go/loafer.git
```
***配置机器人地址***:
![image.png](https://upload-images.jianshu.io/upload_images/18017519-9bca676c01727f5b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
注:其中"CodeAttr"为 基金代码,"Exponent"为大盘指数,"Stock"为股票代码,可在启动前配置,也可启动后,通过http接口添加;

##### 接口说明：
默认端口为8189

###### 查询
实时请求爬虫,并推送至dd群
- 示例: "ip:prot/" or "ip:prot/query"

###### 添加基金代码

- 路径: "ip:prot/add/fund?code="
- 示例: 127.0.0.1:8189/add?code=213001

###### 添加股票代码

- 路径: "ip:prot/add/stock/?code="
- 示例: 127.0.0.1:8189/add/stock?code=sz002905

###### 删除基金代码

- 路径: "ip:prot/delete/fund?code="
- 示例: 127.0.0.1:8189/delete?code=213001

###### 删除股票代码

- 路径:  "ip:prot/delete/stock/?code="
- 示例: 127.0.0.1:8189/delete/stock?code=sz002905
