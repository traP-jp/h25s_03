
<template>
  <v-container>
    {ページタイトル}
    <div class="text-h5 mb-4">
    <v-card
    class="mx-auto my-8"
    elevation="16"
    max-width="600"
    >

    <!--イベント作成エリア-->
    <v-card-item>
      <v-card-title>
        イベント作成
      </v-card-title>
    </v-card-item>
    <v-card-text>
      イベント名を入力してください。
    </v-card-text>
    <div class="pa-4 text-center">
    <v-dialog
      v-model="dialog"
      max-width="600"
    >
      <template v-slot:activator="{ props: activatorProps }">
        <v-btn
          class="text-none font-weight-regular"
          prepend-icon="mdi-account"
          text="イベント作成"
          variant="tonal"
          v-bind="activatorProps"
        ></v-btn>
      </template>

      <v-card
        prepend-icon="mdi-account"
        title="イベント情報入力フォーム"
      >
        <v-card-text>
          <v-row dense>
            <v-col
              cols="12"
              md="4"
              sm="6"
            >
              <v-text-field
                label="イベント名*"
                required
              ></v-text-field>

            </v-col>
            <v-autocomplete
              v-model="selectedManagers"
              label="管理者*"
              :items="userList"
              variant="underlined"
              multiple
              chips
              closable-chips
            ></v-autocomplete>
            <v-col
              cols="12"
              md="4"
              sm="6"
            >
              <v-text-field
                label="開催日*"
                required
              ></v-text-field>
            </v-col>
            <v-col
              cols="12"
              sm="6"
            >
              <v-select
                :items="['自由参加', '招待制']"
                label="参加条件*"
                required
              ></v-select>
            </v-col>

            <v-col
              cols="12"
              sm="6"
            >
            <div class="mb-2">イベント概要</div>
                <v-textarea
                  :counter="10000"
                  class="mb-2"
                  rows="2"
                  variant="outlined"
                  persistent-counter
                ></v-textarea>
            </v-col>
          </v-row>
          
          <small class="text-caption text-medium-emphasis">*indicates required field</small>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="Close"
            variant="plain"
            @click="dialog = false"
          ></v-btn>

          <v-btn
            color="primary"
            text="Save"
            variant="tonal"
            @click="dialog = false"
          ></v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
    </v-card>
    </div>
    <!--イベント一覧テーブル-->
    イベント一覧
    検索
    <v-card
        prepend-icon="mdi-account"
        title="詳細検索"
      >
        <v-card-text>
          <v-row dense>
            <v-col
              cols="12"
              md="4"
              sm="6"
            >
              <v-text-field
                label="イベント名*"
                required
              ></v-text-field>

            </v-col>
            <v-autocomplete
              v-model="selectedManagers"
              label="管理者*"
              :items="userList"
              variant="underlined"
              multiple
              chips
              closable-chips
            ></v-autocomplete>
            <v-col
              cols="12"
              md="4"
              sm="6"
            >
              <v-text-field
                label="開催日*"
                required
              ></v-text-field>
            </v-col>
            <v-col
              cols="12"
              sm="6"
            >
              <v-select
                :items="['自由参加', '招待制']"
                label="参加条件*"
                required
              ></v-select>
            </v-col>

            <v-col
              cols="12"
              sm="6"
            >
            <div class="mb-2">イベント概要</div>
                <v-textarea
                  :counter="10000"
                  class="mb-2"
                  rows="2"
                  variant="outlined"
                  persistent-counter
                ></v-textarea>
            </v-col>
          </v-row>
          
          <small class="text-caption text-medium-emphasis">*indicates required field</small>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            text="Close"
            variant="plain"
            @click="dialog = false"
          ></v-btn>

          <v-btn
            color="primary"
            text="Save"
            variant="tonal"
            @click="dialog = false"
          ></v-btn>
        </v-card-actions>
      </v-card>
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
      :search="tabFilter"
      :custom-filter="customFilter">
      
      <template v-slot:item.title="{ item }">
        <router-link :to="`/${item.event_id}`" class="blue--text text-decoration-underline">
          {{ item.title }}
        </router-link>
      </template>
      
      <template v-slot:item.is_open="{ item }">
        <v-chip
          v-if="item.is_open"
          color="green"
          variant="outlined">
          自由参加
        </v-chip>
        <v-chip
          v-else
          color="orange"
          variant="outlined">
          承認制
        </v-chip>
      </template>
      
      <template v-slot:item.date="{ item }">
        <span v-if="new Date(item.date) > new Date()">
          {{item.date}} (予定)
        </span>
        <span v-else class="text-grey">
          {{item.date}} (終了)
        </span>
      </template>
      
      <template v-slot:item.is_me_admin="{ item }">
        <v-icon v-if="item.is_me_admin" color="primary">
          mdi-crown
        </v-icon>
        <span v-else class="text-grey-lighten-1">
          参加者
        </span>
      </template>
    </v-data-table>
  </div>
  </v-container>
  
</template>
<script setup lang="ts">
  import { shallowRef } from 'vue'
  import { ref } from 'vue'

  import { computed as vueComputed } from 'vue'

const computed = vueComputed

  const dialog = shallowRef(false)
  const tab = ref(0)

  const selectedManagers = ref([])
  const userList = ['ogu_kazemiya','miyamon','ten_ten']

  import { onMounted } from 'vue'
  import { apiClient } from '@/api/apiClient'
  import type { components } from '@/api/schema'

  const users = ref([])
  const newEvent = ref({
    "title": "",
    "description": "",
    "date": new Date(),
    "is_open": true,
    "is_me_admin": true
  })
  

  const fetchUsers = async () => {
      const response = await fetch('https://q.trap.jp/api/v3/users')

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()
      users.value = data
  }

  // const addEvent = async () =>{
  //     const response = (await apiClient.POST('/events',newEvent)).data ?? [];
  // }

  onMounted(() => {
    fetchUsers()
  })

  const events = ref<components['schemas']['EventSummary'][]>([])
  onMounted(async() => {
    events.value = (await apiClient.GET('/events', {
      params: {
        query: {
          ifDeleted: false
        }
      }
    })).data ?? [];
  })

  const currentTab = ref('all_event')
const headers = [
  { title: 'イベント名', value: 'title',sortable: true },
  { title: '開催日', value: 'date',sortable: true },
  { title: '参加条件', value: 'is_open'},
  { title: '権限', value:'is_me_admin'}
]

const tabFilter = computed(() => {
  return currentTab.value
})

const customFilter = (value: any, search: string, item: any) => {
  if (search === 'all_event') {
    return true // すべてのイベントを表示
  } else if (search === 'commingsoon_event') {
    return true //#要変更
  }else if (search === 'attend_event') {
    return true //#要変更
  } else if (search === 'administrate_event') {
    return true //#要変更
  }
  return true
}
</script>
  