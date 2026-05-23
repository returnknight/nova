# Ideas
## WIP
- 基础体验
  - 完善的 Skills 系统
  - /brainstorm /spec /plan /execute /review
  - 章节细纲，从大纲到->章节细纲->章节草稿->章节定稿
- Diff View Accept/Reject: 可以在diff view中点击accept/reject按钮，确认或拒绝当前diff
- Agent能力
  - 支持完善的上下文管理，memory，上下文管理相关代码抽出单独的package，方便统一迭代和管理
  - 考虑：自定义Agent
- ide 日志
- Tantivy / MeiliSearch 全局搜索
- 重构：
  - Piece Tree 文本内核
  - 考虑分支一个 Code Mirror6 编辑器，实验效果
- 用户指引：可以任意提示agent创建新文件，修改文件内容，删除文件等，比如让他跟你多次脑暴，生成两个版本的不同章节等，可以先规划，满意再动笔
- 所有长段后端写死的提示词，集中到一个package下管理，优化整体代码package结构
- 调试模式，开启后可以看到context组成

## NEED FIX
- 不是新对话就一定要重新脑暴，根据当前是否存在大纲来判断是否需要脑暴
- 

# 规划
- 多语言支持
- 剧情分支系统，允许从特定节点开始，分出不同的剧情线延续，允许对比不同的分支然后选择一个合并
- 伏笔能力
- 人物状态时间线
- 版本管理：不用git，自己实现
- 支持导入小说
- 互动创作模式，类似AI Dungeon / 酒馆
- prompt 高级自定义
- 支持在diff view中点击accept/reject按钮，确认或拒绝当前diff
- 角色卡片，支持角色引用
