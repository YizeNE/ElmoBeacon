import { defineStore } from "pinia";
import { ref } from "vue";
import { model } from "../../wailsjs/go/models.ts";
import {GetSetting, GetUserList, SetSetting, DeleteUser} from "../../wailsjs/go/handler/App";
import { NotifyError, NotifySuccess } from "../utils/notify.ts";
import User = model.User;

const getSettingUserId = async () => {
    let settingUserId = 0

    await GetSetting("lastUserId").then(res => {
        if (res) {
            settingUserId = parseInt(res)
        }
    }).catch(err => {
        NotifyError('Error', err)
    })

    return settingUserId
}

export const useUserStore = defineStore('user', () => {
    const userId = ref<number>()
    const userList = ref<User[]>([])

    const updateUserId = async (newUserId: number) => {
        await SetSetting("lastUserId", newUserId.toString()).then(() => {
            userId.value = newUserId
        }).catch(err => {
            NotifyError('Error', err)
        })
    }

    const updateUserList = async () => {
        await GetUserList().then((res) => {
            if (res) {
                userList.value = res;
                if (!userId.value) {
                    userId.value = userList.value[0].id;
                }
            }
        }).catch(err => {
            NotifyError('Error', err)
        })
    }

    const deleteUser = async (id: number, successMsg: string) => {
    await DeleteUser(id).then(async () => {
        NotifySuccess(successMsg, '')
        await GetUserList().then((res) => {
            userList.value = res || []
        }).catch((err) => {
            userList.value = []
            NotifyError('Error', err)
        })
        if (userId.value === id) {
            if (userList.value.length > 0) {
                await updateUserId(userList.value[0].id)
            } else {
                userId.value = undefined
                await SetSetting("lastUserId", "0")
            }
        }
    }).catch((err: any) => {
        NotifyError('Error', err)
    })
}

    const init = async () => {
        userId.value = await getSettingUserId()
        await updateUserList()
    }

    return {userId, userList, updateUserId, updateUserList, deleteUser, init}
})