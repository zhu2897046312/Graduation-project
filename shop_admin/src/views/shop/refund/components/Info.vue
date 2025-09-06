<template>
  <BasicModal :width="800" title="退款详情" @on-close="handleClose" ref="modalRef" ok-text="" cancel-text="关闭" :mask-closable="true">
    <div class="py-2">
      <Descriptions title="退款信息" bordered :column="3" :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="退款单号" :span="1">{{ form?.refund_no || '-' }}</Descriptions.Item>
        <Descriptions.Item label="退款金额" :span="1">{{ form?.refund_amount || '-' }}</Descriptions.Item>
        <Descriptions.Item label="退款状态" :span="1">
          <Tag :color="form?.status === 2 ? 'success' : ''">
            {{ form?.status === 2 ? '处理中' : form?.status === 4 ? '已拒绝' : '已退款' }}
          </Tag>
        </Descriptions.Item>
        <Descriptions.Item label="申请时间" :span="1">{{ form?.created_time || '-' }}</Descriptions.Item>
        <Descriptions.Item label="退款时间" :span="2">{{ form?.refund_time || '-' }}</Descriptions.Item>
      </Descriptions>
      
      <Descriptions bordered :column="1" style="margin-top: 16px;">
        <Descriptions.Item label="退款原因">{{ form?.reason || '无' }}</Descriptions.Item>
      </Descriptions>
      
      <Descriptions bordered :column="1" style="margin-top: 16px;">
        <Descriptions.Item label="凭证图片">
          <template v-if="form?.images?.length > 0">
            <div class="image-container">
              <div 
                v-for="(url, index) in form.images"
                :key="index"
                class="image-wrapper"
              >
                <Image
                  :src="url"
                  :preview="{ src: url }"
                  class="credential-image"
                />
              </div>
            </div>
          </template>
          <template v-else>
            <div class="no-image">暂无图片</div>
          </template>
        </Descriptions.Item>
      </Descriptions>
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { Descriptions, Tag, Image } from 'ant-design-vue';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import api from '/@/api/shop/refund'

const modalRef = ref<any>(null);
const form = ref<any>({
  images: [] // 初始化时确保images是数组
});

const handleOpen = (e: any) => {
  loadingTask(async () => {
    form.value = await api.info(e.id)
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  form.value = {
    images: [] // 重置时保持结构一致
  }
}

defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})
</script>

<style scoped>
.image-container {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.image-wrapper {
  width: 120px;
  height: 120px;
  border-radius: 4px;
  overflow: hidden;
  border: 1px solid #f0f0f0;
  background-color: #f5f5f5;
  flex-shrink: 0;
}

.credential-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  cursor: pointer;
  transition: all 0.3s;
}
.credential-image:hover {
  transform: scale(1.05);
}

.no-image {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 120px;
  color: #999;
  border: 1px dashed #d9d9d9;
  border-radius: 4px;
}
</style>