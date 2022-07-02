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
      <n-grid :cols="3">
        <n-grid-item class="item" v-for="video in videos" :key="video.Name" style="width: 50%;">
          <n-card>
            <template #cover>
              <n-image :src="video.CoverImg" @click="openImg(video.URL)"/>
            </template>
            <div class="video_name"  :href="video.URL">{{ video.Name }}</div>
            <n-tag class="video_tag">{{ video.Source }}</n-tag>
            <n-button @click="openImg(video.URL)" type="info">点击跳转</n-button>
          </n-card>
        </n-grid-item>
      </n-grid>
      </div>
    </n-space>
  </n-config-provider>
</template>

<script lang="ts">
  import { defineComponent, ref} from 'vue'
  import { NConfigProvider, NInput, NSpace, NButton, useMessage, NImage, NCard, NTag, NGrid, NGridItem} from 'naive-ui'
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
      NCard,
      NTag,
      NGrid,
      NGridItem
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
        this.videos = []
        SearchService.getVideos(this.input)
          .then((response: AxiosResponse) => {
            response.data.Videos.forEach((element: Video) => {
              this.videos.push(element)
            });
            console.log(this.videos)
          })
          .catch((e: Error) => {
            console.log(e);
          });
      },
      openImg (URL: string) {
        window.open(URL)
      }
    }
  })
</script>

<style>
.item {
  margin: auto;
  padding: 10px;
}

.video_tag {
  margin: 10px;
}
</style>