<template>
  <BasicModal
    title="新增角色"
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
      <Form.Item name="role_name" label="角色名称">
        <Input v-model:value="formState.role_name" />
      </Form.Item>
      <Form.Item name="role_status" label="角色状态">
        <Radio.Group v-model:value="formState.role_status" >
          <Radio.Button :value="1">启用</Radio.Button>
          <Radio.Button :value="2">停用</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item name="permission" label="权限标识">
        <TreeSelect
          v-model:value="formState.permission"
          show-search
          multiple
          tree-checkable
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          placeholder="请选择..."
          allow-clear
          :field-names = "{
            children: 'children',
            label: 'title',
            value: 'id',
          }"
          tree-default-expand-all
          :tree-data="permisstionTree"
        />
      </Form.Item>
      <Form.Item name="remark" label="备注">
        <Input.TextArea v-model:value="formState.remark" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  role_name: string;
  role_status: number;
  permission: string[];
  remark: string;
}
const defaultFromValues = {
  role_name: '',
  role_status: 1,
  permission: [],
  remark: ''
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, TreeSelect, message, Radio } from 'ant-design-vue';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import { coreRoleCurdApi } from '/@/api/core/role';
  import { apiGetPermissionListTree } from '/@/api/core/permission';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const emit = defineEmits(['on-change']);
  const rules: Record<string, Rule[]> = {
    roleName: [
      { required: true },
    ],
  };

  // 权限列表
  const permisstionTree = ref<any[]>([]);

  const handleOpen = () => {
    loadingTask(async () => {
      permisstionTree.value = await apiGetPermissionListTree();
      modalRef.value && modalRef.value.useOpen();
    })
  }

  const handleClose = () => {
    formState.value = {...defaultFromValues}
  }

  const handleSubmit = async () => {
    const values: any = await formRef.value?.validateFields();
    loadingTask(async () => {
      await coreRoleCurdApi('CREATE', values);
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
