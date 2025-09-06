

<template>
  <ConfigProvider :locale="locale">
    <RouterView />
  </ConfigProvider>
</template>

<script setup lang="ts">
  import { ref } from 'vue';
  import { ConfigProvider } from 'ant-design-vue';
  import zhCN from 'ant-design-vue/es/locale/zh_CN';
  import dayjs from 'dayjs';
  import 'dayjs/locale/zh-cn';
  dayjs.locale('zh-cn');

  const locale = ref<any>(zhCN);

  ConfigProvider.config({
    theme: {
      primaryColor: '#d74545'
    }
  })
</script>

