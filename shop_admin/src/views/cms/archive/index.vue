<template>
  <PageLayout>
    <Card title="图文管理">
      <template #extra>
        <Button type="primary" @click="() => { router.push({ name: 'ShopBlogsAdd' }) }">新增文图文</Button>
      </template>
      <TableList ref="tableRef"
        @on-edit="(e: number) => { router.push({ name: 'ShopBlogsModify', query: { id: e } }) }" />
    </Card>
  </PageLayout>
</template>

<script lang="ts" setup>
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import { Card } from 'ant-design-vue';
import TableList from './components/TableList.vue';
import { ref } from 'vue';
import { Button } from 'ant-design-vue';
import router from "/@/router";
const tableRef = ref<any>(null);

</script>
