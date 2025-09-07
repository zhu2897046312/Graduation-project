<script setup lang="ts">
import api from "../../../api";
import { NButton, NPopconfirm } from "naive-ui";


const { data: siteInfo } = await useAsyncData('siteInfo', async () => {
  return await api.shop.market.siteInfo()
})

const { data, refresh, status } = await useAsyncData('addresses', async () => {
  const res = await api.shop.address.list({ page_no: 1, page_size: 999})
  return { addresses: res.list }
})

const addressFormRef = ref()

const handleDel = (id: number) => {
  api.shop.address.del(id).then(() => {
    refresh()
  })
}

const handleEdit = (id: number) => {
  addressFormRef.value.useOpen(id)
}

const handleAddressChange = () => {
  refresh()
}

</script>

<template>
  <div class="container">

    <template v-if="siteInfo != null">
      <Title>My Account - {{ siteInfo.seo_title ? siteInfo.seo_title : '' }}</Title>
      <Meta name="keywords" content="My Account" />
      <Meta name="description" content="My Account" />
    </template>

    <div class="mt-8 flex items-center justify-between">
      <h2 class="title">My Account</h2>
      <NButton class="btn" type="info" @click="handleEdit(0)">Add Address</NButton>
    </div>
    <NuxtLink class="mt-4 block " style="color: #878787 " to="/account">Back to Account Details</NuxtLink>
    <div class="flex flex-col gap-2 mt-8 items-center" v-if="status == 'success'">
      <h1 style="font-size: 1.5em;">Your Addresses</h1>
      <div class="flex flex-col gap-4" v-if="data && data.addresses.length > 0">
        <div class="flex flex-col" v-for="item in data.addresses">
          <span>{{ item.first_name }} {{ item.last_name }}</span>
          <span>{{ item.detail_address }}</span>
          <span>{{ item.region }}</span>
          <span>{{ item.area_code }} {{ item.city }}</span>
          <span>{{ item.province }}</span>
          <div class="flex gap-2 mt-2">
            <span style="color: #878787;cursor: pointer;" @click="handleEdit(item.id)">Edit</span>
            |
            <NPopconfirm positive-text="Confirm" negative-text="Cancel" @positive-click="handleDel(item.id)">
              <template #trigger>
                <span style="color: #878787;cursor: pointer;">Delete</span>
              </template>
              Are you sure you want to delete this address?
            </NPopconfirm>
          </div>
        </div>
      </div>
      <span v-else>No addresses saved yet</span>
    </div>

    <UserAddressFormModal ref="addressFormRef" @change="handleAddressChange"/>
  </div>
</template>

<style lang="css" scoped>
.title {
  color: #878787;
  font-size: 1.733em;
  font-style: normal;
}
</style>