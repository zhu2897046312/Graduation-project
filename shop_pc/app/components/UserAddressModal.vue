<script setup lang="ts">
import { NModal, NCard, NList, NListItem, NThing, NButton, NButtonGroup, NPopconfirm, NEmpty } from 'naive-ui';
import api from '../../api';
import UserAddressFormModal from './UserAddressFormModal.vue';

const emit = defineEmits(['selected'])

const show = ref(false)

const address_list = ref<any[]>([])

const addressRef = ref()

const handleShow = async () => {
  const res = await api.shop.address.list({page_no: 1, page_size: 50})
  address_list.value = res.list
  show.value = true
}

const handleShowForm = (e: number) => {
  if (addressRef.value) {
    addressRef.value.useOpen(e)
    show.value = false
  }
}

const handleDel = async (e: number) => {
  await api.shop.address.del(e)
  handleShow()
}

const handleSelected = (e: any) => {
  emit('selected', e)
  show.value = false
}


defineExpose({
  useOpen: async() => {
    handleShow()
  }
})

</script>

<template>
  <div>
    <NModal v-model:show="show" style="width: 680px;">
      <NCard title="My Addresses">
        <template #header-extra>
          <NButton type="primary" @click="handleShowForm(0)">Add New Address</NButton>
        </template>
        <NList bordered>
          <template v-if="address_list.length > 0">
            <NListItem v-for="item in address_list">
              <NThing :title="`${item.first_name} ${item.last_name}`" :description="`${item.email} ${item.phone}`">
                <div class="flex items-center gap-1">
                  <span>{{ item.province }}</span>
                  <span>{{ item.city }}</span>
                  <span>{{ item.region }}</span>
                </div>
                <div>{{ item.detail_address }}</div>
              </NThing>
              <template #prefix>
                <NButton type="primary" quaternary  @click="handleSelected(item)">Select</NButton>
              </template>
              <template #suffix>
                <NButtonGroup>
                  <NButton type="info" quaternary  @click="handleShowForm(item.id)">Edit</NButton>
                  <NPopconfirm @positive-click="handleDel(item.id)">
                    <template #trigger>
                      <NButton type="error" quaternary>Delete</NButton>
                    </template>
                    Confirm deletion?
                  </NPopconfirm>
                </NButtonGroup>
              </template>
            </NListItem>
          </template>
          <NEmpty v-else description="No addresses available">
            <template #extra>
              <NButton size="small" @click="handleShowForm(0)">Add New Address</NButton>
            </template>
          </NEmpty>
        </NList>
      </NCard>
    </NModal>
    <UserAddressFormModal ref="addressRef" @change="handleShow()" />
  </div>
</template>