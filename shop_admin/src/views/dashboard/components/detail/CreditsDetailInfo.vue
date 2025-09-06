<template>
  <PageLayout title="算力订单">
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
          <div class="data-content" v-for="dt in data">
                        <span class="data-tit">
                        {{ dt.tit }}
                        </span>
            <span class="data-val">
                        {{ dt.val }}
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
import {getCreditOrderInfoApi} from "/@/api/core/databoard.ts";


const aRouter = useRouter()
const data = ref<any>([
  {
    tit: '算力订单',
    val: 89,
  },
  {
    tit: '算力销售额',
    val: 89,
  },
  {
    tit: '首购算力用户数',
    val: 89,
  },
  {
    tit: '复购算力用户',
    val: 89,
  },
])

const options = {
  tooltip: {
    trigger: 'axis'
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
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
      name: '算力订单数',
      type: 'line',
      data: [11,12,13,14,15,16,17]
    }
  ]
}
onMounted(async () => {
  getRange()
  getLastDay()
  await initChart('detail-analysis-content')
})

const formatData = (res: any) => {
  let idx = 0
  console.log(res, "res")
  Object.keys(res).forEach((d: any) => {
    if(Array.isArray(res[d])){
      return
    }
    if(data.value[idx].tit == "算力销售额"){
      data.value[idx].val = (res[d]/100.0).toFixed(2)
    }else{
      data.value[idx].val = res[d]
    }
    idx ++
  })
}

const formatOptions = (data : any): any => {
  options['xAxis']['data'] = data.map((it: any) => it.date)
  options.series = [
    {
      name: '算力订单数',
      type: 'line',
      data: data.map((it: any) => it.count)
    }
  ]
}
const initChart = async (id: string) => {
  let myEchart = echarts.init(document.getElementById(id));
  const res = await getCreditOrderInfoApi({
    start: registerTime.value.length >= 1 ? registerTime.value[0].format('YYYY-MM-DD') : "",
    end:  registerTime.value.length >= 2 ? registerTime.value[1].format('YYYY-MM-DD') : ""
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
const route = useRoute()
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
  height: 300px;
  margin-top: 20px;
}

.detail-data-box {
  margin-top: 20px;
}

.detail-line {
  width: 100%;
  display: flex;
}

.data-content {
  display: flex;
  flex-direction: column;
  margin-right: 80px;
}

.data-tit {
  font-weight: 600;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}

</style>