<template>
  <PageLayout title="小说订单">
    <Card size="small">
      <div class="top-header">
        <div class="back" @click="goBack">
          <LeftOutlined/>
          <span class="txt">返回数据面板</span>
        </div>
      </div>
      <div class="picker-box">
        <RangePicker
            v-model:value="registerTime"
            :ranges="ranges"
            format="YYYY-MM-DD"
            @change="initChart('detail-analysis-content')"
        />
      </div>
      <div class="detail-data-box">
        <div class="detail-line">
          <div :class="idx == 2? 'data-content1' :'data-content'" v-for="(dt, idx) in data1" :key="idx">
                        <span class='data-tit'>
                        {{ dt.tit }}
                        </span>
            <span class="data-val">
                        {{ dt.val }}
                        </span>
          </div>
        </div>
      </div>
      <div class="detail-data-box">
        <div class="detail-line">
          <div :class="idx == 3? 'data-content2' :'data-content'" v-for="(dt, idx) in data" :key="idx">
                        <span class="data-tit">
                        {{ dt.tit }}
                        </span>
            <span class="data-val">
                        {{ dt.val }}
                        </span>
            <span class="data-count">
                        拉新{{ dt.hasOwnProperty('new_order_count') ? dt.new_order_count : dt.new_order_amount }}
                        </span>
            <span class="data-count">
                        拉失活{{
                dt.hasOwnProperty('inactivation_order_count') ? dt.inactivation_order_count : dt.inactivation_order_amount
              }}
                        </span>
          </div>
        </div>
      </div>
      <div class="analysis-content" id="detail-analysis-content"></div>
    </Card>
  </PageLayout>
</template>

<script lang="ts" setup>
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import {Card, RangePicker} from 'ant-design-vue';
import {LeftOutlined} from '@ant-design/icons-vue';
import {ref, onMounted} from 'vue';
import * as echarts from "echarts";

const registerTime = ref<any>([])
const ranges = ref<any>()
import {useRoute, useRouter} from 'vue-router';
import dayjs, {Dayjs} from "dayjs";
import {getNovelOrderInfoApi} from "/@/api/core/databoard.ts";

const aRouter = useRouter()
const data1 = ref<any>([
  {
    tit: '小说订单',
    val: 89,
    key: 'novel_order_count',
    price: false
  },
  {
    tit: '拉新订单',
    val: 89,
    key: 'new_order_count',
    price: false

  },
  {
    tit: '拉失活订单',
    val: 89,
    key: 'inactivation_order_count',
    price: false

  },
  {
    tit: '小说用户佣金',
    val: 89,
    key: 'novel_order_amount',
    price: true
  },
  {
    tit: '拉新用户佣金',
    val: 89,
    key: 'new_order_amount',
    price: true
  },
  {
    tit: '拉失活用户佣金',
    val: 89,
    key: 'inactivation_order_amount',
    price: true
  },
])
const data = ref<any>([
  {
    tit: '抖音故事订单',
    val: 89,
    new_order_count: 0,
    inactivation_order_count: 0,
    key: 'dygs_order_info'
  },
  {
    tit: '番茄小说订单',
    val: 89,
    new_order_count: 0,
    inactivation_order_count: 0,
    key: 'fqxs_order_info'
  },
  {
    tit: '番茄畅听订单',
    val: 89,
    new_order_count: 0,
    inactivation_order_count: 0,
    key: 'fqct_order_info'
  },
  {
    tit: '今日头条订单',
    val: 89,
    new_order_count: 0,
    inactivation_order_count: 0,
    key: 'jrtt_order_info'
  },
  {
    tit: '抖音故事佣金',
    val: 89,
    new_order_amount: 0,
    inactivation_order_amount: 0,
    key: 'dygs_amount_info'
  },
  {
    tit: '番茄小说佣金',
    val: 89,
    new_order_amount: 0,
    inactivation_order_amount: 0,
    key: 'fqxs_amount_info'
  },
  {
    tit: '番茄畅听佣金',
    val: 89,
    new_order_amount: 0,
    inactivation_order_amount: 0,
    key: 'fqct_amount_info'
  },
  {
    tit: '今日头条佣金',
    val: 89,
    new_order_amount: 0,
    inactivation_order_amount: 0,
    key: 'jrtt_amount_info'
  },
])

const options = {
  tooltip: {
    trigger: 'axis',
  },
  legend: {
    data: ['小说订单', '小说用户佣金'],
    bottom: 0
  },
  grid: {
    left: '3%',
    right: '4%',
    containLabel: true
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: []
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      name: '小说订单',
      type: 'line',
      data: [],
      lineStyle: {
        color: 'red'
      },
      itemStyle: {
        color: 'red'
      }
    },
  ]
}
onMounted(async () => {
  getRange()
  getLastDay()
  await initChart('detail-analysis-content')
})

const formatData = (res: any) => {
  Object.keys(res).forEach((d: any) => {
    for (let i = 0; i < data1.value.length; i++) {
      if (data1.value[i].key == d) {
        if (data1.value[i].price) {
          data1.value[i].val = (res[d] / 100.0).toFixed(2)
        } else {
          data1.value[i].val = res[d]
        }
        break;
      }
    }
  })
  Object.keys(res).forEach((d: any) => {
    for (let i = 0; i < data.value.length; i++) {
      if (data.value[i].key == d) {
        if(res[d]){
          data.value[i].val = res[d].hasOwnProperty('count') ? res[d]['count'] : (res[d]['amount']/100.0).toFixed()
          if(data.value[i].hasOwnProperty('new_order_count')){
            data.value[i].new_order_count =  res[d]['new_order_count']
            data.value[i].inactivation_order_count = res[d]['inactivation_order_count']
          }else{
            data.value[i].new_order_amount =  (res[d]['new_order_amount']/100.0).toFixed()
            data.value[i].inactivation_order_amount = (res[d]['inactivation_order_amount']/100.0).toFixed()
          }
        }
        break;
      }
    }
  })
}

const formatOptions = (data: any) => {
  if(!data) return
  options['xAxis']['data'] = data.map((it: any) => it.date)
  options.series = [
    {
      name: '小说订单',
      type: 'line',
      data: data.map((it: any) => it.count),
      lineStyle: {
        color: 'red'
      },
      itemStyle: {
        color: 'red'
      }
    },
    {
      name: '小说用户佣金',
      type: 'line',
      data: data.map((it: any) => (it.amount / 100.0).toFixed(2)),
      lineStyle: {
        color: 'blue'
      },
      itemStyle: {
        color: 'blue'
      }
    }
  ]
}
const initChart = async (id: string) => {
  let myEchart = echarts.init(document.getElementById(id));
  const res = await getNovelOrderInfoApi({
    start: registerTime.value.length >= 1 ? registerTime.value[0].format('YYYY-MM-DD') : "",
    end: registerTime.value.length >= 2 ? registerTime.value[1].format('YYYY-MM-DD') : ""
  })
  formatData(res)
  formatOptions(res.data)
  myEchart.setOption(options);
  window.onresize = function () {
    myEchart.resize();
  };
}

const formatDate = (date: any) => {
  let newDate = date && ref<Dayjs>(dayjs(date, "YYYY-MM-DD"))
  return newDate.value
}

const getRange = () => {
  ranges.value = {
    '昨天': [formatDate(getDay(-1)), formatDate(getDay(-1))],
    '近7天': [formatDate(getDay(-7)), formatDate(getDay(0))],
    '近30天': [formatDate(getDay(-30)), formatDate(getDay(0))],
  }
}
const route = useRoute();
const getLastDay = () => {
  if(route.query.start){
    registerTime.value = [formatDate(route.query.start), formatDate(route.query.end)]
  }else{
    registerTime.value = [formatDate(getDay(-7)), formatDate(getDay(0))]
  }
}

const getDay = (day: any) => {
  var today = new Date()
  var targetday_milliseconds = today.getTime() + 1000 * 60 * 60 * 24 * day
  today.setTime(targetday_milliseconds);
  var tYear = today.getFullYear()
  var tMonth = today.getMonth()
  var tDate = today.getDate()
  tMonth = doHandleMonth(tMonth + 1)
  tDate = doHandleMonth(tDate)
  return tYear + "-" + tMonth + "-" + tDate
}
const doHandleMonth = (month: any) => {
  var m = month;
  if (month.toString().length == 1) {
    m = "0" + month;
  }
  return m;
}
const goBack = () => {
  aRouter.back()
}
</script>

<style scoped>
.top-header .txt {
  margin-left: 5px;
}

.top-header .back {
  font-size: 16px;
  cursor: pointer;
}

.picker-box {
  margin-top: 15px;
}

.analysis-content {
  width: 100%;
  height: 350px;
  margin-top: 20px;
}

.detail-data-box {
  margin-top: 20px;
}

.detail-line {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
}

.data-content {
  display: flex;
  flex-direction: column;
  width: 100px;
  margin-right: 80px;
}

.data-content1 {
  display: flex;
  flex-direction: column;
  margin-right: 315px;
}

.data-content2{
  display: flex;
  flex-direction: column;
  margin-right: 119px;
}

.data-tit {
  font-weight: 600;
  opacity: 0.6;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}


.data-count {
  font-size: 13px;
  opacity: 0.5;
}

</style>