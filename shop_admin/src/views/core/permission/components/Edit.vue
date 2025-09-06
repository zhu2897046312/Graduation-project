<template>
  <BasicModal
    title="权限信息"
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
      <Form.Item name="pid" label="上级权限">
        <Select v-model:value="formState.pid">
          <Select.Option :value="0">顶级权限</Select.Option>
          <Select.Option v-for="(item, index) in pidList" :key="index" :value="item.id">
            {{ item.title }}
          </Select.Option>
        </Select>
      </Form.Item>
      <Form.Item name="title" label="权限名称">
        <Input v-model:value="formState.title" />
      </Form.Item>
      <Form.Item name="code" label="权限标识">
        <Input v-model:value="formState.code" />
      </Form.Item>
      <Form.Item name="remark" label="备注">
        <Input v-model:value="formState.remark" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  pid: number;
  title: string;
  code: string;
  urls: string[];
  remark: string;
}
const defaultFromValues = {
  pid: 0,
  title: '',
  code: '',
  remark: ''
} as any
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, Select, message } from 'ant-design-vue';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import {
    apiGetPermissionTopList,
    apiUpdatePermission,
    apiGetPermissionInfo,
  } from '/@/api/core/permission';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const emit = defineEmits(['on-change']);
  const rules: Record<string, Rule[]> = {
    title: [
      { required: true },
    ],
  };
  let id = 0;

  // 上级权限列表
  const pidList = ref<any[]>([]);

  const handleOpen = (e: any) => {
    id = e.id;
    loadingTask(async () => {
      pidList.value = await apiGetPermissionTopList();
      formState.value = await apiGetPermissionInfo(id);
      modalRef.value && modalRef.value.useOpen();
    })
  }

  const handleClose = () => {
    formState.value = {...defaultFromValues}
  }

  const handleSubmit = async () => {
    const values: any = await formRef.value?.validateFields();
    console.debug(values)
    values.id = id;
    loadingTask(async () => {
      await apiUpdatePermission(values);
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
