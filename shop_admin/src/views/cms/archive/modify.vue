<template>
  <PageLayout>
    <Card title="编辑文档">
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
                <Input readonly :prefix="`${web_url}/blogs/`" v-model:value="data.code" :maxlength="64"  />
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
import { Card, Form, Input, Radio, DatePicker, InputNumber, message, Button, Modal, Divider  } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form'
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import { ref, onMounted, nextTick } from 'vue';
import dayjs from 'dayjs'
import { useRouter, useRoute } from 'vue-router';
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import { document } from "/@/api/cms";
import {loadingTask} from "/@/utils/helper.ts";
import Editor from '/@/components/Kernel/YexEditor/Editor.vue';

const router = useRouter()
const route = useRoute()


const web_url = ref(import.meta.env.VITE_WEB_URL)

const rules: Record<string, Rule[]> = {
  title: [{ required: true, message: '请输入标题' }],
}
const formRef = ref<any>();

const data = ref<any>({
  title: '',
  code: '',
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

    post_data.id = id
    await document.curdApi("UPDATE", post_data)
    message.success('文档编辑成功');
    handleAsk();
  } catch (e: any) {
    console.warn(e)
    message.warn(e.toString());
  }
}

const handleAsk = () => {
  Modal.confirm({
    title: '文档已保存，是否返回列表页？',
    okText: '返回文档列表',
    cancelText: '继续编辑',
    onOk: () => {
      router.push({ name: 'ShopBlogs' })
    },
    onCancel: () => { }
  })
}

let id = 0

onMounted(async () => {
  loadingTask(async () => {
    const _data = await document.curdApi("INFO", route.query.id)
    const _info = {..._data.document}
    _info.send_time = dayjs(_info.send_time)
    _info.cont = _data.cont.cont
    _info.seo_title = _data.cont.seo_title
    _info.seo_keyword = _data.cont.seo_keyword
    _info.seo_description = _data.cont.seo_description
    data.value = _info
    id = _data.document.id
    console.log(data.value)
  })
  nextTick(() => {
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