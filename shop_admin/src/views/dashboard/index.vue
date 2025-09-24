<template>
  <PageLayout>
    <div class="dashboard-content">
      <Row :gutter="16" style="margin-bottom: 12px;">
        <Col :span="6">
          <Card>
            <div class="stat-card">
              <div class="stat-icon" style="background: #fff2f0;">
                <UserOutlined style="color: #ff4d4f; font-size: 24px;" />
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalUsers || 0 }}</div>
                <div class="stat-label">总用户数</div>
              </div>
            </div>
          </Card>
        </Col>
        <Col :span="6">
          <Card>
            <div class="stat-card">
              <div class="stat-icon" style="background: #f6ffed;">
                <ShoppingOutlined style="color: #52c41a; font-size: 24px;" />
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalProducts || 0 }}</div>
                <div class="stat-label">商品总数</div>
              </div>
            </div>
          </Card>
        </Col>
        <Col :span="6">
          <Card>
            <div class="stat-card">
              <div class="stat-icon" style="background: #f0f5ff;">
                <DollarOutlined style="color: #1890ff; font-size: 24px;" />
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalOrders || 0 }}</div>
                <div class="stat-label">订单总数</div>
              </div>
            </div>
          </Card>
        </Col>
        <Col :span="6">
          <Card>
            <div class="stat-card">
              <div class="stat-icon" style="background: #fff7e6;">
                <BarChartOutlined style="color: #fa8c16; font-size: 24px;" />
              </div>
              <div class="stat-content">
                <div class="stat-value">¥{{ (stats.totalSales || 0).toLocaleString() }}</div>
                <div class="stat-label">总销售额</div>
              </div>
            </div>
          </Card>
        </Col>
      </Row>

      <!-- 图表区域 -->
      <Row :gutter="16">
        <!-- 用户增长图表 -->
        <Col :span="12">
          <GrowthChart 
            title="用户增长趋势" 
            :dataList="userList"
            dateField="created_time"
            mode="cumulative"
            chartColor="#52c41a"
            ref="userGrowthChartRef" 
          />
        </Col>
        
        <!-- 订单增长图表 -->
        <Col :span="12">
          <GrowthChart 
            title="订单增长趋势" 
            :dataList="orderList"
            dateField="created_time"
            mode="daily"
            chartColor="#1890ff"
            ref="orderGrowthChartRef" 
          />
        </Col>
      </Row>

      <!-- 销售额图表（可选） -->
      <Row :gutter="16" style="margin-top: 16px;">
        <Col :span="12">
          <GrowthChart 
            title="销售额趋势（累计）" 
            :dataList="orderList"
            dateField="created_time"
            valueField="pay_amount"
            mode="cumulative"
            chartColor="#fa541c"
            :timeRangeOptions="[
              { label: '近7天', value: '7' },
              { label: '近一个月', value: '30' },
              { label: '近三个月', value: '90' }
            ]"
            ref="salesGrowthChartRef" 
          />
        </Col>
        <Col :span="12">
          <GrowthChart 
            title="销售额趋势（每日）" 
            :dataList="orderList"
            dateField="created_time"
            valueField="pay_amount"
            mode="daily"
            chartColor="#722ed1"
            ref="dailySalesChartRef" 
          />
        </Col>
      </Row>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { 
  UserOutlined, 
  ShoppingOutlined, 
  DollarOutlined, 
  BarChartOutlined 
} from '@ant-design/icons-vue';
import { Card, Row, Col } from 'ant-design-vue';
import PageLayout from '/@/components/Kernel/Layout/PageLayout.vue';
import GrowthChart from './components/GrowthChart.vue'; // 引入通用图表组件
import api from '/@/api/index';

const ordertableRef = ref<any>(null);
const userGrowthChartRef = ref<any>(null); // 用户增长图表引用
const orderGrowthChartRef = ref<any>(null); // 订单增长图表引用
const salesGrowthChartRef = ref<any>(null); // 销售额图表引用
const dailySalesChartRef = ref<any>(null); // 每日销售额图表引用

// 统计数据类型
interface DashboardStats {
  totalUsers: number;
  totalProducts: number;
  totalOrders: number;
  totalSales: number;
}

// 响应式数据
const stats = ref<DashboardStats>({
  totalUsers: 0,
  totalProducts: 0,
  totalOrders: 0,
  totalSales: 0
});

const userList = ref<any[]>([]); // 用户列表数据
const orderList = ref<any[]>([]); // 订单列表数据

// 加载仪表板数据
const loadDashboardData = async () => {
  try {
    const products = await api.shop.product.list({}) as any;
    const orders = await api.shop.order.list({}) as any;
    const users = await api.mp.user.list({}) as any;
    
    console.log('产品数据:', products);
    console.log('订单数据:', orders);
    console.log('用户数据:', users);
    
    stats.value.totalProducts = products.total || 0;
    stats.value.totalOrders = orders.total || 0;
    stats.value.totalUsers = users.total || 0;
    stats.value.totalSales = orders.list?.reduce((total: number, order: any) => total + (order.pay_amount || 0), 0) || 0;
    
    // 保存用户列表和订单列表数据
    userList.value = users.list || [];
    orderList.value = orders.list || [];
    
    console.log('用户数据条数:', userList.value.length);
    console.log('订单数据条数:', orderList.value.length);
  } catch (error) {
    console.error('加载仪表板数据失败:', error);
  }
};

// 响应式调整图表大小
const handleResize = () => {
  userGrowthChartRef.value?.refresh();
  orderGrowthChartRef.value?.refresh();
  salesGrowthChartRef.value?.refresh();
  dailySalesChartRef.value?.refresh();
};

onMounted(() => {
  loadDashboardData();
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

// 暴露刷新方法
defineExpose({
  refresh: () => {
    loadDashboardData();
    ordertableRef.value?.useReload();
  }
});
</script>

<style scoped>
.dashboard-content {
  padding: 24px;
}

.stat-card {
  display: flex;
  align-items: center;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #262626;
  line-height: 1.2;
}

.stat-label {
  font-size: 14px;
  color: #8c8c8c;
  margin-top: 4px;
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