<template>
    <div class="content">
      <div v-if="prop.isMobile">
        <Card style="border-radius: 8px">
          <span class="mobile-card-tit">次日转换统计</span>
          <div class="mobile-card-content">
            <div class="mobile-data-content" v-for="dt in data">
                <span class="mobile-data-tit">
                  {{ dt.tit }}
                </span>
              <span class="mobile-data-val">
                  {{ dt.val }}
          </span>
            </div>
          </div>
        </Card>
      </div>
      <div v-else>
        <Card>
          <template #title>
            <div class="tit-content">
              <span class="card-tit">次日转换统计</span>
              <DatePicker
                  v-model:value="registerTime"
                  format="YYYY-MM-DD"
                  @change="handleChange"
              />
            </div>
          </template>
          <div class="card-content">
            <div class="data-content" v-for="dt in data">
                <span class="data-tit" style="opacity: 0.7;">
                  {{ dt.tit }}
                </span>
              <span class="data-val">
                  {{ dt.val }}
                </span>
            </div>
          </div>
        </Card>
      </div>
    </div>
  </template>

<script lang="ts" setup>
  import {ref, onMounted} from 'vue'
  import {Card, DatePicker} from 'ant-design-vue'
  import dayjs , { Dayjs } from 'dayjs';
  import {getRegUsageApi} from "/@/api/core/databoard.ts";

  const prop = defineProps({
    isMobile: Boolean
  })
  const registerTime = ref<Dayjs>();
  const data = ref<any>([
    {
      tit: '当天注册数',
      val: '-',
      key: 'today_reg_user_count'
    },
    {
      tit: '当天下单数',
      val: '-',
      key: 'today_reg_user_order_count'
    },
    {
      tit: '当天注册用户下单占比',
      val: '-',
      key: 'today_reg_user_order_rate'
    },
    {
      tit: '次日登陆用户数',
      val: '-',
      key: 'next_day_open_app_user_count'
    },
    {
      tit: '次日下单数',
      val: '-',
      key: 'next_day_order_count'
    },
  ])
  const handleChange = () => {
    handleLoadData()
  }
  onMounted(() => {
    getLastDay()
    handleLoadData()
  })

  const getLastDay = () => {
    let day = getDay(-1)
    registerTime.value = formatDate(day)
  }

  const formatDate = (date : any) => {
    let newDate = date && ref<Dayjs>(dayjs(date, "YYYY-MM-DD"))
    return newDate.value
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

const handleLoadData = async () => {
  const payload = {
    start: registerTime.value != undefined ? registerTime.value.format('YYYY-MM-DD') : "",
    end:  registerTime.value != undefined ? registerTime.value.format('YYYY-MM-DD') : "",
  }
  const res = await getRegUsageApi(payload)
  Object.keys(res).forEach((key: any) => {
    let length = data.value.length;
    for (let i = 0;i<length;i++){
      let dataKey = data.value[i].key;
      if (key == dataKey){
        data.value[i].val = res[key]
        break
      }
    }
  })
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
  width: 80%;
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


.mobile-card-tit {
  font-weight: 600;
  font-size: 18px;
  margin-right: 20px;
}

.mobile-card-content {
  width: 100%;
  display: flex;
  justify-content: flex-start;
  align-items: stretch;
  flex-wrap: wrap;
}

.mobile-data-content {
  display: flex;
  flex-direction: column;
  flex: 0 0 calc(50% );
  margin-top: 4px;
  box-sizing: border-box;
}

.mobile-data-tit {
  opacity: 0.7;
  font-weight: 600;
  font-size: 13px;
}

.mobile-data-val {
  font-weight: 500;
  font-size: 22px;
}
</style>