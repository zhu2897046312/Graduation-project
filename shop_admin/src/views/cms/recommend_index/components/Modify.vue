<template>
  <BasicModal title="推荐位内容信息" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef"  v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="title" label="名称">
        <Input.Group compact>
          <Input v-model:value="formState.title"  />
        </Input.Group>
      </Form.Item>
      <Form.Item label="链接" name="link">
        <Input v-model:value="formState.link"  />
      </Form.Item>
      <Form.Item label="封面" name="thumb">
        <UploadImage v-model:value="formState.thumb" />
      </Form.Item>
      <Form.Item label="状态" name="state" >
        <Radio.Group v-model:value="formState.state" button-style="solid">
          <Radio.Button :value="1">发布</Radio.Button>
          <Radio.Button :value="2">不发布</Radio.Button>
        </Radio.Group>
      </Form.Item>
        <Form.Item label="排序" name="sort_num">
          <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="formState.sort_num" />
        </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input,  message, Radio } from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import api from '/@/api/index'
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';


const modalRef = ref<any>(null);
const formState = ref<any>(null);
const formRef = ref<any>(null);
const emit = defineEmits(['on-change']);
const rules: Record<string, Rule[]> = {
  title: [
    { required: true },
  ],
};
const formProp = ref<FormProps | any>({
  labelCol: { flex: '160px' },
  wrapperCol: { span: 20 },
  labelAlign: 'right',
})

let source : any = null
const handleOpen = (e: any) => {
  loadingTask(async () => {
    formState.value = await api.cms.recommendIndex.curdApi('INFO', e.id)
    source = e;
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}


const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  console.debug(values)
  values.id = source.id
  values.recommend_id = formState.value.recommend_id
  values.product_id = formState.value.product_id
  values.document_id = formState.value.document_id
  loadingTask(async () => {
    await api.cms.recommendIndex.curdApi('MODIFY', values);
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