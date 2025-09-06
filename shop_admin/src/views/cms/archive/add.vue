<template>
  <PageLayout>
    <Card title="添加文档">
      <template #extra>
        <Button type="primary" @click="handleSubmit">提交保存</Button>
      </template>
      <div v-if="load" class="full-width">
        <Form ref="formRef" :label-col="{ span: 5 }" :wrapper-col="{ span: 22 }" :model="data" :rules="rules">

          <div class="full-width flex gap-4">
            <div class="w-9/12">
              <Form.Item label="标题" :label-col="{ span: 2 }" name="title">
                <Input v-model:value="data.title" :maxlength="200" />
              </Form.Item>
              <Form.Item label="访问地址" :label-col="{ span: 2 }" name="title">
                <Input :prefix="`${web_url}/blogs/`" v-model:value="data.code" :maxlength="64"  />
              </Form.Item>
              <Form.Item label="图文信息" :label-col="{ span: 2 }" >
                <Editor v-model:value="data.cont" />
              </Form.Item>
            </div>
            <div class="w-3/12">
              <Form.Item label="封面" name="thumb">
                <UploadImage v-model:value="data.thumb" />
              </Form.Item>
              <Divider>SEO设置</Divider>
              <Form.Item label="关键词" name="seo_title">
                <Input v-model:value="data.seo_title" :maxlength="200" />
              </Form.Item>
              <Form.Item label="关键词" name="seo_keyword">
                <Input v-model:value="data.seo_keyword" :maxlength="200" />
              </Form.Item>
              <Form.Item label="描述" name="seo_description">
                <Input.TextArea :rows="3" v-model:value="data.seo_description" :maxlength="200" />
              </Form.Item>
              
              <Divider>其他设置</Divider>
              <Form.Item label="状态" name="state">
                <Radio.Group v-model:value="data.state" button-style="solid">
                  <Radio.Button :value="1">正常发布</Radio.Button>
                  <Radio.Button :value="2">草稿</Radio.Button>
                </Radio.Group>
              </Form.Item>
              <Form.Item label="发布时间" name="send_time">
                <DatePicker v-model:value="data.send_time" show-time />
              </Form.Item>
              <Form.Item label="阅读次数" name="read_num">
                <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="data.read_num" />
              </Form.Item>
              <Form.Item label="点赞数" name="read_num">
                <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="data.like_num" />
              </Form.Item>
              <Form.Item label="排序" name="sort_num">
                <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="data.sort_num" />
              </Form.Item>
            </div>
          </div>
        </Form>
      </div>
    </Card>
  </PageLayout>
</template>

<script setup lang="ts">
import { Card, Form, Input, Radio, DatePicker, InputNumber, message, Button, Divider, Select  } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form'
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import { ref, onMounted, nextTick } from 'vue';
import dayjs from 'dayjs'
import { useRouter } from 'vue-router';
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import {document } from "/@/api/cms";
import Editor from '/@/components/Kernel/YexEditor/Editor.vue';

const router = useRouter()

const generateRandomString = (length: number): string => {
  const characters = 'abcdefghijklmnopqrstuvwxyz0123456789';
  let result = '';
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters[randomIndex];
  }
  return result;
}

const rules: Record<string, Rule[]> = {
  title: [{ required: true, message: '请输入标题' }],
}
const formRef = ref<any>();

const web_url = ref(import.meta.env.VITE_WEB_URL)

const data = ref<any>({
  title: '',
  code: generateRandomString(8),
  thumb: '',
  state: 1,
  link_type: 1,
  send_time: dayjs(),
  author: '',
  source: '',
  read_num: parseInt((Math.random() * 1000).toFixed(0)) + 10,
  link_num: parseInt((Math.random() * 1000).toFixed(0)) + 10,
  sort_num: 50,
  cont: '',
  download_files: [],
  seo_title: '',
  seo_keyword: '',
  seo_description: '',
})

const load = ref(false)


const handleSubmit = async () => {
  try {
    await formRef.value.validate();
    const post_data = { ...data.value }
    post_data.send_time = post_data.send_time.format('YYYY-MM-DD HH:mm:ss');
    post_data.cont = post_data.cont ? post_data.cont : ''
    post_data.cont = post_data.cont.replace(/script/g, '')
      .replace(/alert/g, '')
      .replace(/onerror/g, '')
      .replace(/document/g, '')
      .replace(/window/g, '')
      .replace(/onsuccess/g, '');

    await document.curdApi("CREATE", post_data)
    message.success('文档创建成功');
    router.back()
  } catch (e: any) {
    console.warn(e)
    message.warn(e.toString());
  }
}

onMounted(async () => {
  await nextTick(() => {
    load.value = true
  })
})

</script>

<style scoped>
.archive_submit_box {
  display: flex;
  width: 100%;
  justify-content: center;
}
</style>