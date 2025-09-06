<template>
  <BasicModal :width="1200" title="订单详情" @on-close="handleClose" ref="modalRef" ok-text="" cancel-text="关闭"
    :mask-closable="true">
    <div class="py-2">
      <Descriptions title="订单基础信息" bordered :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="订单编号">{{ form.order.code }}</Descriptions.Item>
        <Descriptions.Item label="下单时间">{{ form.order.created_time }}</Descriptions.Item>
        <Descriptions.Item label="支付金额">{{ form.order.pay_amount }}</Descriptions.Item>
        <Descriptions.Item label="用户昵称">{{ form.order.nickname }}</Descriptions.Item>
        <Descriptions.Item label="邮箱地址">{{ form.order.email }}</Descriptions.Item>
        <Descriptions.Item label="订单状态">
          <EnumLabel :dict-value="form.order.state" dict-code="SpOrderConstant$State" />
        </Descriptions.Item>
      </Descriptions>
    </div>
    <div class="py-2">
      <Descriptions title="商品信息" bordered :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="商品列表" :span="3">
          <Table size="small" :columns="columns" :data-source="form.items" bordered :pagination="false"></Table>
        </Descriptions.Item>
      </Descriptions>
    </div>
    <div class="py-2">
      <Descriptions title="物流信息" bordered :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="收件地址" :span="3">
          国家：{{ form.receive_address.country }}<br>
          省：{{ form.receive_address.province }}<br>
          市：{{ form.receive_address.city }}<br>
          区：{{ form.receive_address.region }}<br>
          详细：{{ form.receive_address.detail_address }}<br>
          电话：{{ form.receive_address.phone }}<br>
          昵称：{{ form.receive_address.first_name }}{{ form.receive_address.last_name }}<br>
        </Descriptions.Item>
        <Descriptions.Item label="快递公司" v-if="form.order.delivery_company">{{ form.order.delivery_company }}
        </Descriptions.Item>
        <Descriptions.Item label="快递单号" v-if="form.order.delivery_company">{{ form.order.delivery_sn }}
        </Descriptions.Item>
        <Descriptions.Item label="发货时间" v-if="form.order.delivery_company">{{ form.order.delivery_time }}
        </Descriptions.Item>

      </Descriptions>
    </div>
    <div class="py-2">
      <Descriptions title="其他信息" bordered :labelStyle="{ width: '110px' }">
        <Descriptions.Item label="备注" :span="3" v-if="form.order.remark">{{ form.order.remark }}</Descriptions.Item>
      </Descriptions>
    </div>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref, h } from 'vue';
import { Descriptions, Table, Image } from 'ant-design-vue';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import EnumLabel from '/@/components/Kernel/EnumLabel/index.vue'
import api from '/@/api/index'


const modalRef = ref<any>(null);
const form = ref<any>(null);

const columns = ref<any[]>([
  {
    dataIndex: 'thumb',
    title: '封面图',
    width: 80,
    customRender: (e: any) => {
      return h(Image, {
        width: 60,
        height: 30,
        src: e.record.thumb,
        fallback: "data:image/png;base64,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"
      })
    }
  },
  {
    dataIndex: 'name', 
    title: '商品名称',
    customRender: (e: any) => {
      const productId = e.record.id; 
      const productDetailUrl = `https://www.earring18.com/product/${productId}`; 
      return h('a', {
        href: productDetailUrl,
        target: '_blank'
      },
        e.record.title
      )
    }
  },
  { title: '规格', dataIndex: 'sku_title', },
  { title: '成本价', dataIndex: 'cost_price', width: 100 },
  { title: '单价', dataIndex: 'price', width: 100 },
  { title: '数量', dataIndex: 'quantity', width: 80 },
  { title: '总价', dataIndex: 'pay_amount', width: 100 },
])


let id = 0
const handleOpen = (e: any) => {
  loadingTask(async () => {
    form.value = await api.shop.order.info(e.id)
    id = e.id;
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