<script setup lang="ts">
import api from '../../api'
import type { Address, AddressCreateParams, AddressModifyParams } from '../../api/type'
import { useAddressData } from '../../composables/useAddressData'

definePageMeta({ layout: 'default' })

const router = useRouter()
const toast = useToast()
const accessToken = useCookie('accessToken')
const isAuthed = computed(() => !!accessToken.value)

if (!isAuthed.value) {
  await navigateTo('/login?redirect=/addresses')
}

const { data: addressList, refresh, pending } = await useAsyncData<Address[]>('address-list', async () => {
  try {
    const res = await api.shop.address.list({})
    return res.list || []
  } catch (error) {
    console.error('Failed to fetch addresses:', error)
    return []
  }
})

// Modal 状态
const modalOpen = ref(false)

// 编辑状态
const editingAddress = ref<Address | null>(null)
const formLoading = ref(false)

// 表单数据
const form = reactive<AddressCreateParams & { id?: number }>({
  title: '',
  first_name: '',
  last_name: '',
  email: '',
  phone: '',
  province: '',
  city: '',
  region: '',
  detail_address: '',
  country: '',
  postal_code: '',
  default_status: 0
})

// 地址数据 composable
const { countryOptions, getProvinceOptions } = useAddressData('address-json')

const provinceOptions = computed(() => getProvinceOptions(form.country))

// 重置表单
const resetForm = () => {
  delete form.id
  form.title = ''
  form.first_name = ''
  form.last_name = ''
  form.email = ''
  form.phone = ''
  form.province = ''
  form.city = ''
  form.region = ''
  form.detail_address = ''
  form.country = ''
  form.postal_code = ''
  form.default_status = 0
}

// 打开 Modal（新增 / 编辑共用）
const openModal = (addr?: Address) => {
  if (addr) {
    // 编辑模式
    editingAddress.value = addr
    form.id = addr.id
    form.title = addr.title || ''
    form.first_name = addr.first_name
    form.last_name = addr.last_name
    form.email = addr.email
    form.phone = addr.phone
    form.province = addr.province
    form.city = addr.city
    form.region = addr.region || ''
    form.detail_address = addr.detail_address
    form.country = addr.country
    form.postal_code = addr.postal_code
    form.default_status = addr.default_status || 0
  } else {
    // 新增模式
    editingAddress.value = null
    resetForm()
  }
  modalOpen.value = true
}

// 保存地址
const handleSave = async () => {
  formLoading.value = true
  try {
    if (editingAddress.value) {
      await api.shop.address.modify(form as AddressModifyParams)
      toast.add({ title: 'Success', description: 'Address updated successfully', color: 'success' })
    } else {
      await api.shop.address.create(form)
      toast.add({ title: 'Success', description: 'Address added successfully', color: 'success' })
    }
    modalOpen.value = false
    resetForm()
    editingAddress.value = null
    await refresh()
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to save address'
    toast.add({ title: 'Error', description: errorMessage, color: 'error' })
  } finally {
    formLoading.value = false
  }
}

// 监听国家变化，清空省份
watch(() => form.country, () => {
  form.province = ''
})

// 设置默认地址
const handleSetDefault = async (addr: Address) => {
  try {
    await api.shop.address.modify({ ...addr, id: addr.id, default_status: 1 })
    await refresh()
    toast.add({ title: 'Success', description: 'Default address updated', color: 'success' })
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to update default address'
    toast.add({ title: 'Error', description: errorMessage, color: 'error' })
  }
}

// 删除地址
const handleDelete = async (id: number) => {
  try {
    await api.shop.address.del(id)
    await refresh()
    toast.add({ title: 'Success', description: 'Address deleted successfully', color: 'success' })
  } catch (error: unknown) {
    const errorMessage = error instanceof Error ? error.message : 'Failed to delete address'
    toast.add({ title: 'Error', description: errorMessage, color: 'error' })
  }
}

useHead({
  title: 'Address Book',
  meta: [{ name: 'description', content: 'Manage your shipping addresses' }]
})
</script>

<template>
  <div class="addresses-page max-w-7xl mx-auto w-full px-4 sm:px-6 lg:px-8 py-8">
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">Address Book</h1>
        <p class="text-gray-600 dark:text-gray-400">Manage your shipping addresses</p>
      </div>
      <UButton color="primary" size="lg" icon="i-lucide-plus" @click="openModal()">
        Add Address
      </UButton>
    </div>

    <div v-if="pending" class="flex justify-center items-center py-20">
      <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
    </div>

    <div v-else-if="addressList?.length" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <UCard
        v-for="addr in addressList"
        :key="addr.id"
        :class="['relative', addr.default_status === 1 ? 'ring-2 ring-primary-500' : '']"
      >
        <template #header>
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-2">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                {{ addr.title || 'Address' }}
              </h3>
              <UBadge v-if="addr.default_status === 1" color="primary" variant="subtle" size="sm">
                Default
              </UBadge>
            </div>
          </div>
        </template>

        <div class="space-y-2 text-sm text-gray-600 dark:text-gray-400">
          <p class="font-medium text-gray-900 dark:text-white">
            {{ addr.first_name }} {{ addr.last_name }}
          </p>
          <p>{{ addr.detail_address }}</p>
          <p>{{ addr.city }}{{ addr.region ? `, ${addr.region}` : '' }} {{ addr.postal_code }}</p>
          <p>{{ addr.province }}, {{ addr.country }}</p>
          <p class="pt-2 flex items-center gap-1">
            <UIcon name="i-lucide-phone" class="w-4 h-4" /> {{ addr.phone }}
          </p>
          <p class="flex items-center gap-1">
            <UIcon name="i-lucide-mail" class="w-4 h-4" /> {{ addr.email }}
          </p>
        </div>

        <template #footer>
          <div class="flex gap-2">
            <UButton variant="ghost" color="neutral" size="sm" icon="i-lucide-edit" @click="openModal(addr)">
              Edit
            </UButton>
            <UButton
              v-if="addr.default_status !== 1"
              variant="ghost"
              color="primary"
              size="sm"
              icon="i-lucide-star"
              @click="handleSetDefault(addr)"
            >
              Set Default
            </UButton>
            <UButton
              variant="ghost"
              color="error"
              size="sm"
              icon="i-lucide-trash-2"
              @click="handleDelete(addr.id)"
            >
              Delete
            </UButton>
          </div>
        </template>
      </UCard>
    </div>

    <UEmpty
      v-else
      icon="i-lucide-map-pin"
      title="No addresses"
      description="Add your first shipping address to get started"
      class="py-12"
    >
      <UButton color="primary" size="lg" icon="i-lucide-plus" @click="openModal()">
        Add Address
      </UButton>
    </UEmpty>

    <!-- 统一的 Modal（新增 & 编辑共用） -->
    <UModal v-model:open="modalOpen" title="Address" :description="editingAddress ? 'Edit your shipping address' : 'Add a new shipping address'">
      <template #body>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <UFormField label="First Name" required>
              <UInput v-model="form.first_name" placeholder="Enter first name" size="lg" />
            </UFormField>
            <UFormField label="Last Name" required>
              <UInput v-model="form.last_name" placeholder="Enter last name" size="lg" />
            </UFormField>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <UFormField label="Email" required>
              <UInput v-model="form.email" type="email" placeholder="Enter email" size="lg" />
            </UFormField>
            <UFormField label="Phone" required>
              <UInput v-model="form.phone" type="tel" placeholder="Enter phone" size="lg" />
            </UFormField>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <UFormField label="Country" required>
              <USelect v-model="form.country" :items="countryOptions" placeholder="Select country" size="lg" />
            </UFormField>
            <UFormField label="Province" required>
              <USelect
                v-model="form.province"
                :items="provinceOptions"
                placeholder="Select province"
                size="lg"
                :disabled="!form.country"
              />
            </UFormField>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <UFormField label="City" required>
              <UInput v-model="form.city" placeholder="Enter city" size="lg" />
            </UFormField>
            <UFormField label="Postal Code" required>
              <UInput v-model="form.postal_code" placeholder="Enter postal code" size="lg" />
            </UFormField>
          </div>

          <UFormField label="Region (Optional)">
            <UInput v-model="form.region" placeholder="Enter region" size="lg" />
          </UFormField>

          <UFormField label="Street Address" required>
            <UInput v-model="form.detail_address" placeholder="Enter street address" size="lg" />
          </UFormField>

          <UCheckbox
            :checked="form.default_status === 1"
            @update:checked="(val: boolean) => form.default_status = val ? 1 : 0"
            label="Set as default address"
          />

        <div class="flex justify-end gap-3">
          <UButton variant="ghost" color="neutral" @click="modalOpen = false">
            Cancel
          </UButton>
          <UButton color="primary" :loading="formLoading" @click="handleSave">
            {{ editingAddress ? 'Update' : 'Add' }} Address
          </UButton>
        </div>
      </template>
    </UModal>
  </div>
</template>

<style scoped>
.addresses-page { width: 100%; }
</style>