<template>
  <div class="content">
    <div v-if="prop.isMobile">
      <DataBoxMobile title="今日数据" :data="componentData.today" :range="componentData.today_range"/>
      <DataBoxMobile title="昨日数据" :data="componentData.yesterday" :range="componentData.yesterday_range"/>
      <DataBoxMobile title="本月数据" :data="componentData.month" :range="componentData.month_range"/>
      <DataBoxMobile title="上月数据" :data="componentData.last_month" :range="componentData.last_month_range"/>
    </div>
    <div v-else>
          <div style="background-color: #ececec;">
            <Row :gutter="24">
              <Col :span="12">
                  <DataBox title="今日数据" :data="componentData.today" :range="componentData.today_range"/>
              </Col>
              <Col :span="12">
                  <DataBox title="本月数据" :data="componentData.month" :range="componentData.month_range"/>
              </Col>
            </Row>
          </div>
          <div style="background-color: #ececec; padding: 2px">
            <Row :gutter="24">
              <Col :span="12">
                  <DataBox title="昨日数据" :data="componentData.yesterday" :range="componentData.yesterday_range"/>
              </Col>
              <Col :span="12">
                  <DataBox title="上月数据" :data="componentData.last_month" :range="componentData.last_month_range"/>
              </Col>
            </Row>
          </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue'
import {Col, Row} from 'ant-design-vue'
import DataBox from "/@/views/dashboard/components/DataBox.vue";
import {getKeyBusinessApi} from "/@/api/core/databoard.ts";
import DataBoxMobile from "/@/views/dashboard/components/DataBoxMobile.vue";

const prop = defineProps({
  isMobile: Boolean
})
const componentData = ref<any>({
  "today": [],
  "today_range": {},
  "yesterday": [],
  "yesterday_range": {},
  "month": [],
  "month_range": {},
  "last_month": [],
  "last_month_range": {}
})

onMounted(() => {
  handleLoadData()
})
const handleLoadData = async () => {
  const res= await getKeyBusinessApi()
  console.log(res, "res")
  pushData("today", res.today, "today_range")
  pushData("yesterday", res.yesterday, "yesterday_range")
  pushData("month", res.month, "month_range")
  pushData("last_month", res.last_month, "last_month_range")
  console.log(componentData.value)

}

const pushData = (index: string, val: any, range: string) => {
  const data = [
    {
      tit: '注册用户',
      val: 89,
      key: 'reg_user_count'
    },
    {
      tit: '算力订单',
      val: 89,
      key: 'credit_order_count',
      route: 'CreditsDetailInfo'
    },
    {
      tit: '算力转化率',
      val: 89,
      key: 'credit_order_rate'
    },
    {
      tit: '算力销售额',
      val: 89,
      key: 'credit_order_price',
      price: true
    },
    {
      tit: '小说订单',
      val: 89,
      key: 'reg_user_count',
      route: 'NovelDetailInfo'
    },
    {
      tit: '小说用户佣金',
      val: 89,
      key: 'novel_order_user_price',
      price: true
    },
    {
      tit: '小说平台佣金',
      val: 89,
      key: 'novel_order_platform_price',
      price: true
    },
  ]
  let idx = 0
  Object.keys(val).forEach((key: any) => {
    if(data[idx]){
      data[idx].val  = data[idx].price? (val[key]/100.0).toFixed(2) : data[idx].val = val[key]
      idx++
    }
  })
  componentData.value[index] = data
  componentData.value[range] = {start: val.start, end: val.end }
}
</script>

<style scoped>
.content {
  margin-bottom: 10px;
}

</style>