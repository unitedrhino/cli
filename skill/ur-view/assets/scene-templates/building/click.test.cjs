/* 文件说明：验证真实鼠标的单击抖动容差、拖动防误选、取消手势和同层选区不重置镜头。
 * 运行：NODE_PATH=<已有 Playwright 的 node_modules> node click.test.cjs
 * 从本地源码直接打开，无服务端依赖；所有测试浏览器在 finally 中关闭。
 */
const { chromium } = require('playwright'); // 复用工作区浏览器依赖。
const { pathToFileURL } = require('node:url'); // 为本地页面生成合法 file URL。
const path = require('node:path'); // 测试只访问当前模板源码。
const assert = require('node:assert/strict'); // 按真实交互结果断言。

/** ready 等待模型与镜头稳定，避免动画未完成时测到不同像素位置。 */
async function ready(page) {
  await page.waitForFunction(() => window.BuildingApp?.scene?.metrics.ready && !BuildingApp.scene.inspect().animating);
}
/** clearSelection 选中三层并清空区域，然后返回避开区域标签的真实画布坐标。 */
async function clearSelection(page) {
  await page.click('[data-floor="3F"]');
  await ready(page);
  return page.evaluate(() => BuildingApp.scene.projectZone('3F/level-3-b', .3));
}
/** run 覆盖零位移、轻微抖动和拖出后返回；跨尺寸复核 CSS 像素容差。 */
async function run() {
  const browser = await chromium.launch({ headless: true, args: ['--no-sandbox', '--enable-unsafe-swiftshader', '--use-angle=swiftshader'] });
  try {
    for (const width of [1920, 1239]) {
      const page = await browser.newPage({ viewport: { width, height: 1080 }, deviceScaleFactor: width === 1920 ? 1 : 2 });
      const errors = []; // 不允许输入监听引入页面异常。
      page.on('pageerror', error => errors.push(error.message));
      await page.goto(pathToFileURL(path.join(__dirname, 'index.html')).href);
      await ready(page);
      // 本用例专测画布射线与拖动，区域 DOM 标签的单击另由页面回归覆盖。
      await page.addStyleTag({ content: '.zone-label { pointer-events: none !important; }' });
      for (const offset of [0, 1, 4]) {
        const point = await clearSelection(page);
        await page.mouse.move(point.x, point.y);
        await page.mouse.down();
        await page.mouse.move(point.x + offset, point.y + offset);
        await page.mouse.up();
        assert.equal(await page.evaluate(() => BuildingApp.inspect().zone), '3F/level-3-b', width + 'px / 抖动 ' + offset);
        assert.equal(await page.evaluate(() => BuildingApp.scene.inspect().animating), false, '同层选区不重置镜头');
      }
      const point = await clearSelection(page); // 即便拖回按下点，也必须保留曾经越界的状态。
      await page.mouse.move(point.x, point.y); await page.mouse.down();
      await page.mouse.move(point.x + 25, point.y + 10);
      await page.mouse.move(point.x, point.y); await page.mouse.up();
      assert.equal(await page.evaluate(() => BuildingApp.inspect().zone), null, '拖动返回起点不得误选');
      const next = await clearSelection(page);
      await page.mouse.move(next.x, next.y); await page.mouse.down();
      await page.evaluate(() => document.dispatchEvent(new PointerEvent('pointercancel', { bubbles: true })));
      await page.mouse.up();
      assert.equal(await page.evaluate(() => BuildingApp.inspect().zone), null, '取消手势不得选中');
      await page.click('#reset'); await ready(page);
      // 单击楼层导航立即选中并展开，不需要双击或额外按钮。
      await page.click('[data-floor="3F"]');
      assert.equal(await page.evaluate(() => BuildingApp.inspect().expanded), true);
      assert.equal(await page.evaluate(() => BuildingApp.inspect().floor), '3F');
      // 淡化楼层的外露楼板可直接切换；重叠处仍优先当前楼层区域。
      await page.click('[data-floor="7F"]'); await ready(page);
      for (const zone of ['8F/level-8-b', '7F/level-7-b', '6F/level-6-b']) {
        const target = await page.evaluate(id => BuildingApp.scene.projectZone(id, .3), zone);
        await page.mouse.click(target.x, target.y); await ready(page);
        assert.equal(await page.evaluate(() => BuildingApp.inspect().floor), zone.split('/')[0], '点击淡化楼板切层：' + zone);
        assert.equal(await page.evaluate(() => BuildingApp.inspect().zone), null, '切层不直接选中其他层区域');
      }
      await clearSelection(page);
      const active = await page.evaluate(() => BuildingApp.scene.projectZone('3F/level-3-b', .3));
      await page.mouse.click(active.x, active.y); await ready(page);
      assert.equal(await page.evaluate(() => BuildingApp.inspect().zone), '3F/level-3-b', '淡化楼板不拦截当前区域');
      assert.deepEqual(errors, []);
      await page.close();
      console.log(width + ' 宽度通过：单击、1/4px 抖动、旋转防误选、取消手势、选中即展开、同层镜头稳定、淡化楼层切换');
    }
  } finally { await browser.close(); }
}
run().catch(error => { console.error(error); process.exitCode = 1; });
