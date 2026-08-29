export default {
    window: {
        title: '艾莫信标',
    },
    github: {
        title: '查看GitHub仓库',
        error: '访问仓库失败'
    },
    server: {
        cn: '国服',
        us: '美服',
        intl: '全球服',
        jp: '日服',
        kr: '韩服',
        tw: '亚服'
    },
    gacha: {
        type: {
            1: '常规采购',
            3: '定向采购',
            4: '军备提升',
            5: '初始采购',
            6: '自选采购·人形',
            7: '自选采购·军备',
            8: '神秘箱',
            9: '新装采购'
        },
        statistic: {
            totalCount: '总抽数',
            pityCount: '保底进度',
            rank5Data: '五星数据',
            rank4Data: '四星数据',
            rank3Data: '三星数据',
            rank5Avg: '五星平均抽数',
            upRank5Avg: '限定五星平均抽数',
            nonUpRate: '五星歪率',
        },
        loading: '加载中...',
        records: {
            title: '抽卡记录',
            tip: '官方服务器仅保存近六个月内的记录，请妥善保存本地记录以免丢失'
        }
    },
    sync: {
        button: {
            title: '同步记录',
            tip: '从服务器拉取抽卡记录，匹配到本地数据库最新记录后停止'
        },
        loading: '',
        status: {
            checkingUser: "正在验证用户信息...",
            readingPoolTypes: "正在获取卡池类型...",
            fetchingPool: {
                1: "正在同步常规采购记录...",
                2: "",
                3: "正在同步定向采购记录...",
                4: "正在同步军备提升记录...",
                5: "正在同步初始采购记录...",
                6: "正在同步自选采购·人形记录...",
                7: "正在同步自选采购·军备记录...",
                8: "正在同步神秘箱记录...",
                9: "正在同步新装采购记录...",
                10: ""
            }
        },
        result: {
            success: {
                title: '{server} {uid} 同步成功',
                changed: '{poolType} 新增 {count} 条',
                unchanged: '无新增数据'
            },
            error: {
                cn: '国服同步出错',
                os: '国际服同步出错'
            }
        },
        // 新增弹窗相关的翻译
        dialog: {
            title: '同步记录',
            uid: 'UID',
            uidPlaceholder: '请输入您的UID',
            gameDataDir: '文件目录',
            gameDataDirPlaceholder: '请选择文件目录',
            browse: '浏览',
            tips: '请选择包含用户凭证的文件目录。',
            cancel: '取消',
            startSync: '开始同步',
            syncing: '同步中...',
            uidRequired: 'UID不能为空',
            dirRequired: '文件目录不能为空',
            syncError: '同步错误',
            selectDirError: '选择目录失败'
        }
    },
    version: {
        checkUpdate: '检查更新',
        update: {
            notify: '发现新版本',
            changelog: '更新日志',
            latest: '已是最新版本',
            confirm: '更新',
            cancel: '取消',
            downloading: "正在下载更新"
        },
    },
    theme: {
        change: "更改主题",
        custom: "自定义主题",
        from: "起始颜色",
        to: "结束颜色",
        preview: "预览效果",
        popular: "流行渐变",
        apply: "应用",
        cancel: "取消"
    },
    user:{
        change:"切换用户"
    },
    language:{
        change:"更改语言"
    }
}