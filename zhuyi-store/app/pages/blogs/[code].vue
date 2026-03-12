<script setup lang="ts">
import api from '../../api'
import type { DocumentInfoResponse } from '../../api/type'

// 使用默认布局
definePageMeta({
  layout: 'default'
})

const route = useRoute()
const code = route.params.code as string

// 获取文档信息
const { data: info, status, pending } = await useAsyncData<DocumentInfoResponse>(
  `document:${code}`,
  async () => {
    try {
      const res = await api.blogs.document.info(code)
      return res
    } catch (error) {
      console.error('Failed to fetch document:', error)
      throw createError({ statusCode: 404, message: 'Document not found' })
    }
  }
)

// 404 处理
if (!info.value) {
  throw createError({ statusCode: 404, message: 'Page not found' })
}

// 格式化日期
const formatDate = (dateString: string | null) => {
  if (!dateString) return ''
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    })
  } catch (error) {
    return dateString
  }
}

// SEO
useHead({
  title: info.value?.archive?.seo_title || info.value?.document?.title || 'Blog Post',
  meta: [
    {
      name: 'keywords',
      content: info.value?.archive?.seo_keyword || info.value?.document?.keyword || ''
    },
    {
      name: 'description',
      content: info.value?.archive?.seo_description || info.value?.document?.description || ''
    }
  ]
})
</script>

<template>
  <div class="blog-detail-page">
    <!-- 加载状态 -->
    <div v-if="pending" class="flex justify-center items-center py-20">
      <UIcon name="i-lucide-loader-2" class="w-8 h-8 animate-spin text-primary-600" />
    </div>

    <!-- 文档内容 -->
    <div v-else-if="status === 'success' && info" class="max-w-4xl mx-auto">
      <!-- 文章头部 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4 leading-tight">
          {{ info.document?.title }}
        </h1>
        
        <!-- 文章元信息 -->
        <div class="flex items-center gap-4 text-sm text-gray-500 dark:text-gray-400 mb-6">
          <div v-if="info.document?.send_time" class="flex items-center gap-2">
            <UIcon name="i-lucide-calendar" class="w-4 h-4" />
            <span>{{ formatDate(info.document.send_time) }}</span>
          </div>
          <div v-if="info.document?.author" class="flex items-center gap-2">
            <UIcon name="i-lucide-user" class="w-4 h-4" />
            <span>{{ info.document.author }}</span>
          </div>
          <div v-if="info.document?.read_num !== undefined" class="flex items-center gap-2">
            <UIcon name="i-lucide-eye" class="w-4 h-4" />
            <span>{{ info.document.read_num }} views</span>
          </div>
        </div>
      </div>

      <!-- 文章内容卡片 -->
      <UCard class="blog-content-card">
        <div
          class="blog-content prose prose-lg dark:prose-invert max-w-none"
          v-html="info.archive?.cont"
        />
      </UCard>

      <!-- 文章底部信息 -->
      <div v-if="info.document?.source || info.document?.like_num !== undefined" class="mt-6 flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
        <div v-if="info.document?.source" class="flex items-center gap-2">
          <UIcon name="i-lucide-link" class="w-4 h-4" />
          <span>Source: {{ info.document.source }}</span>
        </div>
        <div v-if="info.document?.like_num !== undefined" class="flex items-center gap-2">
          <UIcon name="i-lucide-heart" class="w-4 h-4" />
          <span>{{ info.document.like_num }} likes</span>
        </div>
      </div>
    </div>

    <!-- 错误状态 -->
    <UEmpty
      v-else
      icon="i-lucide-file-question"
      title="Document Not Found"
      description="The requested document could not be found"
    />
  </div>
</template>

<style scoped>
.blog-detail-page {
  width: 100%;
}

.blog-content-card {
  padding: 2rem;
}

/* 文章内容样式 */
.blog-content {
  line-height: 1.8;
  color: #374151;
}

.dark .blog-content {
  color: #e5e7eb;
}

/* 文章内容中的图片 */
.blog-content :deep(img) {
  max-width: 100%;
  height: auto;
  border-radius: 8px;
  margin: 1.5rem 0;
}

/* 文章内容中的标题 */
.blog-content :deep(h1),
.blog-content :deep(h2),
.blog-content :deep(h3),
.blog-content :deep(h4),
.blog-content :deep(h5),
.blog-content :deep(h6) {
  margin-top: 2rem;
  margin-bottom: 1rem;
  font-weight: 600;
  color: #111827;
}

.dark .blog-content :deep(h1),
.dark .blog-content :deep(h2),
.dark .blog-content :deep(h3),
.dark .blog-content :deep(h4),
.dark .blog-content :deep(h5),
.dark .blog-content :deep(h6) {
  color: #f9fafb;
}

/* 文章内容中的段落 */
.blog-content :deep(p) {
  margin-bottom: 1.25rem;
}

/* 文章内容中的链接 */
.blog-content :deep(a) {
  color: #3b82f6;
  text-decoration: underline;
  transition: color 0.2s;
}

.blog-content :deep(a:hover) {
  color: #2563eb;
}

.dark .blog-content :deep(a) {
  color: #60a5fa;
}

.dark .blog-content :deep(a:hover) {
  color: #93c5fd;
}

/* 文章内容中的列表 */
.blog-content :deep(ul),
.blog-content :deep(ol) {
  margin: 1.25rem 0;
  padding-left: 2rem;
}

.blog-content :deep(li) {
  margin-bottom: 0.5rem;
}

/* 文章内容中的引用 */
.blog-content :deep(blockquote) {
  border-left: 4px solid #3b82f6;
  padding-left: 1rem;
  margin: 1.5rem 0;
  font-style: italic;
  color: #6b7280;
}

.dark .blog-content :deep(blockquote) {
  border-left-color: #60a5fa;
  color: #9ca3af;
}

/* 文章内容中的代码 */
.blog-content :deep(code) {
  background-color: #f3f4f6;
  padding: 0.125rem 0.375rem;
  border-radius: 0.25rem;
  font-size: 0.875em;
  font-family: 'Courier New', monospace;
}

.dark .blog-content :deep(code) {
  background-color: #374151;
  color: #f9fafb;
}

.blog-content :deep(pre) {
  background-color: #1f2937;
  color: #f9fafb;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow-x: auto;
  margin: 1.5rem 0;
}

.blog-content :deep(pre code) {
  background-color: transparent;
  padding: 0;
  color: inherit;
}

/* 清除浮动 */
.blog-content::after {
  content: '';
  display: block;
  clear: both;
  visibility: hidden;
  width: 0;
  height: 0;
}

@media (max-width: 768px) {
  .blog-content-card {
    padding: 1.5rem;
  }

  .blog-detail-page h1 {
    font-size: 2rem;
  }
}
</style>