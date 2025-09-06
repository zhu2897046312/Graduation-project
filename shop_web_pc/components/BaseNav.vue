<script setup lang="ts">
import api from '~/api';


// 获取类目树
const { data: categoryTree } = await useAsyncData(() => {
  return api.shop.category.tree();
});
</script>

<template>
 <nav class="nav">
  <ul class="container nav_cont">
    <li>
      <NuxtLink class="flink ami_link" href="/">Home</NuxtLink>
    </li>
    <li class="nav_item" v-for="(item,index) in categoryTree" :key="index">
      <NuxtLink class="flink ami_link" :href="`/collections/${item.node.code}`">
        {{ item.node.title }}
        <!-- <NIcon :size="16" v-if="item.children && item.children.length > 0"><ChevronDown /></NIcon> -->
        <svg 
            v-if="item.children && item.children.length > 0"
            class="w-4 h-4 transition-transform group-hover:rotate-180"
            viewBox="0 0 24 24"
          >
            <path fill="currentColor" d="M7 10l5 5 5-5z"/>
          </svg>
      </NuxtLink>
      <template v-if="item.children && item.children.length > 0">
        <div class="sub_nav">
          <div class="sub_nav_item" v-for="(child,cidx) in item.children" :key="cidx">
            <NuxtLink class="slink ami_link text-[16px]" :href="`/collections/${child.node.code}`">
              {{ child.node.title }}
              <!-- <NIcon :size="16" v-if="child.children && child.children.length > 0"><ChevronForward /></NIcon> -->
              <svg 
                v-if="child.children && child.children.length > 0"
                class="w-4 h-4 transition-transform group-hover:rotate-180"
                viewBox="0 0 24 24"
              >
                <path fill="currentColor" d="M7 10l5 5 5-5z"/>
              </svg>
            </NuxtLink>
            <template v-if="child.children && child.children.length > 0">
              <div class="grandson_nav">
                <div class="grandson_nav_item" v-for="(grandson,gidx) in child.children" :key="gidx">
                  <NuxtLink class="tlink text-[16px]" :href="`/collections/${grandson.node.code}`">{{ grandson.node.title }}</NuxtLink>
                </div>
              </div>
            </template>
          </div>
        </div>
      </template>
    </li>
  </ul>
 </nav>
</template>

<style lang="css" scoped>
.nav {
  margin-top: 20px;
  margin-bottom: 20px;
  position: relative;
  z-index: 20;
}
.nav_cont {
  display: flex;
  align-items: center;
  gap: 8px;
}
.nav_item {
  padding: 8px 12px;
  position: relative;
  transition: all 0.3s;
}
.sub_nav {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background: #FB7F86;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  flex-direction:column;
  transition: all 0.6s;
}
.nav_item:hover > .sub_nav {
  display: flex;
}
.nav_item:hover {
  color: #FB7F86;
}
.sub_nav_item {
  padding: 12px 8px;
  color: #fff;
  transition: all 0.3s;
  position: relative;
}
/* .sub_nav_item:hover {
  opacity: 0.8;
} */
.grandson_nav {
  display: none;
  position: absolute;
  top: 0;
  left: 100%;
  background: #FB7F86;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  flex-direction:column;
  transition: all 0.6s;
}
.sub_nav_item:hover > .grandson_nav {
  display: flex;
}
.grandson_nav_item {
  padding: 12px 8px;
  background-color: #FB7F86;
  color: #fff;
  transition: all 0.3s;
  position: relative;
  min-width: 160px;
}
.flink {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 16px;
  color: #fff;
}
.flink:hover {
  color: #707070;
}
.slink {
  display: flex;
  align-items: center;
  gap: 6px;
}
.slink:hover {
  color: #707070;
}
.tlink:hover {
  color: #707070;
}
</style>