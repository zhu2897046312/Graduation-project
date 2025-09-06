<template>
  <BasicModal title="推荐位信息" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef"  v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="title" label="名称">
        <Input.Group compact>
          <Input v-model:value="formState.title"  />
        </Input.Group>
      </Form.Item>
      <Form.Item label="编码" name="code">
        <Input v-model:value="formState.code"  />
      </Form.Item>
      <Form.Item label="更多链接" name="more_link">
        <Input v-model:value="formState.more_link"  />
      </Form.Item>
      <Form.Item label="状态" name="state" >
        <Radio.Group v-model:value="formState.state" button-style="solid">
          <Radio.Button :value="1">发布</Radio.Button>
          <Radio.Button :value="2">不发布</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item label="图标" name="thumb">
        <UploadImage v-model:value="formState.thumb" />
      </Form.Item>
      <Form.Item label="描述" name="description">
        <Textarea :rows="3" v-model:value="formState.description"  />
      </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input,  message, Radio, Textarea } from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import api from '/@/api/index'


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

let id = 0
const handleOpen = (e: any) => {
  loadingTask(async () => {
    formState.value = await api.cms.recommend.curdApi('INFO', e.id)
    id = e.id;
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}


const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  console.debug(values)
  values.id = id
  loadingTask(async () => {
    await api.cms.recommend.curdApi('MODIFY', values);
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