export default {
    window: {
        title: 'ElmoBeacon'
    },
    github: {
        title: 'GitHub リポジトリを表示',
        error: 'リポジトリへのアクセスに失敗しました'
    },
    server: {
        cn: 'China',
        us: 'America',
        intl: 'Global',
        jp: 'Japan',
        kr: 'Korea',
        tw: 'Asia'
    },
    gacha: {
        type: {
            1: '常駐訪問',
            3: '限定訪問',
            4: '軍備拡張',
            5: 'スタートダッシュ訪問',
            6: '選択訪問・人形',
            7: '選択訪問・装備',
            8: 'ミステリーボックス',
            9: '新私服配給',
        },
        statistic: {
            totalCount: 'Total Counter',
            pityCount: 'Pity Counter',
            rank5Data: '5-star Data',
            rank4Data: '4-star Data',
            rank3Data: '3-star Data',
            rank5Avg: 'Avg Pulls per 5-star',
            upRank5Avg: 'Avg Pulls per Up 5-star',
            nonUpRate: 'NonUp 5-star Rate',
        },
        records: {
            title: 'Pull Records'
        }
    },
    sync: {
        button: {
            title: 'Synchronize Records',
            tip: 'Pull records from the server and stops when it matches the latest record in the local database'
        },
        loading: '',
        status: {
            checkingUser: "ユーザー情報を確認中...",
            readingPoolTypes: "カードプールタイプを取得中...",
            fetchingPool: {
                1: "常駐訪問の記録を同期中...",
                2: "",
                3: "限定訪問の記録を同期中...",
                4: "軍備拡張の記録を同期中...",
                5: "スタートダッシュ訪問の記録を同期中...",
                6: "選択訪問・人形の記録を同期中...",
                7: "選択訪問・装備の記録を同期中...",
                8: "ミステリーボックスの記録を同期中...",
                9: "新私服配給の記録を同期中...",
                10: ""
            }
        },
        result: {
            success: {
                title: '{server} {uid} Synchronization Success',
                changed: '{count} new records added for {poolType}',
                unchanged: 'No new records added'
            },
            error: {
                cn: 'Synchronization Error(CN)',
                os: 'Synchronization Error(OS)'
            }
        }
    },
    import: {
        button: {
            title: "インポート記録"
        },
        loading: "インポートしてマージ中...",
        dialog: {
            title: "インポート記録",
            uid: "UID",
            uidPlaceholder: "UIDを入力してください",
            server: "サーバー",
            file: "記録ファイル",
            filePlaceholder: "インポートする記録ファイルを選択してください",
            fileSelected: "ファイル選択済み",
            browse: "参照",
            tips: "現在GFL2 HELPの記録のみインポートに対応しています。データ破損を防ぐため、無関係なファイルを選択しないでください。",
            cancel: "キャンセル",
            startImport: "インポート開始",
            importing: "インポート中...",
            uidRequired: "UIDは必須です",
            fileRequired: "インポートファイルを選択してください",
            importError: "インポートエラー",
            invalidJson: "ファイル形式が不正です。正しい記録ファイルを選択してください"
        },
        result: {
            success: {
                title: "{server} {uid} インポート成功",
                changed: "{poolType}：マージ成功 {successCount} 件、失敗 {failCount} 件",
                unchanged: "新規データなし"
            }
        }
    },
    theme: {
        change: "テーマを変更",
        custom: "カスタムテーマ",
        from: "開始色",
        to: "終了色",
        preview: "プレビュー",
        popular: "人気グラデーション",
        apply: "適用",
        cancel: "キャンセル"
    },
    user: {
        change: "ユーザー切替",
        deleteTitle: "ユーザー削除",
        deleteHint: "このユーザーとすべてのガチャ記録を削除しますか？",
        confirmBtn: "確認",
        cancelBtn: "キャンセル",
        deleteSuccess: "削除しました"
    },
    language: {
        change: "言語変更"
    },
}