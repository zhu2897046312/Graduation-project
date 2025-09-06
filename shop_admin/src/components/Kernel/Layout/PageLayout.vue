<template>
  <div class="page_layout">
    <PageHeader
      :title="title"
      :sub-title="subTitle"
      :ghost="false"
      style="background-color: #fff">
      <slot name="extra" />
      <slot name="footer" />
    </PageHeader>
    <div class="page_body">
      <slot />
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';
  import { PageHeader } from 'ant-design-vue';
  import { useRoute } from 'vue-router';
  const route = useRoute();

  const prop = defineProps({
    title: String,
    subTitle: String,
  });

  const title = computed<string>(() => {
    if (prop.title) {
      return prop.title;
    }
    if (route.meta.title) {
      return route.meta.title.toString();
    }
    return '';
  });

  const subTitle = computed<string>(() => {
    if (prop.subTitle) {
      return prop.subTitle;
    }
    if (route.meta.subTitle) {
      return route.meta.subTitle.toString();
    }
    return '';
  });

</script>

<style scoped>
  .page_body {
    padding: 16px;
    position: relative;
    flex-direction: column;
    flex: 1 1 auto;
    display: flex;
  }
  .page_layout {
    display: flex;
    flex-direction: column;
    flex: 1 1 auto;
  }
</style>