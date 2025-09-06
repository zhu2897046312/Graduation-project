<template>
  <div>
    <div class="p-2 rounded bg-gray-100">
      <div class="mb-4">
        <div class="flex gap-4 items-center" style="width: 460px;">
          <Select v-model:value="product_prod_list" :options="prod_options" mode="multiple" style="width: 98%;" @change="handleProdChange(false)"></Select>
          <Button @click="() => { add_prod_attr && add_prod_attr.useOpen() }">添加商品属性</Button>
        </div>
      </div>
      <div class="mb-2">
        <div class="flex gap-4 items-center mb-2" v-for="(item, index) in prod_value_list">
          <div>{{ item.prod.title }}:</div>
          <Checkbox.Group v-model:value="prod_value_list[index].checked" :options="item.list">
          </Checkbox.Group>
          <div class="flex gap-1 items-center">
            <Input style="width: 98px;" v-model:value="prod_value_list[index].add" size="small" placeholder="新增属性值" />
            <Button size="small" @click="handleAddProdAttrValue(index)">添加</Button>
          </div>
        </div>
      </div>
    </div>


    <div class="mb-2 mt-3">
      <Button.Group>
        <Popconfirm :disabled="prod_value_list.length == 0" title="刷新列表将导致sku信息重新生成，是否要刷新?" @confirm="handleReloadSkuList">
          <Button :disabled="prod_value_list.length == 0" type="link">刷新SKU列表</Button>
        </Popconfirm>
        <Popconfirm :disabled="list.length < 2" title="将同步第一个sku的【当前价格】到所有sku,是否继续?" @confirm="handleSetSkuProperty('price')">
          <Button :disabled="list.length < 2" type="link">同步当前价格</Button>
        </Popconfirm>
        <Popconfirm :disabled="list.length < 2" title="将同步第一个sku的【原价格】到所有sku,是否继续?" @confirm="handleSetSkuProperty('original_price')">
          <Button :disabled="list.length < 2" type="link">同步原价格</Button>
        </Popconfirm>
        <Popconfirm :disabled="list.length < 2" title="将同步第一个sku的【成本价】到所有sku,是否继续?" @confirm="handleSetSkuProperty('cost_price')">
          <Button :disabled="list.length < 2" type="link">同步成本价</Button>
        </Popconfirm>
        <Popconfirm :disabled="list.length < 2" title="将同步第一个sku的【库存】到所有sku,是否继续?" @confirm="handleSetSkuProperty('stock')" >
          <Button :disabled="list.length < 2" type="link">同步库存</Button>
        </Popconfirm>
      </Button.Group>
    </div>
    <Table size="small" :pagination="false" :columns="columns" bordered row-key="code" :dataSource="list">
      <template #bodyCell="{ column, index }: { column: any, index: number, text: any }">
        <InputNumber prefix="￥" :min="0" :max="999999" :precision="2" v-if="column.key === 'price'" v-model:value="list[index]['price']" />
        <InputNumber prefix="￥" :min="0" :max="999999" :precision="2" v-else-if="column.key === 'original_price'" v-model:value="list[index]['original_price']" />
        <InputNumber prefix="￥" :min="0" :max="999999" :precision="2" v-else-if="column.key === 'cost_price'" v-model:value="list[index]['cost_price']" />
        <InputNumber :min="0" :max="999999" :precision="0" v-else-if="column.key === 'stock'" v-model:value="list[index]['stock']" />
        <Switch 
          v-else-if="column.key === 'state'" 
          checked-children="启用" 
          un-checked-children="停用" 
          :checked-value="1" 
          :un-checked-value="2" 
          v-model:checked="list[index]['state']" />
          <Switch 
            v-else-if="column.key === 'default_show'" 
            checked-children="默认" 
            un-checked-children="非默认" 
            :checked-value="1" 
            :un-checked-value="2" 
            v-model:checked="list[index]['default_show']" />
      </template>
    </Table>
    <Add ref="add_prod_attr" @on-change="handleAddProdAttr" />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Table, Button, InputNumber, Popconfirm, message, Input, Form, Select, Checkbox, Switch } from 'ant-design-vue';
import api from '/@/api';
import Add from '/@/views/shop/prod_attributes/components/Add.vue'
import { loadingTask } from '/@/utils/helper';

const porps = defineProps<{ value: any[], skuConfig: Map<number, number[]> }>()

const emit = defineEmits(['update:value', 'change'])

const add_prod_attr = ref<any>(null)

const columns: any[] = [
  { dataIndex: 'title', key: 'title', title: '规格名称', width: 268,  },
  { dataIndex: 'price', key: 'price', title: '当前价格', width: 168, },
  { dataIndex: 'original_price', key: 'original_price', title: '原价格', width: 168, },
  { dataIndex: 'cost_price', key: 'cost_price', title: '成本价', width: 168, },
  { dataIndex: 'stock', key: 'stock', title: '库存', width: 168, },
  { dataIndex: 'state', key: 'state', title: '状态', width: 98, },
  { dataIndex: 'default_show', key: 'default_show', title: '默认显示', width: 168, },
]


const list = computed<any[]>({
  get: () => porps.value,
  set: (e: any[]) => {
    emit('update:value', e)
    emit('change', e)
    formItemContext && formItemContext.onFieldChange()
  }
})

const product_prod_list = ref<number[]>([])

// 商品属性列表
const prod_options = ref<any[]>([])

// 商品属性值
const prod_value_list = ref<{prod: any, list: any[], checked: any[], add: ''}[]>([])


const source_sku_list = new Map<string, any>()


let formItemContext: any = null
onMounted(async () => {
  formItemContext = Form.useInjectFormItemContext();
  handleAddProdAttr()
})

const handleAddProdAttr = async () => {
  prod_options.value = (await api.shop.prodAttributes.list({page_no: 1, page_size: 1000}) as any).list.map(it => {
    return { label: it.title, value: it.id }
  })

  // 这里进行SKU配置初始化
  const _product_prod_list: number[] = [];
  for(let prod_id of porps.skuConfig.keys()) {
    _product_prod_list.push(prod_id)
  }
  product_prod_list.value = _product_prod_list
  await handleProdChange(true)

  // 这里记录一些旧的sku信息，用于后面刷新用
  for(let item of list.value) {
    source_sku_list.set(item.sku_code, JSON.parse(JSON.stringify(item)))
  }
  console.log(source_sku_list)
}

const handleProdChange = async (is_first :boolean = false) => {
  console.log('handleProdChange', '12312312')
  // 移除prod_value_list中已经不存在的商品属性
  prod_value_list.value = prod_value_list.value.filter(it => {
    return product_prod_list.value.includes(it.prod.id)
  })
  for(let prod_id of product_prod_list.value) {
    const res_prod: any = await api.shop.prodAttributes.info(prod_id)
    const res_list: any = (await api.shop.prodAttributesValue.getProdValues(prod_id) as any[]).map(it => {
      return { label: it.title, value: it.id }
    })
    let hit = false
    for(let idx in prod_value_list.value) {
      if(prod_value_list.value[idx].prod.id == res_prod.id) {
        prod_value_list.value[idx].list = res_list
        prod_value_list.value[idx].add = ''
        hit = true
        break
      }
    }
    if (!hit) {
      let _checked: number[] = []
      if (is_first && porps.skuConfig.has(res_prod.id)) {
        _checked = porps.skuConfig.get(res_prod.id)!
      }
      prod_value_list.value.push({
        prod: res_prod,
        list: res_list,
        checked: _checked,
        add: ''
      })
    }
  }
}

/**
 * 刷新SKU列表
 */
const handleReloadSkuList = () => {
  const generateSKUsRecursive = (cur_sku: any, cur_idx: number, sku_list: any[]) => {
    const values = prod_value_list.value[cur_idx].list.filter(it => {
      return prod_value_list.value[cur_idx].checked.includes(it.value)
    })
    values.forEach(one => {
      const new_sku = {
        id: 0,
        _pord: [one],
        sku_code: '',
        title: '',
        price: 0.00,
        original_price: 0.00,
        cost_price: 0.00,
        stock: 99,
        default_show: 2,
        state: 1,
      }
      if (cur_sku) {
        for (const it of cur_sku._pord) {
          new_sku._pord.push(it)
        }
      }
      if (cur_idx < prod_value_list.value.length - 1) {
        generateSKUsRecursive(new_sku, cur_idx + 1, sku_list)
      } else {
        sku_list.push(new_sku)
      }
    })
    return sku_list
  }

  const res = generateSKUsRecursive(null, 0, [])
  const out = res.map(it => {
    it._pord.sort((a, b) => a.value - b.value);
    it.sku_code = it._pord.map(it => it.value).join(';');
    it.title = it._pord.map(it => it.label).join(';');
    const source_info = source_sku_list.get(it.sku_code)
    if (source_info) {
      it.id = source_info.id
      it.price = source_info.price
      it.original_price = source_info.original_price
      it.cost_price = source_info.cost_price
      it.stock = source_info.stock
      it.state = source_info.state
      it.default_show = source_info.default_show
    }
    return it;
  })
  list.value = out
  console.log('dasdasd',out)
}

const handleSetSkuProperty = (e: any) => {
  if (list.value.length < 2) {
    message.info('当前规格列表不足2个，无法设置价格')
    return;
  }
  const val = list.value[0][e]
  for(let i = 1; i < list.value.length; i++) {
    list.value[i][e] = val
  }
}

/**
 * 添加属性值
 */
const handleAddProdAttrValue = async (index: number) => {
  loadingTask(async () => {
    const payload = {
      prod_attributes_id: prod_value_list.value[index].prod.id,
      title: prod_value_list.value[index].add,
      sort_num: 10,
    }
    const res: any = await api.shop.prodAttributesValue.create(payload)
    prod_value_list.value[index].list.push({
      label: res.title,
      value: res.id,
    })
    prod_value_list.value[index].add = ''
  })
}

</script>