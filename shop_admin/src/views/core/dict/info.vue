<template>
  <PageLayout>
    <Card 
      size="small"
      title="字典属性列表">
      <template #extra>
        <Button type="primary" @click="addRef.useOpen(info ? info.id : 0)">新增</Button>
      </template>
      <ItemList 
        v-if="info != null" 
        ref="tableRef" 
        :dict-id="info.id"
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
  import ItemList from './components/ItemList.vue';
  import Add from './components/AddItem.vue';
  import Edit from './components/EditItem.vue';
  import { Button } from 'ant-design-vue';
  import { ref, onMounted } from 'vue';
  import { coreDictCurdApi } from '/@/api/core/dict';
  import { loadingTask } from '/@/utils/helper';
  import { useRoute } from 'vue-router';

  const route = useRoute();

  const info = ref<any>(null);
  const addRef = ref<any>(null);
  const editRef = ref<any>(null);
  const tableRef = ref<any>(null);

  onMounted(() => {
    loadingTask(async () => {
      const id = route.query.dictId;
      info.value = await coreDictCurdApi('INFO', id);
    })
  })

</script>
