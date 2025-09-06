<template>
  <Menu
      mode="inline"
      v-model:selected-keys="currentKey"
      v-model:open-keys="openKey"
      @click="handleMenuClick">
    <template v-for="(one) in menus">
      <template v-if="one.children && one.children.length > 0">
        <Menu.SubMenu :key="one.name">
          <template #icon>
            <YexIcon :icon="(one.meta?.icon as string)" color="#333333"/>
          </template>
          <template #title v-if="one.name === 'TaskFlow'">{{ one.meta!.title }} {{ peopleNotAssignMission }}</template>
          <template #title v-else><span style="user-select: none;">{{ one.meta!.title }}</span></template>
          <Menu.Item v-for="(second) in one.children" :key="second.name">
            <template #icon>
              <YexIcon :icon="(second.meta?.icon as string)" color="#333333"/>
            </template>
            <span  style="user-select: none;">{{ second.meta!.title }}</span>
          </Menu.Item>
        </Menu.SubMenu>
      </template>
      <template v-else>
        <Menu.Item :key="one.name">
          <template #icon>
            <YexIcon :icon="(one.meta?.icon as string)" color="#333333"/>
          </template>
          <span style="user-select: none;">{{ one.meta!.title }}</span>
        </Menu.Item>
      </template>
    </template>
  </Menu>
</template>

<script lang="ts">
interface MenuItemInfo {
  name: string;
  meta: {
    icon: string;
    title: string;
    order: number;
  };
  children?: MenuItemInfo[];
};
</script>

<script lang="ts" setup>
import {ref, onMounted, onUnmounted} from 'vue';
import {Menu} from 'ant-design-vue';
import YexIcon from '../../Kernel/YexIcon/index.vue';
import {useRouter, useRoute} from 'vue-router';
import {useAuthStore} from "/@/store/authStore.ts";

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();
const menus = ref<MenuItemInfo[] | any>([]);
const currentKey = ref<string[]>([]);
const openKey = ref<string[]>([]);
const peopleNotAssignMission = ref(0)

var timerRef: any

onMounted(() => {
  currentKey.value = [route.name as string];
  handleLoadMenu();
});

onUnmounted(() => {
  if (timerRef) {
    clearInterval(timerRef)
  }
})


const handleLoadMenu = () => {
  const isMobile = /iPhone|iPad|iPod|Android/i.test(navigator.userAgent);
  
  menus.value = router.options.routes[0].children?.filter(item => {
    return item.meta && !(item.meta.hideMenu == true);
  }).filter(item => {
    if (isMobile) {
      return item.path === 'dashboard'
    } else {
      console.log('PC端权限检查 - 路由:', item.name);
      console.log('用户信息:', authStore.user);
      console.log(item, "item")
      // 这里过滤权限
      if (item && item.meta && item.meta.permission && authStore.user) {
        const permissions = item.meta.permission;
        console.log('路由需要的权限:', permissions);
        console.log('用户拥有的权限:', authStore.user.permission);
        const piers: any = item.meta.permission;
        console.log((piers))
        for (const ic of piers) {
          if (authStore.user.permission.includes(ic)) {
            return true;
          }
        }
        return false;
      }
      return true
    }
  }).map(item => {
    const has_name = item.children?.filter(ic => {
      return ic.name === route.name
    })
    if (has_name && has_name.length > 0) {
      openKey.value = [item.name as string];
    }
    return {
      name: item.name,
      meta: item.meta,
      children: item.children?.filter(ic => {
        return ic.meta && !(ic.meta.hideMenu == true);
      }).filter(ic => {
        // 这里过滤权限
        return ic.meta
      }).map(ic => {
        return {
          name: ic.name,
          meta: ic.meta,
        }
      })
    }
  });

  console.log(menus.value, "menus")
}

const handleMenuClick = (e: any) => {
  router.push({name: e.key})
}

</script>
