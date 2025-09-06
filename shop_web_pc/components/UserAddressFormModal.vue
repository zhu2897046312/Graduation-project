<script setup lang="ts">
import { NModal, NCard, NForm, NFormItem, NInput, NSelect, NButton, useMessage } from 'naive-ui';
import api from '~/api';
import type { FormRules } from 'naive-ui';
import addressJson from '@/data/address.json'
const checkoutRules = ref<FormRules>({
    email: [
      {
        required: true,
        message: 'Email is required'
      },
      {
        pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
        message: 'Please enter a valid email address'
      }
    ],
    first_name: [{
      required: true,
      message: 'First name is required'
    }],
    last_name: [{
      required: true,
      message: 'Last name is required'
    }],
    phone: [{
      required: true,
      message: 'Phone is required'
    }],
    country: [{
      required: true,
      message: 'Country is required'
    }],
    province: [{
      required: true,
      message: 'region is required'
    }],
    postal_code: [{
      required: true,
      message: 'Postal code is required'
    }],
    detail_address: [{
      required: true,
      message: 'Address is required'
    }],
    city: [{
      required: true,
      message: 'City is required'
    }],
})
const formRef = ref<any>(null)
const show = ref(false)
const message = useMessage()
const emit = defineEmits(['change'])

const info = ref<any>({
  title: new Date().toLocaleString(),
  default_status: 2,
  first_name: '',
  last_name: '',
  email: '',
  phone: '',
  country: null,
  province: null,
  city: '',
  region: '',
  detail_address: '',
  postal_code: '',
})

let id = 0
const countryOptions = computed(() => {
 return addressJson.countries.map(country => ({
  label: country.name,
  value:country.name
 }))
})
const provinceOptions = computed(() => {
  const selectCountry = addressJson.countries.find(c => c.name === info.value.country)
  return selectCountry?.provinces?.map(p => ({
    label: p,
    value:p
  })) || []
})
const handleValidate = async () => {
  console.log(info.value)
  console.log(formRef.value)
  formRef.value.validate((errors: any) => {
    console.log(errors)
    if (!errors) {
      handleSubmit()
    } else {
      window.scrollTo(0, 70);
    }
  })
}

const handleSubmit = async () => {
  const values = { ...info.value }
  try {
    if (id > 0) {
      values.id = id
      await api.shop.address.modify(values)
    } else {
      await api.shop.address.create(values)
    }
    show.value = false
    emit('change')
  } catch (error: any) {
    message.error(error.message)
  }
}

defineExpose({
  useOpen: async (e: number) => {
    if (e > 0) {
      id = e
      info.value = await api.shop.address.info(e)
    } else {
      id = 0
      info.value = {
        title: new Date().toLocaleString(),
        default_status: 2,
        first_name: '',
        last_name: '',
        email: '',
        phone: '',
        country: null,
        province: null,
        city: '',
        region: '',
        detail_address: '',
        postal_code: '',
      }
    }
    console.log('cccc')
    show.value = true
  }
})

</script>

<template>
  <NModal v-model:show="show" style="width: 680px;" @close="emit('change')">
    <NCard style="width: 723px;" :bordered="false" >
          <h1 class="text-xl font-bold mb-4">Address Information</h1>
          <NForm size="large" ref="formRef" :rules="checkoutRules" :model="info">
            <NFormItem label="Email" path="email">
              <div class="flex flex-col gap-2 w-full">
                <NInput v-model:value="info.email" placeholder="Enter email address" />
                <span class="text-sm text-gray-500 pt-2 px-1">Email me with news and offers</span>
              </div>
            </NFormItem>
            <div class="flex items-center gap-2 w-full">

              <NFormItem class="flex-1" label="Last Name" path="last_name">
                <NInput v-model:value="info.first_name" placeholder="Enter last name"/>
              </NFormItem>
              <NFormItem class="flex-1" label="First Name" path="first_name">
                <NInput v-model:value="info.last_name" placeholder="Enter first name" />
              </NFormItem>
            </div>
            <NFormItem label="Phone Number" path="phone">
              <NInput v-model:value="info.phone" placeholder="Enter phone number"/>
            </NFormItem>
            
            <NFormItem label="Country" path="country">
              <NSelect 
                v-model:value="info.country" 
                placeholder="Select Country" 
                :options="countryOptions"
                label-field="label"
                value-field="value"
                clearable
              />
            </NFormItem>
            <NFormItem label="Region" path="province">
              <NSelect 
              v-model:value="info.province" 
              placeholder="Select region/province" 
              :options="provinceOptions"
              :disabled="!info.country"
              clearable
              />
            </NFormItem>
            <NFormItem label="City" path="city">
              <NInput v-model:value="info.city" placeholder="Enter city"/>
            </NFormItem>
            <NFormItem label="Address" path="detail_address">
              <NInput v-model:value="info.detail_address" placeholder="Enter address"/>
            </NFormItem>
            <NFormItem label="ZIP Code" path="postal_code" >
              <NInput v-model:value="info.postal_code" placeholder="Enter postal_code"/>
            </NFormItem>
          </NForm>
          <template #footer>
        <NButton block size="large" type="info" @click="handleValidate">Save</NButton>
      </template>
        </NCard>
  </NModal>
</template>