declare module '*.scss' {
  const content: { [className: string]: string }
  export default content
}

declare module '*.css' {
  const content: { [className: string]: string }
  export default content
}

declare module '*.svg' {
  const content: any
  export default content
}

declare module '*.png' {
  const content: any
  export default content
}

declare module '*.jpg' {
  const content: any
  export default content
}

declare module '*.jpeg' {
  const content: any
  export default content
}

declare module '*.gif' {
  const content: any
  export default content
}

declare module '*.webp' {
  const content: any
  export default content
}

declare module '*.ico' {
  const content: any
  export default content
}

declare module '*.json' {
  const content: any
  export default content
}

// Element Plus 类型增强
declare module '@vue/runtime-core' {
  export interface GlobalComponents {
    ElButton: typeof import('element-plus')['ElButton']
    ElInput: typeof import('element-plus')['ElInput']
    ElForm: typeof import('element-plus')['ElForm']
    ElFormItem: typeof import('element-plus')['ElFormItem']
    ElMessage: typeof import('element-plus')['ElMessage']
    ElMessageBox: typeof import('element-plus')['ElMessageBox']
    // 根据需要添加更多组件...
  }
}

// 第三方库类型声明
declare module '@wangeditor/editor-for-vue'
declare module '@form-create/element-ui'
declare module '@form-create/designer'
declare module 'vform3-builds'
declare module 'vue-qr'
declare module 'vue-cropper'
declare module 'vue3-ace-editor'
declare module 'vue3-sfc-loader'
declare module 'vuedraggable'
declare module 'sortablejs'
declare module 'screenfull'
declare module 'spark-md5'
declare module 'marked'
declare module 'marked-highlight'
declare module 'nprogress'
declare module 'mitt'
declare module 'vite-auto-import-svg'

export { }
