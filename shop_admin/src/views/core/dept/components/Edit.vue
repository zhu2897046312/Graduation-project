<template>
  <BasicModal
    title="更新部门信息"
    @on-ok="handleSubmit"
    @on-close="handleClose"
    ref="modalRef"
  >
    <Form ref="formRef" v-bind="formProp" :model="formState" :rules="rules">
      <Form.Item name="pid" label="顶级部门">
        <TreeSelect
          v-model:value="formState.pid"
          show-search
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          placeholder="Please select"
          allow-clear
          tree-default-expand-all
          :tree-data="tree"
        />
      </Form.Item>
      <Form.Item name="dept_name" label="部门名称">
        <Input v-model:value="formState.dept_name" />
      </Form.Item>
      <Form.Item name="connect_name" label="负责人">
        <Input v-model:value="formState.connect_name" />
      </Form.Item>
      <Form.Item name="connect_mobile" label="电话">
        <Input v-model:value="formState.connect_mobile" />
      </Form.Item>
      <Form.Item name="sort_num" label="排序">
        <InputNumber :min="0" :max="1000000" v-model:value="formState.sort_num" />
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
  dept_name: string;
  connect_name: string;
  connect_mobile: string;
  sort_num: number;
  remark: string;
}
const defaultFromValues = {
  pid: 0,
  dept_name: '',
  connect_name: '',
  connect_mobile: '',
  sort_num: 10,
  remark: ''
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, TreeSelect, InputNumber, message } from 'ant-design-vue';
  import type { FormInstance, FormProps } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import { apiFormatApiTree, apiUpdateDept, apiGetDeptInfo } from '/@/api/core/dept';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const tree = ref<any[]>([]);
  const emit = defineEmits(['on-change']);
  let id = 0;
  const rules: Record<string, Rule[]> = {
    deptName: [
      { required: true },
    ],
  };
  const formProp = ref<FormProps|any>({
    labelCol: { span: 4 },
    wrapperCol: { span: 20 },
    labelAlign: 'right',
  })

  const handleOpen = (e: any) => {
    id = e.id;
    loadingTask(async () => {
      tree.value = await apiFormatApiTree('顶级部门');
      formState.value = await apiGetDeptInfo(id);
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
      await apiUpdateDept(values);
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
