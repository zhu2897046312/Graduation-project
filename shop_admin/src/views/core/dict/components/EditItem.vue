<template>
  <BasicModal
    title="字典信息"
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
      <Form.Item name="label" label="属性名称">
        <Input v-model:value="formState.label" />
      </Form.Item>
      <Form.Item name="value" label="属性值">
        <Input v-model:value="formState.value" />
      </Form.Item>
      <Form.Item name="state" label="状态">
        <Radio.Group v-model:value="formState.state" >
          <Radio.Button :value="1">启用</Radio.Button>
          <Radio.Button :value="2">停用</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item name="sort_num" label="排序">
        <InputNumber :min="0" :max="100000" v-model:value="formState.sort_num" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  label: string;
  value: string;
  state: number;
  sort_num: number;
}
const defaultFromValues = {
  label: '',
  value: '',
  state: 1,
  sort_num: 10
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, message, Radio, InputNumber } from 'ant-design-vue';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import { coreDictItemCurdApi } from '/@/api/core/dictitem';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const emit = defineEmits(['on-change']);
  const rules: Record<string, Rule[]> = {
    label: [
      { required: true },
    ],
    value: [
      { required: true },
    ],
  };

  let id = 0;

  const handleOpen = (e: any) => {
    id = e.id;
    loadingTask(async () => {
      formState.value = await coreDictItemCurdApi('INFO', id);
      modalRef.value && modalRef.value.useOpen();
    })
  }

  const handleClose = () => {
    formState.value = {...defaultFromValues}
  }

  const handleSubmit = async () => {
    const values: any = await formRef.value?.validateFields();
    values.id = id;    
    loadingTask(async () => {
      await coreDictItemCurdApi('UPDATE', values);
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
