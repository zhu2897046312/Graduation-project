<template>
  <PageLayout>
    <Card  title="属性列表">
      <template #extra>
        <Button type="primary" @click="() => { addRef.useOpen({ pid: 0 }) }">新增</Button>
      </template>
      <TableList ref="tableRef" 
        @on-add="(e: number) => { addRef.useOpen({ pid: e }); }"
        @on-edit="(e: number) => { editRef.useOpen({ id: e }); }"
        @on-value="(e: number) => { router.push({name: 'ShopCoreProdAttributesValue', query : {id: e}}) }"
      />
    </Card>
  </PageLayout>
  <Add ref="addRef" @on-change="() => { tableRef.useReload() }" />
  <Modify ref="editRef" @on-change="() => { tableRef.useReload() }" />
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TableList from './components/TableList.vue';
  import Add from './components/Add.vue';
  import Modify from './components/Modify.vue';
  import { ref } from 'vue';
  import { Button } from 'ant-design-vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();

  const addRef = ref<any>(null);
  const editRef = ref<any>(null);
  const tableRef = ref<any>(null);


</script>