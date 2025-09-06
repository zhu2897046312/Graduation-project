<template>
  <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-page="false">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record .id)">编辑</Button>
        <Button v-if="record.pid === 0" type="link" @click="emit('on-add', record.id)">添加子权限</Button>
        <Button type="link" danger>删除</Button>
      </Button.Group>
      <p  v-else-if="column.key === 'title'" :style="{ textIndent: record.pid == 0 ? '0em' : '4em' }">
        {{ record.title }}
      </p>
      <template v-else-if="column.key === 'roleStatus'">
          <EnumLabel dictCode="CoreRoleConstant$RoleStatus" :dictValue="record.roleStatus" />
        </template>
    </template>
  </ApiBasicTable>
</template>

<script lang="ts" setup>
  import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
  import { apiGetPermissionList } from '/@/api/core/permission';
  import { Button } from 'ant-design-vue';
  import { ref } from 'vue';

  const tableRef = ref<any>(null);
  const emit = defineEmits(['on-add', 'on-edit'])
  
  const apiLoadData = async (): Promise<any> => {
    const res = await apiGetPermissionList();
    return {
      list: res,
      total: res.length,
    };
  }

  const columns: any[] = [
    { dataIndex: 'title', title: '权限名称', width: 200 , key: 'title' },
    { dataIndex: 'code', title: '标识' },
    { dataIndex: 'remark', title: '备注' },
  ];

  
  defineExpose({
    useReload: () => {
      tableRef.value && tableRef.value.useReload();
    }
  })
</script>
