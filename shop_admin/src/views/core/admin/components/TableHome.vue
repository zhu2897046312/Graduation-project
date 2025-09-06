<template>
  <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-page="true">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record .id)">编辑</Button>
      </Button.Group>
      <template v-else-if="column.key === 'admin_status'">
          <EnumLabel dictCode="CareAdminConstant$Status" :dictValue="record.admin_status" />
        </template>
    </template>
  </ApiBasicTable>
</template>

<script lang="ts" setup>
  import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
  import { coreAdminCurdApi } from '/@/api/core/admin';
  import { Button } from 'ant-design-vue';
  import { ref } from 'vue';
  import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue';

  const tableRef = ref<any>(null);
  const emit = defineEmits(['on-add', 'on-edit']);

  const columns: any[] = [
    {
      title: '昵称',
      dataIndex: 'nickname'
    },
    {
      title: '登陆账户',
      dataIndex: 'account'
    },
    {
      title: '状态',
      dataIndex: 'admin_status',
      key: 'admin_status',
      width: 130
    },
  ];

  
  const apiLoadData = (values: any) => {
    return coreAdminCurdApi('LIST', values);
  }

  
  defineExpose({
    useReload: () => {
      tableRef.value && tableRef.value.useReload();
    }
  })
</script>
