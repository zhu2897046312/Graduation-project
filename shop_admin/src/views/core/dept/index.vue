<template>
  <PageLayout>
    <Card size="small" title="部门列表">
      <template #extra>
        <Button type="primary" @click="() => { addRef.useOpen({ pid: 0 }) }">新增部门</Button>
      </template>
      <TreeHome ref="tableRef" 
        @on-add="(e: number) => { addRef.useOpen({ pid: e }); }"
        @on-edit="(e: number) => { editRef.useOpen({ id: e }); }"
      />
    </Card>
  </PageLayout>
  <Add ref="addRef" @on-change="() => { tableRef.useReload() }" />
  <Edit ref="editRef" @on-change="() => { tableRef.useReload() }" />
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TreeHome from './components/TreeHome.vue';
  import Add from './components/Add.vue';
  import Edit from './components/Edit.vue';
  import { ref } from 'vue';
  import { Button } from 'ant-design-vue';

  const addRef = ref<any>(null);
  const editRef = ref<any>(null);
  const tableRef = ref<any>(null);

</script>
