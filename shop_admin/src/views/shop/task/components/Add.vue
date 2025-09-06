<template>
  <BasicModal title="新增采集任务" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef" v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item label="类目" name="category_id">
        <TreeSelect v-model:value="formState.category_id" show-search style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }" tree-default-expand-all :tree-data="category_tree" />
      </Form.Item>
      <Form.Item name="title" label="名称">
        <Input.Group compact>
          <Input v-model:value="formState.title" />
        </Input.Group>
      </Form.Item>
      <Form.Item label="采集间隔(小时)" name="interval">
        <InputNumber v-model:value="formState.interval" />
      </Form.Item>
      <Form.Item label="状态" name="state">
        <Radio.Group v-model:value="formState.state" button-style="solid">
          <Radio.Button :value="1">开启</Radio.Button>
          <Radio.Button :value="2">关闭</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item label="目标平台" name="platform">
        <Select v-model:value="formState.platform" :options="[{ value: 1, label: '淘宝' }]" />
      </Form.Item>
      <Form.Item label="采集类型" name="type">
        <Select v-model:value="formState.type" :options="[{ value: 1, label: '关键词' }, { value: 2, label: '详情' }]" />
      </Form.Item>
      <Divider>任务配置</Divider>
      <Form.Item label="关键词" name="q" v-if="formState.type === 1">
        <Input v-model:value="formState.q" />
      </Form.Item>
      <Form.Item label="采集条数" name="num" v-if="formState.type === 1">
        <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="formState.num" />
      </Form.Item>
      <Form.Item label="商品ID（每行一个）" name="ids" v-if="formState.type === 2">
        <Input.TextArea :rows="3" v-model:value="formState.ids" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, message, Radio, Divider, InputNumber, Select, TreeSelect } from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import api from '/@/api/index'


const defaultFromValues = {
  title: '',
  interval: 1,
  state: 1,
  platform: 1,
  type: 1,
  num: 20,
  ids: '',
}
// 栏目
const category_tree = ref<any[]>([])

const modalRef = ref<any>(null);
const formState = ref<any>(null);
const formRef = ref<any>(null);
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
    let _treeData = (await api.shop.category.tree()) as any
    category_tree.value = _treeData

    formState.value = { ...defaultFromValues }
    formState.value.category_id = _treeData[0].value
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}

const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  if (values.type === 1) {
    values.config = JSON.stringify({
      q: values.q,
      num: values.num,
      categoryId: values.category_id,
    })
  } else {
    values.config = JSON.stringify({
      ids: [...values.ids.split('\n')],
      categoryId: values.category_id,
    })
  }

  console.debug(values)
  loadingTask(async () => {
    await api.shop.task.create(values);
    message.success('操作成功');
    modalRef.value && modalRef.value.useClose();
    emit('on-change');
  })
}


defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})

</script>