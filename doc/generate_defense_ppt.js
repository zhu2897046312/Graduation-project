const pptxgen = require("pptxgenjs");

const pptx = new pptxgen();
pptx.layout = "LAYOUT_WIDE"; // 13.333 x 7.5
pptx.author = "嘉应学院本科毕业设计";
pptx.company = "嘉应学院";
pptx.subject = "毕业答辩";
pptx.title = "基于Gin与Nuxt.js的跨境电商独立站的设计与实现";
pptx.lang = "zh-CN";

const C = {
  navy: "1E2761",
  light: "F4F7FF",
  text: "1F2937",
  sub: "4B5563",
  accent: "3B82F6",
  ok: "16A34A",
  warn: "EA580C",
  white: "FFFFFF",
};

const FONT_TITLE = "Microsoft YaHei";
const FONT_BODY = "Calibri";

function bg(slide, dark = false) {
  slide.background = { color: dark ? C.navy : C.white };
  slide.addShape(pptx.ShapeType.rect, {
    x: 0,
    y: 0,
    w: 13.33,
    h: 0.26,
    fill: { color: dark ? C.accent : C.light },
    line: { color: dark ? C.accent : C.light },
  });
}

function title(slide, t, dark = false) {
  slide.addText(t, {
    x: 0.6,
    y: 0.45,
    w: 12.2,
    h: 0.55,
    fontFace: FONT_TITLE,
    fontSize: 30,
    bold: true,
    color: dark ? C.white : C.navy,
  });
}

function bullets(slide, items) {
  let y = 1.35;
  items.forEach((it) => {
    slide.addShape(pptx.ShapeType.roundRect, {
      x: 0.75,
      y: y + 0.08,
      w: 0.16,
      h: 0.16,
      fill: { color: C.accent },
      line: { color: C.accent },
      radius: 0.03,
    });
    slide.addText(it, {
      x: 1.02,
      y,
      w: 11.7,
      h: 0.45,
      fontFace: FONT_BODY,
      fontSize: 19,
      color: C.text,
    });
    y += 0.62;
  });
}

// 1 封面
{
  const s = pptx.addSlide();
  bg(s, true);
  s.addText("毕业答辩", {
    x: 0.65, y: 0.95, w: 3.2, h: 0.7,
    fontFace: FONT_BODY, fontSize: 24, color: "BFD7FF", bold: true,
  });
  s.addText("基于 Gin 与 Nuxt.js 的\n跨境电商独立站的设计与实现", {
    x: 0.65, y: 1.85, w: 8.6, h: 2.2,
    fontFace: FONT_TITLE, fontSize: 42, bold: true, color: C.white, breakLine: true,
  });
  s.addShape(pptx.ShapeType.roundRect, {
    x: 9.25, y: 1.2, w: 3.5, h: 4.8,
    fill: { color: "2A367A", transparency: 10 },
    line: { color: "6EA8FF", transparency: 20 },
    radius: 0.15,
  });
  s.addText("答辩人：XXX\n专业：计算机科学与技术\n指导教师：XXX\n日期：2026年X月", {
    x: 9.55, y: 1.65, w: 2.9, h: 3.8,
    fontFace: FONT_BODY, fontSize: 17, color: C.white, valign: "top",
  });
}

// 2 背景与意义
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "1. 研究背景与意义");
  bullets(s, [
    "第三方平台存在流量成本高、数据归属弱、品牌自主性不足等问题",
    "跨境场景下，独立站可形成“自有域名 + 数据闭环 + 品牌沉淀”",
    "本课题围绕独立站完整业务链路，验证前后端分离架构落地可行性",
    "研究价值：兼顾工程实践与业务可用，具备可部署、可演示特性",
  ]);
}

// 3 目标与贡献
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "2. 研究目标与论文贡献");
  s.addShape(pptx.ShapeType.roundRect, {
    x: 0.75, y: 1.35, w: 12, h: 5.4, radius: 0.08,
    fill: { color: C.light }, line: { color: "D9E4FF" },
  });
  s.addText("目标", {
    x: 1.05, y: 1.65, w: 1.2, h: 0.4, fontFace: FONT_TITLE, fontSize: 22, bold: true, color: C.navy,
  });
  s.addText("构建一套可部署、可演示、可扩展的跨境电商独立站系统", {
    x: 2.05, y: 1.67, w: 10, h: 0.4, fontFace: FONT_BODY, fontSize: 18, color: C.text,
  });
  s.addText("贡献", {
    x: 1.05, y: 2.45, w: 1.2, h: 0.4, fontFace: FONT_TITLE, fontSize: 22, bold: true, color: C.navy,
  });
  bullets(s, [
    "形成 server + shop_admin + zhuyi-store 三子系统协同架构",
    "后端采用 Handler → Service → Repository 分层与工厂组织方式",
    "实现 Session/JWT/设备指纹的多端认证策略与 PayPal 支付集成",
    "完成需求、设计、实现、测试的完整软件工程流程实践",
  ]);
}

// 4 技术栈
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "3. 技术选型与可行性");
  const cards = [
    ["后端", "Go + Gin + GORM\nMySQL + Redis"],
    ["管理后台", "Vue3 + Vite + TS\nAnt Design Vue + Pinia"],
    ["前台商城", "Nuxt4 + Vue3 + Nuxt UI\nTailwind CSS"],
    ["工程支撑", "RESTful API\nDocker Compose\nPayPal Sandbox"],
  ];
  let x = 0.8;
  cards.forEach((c, i) => {
    const y = i < 2 ? 1.5 : 4.0;
    const xx = i % 2 === 0 ? 0.8 : 6.9;
    s.addShape(pptx.ShapeType.roundRect, {
      x: xx, y, w: 5.6, h: 2.05, radius: 0.08,
      fill: { color: "F8FAFF" }, line: { color: "DCE7FF" },
    });
    s.addText(c[0], {
      x: xx + 0.35, y: y + 0.25, w: 2.2, h: 0.4,
      fontFace: FONT_TITLE, fontSize: 21, bold: true, color: C.navy,
    });
    s.addText(c[1], {
      x: xx + 0.35, y: y + 0.82, w: 4.8, h: 0.95,
      fontFace: FONT_BODY, fontSize: 17, color: C.text,
    });
  });
}

// 5 总体架构
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "4. 系统总体架构");
  const boxes = [
    [0.9, 1.8, 3.6, 1.1, "前台商城\nzhuyi-store"],
    [4.9, 1.8, 3.6, 1.1, "后端 API\nserver"],
    [8.9, 1.8, 3.6, 1.1, "管理后台\nshop_admin"],
    [2.3, 4.0, 3.6, 1.1, "MySQL\n业务数据"],
    [6.9, 4.0, 2.8, 1.1, "Redis\nSession/缓存"],
  ];
  boxes.forEach((b) => {
    s.addShape(pptx.ShapeType.roundRect, {
      x: b[0], y: b[1], w: b[2], h: b[3], radius: 0.08,
      fill: { color: "F4F8FF" }, line: { color: "BFD2FF", pt: 1.5 },
    });
    s.addText(b[4], {
      x: b[0], y: b[1] + 0.2, w: b[2], h: 0.8, align: "center",
      fontFace: FONT_BODY, fontSize: 17, color: C.text, bold: true,
    });
  });
  s.addShape(pptx.ShapeType.line, { x: 4.5, y: 2.35, w: 0.35, h: 0, line: { color: C.accent, pt: 2, beginArrowType: "none", endArrowType: "triangle" } });
  s.addShape(pptx.ShapeType.line, { x: 8.5, y: 2.35, w: 0.35, h: 0, line: { color: C.accent, pt: 2, beginArrowType: "none", endArrowType: "triangle" } });
  s.addShape(pptx.ShapeType.line, { x: 6.7, y: 2.95, w: -2.4, h: 1.0, line: { color: C.accent, pt: 2, beginArrowType: "none", endArrowType: "triangle" } });
  s.addShape(pptx.ShapeType.line, { x: 6.7, y: 2.95, w: 1.6, h: 1.0, line: { color: C.accent, pt: 2, beginArrowType: "none", endArrowType: "triangle" } });
  s.addText("API 分组：/api/manage  与  /api/client", {
    x: 0.9, y: 5.65, w: 8.5, h: 0.35, fontFace: FONT_BODY, fontSize: 16, color: C.sub,
  });
  s.addText("认证：管理员 Session / 用户 JWT / 设备指纹", {
    x: 0.9, y: 6.0, w: 8.5, h: 0.35, fontFace: FONT_BODY, fontSize: 16, color: C.sub,
  });
}

// 6 业务流程
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "5. 核心业务流程（用户侧）");
  const steps = ["商品浏览", "选择SKU加购", "购物车管理", "结算下单", "PayPal支付", "订单查询"];
  let x = 0.65;
  steps.forEach((st, i) => {
    s.addShape(pptx.ShapeType.roundRect, {
      x, y: 2.7, w: 1.9, h: 1.1, radius: 0.08,
      fill: { color: i % 2 === 0 ? "E8F1FF" : "F2F7FF" }, line: { color: "AFC9FF" },
    });
    s.addText(`${i + 1}\n${st}`, {
      x, y: 2.88, w: 1.9, h: 0.8, align: "center",
      fontFace: FONT_BODY, fontSize: 15, bold: true, color: C.text,
    });
    if (i < steps.length - 1) {
      s.addShape(pptx.ShapeType.chevron, {
        x: x + 1.96, y: 3.03, w: 0.35, h: 0.45,
        fill: { color: C.accent }, line: { color: C.accent },
      });
    }
    x += 2.1;
  });
  s.addText("游客链路：设备指纹 + 查询码 支持加购与查单", {
    x: 0.8, y: 5.5, w: 12.0, h: 0.45, fontFace: FONT_BODY, fontSize: 18, color: C.navy, bold: true,
  });
}

// 7 数据模型
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "6. 数据库与关键数据模型");
  const groups = [
    ["商品域", "分类 / 商品 / SKU / 属性 / 标签"],
    ["交易域", "订单 / 订单项 / 收货地址 / 退款"],
    ["内容域", "文档 / 推荐位 / 推荐项"],
  ];
  groups.forEach((g, i) => {
    const y = 1.5 + i * 1.65;
    s.addShape(pptx.ShapeType.roundRect, {
      x: 1.0, y, w: 11.2, h: 1.2, radius: 0.07,
      fill: { color: i === 1 ? "EEF6FF" : "F8FAFF" }, line: { color: "CFDFFF" },
    });
    s.addText(g[0], {
      x: 1.35, y: y + 0.35, w: 2.0, h: 0.4, fontFace: FONT_TITLE, fontSize: 22, bold: true, color: C.navy,
    });
    s.addText(g[1], {
      x: 3.2, y: y + 0.38, w: 8.6, h: 0.45, fontFace: FONT_BODY, fontSize: 18, color: C.text,
    });
  });
  s.addText("设计原则：状态字段驱动流程、主外键关系清晰、便于扩展", {
    x: 1.0, y: 6.3, w: 11.6, h: 0.35, fontFace: FONT_BODY, fontSize: 16, color: C.sub,
  });
}

// 8 后端设计
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "7. 后端关键设计");
  bullets(s, [
    "分层架构：Handler → Service → Repository，职责清晰、可测试",
    "Factory 统一依赖注入，降低模块耦合，便于替换与扩展",
    "中间件体系：CORS、管理员认证、JWT、设备指纹、统一错误处理",
    "核心能力：商品/订单/CMS/支付模块按业务域解耦组织",
    "统一响应结构提升前后端联调效率与可维护性",
  ]);
}

// 9 前后端实现
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "8. 前后端关键实现");
  s.addShape(pptx.ShapeType.roundRect, {
    x: 0.9, y: 1.5, w: 5.9, h: 4.8, radius: 0.08,
    fill: { color: "F5F9FF" }, line: { color: "CFE1FF" },
  });
  s.addShape(pptx.ShapeType.roundRect, {
    x: 6.55, y: 1.5, w: 5.9, h: 4.8, radius: 0.08,
    fill: { color: "F8FBFF" }, line: { color: "D9E7FF" },
  });
  s.addText("管理后台（shop_admin）", {
    x: 1.2, y: 1.85, w: 5.2, h: 0.4, fontFace: FONT_TITLE, fontSize: 21, bold: true, color: C.navy,
  });
  s.addText("• 商品/分类/SKU管理\n• 订单处理与退款\n• CMS文档与推荐位\n• 系统与站点配置", {
    x: 1.2, y: 2.35, w: 5.3, h: 2.3, fontFace: FONT_BODY, fontSize: 17, color: C.text,
  });
  s.addText("前台商城（zhuyi-store）", {
    x: 6.9, y: 1.85, w: 5.2, h: 0.4, fontFace: FONT_TITLE, fontSize: 21, bold: true, color: C.navy,
  });
  s.addText("• 首页/分类/详情\n• 购物车与结算\n• PayPal支付与订单查询\n• 登录注册与地址管理", {
    x: 6.9, y: 2.35, w: 5.2, h: 2.3, fontFace: FONT_BODY, fontSize: 17, color: C.text,
  });
  s.addText("前端通过统一 API 封装与状态管理实现高效联调", {
    x: 0.95, y: 6.45, w: 11.8, h: 0.35, fontFace: FONT_BODY, fontSize: 16, color: C.sub,
  });
}

// 10 测试结果
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "9. 测试方案与结果");
  const rows = [
    ["管理端功能", "通过", C.ok],
    ["客户端功能", "通过", C.ok],
    ["接口认证逻辑", "通过", C.ok],
    ["PayPal沙箱支付/退款", "通过", C.ok],
    ["性能压测与渗透测试", "待完善", C.warn],
  ];
  s.addText("测试环境：Go + MySQL + Redis + Node + PayPal Sandbox", {
    x: 0.85, y: 1.3, w: 11.8, h: 0.4, fontFace: FONT_BODY, fontSize: 17, color: C.sub,
  });
  let y = 1.9;
  rows.forEach((r) => {
    s.addShape(pptx.ShapeType.roundRect, {
      x: 0.85, y, w: 11.8, h: 0.7, radius: 0.05,
      fill: { color: "F9FBFF" }, line: { color: "E0EAFF" },
    });
    s.addText(r[0], {
      x: 1.15, y: y + 0.18, w: 8.3, h: 0.35, fontFace: FONT_BODY, fontSize: 17, color: C.text,
    });
    s.addText(r[1], {
      x: 10.1, y: y + 0.16, w: 1.9, h: 0.35, align: "center",
      fontFace: FONT_BODY, fontSize: 16, bold: true, color: r[2],
    });
    y += 0.84;
  });
}

// 11 亮点
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "10. 项目亮点与创新点");
  bullets(s, [
    "双端协同：管理后台与前台商城共同支撑完整电商闭环",
    "认证分层：管理员 Session、用户 JWT、游客设备指纹并行设计",
    "支付闭环：PayPal 创建订单、捕获支付与退款链路打通",
    "工程化：从需求到测试的完整论文与系统交付",
  ]);
}

// 12 不足与改进
{
  const s = pptx.addSlide();
  bg(s);
  title(s, "11. 不足与改进方向");
  s.addShape(pptx.ShapeType.roundRect, {
    x: 0.9, y: 1.45, w: 5.7, h: 4.9, radius: 0.08,
    fill: { color: "FFF8F2" }, line: { color: "FED7AA" },
  });
  s.addShape(pptx.ShapeType.roundRect, {
    x: 6.75, y: 1.45, w: 5.7, h: 4.9, radius: 0.08,
    fill: { color: "F0FDF4" }, line: { color: "BBF7D0" },
  });
  s.addText("当前不足", {
    x: 1.2, y: 1.8, w: 2, h: 0.4, fontFace: FONT_TITLE, fontSize: 22, bold: true, color: "9A3412",
  });
  s.addText("• 缺少系统化压测与高并发验证\n• 安全审计与渗透测试不足\n• 自动化测试覆盖较低", {
    x: 1.2, y: 2.35, w: 5.0, h: 2.0, fontFace: FONT_BODY, fontSize: 17, color: C.text,
  });
  s.addText("后续优化", {
    x: 7.05, y: 1.8, w: 2, h: 0.4, fontFace: FONT_TITLE, fontSize: 22, bold: true, color: "166534",
  });
  s.addText("• 增加缓存策略与索引优化\n• 建立CI/CD与回归测试链路\n• 强化日志审计、限流与风控能力", {
    x: 7.05, y: 2.35, w: 5.0, h: 2.0, fontFace: FONT_BODY, fontSize: 17, color: C.text,
  });
}

// 13 总结与致谢
{
  const s = pptx.addSlide();
  bg(s, true);
  title(s, "12. 总结与致谢", true);
  s.addText("完成了跨境电商独立站从需求到测试的全流程实践", {
    x: 0.9, y: 2.15, w: 11.5, h: 0.55, fontFace: FONT_BODY, fontSize: 24, color: C.white, align: "center",
  });
  s.addText("系统已具备可部署、可演示能力，并具备后续扩展空间", {
    x: 0.9, y: 2.85, w: 11.5, h: 0.55, fontFace: FONT_BODY, fontSize: 22, color: "DCE8FF", align: "center",
  });
  s.addText("感谢各位老师聆听，恳请批评指正", {
    x: 0.9, y: 4.6, w: 11.5, h: 0.65, fontFace: FONT_TITLE, fontSize: 30, bold: true, color: "FFFFFF", align: "center",
  });
}

pptx.writeFile({
  fileName: "E:/WorkSpace/Graduation-project/doc/毕业答辩PPT-跨境电商独立站-7分钟.pptx",
});
