# TypeScript 迁移指南

## 已完成的工作

### ✅ 1. 安装 TypeScript 依赖
已安装以下依赖包：
- `typescript` - TypeScript 编译器
- `@types/node` - Node.js 类型定义
- `vue-tsc` - Vue 3 TypeScript 类型检查工具
- `@vue/tsconfig` - Vue 官方 TypeScript 配置
- `vite-plugin-checker` - Vite TypeScript 检查插件

### ✅ 2. 创建 TypeScript 配置文件
- `tsconfig.json` - 主要的 TypeScript 配置
- `tsconfig.node.json` - Node.js 环境（Vite 配置文件等）的 TypeScript 配置

### ✅ 3. 迁移核心文件
- ✅ `vite.config.js` → `vite.config.ts` (已完成并修复)
- ✅ `src/main.js` → `src/main.ts` (已完成)
- ✅ `index.html` 中的引用已更新为 `main.ts`

### ✅ 4. 创建类型声明文件
- `src/vite-env.d.ts` - Vite 环境类型声明
- `src/env.d.ts` - 环境变量类型定义
- `src/types.d.ts` - 全局类型声明（包括第三方库、资源文件等）

### ✅ 5. 更新 package.json
添加了新的脚本命令：
- `pnpm type-check` - 执行 TypeScript 类型检查
- `pnpm build` - 构建前先进行类型检查

## 当前状态

项目已成功从 JavaScript 迁移到 TypeScript。核心配置文件和入口文件都已转换。

## 下一步建议

### 📝 逐步迁移策略

由于这是一个大型项目，建议采用**渐进式迁移**策略：

#### 阶段 1: 核心工具类（优先级：高）
```bash
src/utils/
├── request.ts        # API 请求封装
├── format.ts         # 数据格式化
├── dictionary.ts     # 字典工具
└── asyncRouter.ts    # 路由工具
```

#### 阶段 2: API 接口层（优先级：高）
```bash
src/api/
├── user.ts          # 用户相关接口
├── system.ts        # 系统接口
└── ...              # 其他 API 文件
```

#### 阶段 3: 状态管理（优先级：中）
```bash
src/pinia/modules/
├── user.ts
├── router.ts
└── ...
```

#### 阶段 4: 组件和视图（优先级：低）
```bash
src/components/
src/view/
```

### 🔧 迁移单个文件的步骤

1. **重命名文件**：将 `.js` 改为 `.ts`
2. **添加类型注解**：
   - 函数参数和返回值
   - 变量声明
   - 接口定义
3. **修复类型错误**：根据 IDE 提示修复错误
4. **测试功能**：确保功能正常

### 📋 创建类型定义文件

建议创建以下类型定义文件：

```typescript
// src/types/api.d.ts - API 响应类型
export interface ApiResponse<T = any> {
  code: number
  data: T
  msg: string
}

// src/types/user.d.ts - 用户相关类型
export interface UserInfo {
  id: number
  username: string
  // ...
}

// src/types/router.d.ts - 路由类型
export interface RouteMetaInfo {
  title: string
  icon?: string
  // ...
}
```

### ⚙️ 配置选项

当前 TypeScript 配置较为严格，如果遇到太多错误，可以临时调整 `tsconfig.json`：

```json
{
  "compilerOptions": {
    "strict": false,           // 关闭严格模式
    "noUnusedLocals": false,   // 允许未使用的局部变量
    "noUnusedParameters": false // 允许未使用的参数
  }
}
```

待迁移稳定后，再逐步开启严格检查。

### 🚀 运行和测试

- **开发模式**：`pnpm dev`
- **类型检查**：`pnpm type-check`
- **生产构建**：`pnpm build`

## 常见问题

### Q1: Vue 文件中如何使用 TypeScript？
```vue
<script setup lang="ts">
import { ref } from 'vue'

const count = ref<number>(0)
const handleClick = (): void => {
  count.value++
}
</script>
```

### Q2: 如何处理第三方库没有类型定义？
在 `src/types.d.ts` 中添加：
```typescript
declare module 'some-package'
```

### Q3: 如何定义组件 Props？
```vue
<script setup lang="ts">
interface Props {
  title: string
  count?: number
}

const props = defineProps<Props>()
</script>
```

## 迁移进度跟踪

- [x] 安装依赖
- [x] 配置文件
- [x] 核心入口文件
- [x] 类型声明
- [ ] utils 工具类
- [ ] API 接口层
- [ ] Pinia 状态管理
- [ ] 路由配置
- [ ] 组件迁移
- [ ] 视图页面迁移

## 参考资源

- [Vue 3 + TypeScript 官方文档](https://cn.vuejs.org/guide/typescript/overview.html)
- [Vite TypeScript 指南](https://cn.vitejs.dev/guide/features.html#typescript)
- [Element Plus TypeScript 支持](https://element-plus.org/zh-CN/guide/typescript.html)

---

**注意**：TypeScript 迁移是一个渐进的过程，不需要一次性完成。建议每次迁移一小部分，确保功能正常后再继续下一部分。
