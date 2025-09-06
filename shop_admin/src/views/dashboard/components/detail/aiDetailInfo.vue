<template>
  <PageLayout title="AI动画数据">
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
          <div :class="idx == 2? 'data-content1' :'data-content'" v-for="(dt, idx) in data" :key="idx">
                        <span class="data-tit">
                        {{ dt.tit }}
                        </span>
            <span class="data-val">
                        {{ dt.val }}
                        </span>
          </div>
        </div>
        <div class="detail-data-box">
          <div class="detail-line">
            <div :class="idx == 1? 'data-content2' :'data-content'" v-for="(dt, idx) in data1" :key="idx">
                        <span class="data-tit">
                        {{ dt.tit }}
                        </span>
              <span class="data-val">
                        {{ dt.val }}
                        </span>
            </div>
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
import {useRouter} from 'vue-router';
import dayjs, {Dayjs} from "dayjs";
import {getVideoProjectInfoApi} from "/@/api/core/databoard.ts";

const aRouter = useRouter()
const data = ref<any>([
  {
    tit: 'AI动画生成',
    val: 89,
    key: 'video_project_count'
  },
  {
    tit: '生成用户',
    val: 89,
    key: 'video_project_user_count'
  },
  {
    tit: '未完成用户',
    val: 89,
    key: 'video_project_no_complete_user_count'
  },
  {
    tit: 'AI动画二创生成',
    val: 89,
    key: 'derivative_work_video_count'
  },
  {
    tit: '生成用户',
    val: 89,
    key: 'derivative_work_video_complete_count'
  },
  {
    tit: '未完成用户',
    val: 89,
    key: 'derivative_work_video_no_complete_user_count'
  },
])
const data1 = ref<any>([
  {
    tit: '未完成',
    val: 89,
    key: 'video_project_no_complete_count'
  },
  {
    tit: '完成生图',
    val: 89,
    key: 'video_project_complete_count'
  },
  {
    tit: '未完成',
    val: 89,
    key: 'derivative_work_video_no_complete_count'
  },
  {
    tit: '完成生图',
    val: 89,
    key: 'derivative_work_video_count'
  },
])
const options = {
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['AI动画', 'AI动画二创'],
    bottom: 0
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '6%',
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
  series : [
    {
      name: 'AI动画',
      type: 'line',
      data: [],
      lineStyle: {
        color: 'red' // 这里将折线颜色设置为红色
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
    for (let i = 0; i < data.value.length; i++) {
      if (data.value[i].key == d) {
        data.value[i].val = res[d]
        break;
      }
    }
  })
  Object.keys(res).forEach((d: any) => {
    for (let i = 0; i < data1.value.length; i++) {
      if (data1.value[i].key == d) {
        data1.value[i].val = res[d]
        break;
      }
    }
  })
}

const formatOptions = (data: any) => {
  options['xAxis']['data'] = data.map((it: any) => it.date)
  options.series = [
    {
      name: 'AI动画',
      type: 'line',
      data: data.map((it: any) => it.video_project_count),
      lineStyle: {
        color: 'red' // 这里将折线颜色设置为红色
      },
      itemStyle: {
        color: 'red'
      }
    },
    {
      name: 'AI动画二创',
      type: 'line',
      data: data.map((it: any) => it.derivative_work_video_count),
      lineStyle: {
        color: 'blue' // 这里将折线颜色设置为红色
      },
      itemStyle: {
        color: 'blue'
      }
    }
  ]
}
const initChart = async (id: string) => {
  let myEchart = echarts.init(document.getElementById(id));
  const res = await getVideoProjectInfoApi({
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

const getLastDay = () => {
  registerTime.value = [formatDate(getDay(-7)), formatDate(getDay(0))]
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
  height: 380px;
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
  width: 100px;
  display: flex;
  flex-direction: column;
  margin-right: 80px;
}

.data-tit {
  opacity: 0.6;
  font-weight: 600;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}

.data-content1 {
  width: 100px;
  display: flex;
  flex-direction: column;
  margin-right: 270px;
}

.data-content2 {
  width: 100px;
  display: flex;
  flex-direction: column;
  margin-right: 448px;
}
</style>