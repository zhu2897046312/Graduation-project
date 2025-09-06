<template>
  <BasicModal
    title="管理员信息"
    @on-ok="handleSubmit"
    @on-close="handleClose"
    ref="modalRef"
  >
    <Form ref="formRef" v-bind="formProp" :model="formState" :rules="rules">
      <Form.Item name="admin_status" label="状态">
        <Radio.Group button-style="solid" v-model:value="formState.admin_status">
          <Radio.Button :value="1">启用</Radio.Button>
          <Radio.Button :value="2">停用</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item name="dept_id" label="所属部门">
        <TreeSelect
          v-model:value="formState.dept_id"
          show-search
          style="width: 100%"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          placeholder="请选择部门"
          allow-clear
          tree-default-expand-all
          :tree-data="tree"
        />
      </Form.Item>
      <Form.Item name="nickname" label="昵称">
        <Input v-model:value="formState.nickname" />
      </Form.Item>
      <Form.Item name="account" label="登陆账户">
        <Input v-model:value="formState.account" />
      </Form.Item>
      <Form.Item name="pwd" label="登陆密码">
        <Input.Password placeholder="如不修改密码，则留空！" v-model:value="formState.pwd" />
      </Form.Item>
      <Form.Item name="confrim_pwd" label="确认密码">
        <Input.Password placeholder="如不修改密码，则留空！" v-model:value="formState.confrim_pwd" />
      </Form.Item>
      <Form.Item name="roles" label="分配角色">
        <Select v-model:value="formState.roles" mode="tags" multiple>
          <Select.Option v-for="(item, index) in roles" :title="item.role_name" :key="index" :value="item.id">
            {{ item.role_name }}
          </Select.Option>
        </Select>
      </Form.Item>
      <Form.Item name="super_state" label="是否超级管理员">
        <Select v-model:value="formState.super_state">
          <Select.Option :value="1">超级管理员</Select.Option>
          <Select.Option :value="2">镇联络站管理员</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item v-if="formState.super_state == 2" name="liaison_station_list" label="分配联络站">
        <Select v-model:value="formState.liaison_station_list" mode="multiple">
          <Select.Option v-for="(it, idx) in liaison_station_list" :key="idx" :value="it.id">{{ it.title }}</Select.Option>
        </Select>
      </Form.Item>
      <Form.Item name="mobile" label="联系电话">
        <Input v-model:value="formState.mobile" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts">
interface FormState {
  dept_id: number;
  nickname: string;
  account: string;
  pwd: string;
  confrim_pwd: string;
  mobile: string;
  admin_status: number,
  roles: number[];
  super_state: number;
  liaison_station_list: number[];
}
const defaultFromValues: FormState = {
  dept_id: 0,
  nickname: '',
  account: '',
  pwd: '',
  confrim_pwd: '',
  mobile: '',
  admin_status: 1,
  roles: [],
  super_state: 1,
  liaison_station_list: [],
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { Form, Input, TreeSelect, Radio, message, Select } from 'ant-design-vue';
  import type { FormInstance, FormProps } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { loadingTask } from '/@/utils/helper';
  import { apiFormatApiTree } from '/@/api/core/dept';
  import { coreRoleCurdApi } from '/@/api/core/role';
  import { coreAdminCurdApi } from '/@/api/core/admin';
  import BasicModal from '/@/components/Kernel/BasicModal/index.vue';

  const modalRef = ref<any>(null);
  const formState = ref<FormState>({...defaultFromValues});
  const formRef = ref<FormInstance|null>(null);
  const tree = ref<any[]>([]);
  const roles = ref<any[]>([]);
  const emit = defineEmits(['on-change']);
  const liaison_station_list = ref<any[]>([]);
  let id = 0;

  const rules: Record<string, Rule[]> = {
    dept_id: [
      { required: true },
    ],
    nickname: [
      { required: true },
    ],
    account: [
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
      tree.value = await apiFormatApiTree('请选择部门');
      const res = await coreRoleCurdApi('LIST', { page_no: 1, page_size: 300})
      roles.value = res.list;
      formState.value = await coreAdminCurdApi('INFO', id);
      formState.value.roles = formState.value.roles.map((it: any) => { return Number(it.id) });
      modalRef.value && modalRef.value.useOpen();
    })
  }

  const handleClose = () => {
    formState.value = {...defaultFromValues}
  }

  const handleSubmit = async () => {
    const values: any = await formRef.value?.validateFields();
    console.log(values)
    if (values.pwd && values.pwd.length > 0 && values.pwd != values.confrim_pwd) {
      message.warn('两次密码不一致');
      return;
    }
    values.id = id;
    loadingTask(async () => {
      await coreAdminCurdApi('UPDATE', values);
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
