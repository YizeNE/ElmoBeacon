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
    import: {
        button: {
            title: "가져오기 기록"
        },
        loading: "가져와서 병합 중...",
        dialog: {
            title: "가져오기 기록",
            uid: "UID",
            uidPlaceholder: "UID를 입력하세요",
            server: "서버",
            file: "기록 파일",
            filePlaceholder: "가져올 기록 파일을 선택하세요",
            fileSelected: "파일 선택됨",
            browse: "찾아보기",
            tips: "현재 GFL2 HELP 기록만 가져올 수 있습니다. 데이터 손상을 방지하기 위해 관련 없는 파일을 선택하지 마세요.",
            cancel: "취소",
            startImport: "가져오기 시작",
            importing: "가져오는 중...",
            uidRequired: "UID는 필수입니다",
            fileRequired: "가져올 파일을 선택하세요",
            importError: "가져오기 오류",
            invalidJson: "파일 형식이 올바르지 않습니다. 올바른 기록 파일을 선택하세요"
        },
        result: {
            success: {
                title: "{server} {uid} 가져오기 성공",
                changed: "{poolType}: 병합 성공 {successCount}건, 실패 {failCount}건",
                unchanged: "새 데이터 없음"
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
    },
    user: {
        change: "사용자 전환",
        deleteTitle: "사용자 삭제",
        deleteHint: "이 사용자와 모든 가챠 기록을 삭제하시겠습니까?",
        confirmBtn: "확인",
        cancelBtn: "취소",
        deleteSuccess: "삭제됨"
    },
    language: {
        change: "언어 변경"
    }
}