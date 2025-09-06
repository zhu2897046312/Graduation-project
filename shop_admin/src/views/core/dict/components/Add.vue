<template>
  <BasicModal
    title="新增字典"
    @on-ok="handleSubmit"
    @on-close="handleClose"
    ref="modalRef"
  >
    <Form
      ref="formRef"
      :labelCol="{ span: 4 }"
      :wrapperCol="{ span: 20 }"
      labelAlign="right"
      :model="formState"
      :rules="rules">
      <Form.Item name="title" label="字典名称">
        <Input v-model:value="formState.title" />
      </Form.Item>
      <Form.Item name="code" label="字典标识">
        <Input v-model:value="formState.code" />
      </Form.Item>
      <Form.Item name="state" label="状态">
        <Radio.Group v-model:value="formState.state" >
          <Radio.Button :value="1">启用</Radio.Button>
          <Radio.Button :value="2">停用</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item name="remark" label="备注">
        <Input.TextArea v-model:value="formState.remark" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  title: string;
  state: number;
  code: string;
  remark: string;
}
const defaultFromValues = {
  title: '',
  state: 1,
  code: '',
  remark: ''
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, message, Radio } from 'ant-design-vue';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import { coreDictCurdApi } from '/@/api/core/dict';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const emit = defineEmits(['on-change']);
  const rules: Record<string, Rule[]> = {
    title: [
      { required: true },
    ],
    code: [
      { required: true },
    ],
  };

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
      await coreDictCurdApi('CREATE', values);
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
