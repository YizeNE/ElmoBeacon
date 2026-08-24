<p align="center">
  <a href="https://github.com/YizeNE/ElmoBeacon/actions"><img src="https://img.shields.io/github/actions/workflow/status/YizeNE/ElmoBeacon/release.yml?style=flat-square" alt="Build"></a> <a href="https://github.com/YizeNE/ElmoBeacon/releases"><img src="https://img.shields.io/github/v/release/YizeNE/ElmoBeacon?style=flat-square&sort=semver" alt="Release"></a> <a href="https://github.com/YizeNE/ElmoBeacon/releases"><img src="https://img.shields.io/github/downloads/YizeNE/ElmoBeacon/total?style=flat-square" alt="Downloads"></a> <a href="https://github.com/YizeNE/ElmoBeacon"><img src="https://img.shields.io/github/stars/YizeNE/ElmoBeacon?style=flat-square&logo=github&logoColor=white" alt="Stars"></a>
</p>

<p align="center">Languages: 简体中文 | <a href="./README.en.md">English</a></p>

## 图片预览

![image](preview.png)

## 项目介绍

<p><a href="https://github.com/YizeNE/ElmoBeacon" rel="nofollow">ElmoBeacon</a> 是一款用以存储和分析《少女前线2：追放》抽卡记录的工具，兼容暗冬服务器（中国大陆、北美）和好玩服务器（全球、亚洲、日本、韩国）。</p>
<p>这是一个开源项目，由<a href="https://gf2.mcc.wiki" rel="nofollow">MccWiki</a> 和社区共同维护。如果您对该项目有任何想法，欢迎提交PR！</p>

## 运行需求

- 少前2客户端
- Fiddler
- Go 1.18+ （仅开发）
- Wails v2 CLI（仅开发）
- Node.js 18+（仅开发）

## 使用教程

1. 由于本地游戏日志不再储存访问令牌，因此需要使用抓包软件来获取[下载Fiddler](https://api.getfiddler.com/fc/latest)（默认情况下Fiddler只能抓取HTTP，因此还需要进行配置[参考](https://developer.aliyun.com/article/1342462)）
1. 配置好Fiddler后确保Fiddler正在运行，然后在游戏内查看抽卡记录，此时Fiddler中会出现一条Host为`gf2-gacha-record-xxx`，URL为`/list?xxx`的记录，右键此条记录**Save>Selected Sessions>as Text**将该条记录保存至本地
1. 打开本软件，点击同步记录后输入**自己的UID**并选择**步骤2中保存至本地的记录**即可开始获取抽卡记录（**⚠本地的记录具有时效性，请定期按步骤2重新获取）**

## 本地开发

1. 克隆本项目到本地
2. 在项目根目录运行`wails dev`命令
