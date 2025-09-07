<script setup>
import { NConfigProvider, NDialogProvider, NMessageProvider, NModalProvider } from 'naive-ui';
</script>

<template>
  <NConfigProvider>
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
