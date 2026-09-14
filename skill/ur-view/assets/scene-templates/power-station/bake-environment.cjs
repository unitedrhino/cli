/* 文件说明：在本地软件 GPU 上离线生成室内 CubeUV 反射图。
 * 输出 js/environment.js；运行时只解码 PNG，不执行 PMREM。
 * 依赖现有 Playwright，运行方式同 verify.cjs。
 */
const { chromium } = require('playwright');
const fs = require('node:fs');
const path = require('node:path');
const zlib = require('node:zlib');
const { pathToFileURL } = require('node:url');

/** chunk 生成带长度与 CRC32 的 PNG 数据块，避免引入额外图片编码依赖。 */
function chunk(name, data) {
  const type = Buffer.from(name);
  const content = Buffer.concat([type, data]);
  let crc = 0xffffffff;
  for (const byte of content) {
    crc ^= byte;
    for (let bit = 0; bit < 8; bit++) crc = (crc >>> 1) ^ ((crc & 1) ? 0xedb88320 : 0);
  }
  const length = Buffer.alloc(4);length.writeUInt32BE(data.length);
  const checksum = Buffer.alloc(4);checksum.writeUInt32BE((crc ^ 0xffffffff) >>> 0);
  return Buffer.concat([length, content, checksum]);
}

/** bake 烘焙无外部素材的环境反射图，完成后释放浏览器。 */
async function bake() {
  const browser = await chromium.launch({ headless: true, args: ['--no-sandbox', '--enable-unsafe-swiftshader', '--use-angle=swiftshader'] });
  try {
    const page = await browser.newPage({ viewport: { width: 320, height: 240 } });
    await page.addInitScript(() => {
      let scene;
      Object.defineProperty(window, 'StationScene', { get: () => scene, set(value) {
        scene = value;const build = value.build;
        value.build = (...args) => { window.__api = build(...args);return window.__api; };
      } });
    });
    await page.goto(pathToFileURL(path.join(__dirname, 'index.html')).href);
    await page.waitForFunction(() => document.getElementById('viewport').dataset.stage === 'ready');
    const atlas = await page.evaluate(() => {
      const original = THREE.PMREMGenerator.prototype.fromEquirectangular;
      let target;
      THREE.PMREMGenerator.prototype.fromEquirectangular = function (...args) { target = original.apply(this, args);return target; };
      const release = StationFinish.bakeEnvironment(window.__api.renderer, window.__api.scene);
      const pixels = new Uint8Array(target.width * target.height * 4);
      window.__api.renderer.readRenderTargetPixels(target, 0, 0, target.width, target.height, pixels);
      const result = { width: target.width, height: target.height, encoding: target.texture.encoding, type: target.texture.type, pixels: Array.from(pixels) };
      release();THREE.PMREMGenerator.prototype.fromEquirectangular = original;
      return result;
    });
    const header = Buffer.alloc(13);
    header.writeUInt32BE(atlas.width);header.writeUInt32BE(atlas.height, 4);header[8] = 8;header[9] = 6;
    const raw = Buffer.from(atlas.pixels);const rows = [];
    for (let y = 0; y < atlas.height; y++) rows.push(Buffer.from([0]), raw.subarray(y * atlas.width * 4, (y + 1) * atlas.width * 4));
    const png = Buffer.concat([Buffer.from([137, 80, 78, 71, 13, 10, 26, 10]), chunk('IHDR', header), chunk('IDAT', zlib.deflateSync(Buffer.concat(rows), { level: 9 })), chunk('IEND', Buffer.alloc(0))]);
    delete atlas.pixels;atlas.image = 'data:image/png;base64,' + png.toString('base64');
    fs.writeFileSync(path.join(__dirname, 'js/environment.js'), '/* 文件说明：离线预滤波室内反射图（r128 CubeUV布局），免除运行期PMREM。 */\nwindow.StationEnvironment = ' + JSON.stringify(atlas) + ';\n');
    console.log('已更新环境图，PNG 字节数：', png.length);
  } finally { await browser.close(); }
}
bake().catch(error => { console.error(error);process.exitCode = 1; });
