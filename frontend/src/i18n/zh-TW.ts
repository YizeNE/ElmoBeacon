export default {
    window: {
        title: '艾莫信標'
    },
    github: {
        title: '查看 GitHub 儲存庫',
        error: '存取儲存庫失敗'
    },
    server: {
        cn: '陸服',
        us: '美服',
        intl: '全球服',
        jp: '日服',
        kr: '韓服',
        tw: '亞服'
    },
    gacha: {
        type: {
            1: '常規採購',
            3: '定向採購',
            4: '軍備提升',
            5: '初始採購',
            6: '自選採購·人形',
            7: '自選採購·軍備',
            8: '神秘箱',
            9: '新裝採購',
        },
        statistic: {
            totalCount: '總抽數',
            pityCount: '保底進度',
            rank5Data: '五星數據',
            rank4Data: '四星數據',
            rank3Data: '三星數據',
            rank5Avg: '五星平均抽數',
            upRank5Avg: '限定五星平均抽數',
            nonUpRate: '五星歪率',
        },
        loading: '載入中...',
        records: {
            title: '抽卡記錄',
            tip: '官方服務器僅保存近六個月內的記錄，請妥善保存本地記錄以免丟失'
        }
    },
    sync: {
        button: {
            title: '同步記錄',
            tip: '從伺服器拉取抽卡記錄，匹配到本地數據庫最新記錄後停止'
        },
        loading: '',
        status: {
            checkingUser: "正在驗證使用者資訊...",
            readingPoolTypes: "正在取得卡池類型...",
            fetchingPool: {
                1: "正在同步常規採購紀錄...",
                2: "",
                3: "正在同步定向採購紀錄...",
                4: "正在同步軍備提升紀錄...",
                5: "正在同步初始採購紀錄...",
                6: "正在同步自選採購·人形紀錄...",
                7: "正在同步自選採購·軍備紀錄...",
                8: "正在同步神秘箱紀錄...",
                9: "正在同步新裝採購紀錄...",
                10: ""
            }
        },
        result: {
            success: {
                title: '{server} {uid} 同步成功',
                changed: '{poolType} 新增 {count} 條',
                unchanged: '无無新增數據'
            },
            error: {
                cn: '陸服同步出錯',
                os: '國際服同步出錯'
            }
        },
        dialog: {
            "title": "同步記錄",
            "uid": "UID",
            "uidPlaceholder": "請輸入您的UID",
            "gameDataDir": "檔案目錄",
            "gameDataDirPlaceholder": "請選擇檔案目錄",
            "browse": "瀏覽",
            "tips": "請選擇包含用戶憑證的檔案目錄。",
            "cancel": "取消",
            "startSync": "開始同步",
            "syncing": "同步中...",
            "uidRequired": "UID不能為空",
            "dirRequired": "檔案目錄不能為空",
            "syncError": "同步錯誤",
            "selectDirError": "選擇目錄失敗"
        }
    },
    version: {
        checkUpdate: '檢查更新',
        update: {
            notify: '發現新版本',
            changelog: '更新日誌',
            latest: '已是最新版本',
            confirm: '更新',
            cancel: '取消'
        },
    },
    theme: {
        change: "更改主題",
        custom: "自訂主題",
        from: "起始顏色",
        to: "結束顏色",
        preview: "預覽效果",
        popular: "流行漸變",
        apply: "應用",
        cancel: "取消"
    }
}