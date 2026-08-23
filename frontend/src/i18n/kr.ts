export default {
    window: {
        title: 'ElmoBeacon'
    },
    github: {
        title: 'GitHub 저장소 보기',
        error: '저장소 접근 실패'
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
            1: '일반 발주',
            3: '지정 발주',
            4: '군비 강화',
            5: '시작 발주',
            6: '선택 발주·인형',
            7: '선택 발주·군비',
            8: '미스터리 박스',
            9: '신규 의상 발주'
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
            checkingUser: "사용자 정보 확인 중...",
            readingPoolTypes: "카드풀 유형 불러오는 중...",
            fetchingPool: {
                1: "일반 발주 기록 동기화 중...",
                2: "",
                3: "지정 발주 기록 동기화 중...",
                4: "군비 강화 기록 동기화 중...",
                5: "시작 발주 기록 동기화 중...",
                6: "선택 발주·인형 기록 동기화 중...",
                7: "선택 발주·군비 기록 동기화 중...",
                8: "미스터리 박스 기록 동기화 중...",
                9: "신규 의상 발주 기록 동기화 중...",
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
    theme: {
        change: "테마 변경",
        custom: "사용자 정의 테마",
        from: "시작 색상",
        to: "종료 색상",
        preview: "미리보기",
        popular: "인기 그라데이션",
        apply: "적용",
        cancel: "취소"
    }
}