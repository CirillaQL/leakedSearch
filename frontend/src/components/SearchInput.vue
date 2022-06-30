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
      <div class="video-list" style="list-style: none;">
        <div class="item" v-for="video in videos" :key="video.Name" style="width: 50%;">
          <n-card>
            <template #cover>
              <n-image :src="video.CoverImg" :href="video.URL"/>
            </template>
            <div class="video_name">{{ video.Name }}</div>
          </n-card>
        </div>
      </div>
    </n-space>
  </n-config-provider>
</template>

<script lang="ts">
  import { defineComponent, ref} from 'vue'
  import { NConfigProvider, NInput, NSpace, NButton, useMessage, NImage, NCard} from 'naive-ui'
  import SearchService from '@/services/SearchInput';
  import Video from '@/types/Video';
  import {AxiosResponse} from 'axios';

  export default defineComponent({
    components: {
      NConfigProvider,
      NInput,
      NSpace,
      NButton,
      NImage,
      NCard
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
            console.log(response.data.Porntn)
            response.data.Porntn.forEach((element: Video) => {
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