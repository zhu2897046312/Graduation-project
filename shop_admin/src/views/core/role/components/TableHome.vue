<template>
  <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-page="true">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record .id)">编辑</Button>
        <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record .id) }">
          <Button type="link" danger>删除</Button>
        </Popconfirm>
      </Button.Group>
      <template v-else-if="column.key === 'roleStatus'">
          <EnumLabel dictCode="CoreRoleConstant$RoleStatus" :dictValue="record.role_status" />
        </template>
    </template>
  </ApiBasicTable>
</template>

<script lang="ts" setup>
  import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
  import { coreRoleCurdApi } from '/@/api/core/role';
  import { Button, message, Popconfirm } from 'ant-design-vue';
  import { ref } from 'vue';
  import { loadingTask } from '/@/utils/helper';
  import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue';

  const tableRef = ref<any>(null);
  const emit = defineEmits(['on-add', 'on-edit']);

  const columns: any[] = [
    {
      title: '角色名称',
      dataIndex: 'role_name'
    },
    {
      title: '状态',
      dataIndex: 'role_status',
      key: 'roleStatus',
      width: 130
    },
    {
      title: '备注',
      dataIndex:'remark'
    },
  ];

  
  const apiLoadData = (values: any) => {
    return coreRoleCurdApi('LIST', values);
  }

  const handleDel = (e: number) => {
    loadingTask(async () => {
      await coreRoleCurdApi('DELETE', e);
      message.success('删除成功');
      tableRef.value && tableRef.value.useReload();
    }, {
      msg: '删除中...'
    })
  }

  
  defineExpose({
    useReload: () => {
      tableRef.value && tableRef.value.useReload();
    }
  })
</script>
