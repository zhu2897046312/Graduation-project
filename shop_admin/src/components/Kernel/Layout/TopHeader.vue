<template>
  <div class="top_header">
    <div class="side">
      <div class="side_menu" @click="() => { open = !open }">
        <YexIcon :icon="open ? 'menu-unfold-outlined' : 'menu-fold-outlined'" color="#666" :size="24" />
      </div>
    </div>
    <div class="title"></div>
    <div class="main">
      <Dropdown>
        <div class="user_box">
<!--          <Avatar :size="23">A</Avatar>-->
          <span class="nickname">{{authStore && authStore.user && authStore.user.nickname}}</span>
        </div>
        <template #overlay>
          <Menu @click="handleUserMenuClick">
            <Menu.Item key="change_pwd">
              <div class="menu_item">
                <YexIcon icon="lock-filled" :size="18" color="#666" />
                修改密码
              </div>
            </Menu.Item>
            <Menu.Divider />
            <Menu.Item key="logout">
              <div class="menu_item">
                <YexIcon icon="logout-outlined" :size="18" color="#666" />
                退出
              </div>
            </Menu.Item>
          </Menu>
        </template>
      </Dropdown>
    </div>
    <ChangePwd ref="changePwdRef" />
  </div>
</template>

<script lang="ts" setup>
  import YexIcon from '/@/components/Kernel/YexIcon/index.vue';
  import { Dropdown, Menu, message } from 'ant-design-vue';
  import { useAuthStore } from '/@/store/authStore';
  import { useRouter } from 'vue-router';
  import { computed, ref } from 'vue';
  import ChangePwd from './ChangePwd.vue';

  const changePwdRef = ref<any>(null);

  const authStore = useAuthStore();
  const router = useRouter();

  const prop = defineProps({
    open: {
      type: Boolean,
      default: true,
    }
  });
  const emit = defineEmits(['update:open'])

  const open = computed<boolean>({
    get: () => {
      return prop.open;
    },
    set: (value: boolean) => {
      emit('update:open', value);
    }
  })

  const handleUserMenuClick = (e: any) => {
    switch (e.key) {
      case 'change_pwd':
        changePwdRef.value && changePwdRef.value.useOpen()
        break;
      case 'logout':
        const hide = message.loading('退出中...', 0);
        authStore.loginOut();
        setTimeout(() =>{
          hide();
          router.push({ name: 'Login' })
        }, 300);
        break;
      default:
        break;
    }
  }
</script>

<style scoped>
  .top_header {
    display: flex;
    align-items: center;
    height: 60px;
  }
  .side {
    padding-left: 10px;
  }
  .side_menu {
    cursor: pointer;
    height: 60px;
  }
  .side_menu:hover {
    opacity: 0.8;
  }
  .title {
    flex: 1;
  }
  .main {
    padding-right: 10px;
  }
  .user_box {
    padding-left: 8px;
    padding-right: 8px;
  }
  .user_box:hover {
    background-color: #f1f1f1;
  }
  .nickname {
    padding-left: 5px;
    user-select: none;
  }
  .menu_item {
    display: flex;
    align-items: center;
    user-select: none;
  }
</style>