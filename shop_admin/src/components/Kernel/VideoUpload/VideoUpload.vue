<template>
  <div class="full-width">
    <div v-if="props.value.length == 0" class="upload_box">
      <span>上传视频文件</span>
      <input
        class="upload_input"
        type="file"
        accept="video/*"
        @change="handleFileChange"
      />
    </div>
    <div v-else class="video_preview">
      <div class="remove" @click="handleRemove">移除视频</div>
      <video class="video_play" ref="videoRef" controls />
    </div>
  </div>
  <Spin tips="上传中" v-if="uploading"/>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue';
import { Spin } from 'ant-design-vue';
import { oss } from '/@/api/cms'


const props = defineProps({
  value: {
    type: String,
    required: true
  }
})



const emit = defineEmits(['update:value', 'change-video-mate'])

const videoRef = ref<HTMLVideoElement|null>(null)
const uploading = ref<boolean>(false)

const readVideoInfo = async () => {
  console.info('props', props.value)
  if (props.value.length == 0) {
    return
  }
  if (videoRef.value) {
    videoRef.value.addEventListener('loadedmetadata', () => {
      if (videoRef.value) {
        console.log(videoRef.value.duration)
        console.log(videoRef.value.videoWidth, videoRef.value.videoHeight)
        let w = 0, h = 0;
        if (videoRef.value.videoWidth > videoRef.value.videoHeight) {
          // 横板
          w = 638
          h = 638 * videoRef.value.videoHeight / videoRef.value.videoWidth
        } else {
          // 竖版
          h = 438
          w = 438 * videoRef.value.videoWidth / videoRef.value.videoHeight
        }
        videoRef.value.style.width = w + 'px'
        videoRef.value.style.height = h + 'px'
        videoRef.value.style.opacity = '1'
        emit('change-video-mate', {
          duration: Math.floor(videoRef.value.duration),
          width: videoRef.value.videoWidth,
          height: videoRef.value.videoHeight
        })
      }
    })
    videoRef.value.setAttribute('src', props.value)
  }
}

watch(props, () => {
  nextTick(() => {
    readVideoInfo()
  })
})


const handleRemove = () => {
  emit('update:value', '')
}

const handleFileChange = async (e: any) => {
  const file = e.target.files[0]
  const objectName = `journey-vlogger/video/${file.name.split('.')[0]}-${Date.now()}.mp4`
  uploading.value = true
  let uploadUrl = (await oss.presignedUrl(objectName)) as any as string
  await oss.fileUpload(file, uploadUrl, (progress: number) => {console.log(progress)})
  uploading.value = false
  emit('update:value', `https://assets.hywbbs.com/${objectName}`)

}
</script>

<style lang="scss" scoped>
.upload_box {
  position: relative;
  display: flex;
  background-color: #f5f5f5;
  border: 1px dashed #6c9dca;
  border-radius: 5px;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 128px;
  cursor: pointer;
  font-size: 18px;
  color: #333;
  .upload_input {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    opacity: 0;
    cursor: pointer;
    z-index: 2;
  }
}
.video_play {
  display: block;
  width: 90%;
  height: 238px;
  opacity: 0;
  transition: all 0.6s
}
.remove {
  display: block;
  width: 138px;
  text-align: center;
  padding: 6px 10px;
  margin-bottom: 10px;
  border: 1px solid #f50;
  color: #f50;
  font-size: 14px;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.2s;
  :hover {
    background-color: #f50;
    color: #fff;
  }
}
</style>