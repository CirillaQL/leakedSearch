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
    <n-space vertical>
      <ul class="video-list">
        <li class="item" v-for="video in videos" :key="video.Name">
          {{ video.Name }}
          <n-image v-model:src="video.CoverImg"/>
        </li>
      </ul>
    </n-space>
  </n-config-provider>
</template>

<script lang="ts">
  import { defineComponent, ref} from 'vue'
  import { NConfigProvider, NInput, NSpace, NButton, useMessage, NImage} from 'naive-ui'
  import SearchService from '@/services/SearchInput';
  import Video from '@/types/Video';
  import {AxiosResponse} from 'axios';

  export default defineComponent({
    components: {
      NConfigProvider,
      NInput,
      NSpace,
      NButton,
      NImage
    },
    setup() {
      const message = useMessage();
      const videos = ref([] as Video[]);
      return {
        message,
        input: ref(""),
        videos
      }
    },
    methods: {
      click () {
        SearchService.getVideos(this.input)
          .then((response: AxiosResponse) => {
            console.log(response.data.Dirtyship)
            response.data.Dirtyship.forEach((element: Video) => {
              this.videos.push(element)
            });
            console.log(this.videos)
          })
          .catch((e: Error) => {
            console.log(e);
          });
      }
    }
  })
</script>