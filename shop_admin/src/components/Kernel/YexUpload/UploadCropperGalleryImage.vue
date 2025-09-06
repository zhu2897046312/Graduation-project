<template>
  <div>
    <div class="img_box">
      <div v-for="(item, index) in pics" class="img_item">
        <Image :src="item" width="80px" height="80px" />
        <Button size="small" danger @click="handleRemove(index)">移除</Button>
      </div>
    </div>
    <div v-if="pics.length < prop.maxNum" class="relative">
      <Button>
        <UploadOutlined />
        {{ prop.txt ? prop.txt : '点击上传' }}</Button>
        <input ref="fileRef" type="file" accept="image/*" class="absolute left-0 top-0 right-0 bottom-0 opacity-0 z-20" @change="handleChange" />
    </div>
    <CropperImage ref="cropperImageRef" @change="handleCrop" :aspect-ratio="prop.aspectRatio ?? 1" />
  </div>
</template>


<script setup lang="ts">
  import { computed, onMounted, ref } from 'vue';
  import { Button, message, Image } from 'ant-design-vue';
  import { UploadOutlined } from '@ant-design/icons-vue';
  import CropperImage from './CropperImage.vue';
  import http from '/@/utils/http';



  const cropperImageRef = ref<any>(null)

  const prop = defineProps<{
    txt: string,
    value: string[],
    aspectRatio: number,
    maxNum: number,
  }>();

  onMounted(() => {
    console.log(prop.value)
  })

  const emit = defineEmits(['update:value', 'change']);

  const pics = computed<string[]>({
    get: () => {
      console.log('get', prop.value);
      if (prop.value && prop.value.length > 0) {
        return JSON.parse(JSON.stringify(prop.value));
      }
      return [];
    },
    set: (val: any) => {
      console.log('set', val);
      emit('update:value', val);
      emit('change', val);
    },
  });

  const fileRef = ref<any>(null)



  const handleChange = (e: any) => {
    if (e.target.files.length === 0) {
      message.warn('请选择文件')
      return
    }
    const render = new FileReader();
    render.onload = function () {
      if (this.result) {
        const baseUrl = render.result;
        cropperImageRef.value.useOpen(baseUrl)
        fileRef.value.value = ''
      } else {
        message.warn('文件读取失败')
      }
    };
    render.readAsDataURL(e.target.files[0]);
  };
  
  const handleRemove = (e: number) => {
    const s: string[] = JSON.parse(JSON.stringify(pics.value))
    s.splice(e, 1);
    pics.value = s;
  };

  const base64ToFile = (base64: string, mine:string, filename: string) => {
    const bstr = atob(base64)
    let n = bstr.length
    let u8arr = new Uint8Array(n)
    while (n--) {
      u8arr[n] = bstr.charCodeAt(n)
    }
    return new File([u8arr], filename, {
      type: mine
    })
  }

  const handleCrop = async (base64Url: string) => {
    const hide = message.loading('上传中,请稍后...', 0)
    try {
      const file = base64ToFile(base64Url, 'image/jpeg', `${new Date().getTime()}.jpg`);
      await customRequest({ 
        file: file,
        onProgress: (process: number) =>{
          console.log(process)
        },
        onSuccess: (url: string) => {
          console.log(url)
          const s: string[] = JSON.parse(JSON.stringify(pics.value))
          s.push(url)
          pics.value = s
        },
      })
      message.success('上传成功')
      cropperImageRef.value.useClose()
    } catch (e) {
      console.error(e)
      message.warn('上传失败')
    } finally {
      hide()
    }
  };

  const customRequest = async (e: any) => {
    console.log(e)
    const file = e.file
    const formData = new FormData();
    formData.append('file', file);
    const res = await http.post('/core/oss/uploadFile', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    const imageUrl = import.meta.env.VITE_API_FILE_URL
    e.onSuccess(`${imageUrl}${res}`)
  }
</script>

<style scoped>
  .img_box {
    padding: 6px;
    display: flex;
    align-items: center;
    justify-items: center;
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 5px;
  }
  .img_item {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 5px;
  }
</style>