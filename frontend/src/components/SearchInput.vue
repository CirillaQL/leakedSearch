<template>
  <n-config-provider>
    <n-space vertical>
      <n-input-group>
        <n-input placeholder="搜索内容" style="width: 50%;" v-model:value="input">
        </n-input>
        <n-button @click="click">
          搜索
        </n-button>
      </n-input-group>
    </n-space>
  </n-config-provider>
</template>

<script lang="ts">
  import { defineComponent, ref} from 'vue'
  import { NConfigProvider, NInput, NSpace, NButton, useMessage } from 'naive-ui'
  import SearchService from '@/services/SearchInput';
  import {AxiosResponse} from 'axios';

  export default defineComponent({
    components: {
      NConfigProvider,
      NInput,
      NSpace,
      NButton
    },
    setup() {
      const message = useMessage();
      return {
        message,
        input: ref("")
      }
    },
    methods: {
      click () {
        SearchService.getVideos(this.input)
          .then((response: AxiosResponse) => {
            console.log(response.data)
          })
      }
    }
  })
</script>