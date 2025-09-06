<template>
  <Select
    v-model:value="value"
    show-search
    option-filter-prop="title"
    style="width: 188px;"
    placeholder="请选择">
    <Select.Option v-for="item in list" :title="item.label" :key="item.id" :value="item.value">
      {{ item.label }}
    </Select.Option>
  </Select>
</template>

<script setup lang="ts">
  import { ref, computed, PropType, onMounted } from 'vue';
  import { Select } from 'ant-design-vue';
  import { apiGetListByCode } from '/@/api/core/dictitem';

  const prop = defineProps({
    value: {
      type: Object as PropType<any>,
      required: true,
    },
    dictCode: {
      type: String,
      required: true,
    },
    dictLabel: {
      type: String,
      required: true,
    }
  });
  const emit = defineEmits(['update:value', 'change']);

  const list = ref<any[]>([]);

  onMounted(async () => {
    list.value = await apiGetListByCode(prop.dictCode);
    let isFind = false
    if (list.value.length > 0) {
      for (const it of list.value) {
        if (it.label == prop.dictLabel) {
          value.value = it.value
          isFind = true;
          break;
        }
      }
    }
    if (!isFind) {
      value.value = ''
    }
  })

  const value = computed({
    get: function() {
      return prop.value;
    },
    set: function (val: any) {
      emit('update:value', val);
      emit('change', val);
    }
  })

</script>