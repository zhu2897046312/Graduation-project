<template>
  <PageLayout>
    <Card title="站点设置">
      <Form>
        <Form.Item label="站点名称">
          <Input style="width: 230px" v-model:value="info.title"  />
        </Form.Item>
        <Form.Item label="站点LOGO">
          <UploadImage v-model:value="info.logo"  />
        </Form.Item>
        <Form.Item label="seo标题">
          <Input style="width: 230px" v-model:value="info.seo_title"  />
        </Form.Item>
        <Form.Item label="seo关键词">
          <Input style="width: 230px" v-model:value="info.seo_keyword"  />
        </Form.Item>
        <Form.Item label="seo描述">
          <Textarea :rows="4" v-model:value="info.seo_description"  />
        </Form.Item>
        <Form.Item >
          <Button type="primary" @click="handleSubmit">保存</Button>
        </Form.Item>
      </Form>
    </Card>
  </PageLayout>
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card, Form, Input, Button, message, Textarea } from 'ant-design-vue';
  import { ref, onMounted } from 'vue';
  import api  from '/@/api/index';
  import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';

  const info = ref<any>({
    title: '',
    logo: '',
    seo_title: '',
    seo_keyword: '',
    seo_description: '',
  })



  const handleSubmit = async () => {
    const hide = message.loading('正在保存中', 0)
    await api.shop.market.saveSiteInfo(info.value)
    setTimeout(() => {
      hide()
      message.success('保存成功')
    }, 800)
  }


  onMounted(async () => {
    info.value = await api.shop.market.siteInfo()
  })


</script>