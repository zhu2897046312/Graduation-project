<script setup lang="ts">
const props = defineProps<{ 
  list?: any[],
  defaultSelected?: string 
}>()
const emit = defineEmits(['change'])

// 关键修复：确保list始终是数组（空数组兜底）
const safeList = computed(() => props.list || [])

const current_list = computed(() => {
  // 使用安全的数组进行操作
  const list: any[] = JSON.parse(JSON.stringify(safeList.value))
  return list.map(warp => {
    warp.value = warp.value.map((item: any) => {
      if (current_selected.value.find(ic => ic.prop_id === warp.id && ic.value_id === item.id)) {
        item.selected = true;
      } else {
        item.selected = false;
      }
      return item;
    })
    return warp;
  })
})

const current_selected = useState<{prop_id: number, value_id: number}[]>('current_selected', () => {
  if (props.defaultSelected) {
    const valueIds = props.defaultSelected.split(';').map(Number)
    const selectedItems: {prop_id: number, value_id: number}[] = []
    
    // 关键修复：使用safeList代替props.list，并检查warp.value是否存在
    safeList.value.forEach(warp => {
      // 额外检查warp.value是否存在，避免嵌套报错
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
    return selectedItems
  }
  return []
})

// 其余代码保持不变...
const handleSelect = (prop_id: number, value_id: number) => {
  let hit = false
  for(var i = 0; i < current_selected.value.length; i++) {
    if (current_selected.value[i].prop_id === prop_id) {
      hit = true
      current_selected.value[i].value_id = value_id;
      break;
    }
  }
  if (hit === false) {
    current_selected.value.push({
      prop_id,
      value_id
    })
  }
  emitChange()
}

const emitChange = () => {
  let ids: number[] = []
  for(var i = 0; i < current_selected.value.length; i++) {
    ids.push(current_selected.value[i].value_id)
  }
  emit('change', ids.join(';'))
}
</script>

<template>
  <!-- 模板部分保持不变 -->
  <div class="sku-box py-2">
    <div class="sku-item flex gap-2 mb-4" v-for="item in current_list" :key="item.id">
      <div class="sku-item-title text-[#333] text-sm py-4">{{ item.title }}:</div>
      <div class="sku-item-value flex items-center flex-wrap gap-2">
        <div
          class="able"
          :class="{'selected': child.selected}"
          v-for="child in item.value"
          :key="child.id"
          @click="handleSelect(item.id, child.id)">
          {{ child.title }}
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
/* 样式部分保持不变 */
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
    