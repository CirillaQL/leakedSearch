<template>
  <div class="Result">
    <div class="search-input-result">
      <n-input-group>
        <n-input placeholder="输入要搜索的内容" v-model:value="searchContent"/>
        <n-button type="info" @click="search">
          <template #icon>
            <n-icon>
              <ios-search />
            </n-icon>
          </template>
          搜索
        </n-button>
      </n-input-group>
    </div>
    <n-divider />
    <n-space vertical>
      <div class="video-list">
        <n-grid :cols="4">
          <n-grid-item class="item" v-for="video in videos" :key="video.Name">
            <n-card>
              <template #cover>
                <n-image :src="video.CoverImg" @click="openImg(video.URL)"/>
              </template>
              <div class="video_name" :href="video.URL">{{ video.Name }}</div>
              <n-tag class="video_tag">{{ video.Source }}</n-tag>
              <n-button @click="openImg(video.URL)" type="info">点击跳转</n-button>
            </n-card>
          </n-grid-item>
        </n-grid>
      </div>
    </n-space>
  </div>
</template>

<script>
import { IosSearch } from '@vicons/ionicons4'
import router from "@/router";
import axios from "axios";
import {reactive, ref} from "vue";

export default {
  name: 'SearchView',
  components: {
    IosSearch
  },
  setup() {
    let videos = reactive([])
    let searchContent = router.currentRoute.value.query.search
    const fetchVideos = () => {
      let url = 'http://localhost:33333/videos/'+searchContent
      axios
          .get(url)
          .then(response => (
              response.data.Videos.forEach(value => {
                videos.push(value);
              })
          ))
    };
    fetchVideos()
    return {
      videos,
      searchContent: ref("")
    }
  },
  methods: {
    search() {
      console.log(router)
      router.push({name: 'result', query: {search: this.searchContent}})
    },
    openImg (URL) {
      window.open(URL)
    }
  }
}
</script>

<style>
.search-input-result {
  width: 40%;
  margin-top: 30px;
  margin-left: 30px;
}
.video-list {
  margin: 20px;
}
.video_tag {
  margin-right: 10px;
}
</style>
