<template>
  <div class="content">
    <Card style="border-radius: 8px">
      <span class="card-tit">{{prop.title}}</span>
      <div class="card-content">
        <div class="data-content" v-for="dt in data">
                <span class="data-tit">
                  {{ dt.tit }}
                </span>
          <span class="data-val">
                  {{ dt.val }}
          </span>
        </div>
      </div>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import {PropType} from 'vue'
import {Card} from 'ant-design-vue'
const prop = defineProps({
  title: String,
  data: Object as PropType<any>,
  range: Object as PropType<any>
})
</script>
<style scoped>
.content {
  margin-bottom: 10px;
  width: 100%;
}

.card-tit {
  font-weight: 600;
  font-size: 18px;
  margin-right: 20px;
}

.card-content {
  width: 100%;
  display: flex;
  justify-content: flex-start;
  align-items: stretch;
  flex-wrap: wrap;
}

.data-content {
  display: flex;
  flex-direction: column;
  flex: 0 0 calc(50% );
  margin-top: 4px;
  box-sizing: border-box;
}

.data-tit {
  opacity: 0.7;
  font-weight: 600;
  font-size: 13px;
}

.data-val {
  font-weight: 500;
  font-size: 22px;
}
</style>