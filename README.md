# bootstra/log-logrus

bottstrap 系列 log-logrus 日志适配器

## 使用说明

1. 通过 `import (_ "bootstrap/log-logrus/logrus)"` 导入适配器
2. 通过 `import "bootstrap/global/log"` 导入门面
3. 通过 `err := log.InitGlobalLogger("logrus")` 进行初始化
4. 代码中调用 `log.Info("message xxx", "foo", "boo")`

## 开发说明

bootstrap 系列是为了提供统一的门面和依赖管理而产生的，期望我们在项目中可以屏蔽底层的实现，可以通过切换适配器快速接入、切换到其他实现。

### TODO

- [ ] format 配置支持
- [ ] 各适配器输出格式统一