<p>
  <a href="https://github.com/YizeNE/ElmoBeacon/blob/main/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/YizeNE/ElmoBeacon"/>
  </a>
  <a href="https://github.com/YizeNE/ElmoBeacon/issues">
    <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg"/>
  </a>
  <a href="https://github.com/YizeNE/ElmoBeacon/actions/workflows/release.yml" rel="nofollow">
    <img src="https://img.shields.io/github/actions/workflow/status/YizeNE/ElmoBeacon/release.yml"/>
  </a>
  <a href="https://github.com/YizeNE/ElmoBeacon/releases" rel="nofollow">
    <img src="https://img.shields.io/github/v/release/YizeNE/ElmoBeacon"/>
  </a>
</p>


## 图片预览

![image](preview.png)

## 项目介绍

<p><a href="https://github.com/YizeNE/ElmoBeacon" rel="nofollow">ElmoBeacon</a> 是一款用以存储和分析《少女前线2：追放》抽卡记录的工具，兼容暗冬服务器（中国大陆、北美）和好玩服务器（全球、亚洲、日本、韩国）。</p>
<p>这是一个开源项目，由<a href="https://gf2.mcc.wiki" rel="nofollow">MccWiki</a> 和社区共同维护。如果您对该项目有任何想法，欢迎提交PR！</p>

## 项目需求

- 少前2客户端
- Fiddler
- Go 1.18+ (仅开发)
- Wails (仅开发)

## 使用教程

1. 由于游戏日志不再存储用户凭证，因此需要使用抓包软件来获取[点此下载Fiddler](https://api.getfiddler.com/fc/latest)（默认情况下Fiddler只能抓取HTTP请求，因此还需要进行配置[参考](https://developer.aliyun.com/article/1342462)）
1. 配置好Fiddler后确保Fiddler正在运行，然后在游戏内点击访问记录，此时Fiddler记录中会出现一条Host为`gf2-gacha-record-xxx`，URL为`/list?xxx`的记录，右键此条记录**Save>Selected Sessions>as Text**将该条记录保存至本地
1. 打开本软件，点击同步记录后输入自己的UID并选择步骤2中保存至本地的记录即可开始获取抽卡记录（**⚠进行该步骤时请关闭Fiddler以及其他代理软件以免发生意料之外的错误，此外本地的记录具有时效性，请定期按步骤2重新获取）**

## 本地开发

1. clone the project to local
2. run `wails dev` to dev,`wails build` to build
