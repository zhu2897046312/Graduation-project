<script setup lang="ts">
import { NDialogProvider, NMessageProvider, NModalProvider, NConfigProvider, type GlobalThemeOverrides } from 'naive-ui';

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColorHover: '#FB7F86',
    iconColor: '#FB7F86',
    successColor: '#FB7F86',
  },
  Button: {
    textColor:'#878787',
    borderPressed: '#FB7F86',
    rippleColor: '#FB7F86',
    textColorPressed: '#FB7F86',
    colorPrimary: '#FB7F86',
    rippleColorPrimary: '#FB7F86',
    borderPressedPrimary: '#FB7F86',
    colorPressedPrimary: '#FB7F86',
    borderPrimary: '#FB7F86',
    colorSuccess: '#FB7F86',
    rippleColorSuccess: '#FB7F86',
    borderPressedSuccess: '#FB7F86',
    colorPressedSuccess: '#FB7F86',
    borderSuccess: '#FB7F86',
    colorFocusSuccess: '#FB7F86',

    //type="info"
    textColorInfo: '#fff',         // 文字颜色
    rippleColorInfo: '#FB7F86',
    borderInfo:'#FB7F86',
    borderHoverInfo:'#FB7F86',
    borderFocusInfo: 'transparent',
    borderPressedInfo:'#FB7F86',
    borderDisabledInfo:'#FB7F86',
    colorInfo: '#FB7F86',          
    colorHoverInfo:'#FB7F86',
    colorFocusInfo: '#FB7F86',
    colorPressedInfo:'#FB7F86',
    colorDisabledInfo: '#FB7F86',
  },  
  Pagination: {
    itemBorder: '1px solid #FB7F86', // 边框颜色
    itemTextColor: '#FB7F86', // 文字颜色
    itemBorderHover: '1px solid #FB7F86', // 悬停时边框颜色
    itemTextColorHover: '#FB7F86', // 悬停时文字颜色
    itemBorderPressed: '1px solid #FB7F86', // 点击时边框颜色
    itemBorderActive: '1px solid #FB7F86', // 激活状态边框颜色
    itemTextColorPressed: '#fff', // 点击时文字颜色
    itemTextColorActive: '#FB7F86', // 激活状态文字颜色
    itemColorActive: '#f4b3c2', // 激活状态背景色
    itemSizeMedium: '32px', // 中等尺寸
    itemBorderRadius: '4px', // 边框圆角
  }
  // ...
}
</script>
<template>
  <NConfigProvider :theme-overrides="themeOverrides">
    <NModalProvider to="body">
      <NMessageProvider>
        <NDialogProvider>
          <NuxtLayout>
            <NuxtPage />
          </NuxtLayout>
        </NDialogProvider>
      </NMessageProvider>
    </NModalProvider>
  </NConfigProvider>
</template>
