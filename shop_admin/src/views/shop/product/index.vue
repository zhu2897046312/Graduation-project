<template>
  <PageLayout>
    <Card  title="商品列表">
      <template #extra>
        <Button type="primary" @click="() => { router.push({ name: 'ShopCoreProductCreate' });  }">新增</Button>
      </template>
      <TableList ref="tableRef" 
        @on-edit="(e: number) => { router.push({ name: 'ShopCoreProductModify', query: { id: e } }); }"
      />
    </Card>
  </PageLayout>
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card } from 'ant-design-vue';
  import TableList from './components/TableList.vue';
  import { Button } from 'ant-design-vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();


</script>