<template>
  <div class="home">
    <img class="logo" alt="leakedSearch logo" src="../assets/leakedSearchLogo.png">
    <div class="search-input">
      <n-input-group>
        <n-input :style="{ width: '100%' }" placeholder="输入要搜索的内容" v-model:value="searchContent"/>
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
  </div>
</template>

<script>
import {IosSearch} from '@vicons/ionicons4'
import {ref} from "vue";
import router from "@/router";

// import router from "@/router";

export default {
  name: 'HomeView',
  components: {
    IosSearch
  },
  setup() {
    return {
      searchContent: ref("")
    }
  },
  methods: {
    search() {
      router.push({name: 'result', query: {search: this.searchContent}})
    }
  }
}
</script>

<style>
.search-input {
  width: 50%;
  margin: auto;
}
.logo {
  margin: 100px;
}
</style>
