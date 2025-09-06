<template>
  <Modal
    :title="prop.title"
    :width="prop.width"
    v-model:open="visible"
    :destroy-on-close="true"
    :mask-closable="prop.maskClosable"
    :z-index="prop.zIndex"
    @cancel="handleClose"
    :centered="true"
    :mask="true"
    >
    <div class="modal_body" :style="body_style">
      <slot />
    </div>
    <template #footer>
      <slot name="footer">
          <Button.Group>
            <Button v-if="prop.cancelText"  @click="handleClose">{{ prop.cancelText }}</Button>
            <Button v-if="prop.otherText"  @click="() => { emit('on-ok', 0)}">{{ prop.otherText }}</Button>
            <Button v-if="prop.okText" type="primary" @click="() => { emit('on-ok', 1) }">{{ prop.okText }}</Button>
          </Button.Group>
      </slot>
    </template>
  </Modal>
</template>

<script lang="ts" setup>
  import { ref, onMounted, onUnmounted, watch } from 'vue';
  import { Modal, Button } from 'ant-design-vue';

  const emit = defineEmits(['on-ok', 'on-close'])
  const prop = defineProps({
    title: String,
    width: {
      type: Number,
      default: 980,
    },
    zIndex: {
      type: Number,
      default: 1000,
    },
    maskClosable: {
      type: Boolean,
      default: false,
    },
    otherText:{
      type: String,
      default: '',
    },
    okText: {
      type: String,
      default: '提交',
    },
    cancelText: {
      type: String,
      default: '取消',
    },
    customFooter: {
      type: Boolean,
      default: false,
    }
  });

  const visible = ref<boolean>(false);
  const body_style = ref<any>({
    maxHeight: 600,
  });

  onMounted(() => {
    handleCmpHeight();
    window.addEventListener('resize', handleCmpHeight)
  });

  onUnmounted(() => {
    window.removeEventListener('resize', handleCmpHeight)
  });

  watch(visible, (newVal, oldVal) => {
    if (oldVal != newVal && !newVal) {
      emit('on-close');
    }
  })

  const handleCmpHeight = () => {
    body_style.value.maxHeight = `${window.innerHeight - 330}px`;
  }

  const handleClose = () => {
    visible.value = false;
  }

  const handleOpen = () => {
    visible.value = true;
  }

  defineExpose({
    useOpen: handleOpen,
    useClose: handleClose,
  })
</script>

<style scoped>
  .modal_body {
    overflow-y: auto;
    padding: 6px;
  }
</style>