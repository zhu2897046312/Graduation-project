<template>
  <BasicModal title="新增标签" @on-ok="handleSubmit" @on-close="handleClose" ref="modalRef">
    <Form ref="formRef"  v-if="formState" v-bind="formProp" :model="formState" :rules="rules" :label-col="{ flex: '160px' }"
      :wrapper-col="{ span: 24 }">
      <Form.Item name="title" label="名称">
        <Input.Group compact>
          <Input v-model:value="formState.title"  />
        </Input.Group>
      </Form.Item>
      <Form.Item label="编码" name="code">
        <Input v-model:value="formState.code"  />
      </Form.Item>
      <Form.Item label="匹配词，空格分割" name="match_word">
        <Input.TextArea v-model:value="formState.match_word"  />
      </Form.Item>
      <Form.Item label="状态" name="state" >
        <Radio.Group v-model:value="formState.state" button-style="solid">
          <Radio.Button :value="1">发布</Radio.Button>
          <Radio.Button :value="2">不发布</Radio.Button>
        </Radio.Group>
      </Form.Item>
      <Form.Item label="缩略图" name="thumb">
        <UploadImage v-model:value="formState.thumb" />
      </Form.Item>
      <Divider>SEO设置</Divider>
      <Form.Item label="关键词" name="seo_title">
        <Input v-model:value="formState.seo_title" :maxlength="200" />
      </Form.Item>
      <Form.Item label="关键词" name="seo_keyword">
        <Input v-model:value="formState.seo_keyword" :maxlength="200" />
      </Form.Item>
      <Form.Item label="描述" name="seo_description">
        <Input.TextArea :rows="3" v-model:value="formState.seo_description" :maxlength="200" />
      </Form.Item>

      <Divider>其他设置</Divider>
      <Form.Item label="阅读次数" name="read_num">
        <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="formState.read_num" />
      </Form.Item>
      <Form.Item label="排序" name="sort_num">
        <InputNumber style="width: 200px;" :min="0" :max="999999" v-model:value="formState.sort_num" />
      </Form.Item>
    </Form>
  </BasicModal>
</template>


<script lang="ts" setup>
import { ref } from 'vue';
import { Form, Input, message, Radio, Divider, InputNumber } from 'ant-design-vue';
import type { FormProps } from 'ant-design-vue';
import type { Rule } from 'ant-design-vue/es/form';
import { loadingTask } from '/@/utils/helper';
import BasicModal from '/@/components/Kernel/BasicModal/index.vue';
import UploadImage from '/@/components/Kernel/YexUpload/UploadImage.vue';
import api from '/@/api/index'


const defaultFromValues = {
  title: '',
  code: '',
  match_word: '',
  thumb: '',
  seo_title: '',
  seo_keyword: '',
  seo_description: '',
  state: 1,
  read_num: '',
  sort_num: '',
}


const modalRef = ref<any>(null);
const formState = ref<any>(null);
const formRef = ref<any>(null);
const emit = defineEmits(['on-change']);
const rules: Record<string, Rule[]> = {
  title: [
    { required: true },
  ],
  code: [
    { required: true },
  ],
};
const formProp = ref<FormProps | any>({
  labelCol: { flex: '160px' },
  wrapperCol: { span: 20 },
  labelAlign: 'right',
})

const handleOpen = (e: any) => {
  loadingTask(async () => {
    formState.value = {...defaultFromValues}
    modalRef.value && modalRef.value.useOpen();
  })
}

const handleClose = () => {
  formState.value = null
}

const handleSubmit = async () => {
  const values: any = await formRef.value?.validateFields();
  console.debug(values)
  loadingTask(async () => {
    await api.cms.tag.curdApi('CREATE', values);
    message.success('操作成功');
    modalRef.value && modalRef.value.useClose();
    emit('on-change');
  })
}


defineExpose({
  useOpen: handleOpen,
  useClose: handleClose,
})

</script>