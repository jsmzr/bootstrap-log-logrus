# bootstra/log-logrus

bottstrap 系列 log-logrus 日志适配器

## 使用说明

1. 通过 `import (_ "github.com/jsmzr/bootstrap-log-logrus/logrus)"` 导入适配器
2. 通过 `import "github.com/jsmzr/bootstrap-global/log"` 导入门面
4. 通过 `err := log.InitLogger("logrus")` 进行初始化
3. 代码引入`import "github.com/jsmzr/bootstrap-log/log` 调用 `log.Info("message %s", key)` 即可

## 开发说明

bootstrap 系列是为了提供统一的门面和依赖管理而产生的，期望我们在项目中可以屏蔽底层的实现，可以通过切换适配器快速接入、切换到其他实现。

### TODO

- [ ] format 配置支持
- [ ] 各适配器输出格式统一