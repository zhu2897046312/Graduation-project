<template>
  <Modal
    :title="prop.title"
    :width="prop.width"
    v-model:visible="visible"
    :destroy-on-close="true"
    :mask-closable="prop.maskClosable"
    :z-index="prop.zIndex"
    @cancel="handleClose"
    :mask="true"
  >
    <div>
      <iframe class="tx-map" width="100%" height="100%" src="https://apis.map.qq.com/tools/locpicker?search=1&type=1&key=SBWBZ-UACRP-VSGDC-LDDQG-D7MBT-45FIM&referer=wurd">
      </iframe>
    </div>
    <template #footer>
      <slot name="footer">
        <Button.Group>
          <Button v-if="prop.cancelText"  @click="handleClose">{{ prop.cancelText }}</Button>
          <Button v-if="prop.okText" type="primary" @click="onOk">{{ prop.okText }}</Button>
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
    default: 680,
  },
  zIndex: {
    type: Number,
    default: 1000,
  },
  maskClosable: {
    type: Boolean,
    default: false,
  },
  okText: {
    type: String,
    default: '确认',
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
  minHeight: '200px',
});
let mapObj: {lat: number, lng: number, address: string} = {
  lat: 0,
  lng: 0,
  address: ''
};

onMounted(() => {
  handleCmpHeight();
  window.addEventListener('resize', handleCmpHeight)
  window.addEventListener('message', handleMapMessage, false);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleCmpHeight)
  window.removeEventListener('message', handleMapMessage);
});

watch(visible, (newVal, oldVal) => {
  if (oldVal != newVal && !newVal) {
    emit('on-close');
  }
})

const onOk = () => {
  emit('on-ok', mapObj);
  handleClose();
}

const handleMapMessage = (event: any) => {
  // 接收位置信息，用户选择确认位置点后选点组件会触发该事件，回传用户的位置信息
  var loc = event.data;
  if (loc && loc.module == 'locationPicker') {//防止其他应用也会向该页面post信息，需判断module是否为'locationPicker'
    console.log('location', loc);
    mapObj = {
      lat: loc.latlng.lat,
      lng: loc.latlng.lng,
      address: loc.poiaddress
    }
  }
}

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
.tx-map {
  border: 0;
  margin: 0 auto;
  padding: 0;
  width: 600px;
  height: 560px;
}
</style>