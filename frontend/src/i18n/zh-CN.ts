export default {
    window: {
        title: '艾莫信标'
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
            8: '神秘箱'
        },
        statistic: {
            totalCount: '记录总数',
            pityCount: '保底进度',
            rank5Data: '五星数据',
            rank4Data: '四星数据',
            rank3Data: '三星数据',
            rank5Avg: '五星平均抽数',
            upRank5Avg: 'Up五星平均抽数',
            nonUpRate: '五星歪率',
        },
        records: {
            title: '抽卡记录',
            tip: '记录名称读取自游戏本体，因此不管你选择什么语言，国服只会显示简体中文，国际服无法显示简体中文(如果你选择了简体中文，将使用繁体中文代替)'
        }
    },
    sync: {
        button: {
            title: '同步记录',
            tip: '从服务器拉取抽卡记录，匹配到本地数据库最新记录后停止'
        },
        loading: '正在同步记录...',
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
        update: {
            notify: '有新的版本可用，是否更新？',
            latest: '已是最新版本',
            confirm: '是',
            cancel: '否'
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
  }
}