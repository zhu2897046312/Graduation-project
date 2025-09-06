<template>
  <div class="content">
    <Card>
      <span class="card-tit">{{prop.title}}</span>
      <div class="card-content">
        <div class="data-content" v-for="dt in data">
                <span class="data-tit">
                  {{ dt.tit }}
                </span>
          <span class="data-val">
                  {{ dt.val }}
          </span>
          <div v-if="dt.route">
                     <span class="action" @click="toDetail(dt.route)">
                    查看详情 >
          </span>
          </div>
        </div>
      </div>
    </Card>
  </div>
</template>

<script lang="ts" setup>
import {PropType} from 'vue'
import {Card} from 'ant-design-vue'
import {useRouter} from "vue-router";

const router = useRouter()
const prop = defineProps({
  title: String,
  data: Object as PropType<any>,
  range: Object as PropType<any>
})


const toDetail = (key: string) => {
  router.push({name: key, query: {start: prop.range.start, end: prop.range.end }})
}

</script>

<style scoped>
.content {
  margin-bottom: 10px;
}

.action {
  opacity: 0.5;
  cursor: pointer;
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
}

.data-content {
  display: flex;
  flex-direction: column;
}

.data-tit {
  opacity: 0.7;
  font-weight: 600;
}

.data-val {
  font-weight: 500;
  font-size: 30px;
}
</style>