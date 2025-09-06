<template>
  <BasicModal title="新增类目" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef"  v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="state" label="状态">
        <Radio.Group v-model:value="formState.state" button-style="solid">
          <Radio.Button :value="1">启用</Radio.Button>
          <Radio.Button :value="2">停用</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item name="pid" label="上级类目">
        <TreeSelect v-model:value="formState.pid" show-search style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }" placeholder="Please select" allow-clear
          tree-default-expand-all :tree-data="tree" @change="(a,b,c) => {console.log(a,b,c)}"/>
      </Form.Item>
      <Form.Item name="title" label="类目名称">
        <Input.Group compact>
          <Input v-model:value="formState.title" @blur="handleChangeTitle" style="width: 85%;" />
          <Button @click="handleChangeTitle" style="width: 15%;">自动中译英</Button>
        </Input.Group>
      </Form.Item>
      
      <Form.Item name="code" label="访问地址">
        <Input v-model:value="formState.code">
          <template #addonBefore>
            <div>{{ web_url }}</div>
          </template>
          <template #addonAfter>
            <div>/</div>
          </template>
        </Input>
        <div class="pre_link" :style="{opacity: formState.code.length > 0 ? 100 : 0}">访问地址: 
          <a :href="`${web_url}/${formState.code}/`" target="_blank">
            {{ web_url }}/{{ formState.code }}/</a>
          </div>
      </Form.Item>
      <Form.Item name="picture" label="封面图">
        <UploadImage v-model:value="formState.picture" />
      </Form.Item>
      <Form.Item name="description" label="描述">
        <Input.TextArea :rows="4" v-model:value="formState.description" />
      </Form.Item>
      <Form.Item name="sort_num" label="排序">
        <InputNumber :min="0" :max="1000000" v-model:value="formState.sort_num" />
      </Form.Item>
      <Divider orientation="left">SEO信息</Divider>
      <Form.Item name="seo_title" label="SEO标题">
        <Input v-model:value="formState.seo_title" />
      </Form.Item>
      <Form.Item name="seo_keyword" label="SEO关键词">
        <Input v-model:value="formState.seo_keyword" />
      </Form.Item>
      <Form.Item name="seo_description" label="SEO描述">
        <Input.TextArea :rows="3" v-model:value="formState.seo_description" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, TreeSelect, InputNumber, message, Radio, Button, Divider } from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import { pinyin } from 'pinyin-pro';
import api from '/@/api/index'


const defaultFromValues = {
  pid: 0,
  title: '',
  code: '',
  state: 1,
  icon: '',
  picture: '',
  description: '',
  seo_title: '',
  seo_keyword: '',
  seo_description: '',
  sort_num: 10,
}

const web_url = ref(import.meta.env.VITE_WEB_URL)

const modalRef = ref<any>(null);
const formState = ref<any>(null);
const formRef = ref<any>(null);
const tree = ref<any[]>([]);
const emit = defineEmits(['on-change']);
const rules: Record<string, Rule[]> = {
  title: [
    { required: true },
  ],
  code: [
    { required: true },
  ],
};
const formProp = ref<FormProps | any>({
  labelCol: { flex: '160px' },
  wrapperCol: { span: 20 },
  labelAlign: 'right',
})

const handleOpen = (e: any) => {
  loadingTask(async () => {
    let treeData = [{label: '无', value: 0}]
    let _treeData = (await api.shop.category.tree()) as any
    treeData.push(..._treeData)
    tree.value = treeData

    formState.value = {...defaultFromValues}
    if (e.pid) {
      formState.value.pid = e.pid
    }
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}

const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  console.debug(values)
  loadingTask(async () => {
    await api.shop.category.create(formState.value)
    message.success('操作成功');
    modalRef.value && modalRef.value.useClose();
    emit('on-change');
  })
}

const handleChangeTitle = () => {
  formState.value.title = pinyin(formState.value.title, { toneType: 'none', separator: '' })
  formState.value.code = pinyin(formState.value.title, { toneType: 'none', separator: '' }).replace(/\s/g, '-').toLowerCase()
}

defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})

</script>