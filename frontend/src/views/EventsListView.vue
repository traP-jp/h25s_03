<template>
  <v-app>
    <v-container>
     <v-app-bar app fixed>
  <div>抽選サイト</div>
  <v-spacer></v-spacer>
  <div><v-avatar size="24">
              <v-img :alt="my_id" :src="`https://q.trap.jp/api/v3/public/icon/${my_id}`"></v-img>
            </v-avatar>
          {{ my_id }}</div>
</v-app-bar>
    <v-main>
    
    <div class="text-h5 mb-4">
      <v-card class="mx-auto my-8" elevation="16" max-width="600">
        <!--イベント作成エリア-->
        <v-card-item>
          <v-card-title> イベント作成 </v-card-title>
        </v-card-item>
        <div class="pa-4 text-center">
          <v-dialog v-model="dialog" max-width="600">
            <template v-slot:activator="{ props: activatorProps }">
              <v-btn
                class="text-none font-weight-regular"
                prepend-icon="mdi-account"
                text="イベント作成"
                variant="tonal"
                v-bind="activatorProps"
              ></v-btn>
            </template>

            <v-card prepend-icon="mdi-calendar-plus" title="イベント情報入力フォーム">
              <v-card-text>
                <v-row dense>
                  <v-col cols="12">
                    <v-text-field
                      v-model="newEvent.title"
                      label="イベント名*"
                      variant="outlined"
                      required
                    ></v-text-field>
                  </v-col>

                  <v-col cols="12">
                    <v-autocomplete
                      v-model="newEvent.admins"
                      label="管理者*"
                      :items="userList"
                      variant="outlined"
                      multiple
                      chips
                      closable-chips
                      required
                    ></v-autocomplete>
                  </v-col>

                  <v-col cols="12" md="6">
                    <v-text-field
                      v-model="newEvent.date"
                      label="開催日*"
                      type="date"
                      variant="outlined"
                      required
                    ></v-text-field>
                  </v-col>

                  <v-col cols="12" md="6">
                    <v-select
                      v-model="newEvent.is_open"
                      :items="[
                        { title: '自由参加', value: true },
                        { title: '承認制', value: false },
                      ]"
                      label="参加条件*"
                      variant="outlined"
                      required
                    ></v-select>
                  </v-col>

                  <v-col cols="12">
                    <v-textarea
                      v-model="newEvent.description"
                      label="イベント概要"
                      :counter="10000"
                      rows="3"
                      variant="outlined"
                      persistent-counter
                    ></v-textarea>
                  </v-col>
                </v-row>

                <small class="text-caption text-medium-emphasis">*は必須項目です</small>
              </v-card-text>

              <v-divider></v-divider>

              <v-card-actions>
                <v-spacer></v-spacer>

                <v-btn text="キャンセル" variant="outlined" @click="dialog = false"></v-btn>

                <v-btn color="primary" text="作成" variant="elevated" @click="createEvent"></v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </div>
      </v-card>
    </div>
    <!--イベント一覧テーブル-->
    イベント一覧 検索

    <!-- 検索エリアを追加（イベント一覧の前に） -->
    <div class="mb-4">
      <v-row align="center">
        <v-col cols="12" md="8">
          <v-text-field
            v-model="searchText"
            label="イベント名で検索"
            prepend-inner-icon="mdi-magnify"
            variant="outlined"
            clearable
            hide-details
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-btn
            :variant="showDetailSearch ? 'elevated' : 'outlined'"
            :color="showDetailSearch ? 'primary' : 'default'"
            @click="showDetailSearch = !showDetailSearch"
          >
            詳細検索
            <v-icon :icon="showDetailSearch ? 'mdi-chevron-up' : 'mdi-chevron-down'"></v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </div>

    <!-- 詳細検索フォーム -->
    <v-expand-transition>
      <v-card v-show="showDetailSearch" class="mb-4" elevation="2">
        <v-card-title>詳細検索</v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="detailSearchForm.eventName"
                label="イベント名"
                variant="outlined"
                clearable
              ></v-text-field>
            </v-col>
            <v-col cols="12" md="6">
              <v-autocomplete
                v-model="detailSearchForm.managers"
                label="管理者"
                :items="userList"
                variant="outlined"
                multiple
                chips
                clearable
              ></v-autocomplete>
            </v-col>
            <v-col cols="12" md="4">
              <v-text-field
                v-model="detailSearchForm.dateFrom"
                label="開始日以降"
                type="date"
                variant="outlined"
                clearable
              ></v-text-field>
            </v-col>
            <v-col cols="12" md="4">
              <v-text-field
                v-model="detailSearchForm.dateTo"
                label="終了日以前"
                type="date"
                variant="outlined"
                clearable
              ></v-text-field>
            </v-col>
            <v-col cols="12" md="4">
              <v-select
                v-model="detailSearchForm.dateStatus"
                label="開催ステータス"
                :items="[
                  { title: 'すべて', value: 'all' },
                  { title: '開催予定', value: 'upcoming' },
                  { title: '終了済み', value: 'ended' },
                ]"
                variant="outlined"
              ></v-select>
            </v-col>
            <v-col cols="12" md="4">
              <v-select
                v-model="detailSearchForm.participationType"
                label="参加条件"
                :items="[
                  { title: 'すべて', value: 'all' },
                  { title: '自由参加', value: '自由参加' },
                  { title: '承認制', value: '承認制' },
                ]"
                variant="outlined"
              ></v-select>
            </v-col>
            <v-col cols="12" md="4">
              <v-select
                v-model="detailSearchForm.myParticipation"
                label="自身の参加状況"
                :items="[
                  { title: 'すべて', value: 'all' },
                  { title: '参加中', value: 'participating' },
                  { title: '未参加', value: 'not_participating' },
                ]"
                variant="outlined"
              ></v-select>
            </v-col>
            <v-col cols="12" md="4">
              <v-select
                v-model="detailSearchForm.myRole"
                label="自身の権限"
                :items="[
                  { title: 'すべて', value: 'all' },
                  { title: '管理者', value: 'admin' },
                  { title: '参加者のみ', value: 'participant_only' },
                ]"
                variant="outlined"
              ></v-select>
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            variant="outlined"
            @click="
              detailSearchForm = {
                eventName: '',
                managers: [],
                dateFrom: '',
                dateTo: '',
                dateStatus: 'all',
                participationType: 'all',
                myParticipation: 'all',
                myRole: 'all',
              }
            "
          >
            クリア
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-expand-transition>
    <div>
      <v-tabs v-model="currentTab">
        <v-tab value="all_event">すべてのイベント</v-tab>
        <v-tab value="commingsoon_event">開催予定のイベント</v-tab>
        <v-tab value="attend_event">参加しているイベント</v-tab>
        <v-tab value="administrate_event">管理しているイベント</v-tab>
      </v-tabs>

      <v-data-table
        :items="events"
        :headers="headers"
        :search="tableSearch"
        :custom-filter="customFilter"
      >
        <template v-slot:[`item.title`]="{ item }">
          <router-link :to="`/${item.event_id}`" class="blue--text text-decoration-underline">
            {{ item.title }}
          </router-link>
        </template>

        <template v-slot:[`item.is_me_attendee`]="{ item }">
          <v-chip v-if="item.is_me_attendee" color="green" variant="outlined"> 参加済み </v-chip>
          <v-chip v-else color="orange" variant="outlined"> 未参加 </v-chip>
        </template>

        <template v-slot:[`item.is_open`]="{ item }">
          <v-chip v-if="item.is_open" color="green" variant="outlined"> 自由参加 </v-chip>
          <v-chip v-else color="orange" variant="outlined"> 承認制 </v-chip>
        </template>

        <template v-slot:[`item.date`]="{ item }">
          <span v-if="new Date(item.date).setHours(0, 0, 0, 0) >= new Date().setHours(0, 0, 0, 0)">
            {{ item.date }}
          </span>
          <span v-else class="text-grey"> {{ item.date }} (終了) </span>
        </template>

        <template v-slot:[`item.user_role`]="{ item }">
          <v-icon v-if="my_id in item.admins" color="primary"> mdi-crown </v-icon>
          <span v-else class="text-grey-lighten-1"> 参加者 </span>
        </template>

        <template v-slot:[`item.admins`]="{ item }">
          <div class="d-flex align-center mr-2">
            <v-avatar size="24" v-for="admin in item.admins" :key="admin">
              <v-img :alt="admin" :src="`https://q.trap.jp/api/v3/public/icon/${admin}`"></v-img>
            </v-avatar>
          </div>
        </template>
      </v-data-table>
    </div>
  </v-main>
  </v-container>
  </v-app>
</template>
<script setup lang="ts">
import { shallowRef } from 'vue'
import { ref } from 'vue'

import { computed as vueComputed } from 'vue'

const computed = vueComputed

const dialog = shallowRef(false)

const selectedManagers = ref([])
const userList = ['ogu_kazemiya', 'miyamon', 'ten_ten', 'cp20', 'Eraxyso', 'Naph1', 'hachimitsu']

const my_id = ref('')

import { onMounted } from 'vue'
import { apiClient } from '@/api/apiClient'
import type { components } from '@/api/schema'

const users = ref([])
const newEvent = ref<components['schemas']['EventBase']>({
  title: '',
  description: '',
  date: '',
  is_open: true,
  is_me_attendee: true,
  admins: [],
  attendees: [],
})

// イベント作成関数を追加
const createEvent = async () => {
  // バリデーション
  if (!newEvent.value.title || !newEvent.value.date) {
    // エラーハンドリング（実装は要件に応じて）
    return
  }

  // API呼び出し処理
  await addEvent()

  // フォームリセット
  newEvent.value = {
    title: '',
    description: '',
    date: '',
    is_open: true,
    is_me_attendee: true,
    admins: [my_id.value],
    attendees: [],
  }
  selectedManagers.value = []
  dialog.value = false
}

const fetchUsers = async () => {
  const response = await fetch('https://q.trap.jp/api/v3/users')

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`)
  }

  const data = await response.json()
  users.value = data
}

const addEvent = async () => {
  if (newEvent.value == undefined) return
  await apiClient.POST('/events', { body: newEvent.value })
}

onMounted(() => {
  fetchUsers()
})

const events = ref<components['schemas']['EventSummary'][]>([])
onMounted(async () => {
  events.value =
    (
      await apiClient.GET('/events', {
        params: {
          query: {
            ifDeleted: false,
          },
        },
      })
    ).data ?? []

  my_id.value = (await apiClient.GET('/users/me')).data?.name ?? ''
  newEvent.value.admins = [my_id.value]
})

const currentTab = ref('all_event')
const headers = [
  { title: 'イベント名', value: 'title', sortable: true },
  { title: '開催日', value: 'date', sortable: true },
  { title: '参加状況', value: 'is_me_attendee' },
  { title: '参加条件', value: 'is_open' },
  { title: '管理者', value: 'admins' },
  { title: 'あなたの権限', value: 'user_role' },
]

// 検索関連の状態を追加
const showDetailSearch = ref(false)
const searchText = ref('')
const detailSearchForm = ref({
  eventName: '',
  managers: [],
  dateFrom: '',
  dateTo: '',
  dateStatus: 'all', // 'all', 'upcoming', 'ended'
  participationType: 'all', // 'all', '自由参加', '承認制'
  myParticipation: 'all', // 'all', 'participating', 'not_participating'
  myRole: 'all', // 'all', 'admin', 'participant_only'
})

// v-data-tableのsearchプロパティを修正
const tableSearch = computed(() => {
  // 詳細検索が有効な場合は特別な値を返す
  if (showDetailSearch.value) {
    return 'detail_search'
  }
  // 名前検索のテキストがある場合はそれを返す
  if (searchText.value) {
    return searchText.value
  }
  // それ以外はタブの値を返す
  return currentTab.value
})

// customFilter関数を修正
const customFilter = (
  value: string,
  search: string,
  item?: { raw?: components['schemas']['EventSummary'] },
) => {
  if (!item?.raw) return true
  // 詳細検索が有効な場合
  if (showDetailSearch.value && search === 'detail_search') {
    // まずタブフィルターを適用
    let tabResult = true
    const currentTabValue = currentTab.value

    if (currentTabValue === 'commingsoon_event') {
      tabResult = item.raw?.date ? new Date(item.raw.date) > new Date() : true
    } else if (currentTabValue === 'attend_event') {
      tabResult = item.raw?.is_me_attendee === true
    } else if (currentTabValue === 'administrate_event') {
      tabResult = item.raw?.admins ? item.raw?.admins.includes(my_id.value) : true
    }

    if (!tabResult) return false

    // 詳細検索フィルター（既存のコード）
    const form = detailSearchForm.value

    if (form.eventName && !item.raw?.title.toLowerCase().includes(form.eventName.toLowerCase())) {
      return false
    }

    if (form.managers.length > 0) {
      const hasMatchingAdmin = form.managers.some(selectedManager => 
    item.raw?.admins?.includes(selectedManager)
  );
  if (!hasMatchingAdmin) {
    return false;
  }
    }

    const itemDate = item.raw?.date ? new Date(item.raw?.date) : new Date()
    if (form.dateFrom && itemDate < new Date(form.dateFrom)) {
      return false
    }
    if (form.dateTo && itemDate > new Date(form.dateTo)) {
      return false
    }

    if (form.dateStatus === 'upcoming' && itemDate <= new Date()) {
      return false
    }
    if (form.dateStatus === 'ended' && itemDate > new Date()) {
      return false
    }

    if (form.participationType !== 'all') {
      const isOpen = form.participationType === '自由参加'
      if (item.raw?.is_open !== isOpen) {
        return false
      }
    }

    if (form.myParticipation === 'participating' && !item.raw?.is_me_attendee) {
      return false
    }
    if (form.myParticipation === 'not_participating' && item.raw?.is_me_attendee) {
      return false
    }

    if (form.myRole === 'admin' && !(item.raw?.admins ? item.raw?.admins.includes(my_id.value) : false)) {
      return false
    }
    if (
      form.myRole === 'participant_only' && item.raw?.admins
        ? item.raw?.admins.includes(my_id.value)
        : false
    ) {
      return false
    }

    return true
  }

  // 名前検索の場合（searchTextが入力されている）
  if (
    typeof search === 'string' &&
    search !== 'all_event' &&
    search !== 'commingsoon_event' &&
    search !== 'attend_event' &&
    search !== 'administrate_event' &&
    search !== 'detail_search'
  ) {
    // 名前検索 + 現在のタブフィルター
    const nameMatch = item.raw?.title.toLowerCase().includes(search.toLowerCase())
    if (!nameMatch) return false

    // タブフィルターも適用
    const currentTabValue = currentTab.value
    if (currentTabValue === 'commingsoon_event') {
      return item.raw?.date ? new Date(item.raw.date) > new Date() : true
    } else if (currentTabValue === 'attend_event') {
      return item.raw?.is_me_attendee === true
    } else if (currentTabValue === 'administrate_event') {
      return item.raw?.admins ? item.raw?.admins.includes(my_id.value) : true
    }
    return true
  }

  // タブフィルターのみの場合
  if (search === 'all_event') {
    return true
  } else if (search === 'commingsoon_event') {
    return item.raw?.date ? new Date(item.raw.date) > new Date() : true
  } else if (search === 'attend_event') {
    return item.raw?.is_me_attendee === true
  } else if (search === 'administrate_event') {
    return item.raw?.admins ? item.raw?.admins.includes(my_id.value) : true
  }

  return true
}
</script>

