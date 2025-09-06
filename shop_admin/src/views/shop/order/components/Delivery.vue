<template>
  <BasicModal title="发货信息" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef" v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="delivery_company" label="快递公司">
        <Input.Group compact>
          <Input v-model:value="formState.delivery_company" />
        </Input.Group>
      </Form.Item>
      <Form.Item label="快递单号" name="delivery_sn">
        <Input v-model:value="formState.delivery_sn" />
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
  delivery_company: '',
  delivery_sn: ''
}
const rules: Record<string, Rule[]> = {
  delivery_company: [
    { required: true },
  ],
  delivery_sn: [
    { required: true },
  ],
};
// 栏目
const category_tree = ref<any[]>([])

const modalRef = ref<any>(null);
const formState = ref<any>(null);
const formRef = ref<any>(null);
const emit = defineEmits(['on-change']);
const formProp = ref<FormProps | any>({
  labelCol: { flex: '160px' },
  wrapperCol: { span: 20 },
  labelAlign: 'right',
})

const handleOpen = (e: number) => {
  loadingTask(async () => {
    formState.value = e
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}

const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  values.id = formState.value.id
  console.debug(values)
  loadingTask(async () => {
    await api.shop.order.delivery(values);
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