<template>
  <PageLayout title="推广数据">
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
          <div :class="idx == 1 ?'data-content1' :  idx == 3? 'data-content3':'data-content'" v-for="(dt, idx) in data" :key="idx">
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
            <div :class="idx == 3 || idx == 8? 'data-content2' :'data-content'" v-for="(dt, idx) in data1" :key="idx">
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
import {getPopularizeInfoApi} from "/@/api/core/databoard.ts";

const aRouter = useRouter()
const data = ref<any>([
  {
    tit: '挂载UID申请',
    val: 89,
    key: 'novel_permission_count'
  },
  {
    tit: '挂载申请用户',
    val: 89,
    key: 'novel_permission_user_count'
  },
  {
    tit: '别名申请',
    val: 89,
    key: 'command_count'
  },
  {
    tit: '别名申请用户',
    val: 89,
    key: 'command_user_count'
  },
  {
    tit: '别名回填',
    val: 89,
    key: 'command_back_fill_count'
  },
])
const data1 = ref<any>([
  {
    tit: '审核成功',
    val: 89,
    key: 'novel_permission_success_count'
  },
  {
    tit: '审核失败',
    val: 89,
    key: 'novel_permission_fail_count'
  },
  {
    tit: '待审核',
    val: 89,
    key: 'novel_permission_wait_count'
  },
  {
    tit: '审核中',
    val: 89,
    key: 'novel_permission_doing_count'
  },
  {
    tit: '审核成功',
    val: 89,
    key: 'command_success_count'
  },
  {
    tit: '审核失败',
    val: 89,
    key: 'command_fail_count'
  },
  {
    tit: '待审核',
    val: 89,
    key: 'command_wait_count'
  },
  {
    tit: '审核中',
    val: 89,
    key: 'command_doing_count'
  },
  {
    tit: '审核成功',
    val: 89,
    key: 'command_back_fill_success_count'
  },
  {
    tit: '审核失败',
    val: 89,
    key: 'command_back_fill_fail_count'
  },
])
const options = {
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['挂载申请', '别名申请', '别名回填'],
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
  series: [
    {
      name: '别名回填',
      type: 'line',
      data: [],
      lineStyle: {
        color: 'green' // 这里将折线颜色设置为红色
      },
      itemStyle: {
        color: 'green'
      }
    }
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
      name: '挂载申请',
      type: 'line',
      data: data.map((it: any) => it.novel_permission_count),
      lineStyle: {
        color: 'red'
      },
      itemStyle: {
        color: 'red'
      }
    },
    {
      name: '别名申请',
      type: 'line',
      data: data.map((it: any) => it.command_count),
      lineStyle: {
        color: 'blue'
      },
      itemStyle: {
        color: 'blue'
      }
    },
    {
      name: '别名回填',
      type: 'line',
      data: data.map((it: any) => it.command_back_fill_count),
      lineStyle: {
        color: 'green' // 这里将折线颜色设置为红色
      },
      itemStyle: {
        color: 'green'
      }
    }
  ]
}
const initChart = async (id: string) => {
  let myEchart = echarts.init(document.getElementById(id));
  const res = await getPopularizeInfoApi({
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


.data-tit {
  opacity: 0.6;
  font-weight: 600;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}


.data-content {
  width: 120px;
  display: flex;
  flex-direction: column;
  margin-right: 50px;
}

.data-content1 {
  width: 120px;
  display: flex;
  flex-direction: column;
  margin-right: 360px;
}

.data-content2 {
  width: 120px;
  display: flex;
  flex-direction: column;
  margin-right: 20px;
}
.data-content3 {
  width: 130px;
  display: flex;
  flex-direction: column;
  margin-right: 380px;
}
</style>