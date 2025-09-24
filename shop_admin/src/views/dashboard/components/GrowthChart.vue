<template>
    <Card :title="title" class="chart-card">
      <div class="chart-header">
        <RadioGroup v-model:value="timeRange" @change="handleTimeRangeChange">
          <RadioButton v-for="option in timeRangeOptions" 
                       :key="option.value" 
                       :value="option.value">
            {{ option.label }}
          </RadioButton>
        </RadioGroup>
      </div>
      <div ref="chartRef" style="height: 300px;"></div>
    </Card>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted, onUnmounted, watch, computed } from 'vue';
  import { Card, RadioGroup, RadioButton } from 'ant-design-vue';
  import * as echarts from 'echarts';
  
  interface GrowthData {
    date: string;
    count: number;
  }
  
  interface TimeRangeOption {
    label: string;
    value: string;
  }
  
  interface Props {
    title?: string;
    dataList?: any[];
    // 日期字段名，默认为'created_time'
    dateField?: string;
    // 数值字段名，如果提供则统计该字段的和，否则计数
    valueField?: string;
    // 时间范围选项
    timeRangeOptions?: TimeRangeOption[];
    // 图表模式：累计(cumulative) 或 每日新增(daily)
    mode?: 'cumulative' | 'daily';
    // 图表颜色
    chartColor?: string;
  }
  
  const props = withDefaults(defineProps<Props>(), {
    title: '增长趋势',
    dataList: () => [],
    dateField: 'created_time',
    valueField: '',
    timeRangeOptions: () => [
      { label: '近7天', value: '7' },
      { label: '近一个月', value: '30' },
      { label: '近三个月', value: '90' }
    ],
    mode: 'cumulative',
    chartColor: '#52c41a'
  });
  
  const chartRef = ref<HTMLElement>();
  const timeRange = ref<string>(props.timeRangeOptions[0].value);
  let chart: echarts.ECharts | null = null;
  
  // 计算当前时间范围选项
  const currentTimeRangeOption = computed(() => {
    return props.timeRangeOptions.find(option => option.value === timeRange.value) || props.timeRangeOptions[0];
  });
  
  // 生成指定天数的日期数组
  const generateDateRange = (days: number): string[] => {
    const dates: string[] = [];
    for (let i = days - 1; i >= 0; i--) {
      const date = new Date();
      date.setDate(date.getDate() - i);
      dates.push(date.toISOString().split('T')[0]);
    }
    return dates;
  };
  
  // 分析增长数据
  const analyzeGrowthData = (): GrowthData[] => {
    const days = parseInt(timeRange.value);
    const dateRange = generateDateRange(days);
    const dataCountByDate: { [key: string]: number } = {};
    
    // 初始化每天数据为0
    dateRange.forEach(date => {
      dataCountByDate[date] = 0;
    });
    
    // 统计每天的数据
    props.dataList.forEach(item => {
      const itemDate = item[props.dateField]?.split('T')[0];
      if (itemDate && dateRange.includes(itemDate)) {
        const value = props.valueField ? 
          (parseFloat(item[props.valueField]) || 0) : 1;
        dataCountByDate[itemDate] = (dataCountByDate[itemDate] || 0) + value;
      }
    });
    
    // 根据模式返回数据
    if (props.mode === 'daily') {
      return dateRange.map(date => ({
        date: formatDateLabel(date, days),
        count: dataCountByDate[date] || 0
      }));
    } else {
      // 累计模式
      let cumulativeCount = 0;
      return dateRange.map(date => {
        const dailyCount = dataCountByDate[date] || 0;
        cumulativeCount += dailyCount;
        return {
          date: formatDateLabel(date, days),
          count: cumulativeCount
        };
      });
    }
  };
  
  // 根据时间范围格式化日期标签
  const formatDateLabel = (date: string, days: number): string => {
    if (days <= 30) {
      return date.substring(5); // 显示月-日格式
    } else {
      return date.substring(2, 7); // 显示年-月格式
    }
  };
  
  // 时间范围改变处理
  const handleTimeRangeChange = () => {
    updateChart();
  };
  
  // 更新图表
  const updateChart = () => {
    if (!chart) return;
    
    const growthData = analyzeGrowthData();
    
    chart.setOption({
      xAxis: {
        data: growthData.map(item => item.date)
      },
      series: [{
        data: growthData.map(item => item.count)
      }]
    });
  };
  
  // 初始化图表
  const initChart = () => {
    if (!chartRef.value) return;
    
    chart = echarts.init(chartRef.value);
    const growthData = analyzeGrowthData();
    const days = parseInt(timeRange.value);
    
    chart.setOption({
      tooltip: {
        trigger: 'axis',
        formatter: (params: any) => {
          const param = params[0];
          const modeText = props.mode === 'cumulative' ? '累计' : '新增';
          return `${param.name}<br/>${modeText}: ${param.value}`;
        }
      },
      xAxis: {
        type: 'category',
        data: growthData.map(item => item.date),
        axisLabel: {
          rotate: days > 30 ? 45 : 0
        }
      },
      yAxis: {
        type: 'value'
      },
      series: [{
        data: growthData.map(item => item.count),
        type: 'line',
        smooth: true,
        itemStyle: {
          color: props.chartColor
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [{
              offset: 0,
              color: hexToRgba(props.chartColor, 0.3)
            }, {
              offset: 1,
              color: hexToRgba(props.chartColor, 0.1)
            }]
          }
        }
      }],
      grid: {
        left: '3%',
        right: '4%',
        bottom: days > 30 ? '15%' : '3%',
        containLabel: true
      }
    });
  };
  
  // 将十六进制颜色转换为RGBA
  const hexToRgba = (hex: string, alpha: number): string => {
    const r = parseInt(hex.slice(1, 3), 16);
    const g = parseInt(hex.slice(3, 5), 16);
    const b = parseInt(hex.slice(5, 7), 16);
    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
  };
  
  // 响应式调整图表大小
  const handleResize = () => {
    chart?.resize();
  };
  
  onMounted(() => {
    initChart();
    window.addEventListener('resize', handleResize);
  });
  
  onUnmounted(() => {
    window.removeEventListener('resize', handleResize);
    chart?.dispose();
  });
  
  // 监听数据变化
  watch(() => props.dataList, () => {
    updateChart();
  });
  
  // 监听配置变化
  watch(() => [props.mode, props.chartColor, props.dateField, props.valueField], () => {
    updateChart();
  });
  
  // 暴露方法
  defineExpose({
    refresh: () => {
      updateChart();
    }
  });
  </script>
  
  <style scoped>
  .chart-header {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 16px;
  }
  
  .chart-card {
    margin-bottom: 16px;
  }
  
  :deep(.ant-card-head) {
    border-bottom: 1px solid #f0f0f0;
  }
  
  :deep(.ant-card-body) {
    padding: 24px;
  }
  </style>