<template>
  <div ref="warpRef" style="height: 100%;">
    <Table 
      :data-source="list_data.length ? list_data : tableData"
      bordered
      empty-text="暂无数据"
      :pagination="prop.showPage ? pagination : false"
      :row-key="prop.rowKey"
      :scroll="{
        scrollToFirstRowOnChange: true,
        y: box_height,
        x: prop.width + 'px'
      }"
      size="small"
      @change="handleLoadData(false)"
      :columns="columns">
      <template #bodyCell="data">
        <slot name="bodyCell" v-bind="data || {}"></slot>
      </template>
    </Table>
  </div>
</template>

<script lang="ts">
interface ApiResult {
  list: any[];
  total: number;
}
</script>

<script lang="ts" setup>
  import { ref, PropType, onMounted, onUnmounted, nextTick, computed } from 'vue';
  import { Table } from 'ant-design-vue';

  const prop = defineProps({
    api: { // 网络接口
      type: Function as PropType<(values?: Record<string, any>) => Promise<ApiResult>> | any,
    },
    showPage: {
      type: Boolean,
      default: true,
    },
    rowKey: {
      type: String,
      default: 'id',
    },
    tableData: {
      type: Array,
      default: () => []
    },
    columns: {
      type: Array as PropType<any>,
      default: () => [],
    },
    showId: {
      type: Boolean,
      default: true,
    },
    showAction: {
      type: Boolean,
      default: true,
    },
    actionWidth: {
      type: Number,
      default: 0,
    },
    searchParam: {
      type: Object as PropType<any>,
      default: () => {},
    },
    width: {
      type: Number,
      default: 1000,
    },
    total: {
      type: Number,
      default: 0,
    },
  });

  const warpRef = ref<any>(null);

  // 表格数据
  const list_data = ref<any[]>([]);
  // 加载中状态
  const loading = ref<boolean>(true);
  const emit = defineEmits(["changeTotal", "pageChange"]);

  const pagination = ref<any>({
    current: 1,
    pageSize: 30,
    pageSizeOptions: ['10', '30', '50', '100'],
    showQuickJumper: true,
    total: prop.total ? prop.total : 0,
    onChange: (page, pageSize) => {
      pagination.value.current = page
      pagination.value.pageSize = pageSize
    }
  })

  const columns = computed<any[]>(() => {
    const base_column: any[] = []

    if (prop.showId) {
      base_column.push({dataIndex: 'id', width: 78, align: 'center', title: '#'});
    }
    const cols = base_column.concat(prop.columns);

    if (prop.showAction) {
      const config = { align: 'left', title: '操作', key: '_action', fixed: 'right',}
      if (prop.actionWidth > 0) config['width'] = prop.actionWidth
      cols.push(config)
    }
    return cols;
  })

  // 表格容器的高度
  const box_height = ref<number>(200);
  const box_width = ref<number>(1000);

  onMounted(() => {
    handleCmpHeight();
    handleLoadData(false);
    window.addEventListener('resize', handleCmpHeight)
  });

  onUnmounted(() => {
    window.removeEventListener('resize', handleCmpHeight)
  });
  
  const handleCmpHeight = () => {
    nextTick(() => {
      const tableBodyEl = warpRef.value.querySelector('div.ant-table-body') as HTMLElement;
      const btn = document.querySelector('.ant-table-tbody')?.querySelector('.ant-btn-group');
      let offsetTop = tableBodyEl.offsetTop;
      let pEL = tableBodyEl.offsetParent as HTMLElement | null;
      // 获取 tableBodyEl 元素相对于浏览器的顶部
      while (pEL) {
        offsetTop += pEL.offsetTop;
        pEL = pEL.offsetParent as HTMLElement | null;
      }
      const footerEl = document.querySelector('.ant-layout-footer') as HTMLElement | null;
      const footerHeight = footerEl ? footerEl.clientHeight : 20;
      const bodyHeight = window.innerHeight - offsetTop - footerHeight - (prop.showPage ? 75 : 30) - 20
      tableBodyEl.style.height = `${bodyHeight}px`;
      box_height.value = bodyHeight; 
      box_width.value = (btn?.clientWidth ?  btn?.clientWidth * 4 : 0) + tableBodyEl.clientWidth
    })
  }

  const handleLoadData = async (reload: boolean) => {
    console.log(reload)
    const req: any = Object.assign({}, prop.searchParam);
    if(reload){
      pagination.value.current = 1
    }
    if (prop.showPage) {
      req.page_size = pagination.value.pageSize;
      req.page_no = pagination.value.current;
      emit('pageChange', {
        page_no: pagination.value.current,
        page_size: pagination.value.pageSize
      })
    }
    try {
      loading.value = true;
      const res = await prop.api(req);
      list_data.value = res.list;
      emit('changeTotal', res.total)
      pagination.value.total = res.total;

    } catch (e) {
      console.warn(e);
    } finally {
      loading.value = false;
    }
    handleCmpHeight()
  };

  defineExpose({
    useReload: (reload: boolean) => {
      return handleLoadData(reload);
    }
  })

</script>

<style>
.ant-table-header {
	overflow-y: hidden!important;
}
</style>