/* 文件说明：浏览器验收通用场景、空状态、拾取、巡航、趋势及窄窗口。
 * 仅本测试注入受控绑定夹具，不访问外部环境；运行产物写入 .temp，不进入 ZIP。
 */
const { chromium } = require('playwright');
const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const http = require('node:http');
/** 等待面板过渡结束后核对整个矩形在视口内，避免保存动画中途的截断截图。 */
async function waitForPanel(page) {
  const bounds = await page.locator('#info-panel').evaluate(async element => {
    element.getBoundingClientRect(); // 刷新样式，使 CSS 入场过渡进入动画列表。
    await Promise.all(element.getAnimations().map(animation => animation.finished));
    await new Promise(resolve => requestAnimationFrame(resolve));
    const rect = element.getBoundingClientRect(); // 稳定后的面板完整边界。
    return { left: rect.left, top: rect.top, right: rect.right, bottom: rect.bottom, width: rect.width, height: rect.height, viewportWidth: innerWidth, viewportHeight: innerHeight };
  });
  assert(bounds.width > 0 && bounds.height > 0 && bounds.left >= 0 && bounds.top >= 0 && bounds.right <= bounds.viewportWidth && bounds.bottom <= bounds.viewportHeight, '信息面板必须完整位于视口内');
}
/** 用独立本地静态服务验收，结束后释放服务与浏览器。 */
async function run() {
  const root = __dirname, output = path.join(root, '.temp', 'verification'); // 本模板及产物目录。
  fs.mkdirSync(output, { recursive: true });
  const server = http.createServer((request, response) => {
    const file = path.resolve(root, '.' + (request.url.split('?')[0] === '/' ? '/index.html' : request.url.split('?')[0])); // 限定请求在模板目录。
    if (!file.startsWith(root + path.sep)) { response.writeHead(403); response.end(); return; }
    fs.readFile(file, (error, data) => {
      if (error) { response.writeHead(404); response.end(); return; }
      response.setHeader('Content-Type', ({ '.js': 'application/javascript', '.css': 'text/css', '.html': 'text/html' })[path.extname(file)] || 'application/octet-stream');
      response.end(data);
    });
  });
  await new Promise(resolve => server.listen(0, '127.0.0.1', resolve));
  const browser = await chromium.launch({ headless: true, args: ['--no-sandbox', '--enable-unsafe-swiftshader', '--use-angle=swiftshader'] }); // 软件渲染只用于功能验收。
  try {
    const page = await browser.newPage({ viewport: { width: 1920, height: 1080 } }), errors = []; // 桌面验收窗口。
    page.on('pageerror', error => errors.push(error.message));
    await page.goto('http://127.0.0.1:' + server.address().port + '/');
    await page.waitForFunction(() => document.getElementById('viewport').dataset.stage === 'ready');
    assert.equal(await page.locator('#sum-load').textContent(), '未配置汇总');
    assert.equal(await page.locator('#sum-online').textContent(), '0 / 128');
    await page.locator('[data-mid="A-IN-M"]').click();
    assert.equal(await page.locator('#p-status').textContent(), '未绑定');
    assert((await page.locator('#p-body').textContent()).includes('--'));
    await waitForPanel(page);
    assert.equal(await page.locator('.tree-item .dot[data-state="on"]').count(), 0);
    await page.screenshot({ path: path.join(output, 'unbound.png') });
    await page.locator('#panel-close').click();
    await page.locator('[data-view="s1"]').click();
    await page.waitForTimeout(1400);
    const pick = await page.evaluate(() => {
      const app = window.StationApp, canvas = app.api.renderer.domElement, bounds = canvas.getBoundingClientRect(); // 当前场景可见点。
      const ray = new THREE.Raycaster(); // 只选择确实可点击、未被面板遮挡的仪表。
      for (const [id, meter] of Object.entries(app.api.meterMeshes)) {
        const world = meter.group.localToWorld(new THREE.Vector3(0, 0, 0.024)), projected = world.clone().project(app.api.camera); // 仪表正面中心。
        const x = bounds.left + (projected.x + 1) * bounds.width / 2, y = bounds.top + (1 - projected.y) * bounds.height / 2; // 屏幕坐标。
        if (document.elementFromPoint(x, y) !== canvas) continue;
        ray.setFromCamera(projected, app.api.camera);
        const hits = ray.intersectObjects(Object.values(app.api.meterMeshes).map(item => item.group), true); // 模板拾取对象集合。
        if (!hits.length) continue;
        let group = hits[0].object; while (group && !group.userData.meterId) group = group.parent;
        if (group && group.userData.meterId === id) return { id, x, y };
      }
      return null;
    });
    assert(pick, '场景中应能找到未遮挡的仪表');
    await page.mouse.click(pick.x, pick.y);
    assert((await page.locator('#p-id').textContent()).startsWith(pick.id));
    assert.equal(await page.evaluate(() => !!StationApp.api.scene.getObjectByName('meter-selection')), true);
    await page.locator('#panel-close').click();
    await page.locator('#btn-auto').click();
    assert.equal(await page.evaluate(() => StationApp.api.controls.autoRotate), true);
    await page.locator('#btn-auto').click();
    await page.locator('#btn-patrol').click();
    assert.equal(await page.locator('#info-panel').getAttribute('class'), 'open');
    await page.locator('#btn-patrol').click();
    await page.evaluate(async () => {
      const provider = StationApp.provider, now = Date.now(); // 测试专用夹具，不保存任何实际环境身份。
      provider.config.fields.p.identifier = 'test_power';
      provider.config.fields.q.identifier = 'test_reactive';
      provider.config.summary = { verified: true, meterIds: ['A-IN-M'] };
      provider.sdk = { callApi: async (url, body) => ({ code: 200, data: { list: url.includes('property-log') ? (body.dataID === 'test_power' ? [{ timestamp: now - 3600000, value: -2 }, { timestamp: now, value: 0 }] : []) : [{ dataID: 'test_power', timestamp: now, value: 0 }] } }) };
      provider._handleInit([{ path: 'meter/A-IN-M', productID: 'test-product', deviceName: 'test-device' }]);
      await provider.fetchDetail('A-IN-M'); await provider.fetchTrend('A-IN-M', 'p');
    });
    await page.locator('[data-mid="A-IN-M"]').click();
    assert.equal(await page.locator('#sum-load').textContent(), '0.00 kW');
    assert.equal(await page.locator('#p-status').textContent(), '数据有效');
    await page.locator('.trend-tab').filter({ hasText: '无功' }).click();
    await page.waitForFunction(() => StationApp.provider.get('A-IN-M').trendStatus.q === 'empty');
    assert.equal(await page.locator('#trend-status-title').textContent(), '近24小时暂无数据');
    for (const width of [1239, 900]) {
      await page.setViewportSize({ width, height: 930 });
      await page.waitForTimeout(100);
      assert.equal(await page.evaluate(() => document.documentElement.scrollWidth <= innerWidth), true);
      await waitForPanel(page);
      await page.screenshot({ path: path.join(output, 'layout-' + width + '.png') });
    }
    assert.deepEqual(errors, []);
    const metrics = await page.evaluate(() => ({ firstFrameMs: performance.getEntriesByName('station-first-frame')[0].startTime, readyMs: performance.getEntriesByName('station-ready')[0].startTime, draws: StationApp.api.renderer.info.render.calls, triangles: StationApp.api.renderer.info.render.triangles })); // 功能验收时的加载记录，非硬件性能承诺。
    fs.writeFileSync(path.join(output, 'result.json'), JSON.stringify({ metrics, errors, checks: ['未绑定空状态', '仪表拾取', '选择框', '聚焦', '巡航', '巡检', '零值', '趋势切换', '窄窗口'] }, null, 2));
    console.log(JSON.stringify(metrics));
  } finally { await browser.close(); await new Promise(resolve => server.close(resolve)); }
}
run().catch(error => { console.error(error); process.exitCode = 1; });
