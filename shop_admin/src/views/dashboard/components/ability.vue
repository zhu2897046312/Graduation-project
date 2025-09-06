<template>
    <div class="content">
        <Card>
            <template #title>
                <div class="tit-content">
                    <span class="card-tit">功能使用统计</span>
                    <RangePicker
                      v-model:value="registerTime"
                      :ranges="ranges"
                      format="YYYY-MM-DD"
                      @change="handleChange"
                  />
                </div>
            </template>
            <div class="card-content">
              <div class="data-content" v-for="dt in data">
                <span class="data-tit">
                  {{ dt.tit }}
                </span>
                <span class="data-val">
                  {{ dt.val }}
                </span>
                <span class="action" @click="toDetail(dt.key)">
                    查看详情 >
                </span>
              </div>
            </div>
        </Card>
    </div>
  </template>

<script lang="ts" setup>
  import {ref, onMounted} from 'vue'
  import {Card, RangePicker} from 'ant-design-vue'
  import dayjs , { Dayjs } from 'dayjs';
  import { useRouter } from 'vue-router';
  import {getFunctionUsageApi} from "/@/api/core/databoard.ts";
  const router = useRouter()
  const registerTime = ref<any>([])
  const ranges = ref<any>()
  const data = ref<any>([
    {
      tit: 'AI动画',
      val: 89,
      key: 'videoProjectDetail'
    },
    {
      tit: 'AI动画二创',
      val: 89,
      key: 'videoProjectDetail'
    },
    {
      tit: '智能混剪',
      val: 89,
      key: 'MixInfo'
    },
    {
      tit: '智能混剪二创',
      val: 89,
      key: 'MixInfo'
    },
    {
      tit: '文章提取',
      val: 89,
      key: 'AuxiliaryInfo'
    },
    {
      tit: '文章改编',
      val: 89,
      key: 'AuxiliaryInfo'
    },
    {
      tit: '配音生成',
      val: 89,
      key: 'AuxiliaryInfo'
    },
    {
      tit: 'AI绘画',
      val: 89,
      key: 'AppAiInfo'
    },
    {
      tit: '真人转动漫',
      val: 89,
      key: 'AppAiInfo'
    },
    {
      tit: '别名申请',
      val: 89,
      key: 'PopularizeInfo'
    },
    {
      tit: '挂载申请',
      val: 89,
      key: 'PopularizeInfo'
    },
  ])

  const formatDate = (date : any) => {
  let newDate = date && ref<Dayjs>(dayjs(date, "YYYY-MM-DD"))
  return newDate.value
}
  const handleChange = () => {
    handleLoadData()
  }
onMounted(() => {
  getRange()
  getLastDay()
  handleLoadData()
})

const handleLoadData = async () => {
  const payload = {
    start: registerTime.value.length >= 1 ? registerTime.value[0].format('YYYY-MM-DD') : "",
    end:  registerTime.value.length >= 2 ? registerTime.value[1].format('YYYY-MM-DD') : ""
  }
  const res = await getFunctionUsageApi(payload)
  let idx = 0
  Object.keys(res).forEach((key: any) => {
    if(data.value[idx]){
      data.value[idx].val = res[key]
    }
    idx ++
  })
}


const toDetail = (key: any) => {
      router.push({name: key})
}

const getRange = () => {
    ranges.value = {
    '昨天': [formatDate(getDay(-1)), formatDate(getDay(-1))],
    '近7天': [formatDate(getDay(-7)), formatDate(getDay(0))],
    '近30天': [formatDate(getDay(-30)), formatDate(getDay(0))],
    }
}

const getLastDay = () => {
    let day = getDay(-1)
    registerTime.value = [formatDate(day), formatDate(day)]
}

const getDay = (day: any)  => {
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
.content{
  margin-bottom: 10px;
}
.card-tit{
  font-weight: 600;
  font-size: 18px;
  margin-right: 20px;
}
.card-content{
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.data-content{
  display: flex;
  flex-direction: column;
}
.data-tit{
  font-weight: 600;
}
.data-val{
  font-weight: 500;
  font-size: 30px;
}
.action{
    opacity: 0.5;
    cursor: pointer;
}
</style>