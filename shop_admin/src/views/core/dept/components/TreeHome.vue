<template>
   <ApiBasicTable
    ref="tableRef"
    :api="apiLoadData"
    :columns="columns"
    :show-id="false"
    :show-page="false">
    <template #bodyCell="{column, record }">
      <Button.Group v-if="column.key == '_action'">
        <Button type="link" @click="emit('on-edit', record.value)">编辑</Button>
        <Button type="link" @click="emit('on-add', record.value)">添加子部门</Button>
        <Popconfirm title="是否确定删除?" @confirm="() => { handleDel(record.value) }">
          <Button type="link" danger>删除</Button>
        </Popconfirm>
      </Button.Group>
      <p v-if="column.key == 'label'" :style="{ textIndent: `${record.level * 2}em` }">
        {{ record.label }}
      </p>
    </template>
  </ApiBasicTable>
</template>

<script lang="ts">
function treeToLine(source: any, level: number): any[] {
  let out: any[] = [];
  if (source){
    source.level = level;
    out.push(source)
    if (source.children) {
      for (const child of source.children) {
        const cs = treeToLine(child, level + 1);
        out.push(...cs);
      }
    }
  }
  return out;
}
</script>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { apiGetDeptList } from '/@/api/core/dept';
  import { message, Button, Popconfirm } from 'ant-design-vue';
  import { loadingTask } from '/@/utils/helper';
  import ApiBasicTable from '/@/components/Kernel/BasicTable/ApiBasicTable.vue';

  const emit = defineEmits(['on-add', 'on-edit'])
  const tableRef = ref<any>(null);

  const apiLoadData = async (): Promise<any> => {
    const res = await apiGetDeptList();
    const list: any[] = [];
    for (const it of res) {
      list.push(...treeToLine(it, 0));
    }
    console.log(list)
    return {
      list: list.map(it => {
        if (it.children) {
          Reflect.deleteProperty(it, 'children')
        }
        return it;
      }),
      total: res.length,
    };
  }

  const columns: any[] = [
    { dataIndex: 'label', title: '部门名称', key: 'label' },
    { dataIndex: ['node', 'connect_name'], title: '负责人' },
    { dataIndex: ['node', 'connect_mobile'], title: '电话' },
    { dataIndex: ['node', 'remark'], title: '备注' },
  ];

  const handleDel = (e: number) => {
    console.debug(e);
    loadingTask(async () => {
      message.info('暂未开放删除接口');
    }, {
      msg: '删除中...'
    })
  }

  defineExpose({
    useReload: () => { tableRef.value && tableRef.value.useReload() },
  })

</script>
