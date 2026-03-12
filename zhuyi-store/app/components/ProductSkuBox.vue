<script setup lang="ts">
import type { SpProductProdFrontVo } from '../api/type'

const props = withDefaults(defineProps<{
  list: SpProductProdFrontVo[],
  defaultSelected?: string
}>(), {
  list: () => [] 
})

const emit = defineEmits(['change'])

// 使用 ref 而不是 useState 来避免 SSR 问题
const current_selected = ref<{prop_id: number, value_id: number}[]>([])

// 在 onMounted 中设置初始选中状态，确保只在客户端执行
onMounted(() => {
  if (props.defaultSelected && props.list && props.list.length > 0) {
    const valueIds = props.defaultSelected.split(';').map(Number)
    const selectedItems: {prop_id: number, value_id: number}[] = []
    
    props.list.forEach(warp => {
      // 确保 warp.value 存在且是数组
      if (warp.value && Array.isArray(warp.value)) {
        warp.value.forEach((item: any) => {
          if (valueIds.includes(item.id)) {
            selectedItems.push({
              prop_id: warp.id,
              value_id: item.id
            })
          }
        })
      }
    })
    current_selected.value = selectedItems
    emitChange() // 初始化后触发一次 change 事件
  }
})

const current_list = computed(() => {
  if (!props.list || props.list.length === 0) {
    return []
  }
  
  const list: any[] = JSON.parse(JSON.stringify(props.list))
  return list.map(warp => {
    // 确保 warp.value 存在且是数组
    if (!warp.value || !Array.isArray(warp.value)) {
      warp.value = []
      return warp
    }
    
    warp.value = warp.value.map((item: any) => {
      // 使用可选链操作符避免 SSR 期间的错误
      const isSelected = current_selected.value?.find(ic => 
        ic.prop_id === warp.id && ic.value_id === item.id
      )
      item.selected = !!isSelected
      return item
    })
    return warp
  })
})

const handleSelect = (prop_id: number, value_id: number) => {
  let hit = false
  for(let i = 0; i < current_selected.value.length; i++) {
    const item = current_selected.value[i];
    if (item?.prop_id === prop_id) {
      hit = true;
      item.value_id = value_id;
      break;
    }
  }
  
  if (!hit) {
    current_selected.value.push({
      prop_id,
      value_id
    })
  }
  emitChange()
}

const emitChange = () => {
  if (!current_selected.value || !Array.isArray(current_selected.value)) {
    emit('change', '')
    return
  }
  
  const ids = current_selected.value
    .filter(item => item !== null && item !== undefined)
    .map(item => item.value_id)
  emit('change', ids.join(';'))
}
</script>

<template>
  <div class="sku-box py-2">
    <div v-if="current_list && current_list.length > 0">
      <div class="sku-item flex gap-4 mb-4" v-for="item in current_list" :key="item.id">
        <div class="sku-item-title text-gray-700 dark:text-gray-300 text-sm font-medium py-2 min-w-[80px]">
          {{ item.title }}:
        </div>
        <div class="sku-item-value flex items-center flex-wrap gap-2 flex-1">
          <UButton
            v-for="child in (item.value || [])"
            :key="child.id"
            :variant="child.selected ? 'solid' : 'outline'"
            :color="child.selected ? 'primary' : 'neutral'"
            :disabled="child.disabled"
            size="sm"
            :class="[
              'sku-option',
              child.selected && 'sku-option-selected',
              child.disabled && 'sku-option-disabled'
            ]"
            @click="!child.disabled && handleSelect(item.id, child.id)"
          >
            {{ child.title }}
          </UButton>
        </div>
      </div>
    </div>
    <div v-else class="text-sm text-gray-500 dark:text-gray-400 py-4">
      暂无可用选项
    </div>
  </div>
</template>

<style lang="css" scoped>
.sku-option {
  transition: all 0.2s ease;
}

.sku-option:not(.sku-option-disabled):hover {
  transform: scale(1.02);
}

.sku-option-disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.sku-option-selected {
  background-color: #FB7F86 !important;
  border-color: #FB7F86 !important;
  color: white !important;
}
</style>


