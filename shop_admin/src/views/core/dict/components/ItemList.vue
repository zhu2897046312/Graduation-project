<template>
  <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-page="false">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record .id)">编辑</Button>
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
  import { coreDictItemCurdApi, apiGetList } from '/@/api/core/dictitem';
  import { Button, message, Popconfirm } from 'ant-design-vue';
  import { ref } from 'vue';
  import { loadingTask } from '/@/utils/helper';
  import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue';

  const tableRef = ref<any>(null);
  const emit = defineEmits(['on-add', 'on-edit']);

  const prop = defineProps({
    dictId: {
      type: Number,
      required: true,
    },
  })

  const columns: any[] = [
    {
      title: '属性名称',
      dataIndex: 'label'
    },
    {
      title: '属性值',
      dataIndex: 'value'
    },
    {
      title: '状态',
      dataIndex: 'state',
      key: 'state',
      width: 130
    },
    {
      title: '排序',
      dataIndex:'sort_num'
    },
  ];

  
  const apiLoadData = async () => {
    const res = await apiGetList(prop.dictId);
    return {
      total: res.length,
      list: res,
    };
  }

  const handleDel = (e: number) => {
    loadingTask(async () => {
      await coreDictItemCurdApi('DELETE', e);
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
