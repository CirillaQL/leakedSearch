<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-space vertical>
      <n-input-group>
        <n-input placeholder="搜索内容" style="width: 50%;">
          <template #suffix>
          </template>
        </n-input>
        <n-button :loading="loading" @click="handleClick">
          搜索
        </n-button>
      </n-input-group>
    </n-space>
  </n-config-provider>
</template>

<script lang="ts">
  import { defineComponent, ref } from 'vue'
  import { NConfigProvider, NInput, NSpace, NButton } from 'naive-ui'

  export default defineComponent({
    components: {
      NConfigProvider,
      NInput,
      NSpace,
      NButton
    },
    setup() {
      const loadingRef = ref(false)
      return {
        handleClick(){
          loadingRef.value = true
          setTimeout(() => {
            loadingRef.value = false
          }, 2000)
        },
        loading: loadingRef,
      }
    }
  })
</script>