<template>
  <div style="border: 1px solid #ccc">
    <Toolbar
      style="border-bottom: 1px solid #ccc"
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      mode="default"
    />
    <Editor
      style="height: 500px; overflow-y: hidden;"
      v-model="valueHtml"
      :defaultConfig="editorConfig"
      mode="default"
      @onCreated="handleCreated"
    />
  </div>
</template>

<script setup lang="ts">
  import '@wangeditor/editor/dist/css/style.css';
  import { onBeforeUnmount, shallowRef, computed } from 'vue';
  import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
  import http from "/@/utils/http.ts";

  const emit = defineEmits(['update:value', 'change', 'change-file-name']);
  
  const prop = defineProps({
    placeholder: String,
    value: String,
  });

  const editorRef = shallowRef();
  
  const valueHtml = computed<string>({
    get: () => {
      if (prop.value && prop.value.length > 0) {
        return `${prop.value}`;
      }
      return '';
    },
    set: (val: any) => {
      emit('update:value', val);
      emit('change', val);
    },
  });

  const toolbarConfig = {}
  const editorConfig = {
    placeholder: prop.placeholder,
    MENU_CONF: {
      uploadImage: {
        allowedFileTypes: ['image/*'],
        async customUpload(file: any, insertFn: any) { // 文件上传
          const formData = new FormData();
          formData.set('file', file);
          const result = await http.post('/core/oss/uploadFile', formData, {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          })
          
          const imageUrl = import.meta.env.VITE_API_FILE_URL
          // 插入到富文本编辑器中，主意这里的三个参数都是必填的，要不然控制台报错：typeError: Cannot read properties of undefined (reading 'replace')
          insertFn(`${imageUrl}${result}`, file.name, file.name)
        }
      }
    }
  }

  onBeforeUnmount(() => {
    editorRef.value && editorRef.value.destroy();
  })

  const handleCreated = (editor: any) => {
    editorRef.value = editor
  }
</script>
