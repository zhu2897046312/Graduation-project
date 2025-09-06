<template>
  <div>
    <div class="dash-index" style="padding: 12px;">
        <indicators :isMobile="isMobile"/>
        <reg-usage v-if="!isMobile"/>
        <ability v-if="!isMobile"/>
        <analysis v-if="!isMobile"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import indicators from './components/indicators.vue'
import ability from './components/ability.vue'
import analysis from './components/analysis.vue'
import regUsage from './components/regUsage.vue'

const isMobile = /iPhone|iPad|iPod|Android/i.test(navigator.userAgent);

</script>

<style scoped>
.grid-item {
  display: flex;
  text-align: center;
  align-items: center;
  justify-content: center;
}
.grid-item-tx {
  padding-left: 10px;
  font-size: 18px;
}
.color1 {
  color: #d74545;
}
.color2 {
  color: #1bc4a5
}
.color3 {
  color: #e02433
}
.card-tit{
  font-weight: 600;
  font-size: 18px;
}
</style>
<style>
.dash-index .ant-card-body{
  padding: 12px 24px;
}
</style>