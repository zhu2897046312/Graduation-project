<template>
  <BasicModal :width="1200" title="订单详情" @on-close="handleClose" ref="modalRef" ok-text="" cancel-text="关闭" :mask-closable="true">
    <div class="py-2">
      <Descriptions title="订单基础信息" bordered :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="订单编号">{{ form.code }}</Descriptions.Item>
        <Descriptions.Item label="游客查询码">{{ form.visitor_query_code }}</Descriptions.Item>
        <Descriptions.Item label="下单时间">{{ form.created_time }}</Descriptions.Item>
        <Descriptions.Item label="支付金额">{{ form.pay_amount }}</Descriptions.Item>
        <Descriptions.Item label="用户昵称">{{ form.nickname }}</Descriptions.Item>
        <Descriptions.Item label="邮箱地址">{{ form.email }}</Descriptions.Item>
        <Descriptions.Item label="订单状态"><EnumLabel :dict-value="form.state" dict-code="SpOrderConstant$State" /></Descriptions.Item>
      </Descriptions>
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { Descriptions } from 'ant-design-vue';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'
import api from '/@/api/'

const modalRef = ref<any>(null);
const form = ref<any>(null);

const handleOpen = (e: any) => {
  loadingTask(async () => {
    form.value = await api.shop.order.infoByOrderCode({ code: e.orderCode })
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  form.value = null
}

defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})
</script>