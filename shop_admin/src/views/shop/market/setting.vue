<template>
  <PageLayout>
    <Card title="商店设置">
      <Form>
        <Form.Item label="汇率">
          <InputNumber style="width: 230px" v-model:value="info.exchange" placeholder="请输入汇率" :precision="5" />
        </Form.Item>
        <Form.Item label="运费">
          <InputNumber style="width: 230px" v-model:value="info.freight" placeholder="请输入运费" :precision="2" />
        </Form.Item>
        <!-- <Form.Item label="原价比例" help="原价=销售价*原价比例">
          <InputNumber style="width: 230px" v-model:value="info.original" placeholder="原价比例" :precision="2" />
        </Form.Item> -->
        <Form.Item >
          <Button type="primary" @click="onSubmit">保存</Button>
        </Form.Item>
      </Form>
    </Card>
  </PageLayout>
</template>

<script lang="ts" setup>
  import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
  import { Card, Form, InputNumber, Button, Modal, message } from 'ant-design-vue';
  import { ref, onMounted } from 'vue';
  import api  from '/@/api/index';

  const info = ref<any>({
    freight: '0.00',
    exchange: '0.00',
    original: '1.00'
  })

  const onSubmit = () => {
    Modal.confirm({
      title: '请确认是否修改？',
      content: '修改后，将会设置所有商品（除锁定价格商品外）的销售价格（成本价*汇率）',
      okText: '确认修改',
      cancelText: '取消',
      onOk: () => {
        handleSubmit()
      },
      onCancel: () => { }
    })
  }

  const handleSubmit = async () => {
    const hide = message.loading('正在保存中', 0)
    await api.shop.market.save(info.value)
    setTimeout(() => {
      hide()
      message.success('保存成功')
    }, 800)
  }


  onMounted(async () => {
    info.value = await api.shop.market.info()
  })


</script>