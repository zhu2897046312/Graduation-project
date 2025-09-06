<template>
  <PageLayout>
    <Card :title="title">
      <template #extra>
        <Button type="primary" @click="() => { addRef.useOpen({ prod_attributes_id: prod_id }) }">新增</Button>
      </template>
      <TableList ref="tableRef" 
        :prod-id="prod_id"
        @on-edit="(e: number) => { editRef.useOpen({ id: e }); }"
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
  import { ref, onMounted } from 'vue';
  import { Button } from 'ant-design-vue';
  import { useRoute } from 'vue-router';
  import api  from '/@/api/';

  const route = useRoute();

  const prod_id = ref<any>(route.query.id)

  const addRef = ref<any>(null);
  const editRef = ref<any>(null);
  const tableRef = ref<any>(null);

  const title = ref('')

  onMounted(async () => {
    const info: any = await api.shop.prodAttributes.info(prod_id.value)
    title.value = `【${info.title}】值列表`
  })


</script>