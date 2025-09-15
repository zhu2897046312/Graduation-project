<script setup lang="ts">
import type { SpProductProdFrontVo } from '../../api/type'

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
  if (props.defaultSelected) {
    const valueIds = props.defaultSelected.split(';').map(Number)
    const selectedItems: {prop_id: number, value_id: number}[] = []
    
    props.list.forEach(warp => {
      warp.value.forEach((item: any) => {
        if (valueIds.includes(item.id)) {
          selectedItems.push({
            prop_id: warp.id,
            value_id: item.id
          })
        }
      })
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
  const ids = current_selected.value
    .filter(item => item !== null && item !== undefined)
    .map(item => item.value_id)
  emit('change', ids.join(';'))
}
</script>

<template>
  <div class="sku-box py-2">
    <div class="sku-item flex gap-2 mb-4" v-for="item in current_list" :key="item.id">
      <div class="sku-item-title text-[#333] text-sm py-4">{{ item.title }}:</div>
      <div class="sku-item-value flex items-center flex-wrap gap-2">
        <div
          :class="{'selected': child.selected, 'able': !child.disabled, 'disabled': child.disabled}"
          v-for="child in item.value"
          :key="child.id"
          @click="!child.disabled && handleSelect(item.id, child.id)">
          {{ child.title }}
        </div>
      </div>
    </div>
  </div>
</template>

<!-- 样式保持不变 -->

<style lang="css" scoped>
.sku-item-value>div {
  display: block;
  border: 1px solid #c7c7c7;
  padding: 5px 8px;
  border-radius: 5px;
  font-size: 14px;
  color: #333;
  cursor: pointer;
  transition: all 0.2s;
}
.sku-item-value>div.able:hover {
  opacity: 0.8;
  transform: scale(1.02);
  border-color: #FB7F86;
}
.sku-item-value>div.disabled {
  background-color: #c7c7c7;
}
.sku-item-value>div.selected {
  border-color: #FB7F86;
  background-color: #FB7F86;
  color: #fff;
}
</style>


