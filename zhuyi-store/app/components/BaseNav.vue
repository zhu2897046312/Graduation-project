<script setup lang="ts">
import api from '../api'
import type { CategoryTreeNode, NavMenuItem } from '../types/type'

// 获取类目树
const { data: categoryTree } = await useAsyncData('category-tree', () => {
  return api.shop.category.tree()
})

// 构建下拉菜单项（支持三级菜单）
const buildMenuItems = (children: CategoryTreeNode[] | undefined): NavMenuItem[] => {
  if (!children || children.length === 0) return []

  return children.map((child) => {
    const hasGrandchildren = child.children && child.children.length > 0

    const menuItem: NavMenuItem = {
      label: child.node.title,
      to: `/collections/${child.node.code}`
    }

    // 如果有三级菜单，添加 children
    if (hasGrandchildren) {
      menuItem.children = child.children!.map(grandson => ({
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
    <template
      v-for="(item, index) in categoryTree"
      :key="index"
    >
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
