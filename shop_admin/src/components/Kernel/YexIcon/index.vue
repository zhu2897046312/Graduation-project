<template>
  <span :style="{
    fontSize: `${prop.size}px`
  }" :class="cls"></span>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';

  const prop = defineProps({
    icon: {
      type: String ,
      required: true,
    },
    prefix: {
      type: String,
      default: 'ant-design',
    },
    color: {
      type: String,
      default: '#333',
    },
    size: {
      type: Number,
      default: 20,
    },
  });

  const cls = computed<any>(() => {
    const _cls = `${prop.prefix}:${prop.icon}`
    const domCls: any = {};
    domCls['yex_icon'] = true;
    domCls[_cls] = true;
    return domCls ;
  });
</script>
