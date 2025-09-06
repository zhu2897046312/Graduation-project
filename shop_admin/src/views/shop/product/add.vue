<script setup lang="ts">
import { Card, Form, Input, TreeSelect, Radio, DatePicker, InputNumber, message, Button, Divider, Select, Textarea  } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form'
import UploadImage from '/@/components/Kernel/YexUpload/UploadCropperImage.vue';
import UploadCropperGalleryImage from '/@/components/Kernel/YexUpload/UploadCropperGalleryImage.vue';
import { ref, onMounted, nextTick } from 'vue';
import dayjs from 'dayjs'
import { useRouter } from 'vue-router';
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import api from '/@/api';
import Editor from '/@/components/Kernel/YexEditor/Editor.vue';
import PropertyForm from './components/PropertyForm.vue';
import SkuForm from './components/SkuForm.vue';

const router = useRouter()

const layout = ref<any>({
  mainLabel: { flex: '160px' },
})

const rules: Record<string, Rule[]> = {
  title: [{ required: true, message: '请输入标题' }],
}
const formRef = ref<any>();

const data = ref<any>({
  category_id: '',
  title: '',
  state: 1,
  price: '0.00',
  original_price: '0.00',
  cost_price: '0.00',
  stock: 99,
  picture: '',
  picture_gallery: [],
  description: '',
  sold_num: parseInt((Math.random() * 1000).toFixed(0)) + 10,
  sort_num: 100,
  putaway_time: dayjs(),
  content: '',
  seo_title: '',
  seo_keyword: '',
  seo_description: '',
  property_list: [],
  open_sku: 2,
  sku_list: [],
  tags: [],
  hot: 2,
})

const load = ref(false)

// 栏目
const category_tree = ref<any[]>([])

const sku_config = ref<Map<number, number[]>>(new Map());

// 标签
const tags = ref<any[]>([]);


const handleSubmit = async () => {
  try {
    await formRef.value.validate();
    const post_data = { ...data.value }
    if (post_data.category_id == '' || post_data.category_id <= 0) {
      message.warn('请选择栏目');
      return;
    }
    if (post_data.open_sku == 1) {
      if (post_data.sku_list.length == 0) {
        message.warn('请添加规格配置');
        return;
      }
      const default_sku = post_data.sku_list.find((it: any) => {
        return it.default_show == 1
      })
      if (!default_sku) {
        message.warn('请设置一个默认的规格');
        return;
      }
    }

    post_data.putaway_time = post_data.putaway_time.format('YYYY-MM-DD HH:mm:ss');
    post_data.content = post_data.content ? post_data.content : ''
    post_data.content = post_data.content.replace(/script/g, '')
      .replace(/alert/g, '')
      .replace(/onerror/g, '')
      .replace(/document/g, '')
      .replace(/window/g, '')
      .replace(/onsuccess/g, '');

    if (post_data.description.length == 0) {
      post_data.description = post_data.content.replace(/<[^>]+>/g, '').replace('\n', '').substring(0, 80);
    }
    await api.shop.product.create(post_data)
    message.success('录入商品成功');
    router.back()
  } catch (e: any) {
    console.warn(e)
    if (e.errorFields) {
      message.warn(e.errorFields[0].errors[0]);
    } else {
      message.warn(e.toString());
    }
  }
}

onMounted(async () => {

  let _treeData = (await api.shop.category.tree()) as any
  category_tree.value = _treeData
  if (_treeData.length > 0) {
    data.value.category_id = category_tree.value[0].value
  }

  const res = await api.cms.tag.curdApi('LIST', {page_no: 1, page_size: 600, status: 1})
  tags.value = res.list.map(it => {
    return {
      label: it.title,
      value: it.id
    }
  })
  await nextTick(() => {
    load.value = true
  })
})

</script>

<template>
  <PageLayout>
    <Card title="添加商品">
      <template #extra>
        <Button type="primary" @click="handleSubmit">提交保存</Button>
      </template>
      <div v-if="load" class="full-width">
        <Form ref="formRef" :label-col="{ span: 5 }" :wrapper-col="{ span: 22 }" :model="data" :rules="rules">

          <div class="full-width flex gap-4">
            <div class="w-9/12">
              <Form.Item label="商品名称" :label-col="layout.mainLabel" name="title">
                <Input v-model:value="data.title" :maxlength="200" />
              </Form.Item>
              <Form.Item label="开启SKU" name="open_sku" :label-col="layout.mainLabel" >
                <Radio.Group v-model:value="data.open_sku" button-style="solid">
                  <Radio.Button :value="1">开启</Radio.Button>
                  <Radio.Button :value="2">关闭</Radio.Button>
                </Radio.Group>
              </Form.Item>
              <div v-if="data.open_sku == 2" class="full-width flex gap-6">
                <Form.Item label="商品原价" :label-col="layout.mainLabel" name="original_price">
                  <InputNumber prefix="￥" style="width: 200px;" :min="0" :max="999999" :precision="2" v-model:value="data.original_price" />
                </Form.Item>
                <Form.Item label="当前价格" :label-col="layout.mainLabel" name="price">
                  <InputNumber prefix="￥" style="width: 200px;" :min="0" :max="999999" :precision="2" v-model:value="data.price" />
                </Form.Item>
                <Form.Item label="成本价格" :label-col="layout.mainLabel" name="cost_price">
                  <InputNumber prefix="￥" style="width: 200px;" :min="0" :max="999999" :precision="2" v-model:value="data.cost_price" />
                </Form.Item>
              </div>
              
              <Form.Item v-if="data.open_sku == 1" label="规格配置" name="sku_list" :label-col="layout.mainLabel" >
                <SkuForm v-model:value="data.sku_list"  :sku-config="sku_config" />
              </Form.Item>
              <div class="full-width flex gap-6">
                <Form.Item v-if="data.open_sku == 2" label="库存" :label-col="layout.mainLabel" name="stock">
                  <InputNumber style="width: 200px;" :min="0" :max="999999" :precision="0" v-model:value="data.stock" />
                </Form.Item>
                <Form.Item label="销量" :label-col="layout.mainLabel" name="sold_num">
                  <InputNumber style="width: 200px;" :min="0" :max="999999" :precision="0" v-model:value="data.sold_num" />
                </Form.Item>
              </div>

              
              
              <Form.Item label="商品内容" name="content" :label-col="layout.mainLabel" >
                <Editor v-model:value="data.content" />
              </Form.Item>
              <Form.Item label="商品属性" name="property_list" :label-col="layout.mainLabel" >
                <PropertyForm v-model:value="data.property_list" />
              </Form.Item>
            </div>
            <div class="w-3/12">
              <Form.Item label="类目" name="category_id">
                <TreeSelect v-model:value="data.category_id" show-search style="width: 100%"
                  :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }" tree-default-expand-all :tree-data="category_tree" />
              </Form.Item>
              
              <Form.Item label="封面" name="picture" help="封面图尺寸为：502px * 502px ,展示效果最佳">
                <UploadImage v-model:value="data.picture" :aspect-ratio="1" />
              </Form.Item>
              <Form.Item label="组图" name="picture_gallery" help="封面图尺寸为：502px * 502px ,展示效果最佳">
                <UploadCropperGalleryImage :max-num="6" txt="上传组图" v-model:value="data.picture_gallery" :aspect-ratio="1" />
              </Form.Item>
              <Form.Item label="标签" name="tags">
                <Select
                  v-model:value="data.tags"
                  mode="multiple"
                  style="width: 100%"
                  show-search
                 :filter-option="(input: string, option: any) => {
                    return option.label.indexOf(input) !== -1;
                 }"
                  :options="tags" />
              </Form.Item>
              <Divider>SEO设置</Divider>
              <Form.Item label="SEO标题" name="seo_title">
                <Input v-model:value="data.seo_title" :maxlength="200" />
              </Form.Item>
              <Form.Item label="SEO关键词" name="seo_keyword">
                <Input v-model:value="data.seo_keyword" :maxlength="200" />
              </Form.Item>
              <Form.Item label="SEO描述" name="seo_description">
                <Textarea :rows="4" v-model:value="data.seo_description" :maxlength="200" />
              </Form.Item>
              
              <Divider>其他设置</Divider>
              <Form.Item label="状态" name="state">
                <Radio.Group v-model:value="data.state" button-style="solid">
                  <Radio.Button :value="1">上线</Radio.Button>
                  <Radio.Button :value="2">下线</Radio.Button>
                  <Radio.Button :value="3">待审核</Radio.Button>
                  <Radio.Button :value="4">审核不通过</Radio.Button>
                </Radio.Group>
              </Form.Item>
              <Form.Item label="热门商品" name="hot">
                <Radio.Group v-model:value="data.hot" button-style="solid">
                  <Radio.Button :value="1">热门</Radio.Button>
                  <Radio.Button :value="2">非热门</Radio.Button>
                </Radio.Group>
              </Form.Item>
              <Form.Item label="发布时间" name="putaway_time">
                <DatePicker v-model:value="data.putaway_time" show-time />
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


<style scoped>
.archive_submit_box {
  display: flex;
  width: 100%;
  justify-content: center;
}
</style>