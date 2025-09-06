<template>
  <div>
    <div v-if="prop.value && prop.value.length > 0" class="img_box">
      <a :href="picUrl" target="_blank">
        <img class="img" :src="picUrl" />
      </a>
      <Button size="small" danger @click="handleRemove">移除</Button>
    </div>
    <Upload
      name="file"
      :action="actionUrl"
      accept="image/png, image/jpeg"
      :headers="actionHeader"
      :showUploadList="false"
      @change="handleChange"
      :beforeUpload="<any>beforeUpload"
      :customRequest="customRequest"
    >
      <Button>
        <UploadOutlined />
        {{ prop.txt ? prop.txt : '点击上传' }}</Button>
    </Upload>
  </div>
</template>


<script setup lang="ts">
  import { computed, onMounted } from 'vue';
  import { Upload, Button, message } from 'ant-design-vue';
  import { UploadOutlined } from '@ant-design/icons-vue';
  import type { UploadChangeParam } from 'ant-design-vue';
  import { useAuthStore } from '/@/store/authStore';
import http from '/@/utils/http';

  const authStore = useAuthStore();

  const prop = defineProps({
    txt: String,
    value: String,
    actionUrl: String,
    needLimit: Object, // 需要限制的宽高({ width: xx , height: xx })
  });

  onMounted(() => {
    console.log(prop.value)
  })

  const emit = defineEmits(['update:value', 'change', 'change-file-name']);

  const picUrl = computed<string>({
    get: () => {
      console.log('get', prop.value);
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

  const beforeUpload = (file: any) => {
    return new Promise((resolve) => {
      handleBeforeUpload(file).then(()=> {
      return resolve(true)
      }).catch(res=>{
        message.error(res)
        return false
      })
    })
  }

const handleBeforeUpload = file => checkImageWH(file, prop.needLimit && prop.needLimit.width, prop.needLimit && prop.needLimit.height)

  // 上传图片尺寸限制
const checkImageWH = (file : any, width : any, height : any) => { // 参数分别是上传的file，想要限制的宽，想要限制的高
    return new Promise(function(resolve, reject) {
      let filereader = new FileReader();
      filereader.onload = e => {
        let src = e.target && e.target.result;
        const image = new Image();
        image.onload = function() {
          if(width && (this as any).width != width){
            reject('上传的图片尺寸不符');
          }else if(height && (this as any).height != height){
            reject('上传的图片尺寸不符');
          }else {
            resolve(true);
          }
        };
        image.onerror = reject;
        (image.src as any) = src;
      };
      filereader.readAsDataURL(file);
    });
  }

  const actionUrl = computed(() => {
    if (prop.actionUrl && prop.actionUrl.length > 0) {
      return prop.actionUrl;
    }
    return `${import.meta.env.VITE_API_UPLOAD_IMAGE_URL}`;
  });

  const actionHeader = computed<any>(() => {
    return {
      Authorization: authStore.currentToken,
    };
  });

  const handleChange = (info: UploadChangeParam) => {
    console.log('load===>', info)
    if (info.file.status === 'done') {
      const resp = info.file.response;
      console.log('res===>', resp)
      if (resp) {
        picUrl.value = resp;
        emit('change-file-name', info.file.name);
      } else {
        message.success('文件上传失败');
      }
    } else if (info.file.status === 'error') {
      message.error('文件上传失败');
    }
  };
  
  const handleRemove = () => {
    picUrl.value = '';
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
  }
  .img_box .img {
    display: block;
    max-width: 120px;
    max-height: 120px;
    border: 1px solid #f5f5f5;
    padding: 3px;
    background-color: #fff;
    border-radius: 5px;
    margin-bottom: 2px;
  }
</style>