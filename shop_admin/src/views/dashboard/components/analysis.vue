<template>
  <div class="content">
    <Card>
      <template #title>
        <div class="tit-content">
          <span class="card-tit">趋势分析</span>
          <RangePicker
              v-model:value="registerTime"
              :ranges="ranges"
              format="YYYY-MM-DD"
              @change="initChart('analysis-content')"
          />
        </div>
      </template>
      <div class="card-content">
        <div class="select-content">
          <span style="margin-right: 30px;">数据项</span>
          <Select v-model:value="select_type[0]" @change="handleChange" style="width: 130px;margin-right: 10px;">
            <Select.Option v-for="(item, index) in select_options" :key="index" :value="item.val">
              {{ item.show_title }}
            </Select.Option>
          </Select>
          <Select v-model:value="select_type[1]" @change="handleChange" style="width: 130px;margin-right: 10px;">
            <Select.Option v-for="(item, index) in select_options" :key="index" :value="item.val">
              {{ item.show_title }}
            </Select.Option>
          </Select>
          <Select v-model:value="select_type[2]" @change="handleChange" style="width: 130px;">
            <Select.Option v-for="(item, index) in select_options" :key="index" :value="item.val">
              {{ item.show_title }}
            </Select.Option>
          </Select>
        </div>
        <div class="analysis-content" id="analysis-content"></div>
      </div>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted} from 'vue'
import {Card, RangePicker, Select} from 'ant-design-vue'
import dayjs, {Dayjs} from 'dayjs';
import * as echarts from "echarts";
import {getTrendAnalysisApi} from "/@/api/core/databoard.ts";

const registerTime = ref<any>([])
const ranges = ref<any>()
const select_type = ref<any[]>([0, 1, 2])
const data = ref<any[]>(["注册用户", "算力订单", "算力转化率"])
const analysisData = ref<any>([])

const handleLoadData = async () => {
  const payload = {
    start: registerTime.value.length >= 1 ? registerTime.value[0].format('YYYY-MM-DD') : "",
    end: registerTime.value.length >= 2 ? registerTime.value[1].format('YYYY-MM-DD') : ""
  }
  analysisData.value = await getTrendAnalysisApi(payload)
}

const select_options: any[] = [
  {show_title: '注册用户', val: 0, code: 'reg_user_count'},
  {show_title: '算力订单', val: 1, code: 'credit_order_count'},
  {show_title: '算力转化率', val: 2, code: "credit_order_rate"},
  {show_title: '领算力用户', val: 3, code: 'new_comer_reward_user_count'},
  {show_title: '算力销售额', val: 4, code: 'credit_order_price'},
  {show_title: '购买算力用户', val: 5, code: 'buy_credit_user_count'},
  {show_title: '小说订单', val: 6, code: 'novel_order_count'},
  {show_title: '小说订单佣金', val: 7, code: 'novel_order_price'},
  {show_title: '别名申请', val: 8, code: 'command_count'},
  {show_title: '挂载申请', val: 9, code: 'novel_permission_count'},
  {show_title: 'AI动画', val: 10, code: 'video_project_count'},
  {show_title: 'AI动画二创', val: 11, code: 'derivative_work_video_project_count'},
  {show_title: '智能混剪', val: 12, code: 'alone_mix_task_count'},
  {show_title: '智能混剪二创', val: 13, code: 'alone_derivative_mix_task_count'},
  {show_title: '文章提取', val: 14, code: 'reptile_task_count'},
  {show_title: '文章改编', val: 15, code: 'adapt_task_count'},
  {show_title: '配音生成', val: 16, code: 'audio_task_count'},
  {show_title: '视频提取', val: 17, code: 'video_extract_task_count'},
  {show_title: '音频提取', val: 18, code: 'alone_audio_extract_task_count'},
  {show_title: 'AI绘画', val: 19, code: 'app_text_to_image_count'},
  {show_title: '真人转漫画', val: 20, code: 'app_image_to_image_count'},
]

const handleChange = () => {
  select_type.value.forEach(((it: any, idx: any) => {
    data.value[idx] = select_options.filter((it1: any) => it1.val === it)[0].show_title
  }))
  initChart1("analysis-content")
}
const options = {
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: data.value
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
      name: "",
      type: 'line',
      data: []
    },
    {
      name: "",
      type: 'line',
      data: []
    },
    {
      name: "",
      type: 'line',
      data: []
    }
  ]
}

const initChart = async (id: string) => {
  await handleLoadData()
  initChart1(id)
}
const initChart1 = (id: string) => {
  let myEchart = echarts.init(document.getElementById(id));
  const options = formatOption()
  myEchart.setOption(options);
  console.log(options)
  window.onresize = function () {
    myEchart.resize();
  };
}
const formatOption: any = () => {
  options['legend'] = {data: data.value}
  options['xAxis']['data'] = analysisData.value.map((it: any) => it.date)
  const arr = getSeries_data()
  options['series'] = data.value.map((it: any, idx: any) => {
    return {
      name: it,
      type: 'line',
      data: arr[idx]
    }

  })
  return options
}
const getSeries_data = (): any[] => {
  let arr: any[] = []
  select_type.value.forEach((type: any) => {
    for (let i = 0; i < select_options.length; i++) {
      if (select_options[i].val === type) {
        arr.push(analysisData.value.map((d: any) => {
          if(select_options[i].code == "credit_order_price" || select_options[i].code == "novel_order_price"){
            return (d[select_options[i].code]/100.0).toFixed(2)
          }else if(select_options[i].code == "credit_order_rate"){
            return d[select_options[i].code].toFixed(2)
          }
          return d[select_options[i].code]
        }))
        break
      }
    }
  })
  return arr
}


onMounted(async () => {
  getRange()
  getDefultDay()
  await initChart('analysis-content')
})

const getRange = () => {
  ranges.value = {
    '昨天': [formatDate(getDay(-1)), formatDate(getDay(-1))],
    '近7天': [formatDate(getDay(-7)), formatDate(getDay(0))],
    '近30天': [formatDate(getDay(-30)), formatDate(getDay(0))],
  }
}

const formatDate = (date: any) => {
  let newDate = date && ref<Dayjs>(dayjs(date, "YYYY-MM-DD"))
  return newDate.value
}

const getDefultDay = () => {
  let day = getDay(-30)
  let today = getDay(0)
  registerTime.value = [formatDate(day), formatDate(today)]
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

const doHandleMonth = (month) => {
  var m = month;
  if (month.toString().length == 1) {
    m = "0" + month;
  }
  return m;
}
</script>

<style scoped>
.content {
  margin-bottom: 10px;
}

.card-tit {
  font-weight: 600;
  font-size: 18px;
  margin-right: 20px;
}

.card-content {
  width: 100%;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
}

.analysis-content {
  width: 100%;
  height: 300px;
}

.data-content {
  display: flex;
  flex-direction: column;
}

.data-tit {
  font-weight: 600;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}
</style>