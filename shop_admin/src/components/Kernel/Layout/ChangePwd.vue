<template>
  <BasicModal
    title="修改密码"
    @on-ok="handleSubmit"
    @on-close="handleClose"
    ref="modalRef"
  >
    <Form ref="formRef" v-bind="formProp" :model="formState" :rules="rules">
      <Form.Item name="old_pwd" label="原密码">
        <Input.Password v-model:value="formState.old_pwd" />
      </Form.Item>
      <Form.Item name="new_pwd" label="新密码">
        <Input.Password v-model:value="formState.new_pwd" />
      </Form.Item>
      <Form.Item name="confirm_pwd" label="确认密码">
        <Input.Password v-model:value="formState.confirm_pwd" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  old_pwd: string;
  new_pwd: string;
  confirm_pwd: string;
}
const defaultFromValues: FormState = {
  old_pwd: '',
  new_pwd: '',
  confirm_pwd: '',
}
</script>

<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, message } from 'ant-design-vue';
import type { FormInstance, FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import { apiChangePwd } from '/@/api/auth';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

const modalRef = ref<any>(null);
const formState = ref<FormState>({...defaultFromValues});
const formRef = ref<FormInstance|null>(null);
const emit = defineEmits(['on-change']);

let confirmRule = async (_: any, value: string) => {
  if (value != formState.value.new_pwd) {
    return Promise.reject('两次密码不一致');
  }
  return Promise.resolve();
}

const rules: Record<string, Rule[]> = {
  old_pwd: [
    { required: true },
  ],
  new_pwd: [
    { required: true },
  ],
  confirm_pwd: [
    { required: true, validator: confirmRule },
  ]
};
const formProp = ref<FormProps|any>({
  labelCol: { span: 4 },
  wrapperCol: { span: 20 },
  labelAlign: 'right',
})

const handleOpen = () => {
  loadingTask(async () => {
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = {...defaultFromValues}
}

const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  loadingTask(async () => {
    await apiChangePwd(values);
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
