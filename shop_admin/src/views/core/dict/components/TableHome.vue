<template>
  <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-page="true">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record .id)">编辑</Button>
        <Button type="link" @click="handleToItem(record .id)">字典属性</Button>
        <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record .id) }">
          <Button type="link" danger>删除</Button>
        </Popconfirm>
      </Button.Group>
      <template v-else-if="column.key === 'state'">
          <EnumLabel dictCode="CommonConstant$State" :dictValue="record.state" />
        </template>
    </template>
  </ApiBasicTable>
</template>

<script lang="ts" setup>
  import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';
  import { coreDictCurdApi } from '/@/api/core/dict';
  import { Button, message, Popconfirm } from 'ant-design-vue';
  import { ref } from 'vue';
  import { loadingTask } from '/@/utils/helper';
  import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const tableRef = ref<any>(null);
  const emit = defineEmits(['on-add', 'on-edit']);

  const columns: any[] = [
    {
      title: '字典名称',
      dataIndex: 'title'
    },
    {
      title: '字典标识',
      dataIndex: 'code'
    },
    {
      title: '状态',
      dataIndex: 'state',
      key: 'state',
      width: 130
    },
    {
      title: '备注',
      dataIndex:'remark'
    },
  ];

  
  const apiLoadData = (values: any) => {
    return coreDictCurdApi('LIST', values);
  }

  const handleDel = (e: number) => {
    loadingTask(async () => {
      await coreDictCurdApi('DELETE', e);
      message.success('删除成功');
      tableRef.value && tableRef.value.useReload();
    }, {
      msg: '删除中...'
    })
  }

  const handleToItem = (e: number) => {
    router.push({ name: 'CoreDictItem', query: { dictId: e } })
  }

  
  defineExpose({
    useReload: () => {
      tableRef.value && tableRef.value.useReload();
    }
  })
</script>
