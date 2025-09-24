<template>
  <BasicModal title="申请退款" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef" v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }">
      <Form.Item label="退款金额" name="refund_amount">
        <InputNumber v-model:value="formState.refund_amount"  style="width: 100%" :min="0" :precision="2" />
        <div style="color: #999; font-size: 12px; margin-top: 4px;">
          最多可退：{{ form.order.pay_amount }}
        </div>
      </Form.Item>
      <Form.Item label="退款原因" name="reason">
        <Input.TextArea v-model:value="formState.reason" :rows="3" />
      </Form.Item>
      <Form.Item label="凭证图片" name="images">
        <UploadCropperGalleryImage 
          :max-num="6" 
          txt="上传凭证" 
          v-model:value="formState.images" 
          :aspect-ratio="1" 
        />
      </Form.Item>
    </Form>
  </BasicModal>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, InputNumber, message } from 'ant-design-vue';
import type { FormProps, Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import UploadCropperGalleryImage from '/@/components/Kernel/YexUpload/UploadCropperGalleryImage.vue';
import api from '/@/api/index';

const modalRef = ref<any>(null);
const formRef = ref<any>(null);
const formState = ref<any>(null);
const form = ref<any>(null);
const emit = defineEmits(['on-change']);

const rules: Record<string, Rule[]> = {
  refund_amount: [{ required: true, message: '请输入退款金额' }],
  reason: [{ required: false, message: '请输入退款原因' }]
};

const formProp = ref<FormProps>({
  labelCol: { flex: '160px' },
  wrapperCol: { span: 24 },
  labelAlign: 'right'
});

const handleOpen = async (e: any) => {
  form.value = await api.shop.order.info(e.id)
  formState.value = {
    order_id: e.id,
    refund_amount: form.value?.order?.pay_amount || 0, 
    reason: '',
    images: []
  };
  modalRef.value?.useOpen();
};

const handleClose = () => {
  modalRef.value?.useClose();
};

// Make sure to expose these methods
defineExpose({
  useOpen: handleOpen,
  useClose: handleClose
});

const handleSubmit = async () => {
  try {
    const values = await formRef.value?.validateFields();
    values.order_id = formState.value.order_id;
    
    await loadingTask(async () => {
      await api.shop.refund.create(values);
      message.success('退款申请已提交');
      modalRef.value?.useClose();
      emit('on-change');
    });
  } catch (error) {
    console.error(error);
  }
};
</script>