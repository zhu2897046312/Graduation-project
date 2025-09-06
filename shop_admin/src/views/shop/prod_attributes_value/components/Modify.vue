<template>
  <BasicModal title="SKU属性" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef"  v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="title" label="名称">
        <Input.Group compact>
          <Input v-model:value="formState.title"  />
        </Input.Group>
      </Form.Item>
      <Form.Item name="sort_num" label="排序">
        <InputNumber :min="0" :max="1000000" v-model:value="formState.sort_num" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, InputNumber, message} from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import api from '/@/api/index'


const defaultFromValues = {
  title: '',
  sort_num: 10,
}



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
    formState.value = {...defaultFromValues}
    formState.value = await api.shop.prodAttributesValue.info( e.id)
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
    await api.shop.prodAttributesValue.modify(formState.value)
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