<template>
  <PageLayout>
    <Card  title="类目列表">
      <template #extra>
        <Button type="primary" @click="() => { addRef.useOpen({ pid: 0 }) }">新增类目</Button>
      </template>
      <TableList ref="tableRef" 
        @on-add="(e: number) => { addRef.useOpen({ pid: e }); }"
        @on-edit="(e: number) => { editRef.useOpen({ id: e }); }"
        @on-cont="(e: number) => {  }"
      />
    </Card>
  </PageLayout>
  <Add ref="addRef" @on-change="() => { tableRef.useReload() }" />
  <Modify ref="editRef" @on-change="() => { tableRef.useReload() }" />
  <!-- <Cont ref="contRef" @on-change="() => { tableRef.useReload() }" /> -->
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TableList from './components/TableList.vue';
  import Add from './components/Add.vue';
  import Modify from './components/Modify.vue';
  import { ref } from 'vue';
  import { Button } from 'ant-design-vue';

  const addRef = ref<any>(null);
  const editRef = ref<any>(null);
  // const contRef = ref<any>(null);
  const tableRef = ref<any>(null);


</script>