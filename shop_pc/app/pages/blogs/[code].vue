<script setup lang="ts">
import api from '../../../api';


const route = useRoute()
const code = route.params.code as any

const {data: info, status} = await useAsyncData(`document:${code}`, async() => {
  const res = await api.blogs.document.info(code)
  return res
});
if (!info.value) {
  throw createError({ statusCode: 404, message: 'Page not found' })
}

</script>

<template>
  <div class="container" v-if="status == 'success'">

    <Title>{{ info.cont.seo_title }}</Title>
    <Meta name="keywords" :content="info.cont.seo_keyword" />
    <Meta name="description" :content="info.cont.seo_description" />

    <h1>{{ info.document.title }}</h1>
    <div class="text-sm	text-slate-500 mb-10">{{ info.document.send_time }}</div>
    <div class="blogs-box blogs_box" v-html="info.cont.cont">
    </div>
  </div>
</template>

<style lang="css" scoped>
h1 {
  padding: 20px 0;
  font-size: 23px;
  font-weight: bold;
  color: var(--font-text-color);
  letter-spacing: 2px;
}

.blogs-box {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px dashed #f5f5f5;
  padding-bottom: 20px;
}
.blogs-box::after {
  content: " ";
  display: block;
  clear: both;
  visibility: hidden;
  width: 0;
  height: 0;
}
</style>