<script setup lang="ts">
import { Modal } from 'ant-design-vue';
import { nextTick, ref } from 'vue';
import Cropper from 'cropperjs';
import 'cropperjs/dist/cropper.min.css'

const props = defineProps<{aspectRatio: number}>()
const emit = defineEmits(['change'])

const show = ref(false)

const imageUrl = ref<string>('')

const imageRef = ref<any>(null)

let cropper: Cropper|null = null;
const initCropper = () => {
  cropper = new Cropper(imageRef.value, {
    aspectRatio: props.aspectRatio,
    viewMode: 1,
    dragMode: 'move',
  });
}

const handleSubmit = () => {
  if (cropper) {
    const baseUrl = cropper.getCroppedCanvas().toDataURL('image/jpeg', 0.9)
    emit('change', baseUrl.substring(baseUrl.indexOf(',') + 1))
  }
}

const handleClose = () => {
  cropper = null
}

defineExpose({
  useOpen: (e: string) => {
    imageUrl.value = e;
    show.value = true;
    nextTick(() => {
      initCropper()
    })
  },
  useClose: () => {
    show.value = false;
    cropper = null
    imageUrl.value = ''
  }
})
</script>

<template>
  <Modal
    title="图片裁剪"
    v-model:open="show"
    ok-text="裁剪"
    cancel-text="取消"
    @ok="handleSubmit"
    @cancel="handleClose"
    :destroy-on-close="true"
    :mask-closable="false"
  >
    <img ref="imageRef" :src="imageUrl" style="max-width: 960px; max-height: 680px;display: block;" />
  </Modal>
</template>