<template>
  <span v-if="prop.link">{{ dictInfo.label }}</span>
  <Tag v-else><span v-if="prop.needPoint" class="point" :class="prop.dictValue != 1 ? 'done' : ''"></span>{{ dictInfo.label }}</Tag>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';
  import enumDict from '/@/utils/enum-dict';
  // ------ api接口 ------

  // ------ 组件 ------
  import { Tag } from 'ant-design-vue';

  // ------ 变量定义 ------
  const prop = defineProps({
    dictCode: {
      type: String,
      required: true,
    },
    dictValue: {
      type: [String, Number],
      required: true,
    },
    needPoint: {
      type: Boolean
    },
    link: Boolean,
  });
  const dictInfo = computed(() => {
    const info = enumDict.getDictInfo(prop.dictCode, prop.dictValue);
    if (info == null) {
      return {
        label: prop.dictValue,
        value: prop.dictValue,
        color: '',
      };
    }
    return info;
  });

  // ------ 生命周期 ------

  // ------ 方法区 ------
</script>
<style scoped>
  .point{
    display: inline-block;
    width: 10px;
    height: 10px;
    background: rgb(50, 199, 11);
    border-radius: 10px;
    margin-right: 5px;
  }
  .point.done{
    background: rgb(165, 165, 165);
  }
</style>