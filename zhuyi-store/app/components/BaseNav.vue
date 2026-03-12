<script setup lang="ts">
import api from '../api'

// 获取类目树
const { data: categoryTree } = await useAsyncData('category-tree', () => {
  return api.shop.category.tree()
})

// // 打印分类树数据结构
// watchEffect(() => {
//   if (categoryTree.value) {
//     console.group('🌳 分类树数据')
//     console.log('完整分类树:', JSON.stringify(categoryTree.value, null, 2))
//     console.log('数据类型:', typeof categoryTree.value)
//     console.log('是否为数组:', Array.isArray(categoryTree.value))
//     if (Array.isArray(categoryTree.value)) {
//       console.log('分类数量:', categoryTree.value.length)
//       if (categoryTree.value.length > 0) {
//         console.log('第一个分类:', categoryTree.value[0])
//         console.log('第一个分类的所有属性:', Object.keys(categoryTree.value[0] || {}))
//         console.log('第一个分类完整数据:', JSON.stringify(categoryTree.value[0], null, 2))
        
//         // 检查分类结构
//         const firstCategory = categoryTree.value[0]
//         if (firstCategory) {
//           console.log('第一个分类的 node 属性:', firstCategory.node)
//           console.log('第一个分类的 children 属性:', firstCategory.children)
//           if (firstCategory.children && firstCategory.children.length > 0) {
//             console.log('第一个分类的子分类数量:', firstCategory.children.length)
//             console.log('第一个子分类:', firstCategory.children[0])
//             if (firstCategory.children[0]) {
//               console.log('第一个子分类的 node:', firstCategory.children[0].node)
//               console.log('第一个子分类的 children:', firstCategory.children[0].children)
//             }
//           }
//         }
//       }
//     } else if (typeof categoryTree.value === 'object' && categoryTree.value !== null) {
//       console.log('对象键:', Object.keys(categoryTree.value))
//       console.log('对象完整数据:', JSON.stringify(categoryTree.value, null, 2))
//     }
//     console.groupEnd()
//   }
// })

// 菜单项类型
interface CategoryNode {
  node: {
    code: string
    title: string
    [key: string]: any
  }
  children?: CategoryNode[]
}

// 构建下拉菜单项（支持三级菜单）
const buildMenuItems = (children: CategoryNode[] | undefined): any[] => {
  if (!children || children.length === 0) return []
  
  return children.map((child) => {
    const hasGrandchildren = child.children && child.children.length > 0
    
    const menuItem: any = {
      label: child.node.title,
      to: `/collections/${child.node.code}`
    }
    
    // 如果有三级菜单，添加 children
    if (hasGrandchildren) {
      menuItem.children = child.children!.map((grandson) => ({
        label: grandson.node.title,
        to: `/collections/${grandson.node.code}`
      }))
    }
    
    return menuItem
  })
}
</script>

<template>
  <div class="flex items-center gap-2">
    <!-- 首页链接 -->
    <UButton
      to="/"
      variant="ghost"
      color="neutral"
      size="sm"
      class="font-medium"
    >
      首页
    </UButton>

    <!-- 分类菜单 -->
    <template v-for="(item, index) in categoryTree" :key="index">
      <!-- 有子菜单的分类 -->
      <UDropdownMenu
        v-if="item.children && item.children.length > 0"
        :items="buildMenuItems(item.children)"
        :popper="{ placement: 'bottom-start', offsetDistance: 4 }"
        :ui="{ content: 'min-w-[200px]' }"
      >
        <UButton
          :to="`/collections/${item.node.code}`"
          variant="ghost"
          color="neutral"
          size="sm"
          trailing-icon="i-lucide-chevron-down"
          class="font-medium"
        >
          {{ item.node.title }}
        </UButton>
      </UDropdownMenu>
      
      <!-- 无子菜单的分类 -->
      <UButton
        v-else
        :to="`/collections/${item.node.code}`"
        variant="ghost"
        color="neutral"
        size="sm"
        class="font-medium"
      >
        {{ item.node.title }}
      </UButton>
    </template>
  </div>
</template>
