<template>
  <BasicModal title="更新状态" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef" v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item label="订单状态" name="state">
        <Select v-model:value="formState.state" :options="stateList" />
      </Form.Item>
      <Form.Item label="备注" name="remark">
        <Input.TextArea :rows="3" v-model:value="formState.remark" />
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
import enumDict from '/@/utils/enum-dict';


const defaultFromValues = {
  delivery_company: '',
  delivery_sn: ''
}
const rules: Record<string, Rule[]> = {
  state: [
    { required: true },
  ],
};
// 栏目
const stateList = ref<any[]>([])

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
    await api.shop.order.updateState(values);
    message.success('操作成功');
    modalRef.value && modalRef.value.useClose();
    emit('on-change');
  })
}

console.log(enumDict.getDictList('SpOrderConstant$State'))
stateList.value = enumDict.getDictList('SpOrderConstant$State')


defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})

</script>