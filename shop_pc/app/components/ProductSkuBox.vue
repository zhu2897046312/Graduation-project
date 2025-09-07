<script setup lang="ts">
const props = withDefaults(defineProps<{
  list: any[],
  defaultSelected?: string // 添加默认选中的SKU值
}>(), {
  list: () => [] 
})
const emit = defineEmits(['change'])

const current_list = computed(() => {
  // 添加防御性检查
  if (!props.list || props.list.length === 0) {
    return []
  }
  const list: any[] = JSON.parse(JSON.stringify(props.list))
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
  // 如果有默认选中的SKU，解析并设置初始状态
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
    console.log(selectedItems)
    return selectedItems
  }
  return []
})

const handleSelect = (prop_id: number, value_id: number) => {
  let hit = false
  for(var i = 0; i < current_selected.value.length; i++) {
    // 添加非空断言操作符 (!) 告诉 TypeScript 这个值不会为 null 或 undefined
    if (current_selected.value[i]!.prop_id === prop_id) {
      hit = true
      current_selected.value[i]!.value_id = value_id;
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
};

const emitChange = () => {
  let ids: number[] = []
  for(var i = 0; i < current_selected.value.length; i++) {
    // 同样添加非空断言
    ids.push(current_selected.value[i]!.value_id)
  }
  emit('change', ids.join(';'))
};

</script>

<template>
  <div class="sku-box py-2">
    <div class="sku-item flex gap-2 mb-4" v-for="item in current_list">
      <div class="sku-item-title text-[#333] text-sm py-4">{{ item.title }}:</div>
      <div class="sku-item-value flex items-center flex-wrap gap-2">
        <div
          class=""
          :class="{'selected': child.selected}"
          v-for="child in item.value"
          @click="handleSelect(item.id, child.id)">
          {{ child.title }}
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
/* .sku-item {
  display: flex;
  gap: 8px;
  margin-bottom: 10px;
} */
/* .sku-item-title {
  font-size: 14px;
  color: #333;
  padding: 7px 0px;
} */
/* .sku-item-value {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
} */
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


