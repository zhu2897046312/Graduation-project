<template>
  <Layout class="layout_warp">
    <Layout.Sider class="zy_menu_sider" theme="light" v-model:collapsed="collapsed" collapsible>
      <div class="menu_title">
        <img src="/logo.png" />
        <h3 v-if="collapsed == false" class="menu_title_tx">商城</h3>
      </div>
      <YexMenu />
    </Layout.Sider>
    <Layout class="layout_box">
      <Layout.Header class="layout_top_header">
        <TopHeader v-model:open="collapsed" />
      </Layout.Header>
      <Layout.Content class="base-content" style="display: flex; flex-direction: column; overflow: scroll;">
        <RouterView />
      </Layout.Content>
      <!-- <Layout.Footer class="layout_footer">
        <p class="copyright">Copyright © 2024 广东真格有限公司</p>
      </Layout.Footer> -->
    </Layout>
  </Layout>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { Layout } from 'ant-design-vue';
import YexMenu from './YexMenu.vue';
import TopHeader from './TopHeader.vue';
const collapsed = ref<boolean>((/iPhone|iPad|iPod|Android/i.test(navigator.userAgent)));
</script>

<style scoped>
.layout_warp {
  position: relative;
  width: 100%;
  height: 100%;
  background-color: #fff;
}

.layout_box {
  position: relative;
  height: 100%;
}

.menu_title {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu_title>img {
  width: 38px;
  height: 38px;
  display: block;
  padding: 0;
  margin: 0 5px 0 0;
}

.menu_title_tx {
  color: var(--primary-color);
  font-size: 22px;
  font-weight: 700;
  display: block;
  padding: 0;
  margin: 0;
  user-select: none;
}

.layout_top_header {
  padding: 0;
  background: #fff;
  background: -webkit-linear-gradient(top, #fff 0%, #f9f9f9 100%);
  background: linear-gradient(to bottom, #fff 0%, #f9f9f9 100%);
  z-index: 1000;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.layout_footer {
  padding: 12px 20px;
}

.copyright {
  opacity: 0.9;
  color: #666;
  font-size: 14px;
  text-align: center;
}

.zy_menu_sider{
  overflow-y: auto;
}

.base-content::-webkit-scrollbar{
  width: 0;
}
</style>