/* ============================================================
 * scene.js — 三维场景构建（Three.js r128）
 * 房间 / 高压环网柜 / 干式变压器 / GCS 低压抽屉柜（含 SVG 无功补偿柜）
 * / 母联柜 / 母线桥 / 电表
 * 全部参数化建模，几何体与材质全局复用；中文铭牌用 CanvasTexture，
 * 浮动标签用手写 CSS2D 投影（file:// 直开无需额外渲染器文件）。
 * ============================================================ */

(function () {
  const T = THREE;
  /* 视觉资源工具只依赖包内 Three.js，保持场景离线可运行。 */
  const finish = window.StationFinish;

  /* ---------- 材质与几何缓存 ---------- */
  const M = {};
  function initMaterials() {
    M.cabinet   = new T.MeshStandardMaterial({ color: 0xcbd0d1, metalness: 0.12, roughness: 0.38 });
    M.cabinetHV = new T.MeshStandardMaterial({ color: 0xd5d7d7, metalness: 0.12, roughness: 0.4 });
    M.dark      = new T.MeshStandardMaterial({ color: 0x252d2e, metalness: 0.15, roughness: 0.64 });
    M.plinth    = new T.MeshStandardMaterial({ color: 0x17191d, metalness: 0.4,  roughness: 0.7 });
    M.seam      = new T.MeshStandardMaterial({ color: 0xc4c9ca, metalness: 0.12, roughness: 0.42 });
    M.glass     = new T.MeshStandardMaterial({ color: 0x1b2733, metalness: 0.1, roughness: 0.08, transparent: true, opacity: 0.55 });
    M.copper    = new T.MeshStandardMaterial({ color: 0xc47a3d, metalness: 0.9,  roughness: 0.35 });
    M.busduct   = new T.MeshStandardMaterial({ color: 0xb5bdc5, metalness: 0.72, roughness: 0.3 });
    M.transCase = new T.MeshStandardMaterial({ color: 0xb9c1c5, metalness: 0.16, roughness: 0.42 });
    M.transIn   = new T.MeshStandardMaterial({ color: 0x74402c, metalness: 0.3,  roughness: 0.6 });
    M.lampRed   = new T.MeshStandardMaterial({ color: 0x551111, emissive: 0xff3333, emissiveIntensity: 0.35 });
    M.lampGreen = new T.MeshStandardMaterial({ color: 0x114411, emissive: 0x2ecc40, emissiveIntensity: 0.35 });
    M.lampOff   = new T.MeshStandardMaterial({ color: 0x2c3136, emissive: 0x000000 });
    M.lampYel   = new T.MeshStandardMaterial({ color: 0x554411, emissive: 0xffcc00, emissiveIntensity: 0.35 });
    M.meterBody = new T.MeshStandardMaterial({ color: 0x14181d, metalness: 0.4, roughness: 0.5 });
    M.wall      = new T.MeshStandardMaterial({ color: 0xa6a8a5, metalness: 0.0, roughness: 0.92, side: T.DoubleSide });
    M.warn      = null; /* 画布生成 */
    /* r128 的十六进制颜色不会自动转线性；与 sRGB 输出配对，避免整体泛白。 */
    Object.values(M).forEach(material => {
      if (material && material.color) material.color.convertSRGBToLinear();
    });
  }

  /* ---------- 工具 ---------- */
  function box(w, h, d, mat, x, y, z, parent, cast) {
    const bevel = mat === M.cabinet || mat === M.cabinetHV || mat === M.transCase || mat === M.seam;
    const g = bevel ? finish.beveledBox(w, h, d, mat === M.seam ? 0.002 : 0.012) : CACHE.box(w, h, d);
    const m = new T.Mesh(g, mat);
    m.position.set(x, y, z);
    m.castShadow = cast !== false;
    m.receiveShadow = true;
    parent.add(m);
    return m;
  }
  const CACHE = {
    _box: {},
    box(w, h, d) {
      const k = w.toFixed(3) + '_' + h.toFixed(3) + '_' + d.toFixed(3);
      if (!this._box[k]) this._box[k] = new T.BoxGeometry(w, h, d);
      return this._box[k];
    }
  };

  /* ---------- Canvas 文本贴图 ---------- */
  function makeTextTexture(text, sub, w, h, fontSize, bg, fg) {
    const cv = document.createElement('canvas');
    cv.width = w; cv.height = h;
    const c = cv.getContext('2d');
    c.fillStyle = bg || 'rgba(16,22,30,0.92)';
    c.fillRect(0, 0, w, h);
    c.strokeStyle = 'rgba(195,210,200,0.3)';
    c.lineWidth = 2;
    c.strokeRect(1, 1, w - 2, h - 2);
    c.fillStyle = fg || '#e0e8df';
    c.font = 'bold ' + fontSize + 'px "Microsoft YaHei", sans-serif';
    c.textAlign = 'center';
    c.textBaseline = 'middle';
    c.fillText(text, w / 2, sub ? h * 0.38 : h / 2, w - 10);
    if (sub) {
      c.font = (fontSize * 0.62) + 'px "Microsoft YaHei", sans-serif';
      c.fillStyle = '#adbfb5';
      c.fillText(sub, w / 2, h * 0.72, w - 10);
    }
    const tex = new T.CanvasTexture(cv);
    tex.encoding = T.sRGBEncoding;
    tex.anisotropy = 4;
    return { tex, canvas: cv, ctx: c };
  }

  /* ============================================================ */
  const StationScene = {};

  StationScene.build = function (container, registry) {
    initMaterials();
    /* 准备阶段独立于数据接口；用于加载状态与自动化性能验收。 */
    performance.mark('station-build-start');
    const detailJobs = [];
    const disposers = [];
    let detailsReady = false;
    const surfaces = finish.applySurface(M); // 共用喷涂、墙面与地坪小贴图，不增加素材请求。

    /* 渲染器 */
    const renderer = new T.WebGLRenderer({ antialias: true });
    /* 纯软件 GPU 不具备显卡吞吐量，降低三维内部像素量；DOM 字体仍保持原分辨率。 */
    const gl = renderer.getContext();
    const debugInfo = gl.getExtension('WEBGL_debug_renderer_info');
    const softwareGpu = debugInfo && /swiftshader|llvmpipe|software/i.test(gl.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL));
    renderer.setPixelRatio(softwareGpu ? 0.75 : Math.min(window.devicePixelRatio, 1.5));
    renderer.setSize(container.clientWidth, container.clientHeight);
    renderer.shadowMap.enabled = true;
    renderer.shadowMap.type = T.PCFSoftShadowMap;
    renderer.outputEncoding = T.sRGBEncoding;
    renderer.toneMapping = T.ACESFilmicToneMapping;
    renderer.toneMappingExposure = 0.94;
    renderer.setClearColor(0x111b20);
    container.appendChild(renderer.domElement);

    const scene = new T.Scene();
    scene.fog = new T.Fog(0x111b20, 38, 80);
    /* 用同类型的四字节中性反射占位，避免细节阶段从无环境切换时重编所有标准材质。 */
    const reflectionPlaceholder = new T.DataTexture(new Uint8Array([128, 134, 140, 255]), 1, 1);
    reflectionPlaceholder.mapping = T.CubeUVReflectionMapping;
    reflectionPlaceholder.encoding = T.sRGBEncoding;
    reflectionPlaceholder.needsUpdate = true;
    scene.environment = reflectionPlaceholder;
    disposers.push(() => reflectionPlaceholder.dispose());

    const camera = new T.PerspectiveCamera(45, container.clientWidth / container.clientHeight, 0.1, 200);
    camera.position.set(12.6, 11.8, 17.2);

    const controls = new T.OrbitControls(camera, renderer.domElement);
    controls.enableDamping = true;
    controls.dampingFactor = 0.08;
    controls.minDistance = 2.5;
    controls.maxDistance = 34;
    controls.maxPolarAngle = 1.46;
    controls.target.set(0, 0.9, 0);

    /* 灯光 */
    scene.add(new T.HemisphereLight(0xeaf0f7, 0x70756f, 0.72));
    const sun = new T.DirectionalLight(0xf6f8ff, 0.86);
    sun.position.set(-3, 10, 4);
    sun.castShadow = true;
    sun.shadow.mapSize.set(2048, 2048);
    sun.shadow.camera.left = -12; sun.shadow.camera.right = 12;
    sun.shadow.camera.top = 12; sun.shadow.camera.bottom = -12;
    sun.shadow.bias = -0.001;
    sun.shadow.normalBias = 0.025;
    sun.shadow.radius = 3;
    sun.shadow.camera.near = 0.5; sun.shadow.camera.far = 32;
    scene.add(sun);
    const warm1 = new T.PointLight(0xe5f0ff, 0.22, 14);
    warm1.position.set(-3.2, 3.4, 3.0); scene.add(warm1);
    const warm2 = new T.PointLight(0xe5f0ff, 0.22, 14);
    warm2.position.set(3.2, 3.4, 3.0); scene.add(warm2);

    /* ---------- 地面（含警示涂装） ---------- */
    const roomDepth = 11.4, roomCenterZ = 0.7; // 南墙外移1.4个场景单位，放宽Ⅱ段柜前通道，设备坐标不动。
    const floorCv = document.createElement('canvas');
    floorCv.width = 1500; floorCv.height = 1140;
    (function paintFloor() {
      const c = floorCv.getContext('2d');
      const X = wx => (wx + 7.5) * 100;   /* 世界 x → 画布 */
      const Y = wz => (wz + 5) * 100;     /* 世界 z → 画布（北上） */
      /* 连续环氧地坪与大板伸缩缝，避免棋盘格造成微缩玩具感。 */
      c.fillStyle = '#58766b'; c.fillRect(0, 0, 1500, 1140);
      /* 将可平铺颗粒和轻微拖擦痕烘入现有地坪贴图，警示涂装仍绘制在最上层。 */
      c.save();
      c.globalAlpha = 0.055;
      c.fillStyle = c.createPattern(surfaces.grain, 'repeat');
      c.fillRect(0, 0, 1500, 1140);
      c.restore();
      c.save(); c.strokeStyle = 'rgba(219,232,223,0.13)'; c.lineWidth = 0.8;
      for (let i = 0; i < 32; i++) {
        const x = 190 + (i * 137) % 1100, y = 670 + (i * 31) % 130; // 仅在巡视通道加入稀疏磨痕。
        c.beginPath(); c.moveTo(x, y); c.lineTo(x + 12 + i % 27, y + 2); c.stroke();
      }
      c.restore();
      c.strokeStyle = 'rgba(24,43,36,0.18)'; c.lineWidth = 1;
      for (let x = 250; x < 1500; x += 250) { c.beginPath(); c.moveTo(x, 0); c.lineTo(x, 1140); c.stroke(); }
      for (let z = 250; z < 1140; z += 250) { c.beginPath(); c.moveTo(0, z); c.lineTo(1500, z); c.stroke(); }
      /* 柜前警示线（黄黑斜纹带） */
      function stripes(wx, wz, wLen, d) {
        const x0 = X(wx), y0 = Y(wz), w = wLen * 100, h = d * 100;
        c.save(); c.beginPath(); c.rect(x0, y0, w, h); c.clip();
        c.fillStyle = '#c8a641'; c.fillRect(x0, y0, w, h);
        c.fillStyle = '#1c1c1c';
        for (let i = -h; i < w; i += 40) {
          c.beginPath(); c.moveTo(x0 + i, y0 + h); c.lineTo(x0 + i + h, y0);
          c.lineTo(x0 + i + h + 14, y0); c.lineTo(x0 + i + 14, y0 + h); c.fill();
        }
        c.restore();
      }
      stripes(-2.55, -3.62, 5.1, 0.24);  /* 高压柜前（与变压器齐平） */
      stripes(-5.3, -3.62, 2.4, 0.24);   /* 变压器 A前 */
      stripes(2.9, -3.62, 2.4, 0.24);    /* 变压器 B前 */
      /* 低压柜列坐标（柜宽 0.8、间距 0.85）：Ⅰ段 8 柜 x=-5.4~0.55 + 母联 BUS-LINK x=1.4，
       * 全跨度 -5.8~1.8；Ⅱ段 8 柜 x=5.4~-0.55，全跨度 -0.95~5.8 */
      stripes(-5.65, 2.06, 7.35, 0.26);  /* 低压Ⅰ段柜前（8 柜 + 母联，-5.65~1.70） */
      stripes(-0.95, 4.56, 6.75, 0.26);  /* 低压Ⅱ段柜前（8 柜，-0.95~5.80） */
      /* 电缆沟（与母线/电缆走向一致） */
      c.fillStyle = '#22262b';
      c.fillRect(X(-5.55), Y(-3.55), 30, 505);  /* 变压器 A → Ⅰ段进线（A-IN x=-5.4） */
      c.fillRect(X(5.25), Y(-3.55), 30, 745);   /* 变压器 B → Ⅱ段进线（B-IN x=5.4） */
      c.fillRect(X(-4.15), Y(-4.08), 375, 26);  /* 高压出线沟 */
      c.fillRect(X(0.4), Y(-4.08), 375, 26);
      c.fillRect(X(1.25), Y(2.0), 30, 210);     /* 母联沟（母联柜 BUS-LINK x=1.4 → Ⅱ段排尾） */
    })();
    const floorTex = new T.CanvasTexture(floorCv);
    floorTex.anisotropy = Math.min(8, renderer.capabilities.getMaxAnisotropy());
    floorTex.encoding = T.sRGBEncoding;
    const floor = new T.Mesh(new T.PlaneGeometry(15, roomDepth),
      new T.MeshStandardMaterial({ map: floorTex, roughness: 0.34, metalness: 0.0,
        normalMap: surfaces.normal, normalScale: new T.Vector2(0.08, 0.08), roughnessMap: surfaces.roughness }));
    floor.rotation.x = -Math.PI / 2;
    floor.position.z = roomCenterZ; // 保持北墙与地坪涂装世界坐标不变，新增面积全部位于南侧。
    floor.receiveShadow = true;
    scene.add(floor);

    /* ---------- 房间墙体（南侧压矮便于鸟瞰） ---------- */
    const WH = 3.6;
    const north = box(15, WH, 0.2, M.wall, 0, WH / 2, -5.1, scene, false);
    const west  = box(0.2, WH, roomDepth, M.wall, -7.6, WH / 2, roomCenterZ, scene, false);
    const east  = box(0.2, WH, roomDepth, M.wall, 7.6, WH / 2, roomCenterZ, scene, false);
    const south = box(15, 1.0, 0.2, M.wall, 0, 0.5, 6.5, scene, false);
    [north, west, east, south].forEach(w => w.receiveShadow = true);
    /* 踢脚与顶带 */
    box(15.4, 0.12, 0.24, M.dark, 0, WH + 0.06, -5.1, scene, false);
    box(0.24, 0.12, roomDepth + 0.4, M.dark, -7.7, WH + 0.06, roomCenterZ, scene, false);
    box(0.24, 0.12, roomDepth + 0.4, M.dark, 7.7, WH + 0.06, roomCenterZ, scene, false);

    /* 空间细节分批装配，首帧不等待五金、纹理反射预处理。 */
    detailJobs.push(() => finish.room(scene, M, roomDepth, roomCenterZ));

    /* 安全出口牌（西墙） */
    const exitTex = makeTextTexture('安全出口', 'EXIT →', 256, 96, 44, 'rgba(6,80,40,0.95)', '#d8ffe8');
    const exitSign = new T.Mesh(new T.PlaneGeometry(0.9, 0.34),
      new T.MeshBasicMaterial({ map: exitTex.tex, transparent: false }));
    exitSign.position.set(-7.48, 2.6, 3.5);
    exitSign.rotation.y = Math.PI / 2;
    scene.add(exitSign);

    /* ---------- 电表管理 ---------- */
    const meterMeshes = {};    /* id -> group */

    /* 数显表 LCD 贴图（CanvasTexture，运行期刷新） */
    function makeLcd(w, h) {
      const cv = document.createElement('canvas');
      cv.width = 96; cv.height = 64;
      const tex = new T.CanvasTexture(cv);
      tex.encoding = T.sRGBEncoding;
      tex.anisotropy = 4;
      return { cv, ctx: cv.getContext('2d'), tex };
    }
    function drawLcd(lcd, line1, line2, alarm) {
      const key = line1 + '|' + line2 + '|' + alarm;
      if (lcd.last === key) return;
      lcd.last = key;
      const c = lcd.ctx;
      c.fillStyle = alarm ? '#182000' : '#04120a';
      c.fillRect(0, 0, 96, 64);
      c.fillStyle = '#39ff6a';
      c.font = 'bold 17px Consolas, monospace';
      c.textAlign = 'left'; c.textBaseline = 'top';
      c.fillText(line1, 6, 8, 88);
      c.font = 'bold 15px Consolas, monospace';
      c.fillStyle = alarm ? '#ffd23c' : '#9dffb2';
      c.fillText(line2, 6, 34, 88);
      lcd.tex.needsUpdate = true;
    }

    /* 创建一块电表：body + LCD 前面板。
     * visual=true 时为纯视觉表计（用于高压计量柜等未接入平台的柜）：
     * 不写入 meterMeshes、不带 userData.meterId（不可点击、无数据交互），
     * LCD 画一次静态待机画面 */
    function createMeter(m, size, visual) {
      const g = new T.Group();
      const bw = size.w, bh = size.h, bd = 0.035;
      const body = new T.Mesh(CACHE.box(bw, bh, bd), M.meterBody);
      body.castShadow = true;
      g.add(body);
      const lcd = makeLcd();
      /* z 必须大于边框盒前表面（bd/2+0.001+0.002），否则边框完全遮挡 LCD 画面
       * （表计呈纯黑方块，LCD 内容不可见） */
      const lcdPlane = new T.Mesh(new T.PlaneGeometry(bw * 0.86, bh * 0.62),
        new T.MeshBasicMaterial({ map: lcd.tex }));
      lcdPlane.position.z = bd / 2 + 0.005;
      g.add(lcdPlane);
      /* 边框 */
      const frame = new T.Mesh(CACHE.box(bw * 0.94, bh * 0.7, 0.004), M.dark);
      frame.position.z = bd / 2 + 0.001;
      frame.userData.staticDetail = true; // 表框可实例化，表体和 LCD 保留独立拾取。
      g.add(frame);
      if (visual) {
        drawLcd(lcd, '--', '--', false);
      } else {
        g.userData.meterId = m.id;
        meterMeshes[m.id] = { group: g, lcd };
      }
      return g;
    }

    const lampGeometry = new T.SphereGeometry(0.018, 8, 6); // 全柜共用低面数灯罩。
    /* 指示灯条（红/绿/黄） */
    function lampStrip(parent, x, y, z, on) {
      const l1 = new T.Mesh(lampGeometry, on ? M.lampRed : M.lampOff);
      l1.position.set(x - 0.05, y, z); parent.add(l1);
      const l2 = new T.Mesh(lampGeometry, on ? M.lampGreen : M.lampOff);
      l2.position.set(x, y, z); parent.add(l2);
      const l3 = new T.Mesh(lampGeometry, on ? M.lampYel : M.lampOff);
      l3.position.set(x + 0.05, y, z); parent.add(l3);
    }

    /* 铭牌 */
    function nameplate(parent, text, sub, w) {
      const t = makeTextTexture(text, sub, 512, 88, 40);
      const p = new T.Mesh(new T.PlaneGeometry(w || 0.68, 0.115),
        new T.MeshBasicMaterial({ map: t.tex }));
      return p;
    }

    /* ---------- GCS 低压柜 ---------- */
    const CAB_W = 0.8, CAB_D = 0.8, CAB_H = 2.2;
    function buildGcs(cab, cx, cz, faceSouth) {
      const g = new T.Group();
      g.position.set(cx, 0, cz);
      if (!faceSouth) g.rotation.y = Math.PI;
      scene.add(g);

      box(CAB_W, 0.1, CAB_D, M.plinth, 0, 0.05, 0, g);                 /* 底座 */
      box(CAB_W, CAB_H - 0.1, CAB_D, M.cabinet, 0, 0.1 + (CAB_H - 0.1) / 2, 0, g);
      /* 门缝（前面板内嵌 4mm） */
      box(CAB_W - 0.06, CAB_H - 0.28, 0.012, M.seam, 0, 0.1 + (CAB_H - 0.1) / 2, CAB_D / 2 + 0.004, g, false);
      /* 仪表室观察窗 */
      box(0.5, 0.42, 0.02, M.glass, 0, 1.78, CAB_D / 2 + 0.015, g, false);
      /* 把手 */
      box(0.05, 0.16, 0.03, M.dark, CAB_W / 2 - 0.09, 1.05, CAB_D / 2 + 0.012, g, false);
      /* 柜顶盖板贴合外壳，避免悬空黑块造成积木感；母线接入件独立补齐。 */
      box(CAB_W - 0.035, 0.02, CAB_D - 0.035, M.cabinet, 0, CAB_H + 0.008, 0, g, false);

      const fz = CAB_D / 2 + 0.035;   /* 前脸悬挂面深度 */
      lampStrip(g, 0.16, 1.98, fz, true);

      if (cab.kind === 'incoming') {
        const meter = createMeter({ id: cab.meterId || (cab.id + '-M') }, { w: 0.16, h: 0.2 });
        meter.position.set(0, 1.62, fz + 0.01);
        g.add(meter);
        /* ACB 合闸机构面板，不预置现场整定值 */
        box(0.22, 0.3, 0.03, M.dark, 0, 0.95, fz + 0.005, g, false);
        const acbTex = makeTextTexture('框架断路器', 'ACB', 256, 96, 30);
        const acb = new T.Mesh(new T.PlaneGeometry(0.2, 0.075), new T.MeshBasicMaterial({ map: acbTex.tex }));
        acb.position.set(0, 1.0, fz + 0.022); g.add(acb);
        const np = nameplate(g, cab.name, cab.id + ' · 进线');
        np.position.set(0, 2.12, fz + 0.01); g.add(np);
      } else if (cab.kind === 'svg') {
        /* SVG 无功自动补偿柜（A-COMP/B-COMP，未接入平台）：分层建模——
         * 上部纯视觉小表（不进 meterMeshes）→ 中部观察窗（半透明玻璃后嵌
         * 扁圆柱电容器组）→ 下部电容器室通风格栅。
         * 门缝面板前表面 z≈0.41、玻璃位于 z=0.453~0.473，电容器圆心 z 取
         * 玻璃中心-玻璃厚/2-半径，恰好顶在玻璃背面，透过玻璃可见 */
        const meter = createMeter(null, { w: 0.14, h: 0.17 }, true);
        meter.position.set(0, 1.62, fz + 0.01);
        g.add(meter);
        box(0.5, 0.34, 0.02, M.glass, 0, 1.3, fz + 0.028, g, false);   /* 观察窗 */
        for (let i = 0; i < 3; i++) {
          const cap = new T.Mesh(new T.CylinderGeometry(0.05, 0.05, 0.24, 12), M.seam);
          cap.position.set(-0.14 + i * 0.14, 1.3, fz + 0.028 - 0.01 - 0.05);
          g.add(cap);
        }
        /* 通风格栅：SVG 散热量大，电容器室前脸留 4 条水平散热缝 */
        for (let i = 0; i < 4; i++) {
          box(0.56, 0.02, 0.014, M.dark, 0, 0.44 + i * 0.1, fz + 0.002, g, false);
        }
        const svgTex = makeTextTexture('SVG', '无功自动补偿', 512, 110, 34);
        const sv = new T.Mesh(new T.PlaneGeometry(0.56, 0.12), new T.MeshBasicMaterial({ map: svgTex.tex }));
        sv.position.set(0, 1.06, fz + 0.01); g.add(sv);
        const np = nameplate(g, cab.name, cab.id);
        np.position.set(0, 2.12, fz + 0.01); g.add(np);
      } else if (cab.kind === 'bus') {
        if (cab.meterId) {
          const meter = createMeter({ id: cab.meterId }, { w: 0.14, h: 0.17 });
          meter.position.set(0, 1.6, fz + 0.01);
          g.add(meter);
        }
        const np = nameplate(g, cab.name, cab.id + ' · 母联');
        np.position.set(0, 2.12, fz + 0.01); g.add(np);
      } else { /* feeder：抽屉回路数按台账动态渲染，每抽屉 1 块多功能表计 */
        const d = cab.drawers;
        const n = d.length;
        const colX = [-0.185, 0.185];
        const top = 1.66, bottom = 0.2;
        const rows = Math.max(1, Math.ceil(n / 2));
        const gap = rows > 1 ? Math.min(0.45, (top - bottom) / (rows - 1)) : 0;
        const panelH = rows > 1 ? Math.min(0.4, gap * 0.92) : 0.4;
        for (let i = 0; i < n; i++) {
          const col = i % 2, row = Math.floor(i / 2);
          const dx = colX[col], dy = top - row * gap;
          /* 抽屉面板 */
          box(0.34, panelH, 0.015, M.seam, dx, dy, fz - 0.006, g, false);
          const meter = createMeter({ id: d[i].id }, { w: 0.12, h: Math.min(0.14, panelH * 0.42) });
          meter.position.set(dx, dy + panelH * 0.16, fz + 0.014);
          g.add(meter);
          /* 抽屉把手 */
          box(0.1, 0.025, 0.02, M.dark, dx, dy - panelH * 0.32, fz + 0.008, g, false);
        }
        const np = nameplate(g, cab.name, cab.id + ' · ' + n + ' 回路抽屉');
        np.position.set(0, 2.12, fz + 0.01); g.add(np);
      }
      detailJobs.push(() => finish.cabinet(g, CAB_W, CAB_H, CAB_D, cab.kind));
      return g;
    }

    /* ---------- 高压环网柜 ---------- */
    const HV_W = 0.75, HV_D = 0.9, HV_H = 2.5;
    function buildHv(cab, cx, cz) {
      const g = new T.Group();
      g.position.set(cx, 0, cz);
      scene.add(g);
      box(HV_W, 0.1, HV_D, M.plinth, 0, 0.05, 0, g);
      box(HV_W, HV_H - 0.1, HV_D, M.cabinetHV, 0, 0.1 + (HV_H - 0.1) / 2, 0, g);
      box(HV_W - 0.06, HV_H - 0.34, 0.012, M.seam, 0, 0.1 + (HV_H - 0.1) / 2, HV_D / 2 + 0.004, g, false);
      box(0.06, 0.18, 0.03, M.dark, HV_W / 2 - 0.08, 1.2, HV_D / 2 + 0.012, g, false);
      const fz = HV_D / 2 + 0.035;
      /* 带电指示器 */
      lampStrip(g, 0, 2.1, fz, true);
      if (cab.kind === 'hv-pt') {
        /* 高压计量柜（AH02/AH05，未接入平台）：表计为纯视觉展示（visual，
         * 不进 meterMeshes 不可点击）；此前此处误写 'hv-meter'，与 data.js
         * 的 kind='hv-pt' 不匹配，导致表计与铅封窗从未渲染 */
        const meter = createMeter(null, { w: 0.15, h: 0.18 }, true);
        meter.position.set(0, 1.62, fz + 0.01);
        g.add(meter);
        box(0.4, 0.5, 0.02, M.glass, 0, 1.05, fz + 0.012, g, false);  /* 计量铅封窗 */
      }
      const np = nameplate(g, cab.name, cab.id, 0.62);
      np.position.set(0, 2.32, fz + 0.01); g.add(np);
      detailJobs.push(() => finish.cabinet(g, HV_W, HV_H, HV_D, 'hv'));
      return g;
    }

    /* ---------- 干式变压器（与高压柜同一条线，柜前面对齐） ---------- */
    function buildTransformer(t, cx, cz) {
      const g = new T.Group();
      g.position.set(cx, 0, cz);
      scene.add(g);
      const W = 2.0, D = 1.3, H = 2.3;
      box(W, 0.12, D, M.dark, 0, 0.06, 0, g);                          /* 底座轨道 */
      /* 防护外壳（IP3X），正面百叶 */
      box(W, H, D, M.transCase, 0, 0.12 + H / 2, 0, g);
      box(W - 0.3, H - 0.4, 0.02, M.seam, 0, 0.12 + H / 2, D / 2 + 0.008, g, false);
      /* 百叶片由细节阶段统一添加，避免旧黑条与新百叶重叠。 */
      /* 高压警示牌 */
      const warnTex = makeTextTexture('⚠ 止步 高压危险', t.ratio || '', 512, 110, 34, 'rgba(64,48,4,0.95)', '#ffe28a');
      const warn = new T.Mesh(new T.PlaneGeometry(0.85, 0.18), new T.MeshBasicMaterial({ map: warnTex.tex }));
      warn.position.set(0, 2.05, D / 2 + 0.03); g.add(warn);
      /* 温控仪作为示例结构，仅配置 meterId 时创建可绑定表计 */
      if (t.meterId) {
        const tc = createMeter({ id: t.meterId }, { w: 0.15, h: 0.18 });
        tc.position.set(W / 2 - 0.32, 1.5, D / 2 + 0.05);
        g.add(tc);
      }
      const np = nameplate(g, t.name, [t.model, t.capacity, t.vector].filter(Boolean).join(' · ') || t.ratio || '', 0.95);
      np.position.set(-0.4, 2.29, D / 2 + 0.03); g.add(np);
      detailJobs.push(() => finish.cabinet(g, W, H + 0.12, D, 'transformer'));
      return g;
    }

    /* ---------- 布置设备（北墙一列：变压器—高压柜—变压器，前面对齐） ---------- */
    const HV_Z = -4.15, TR_Z = -4.35;   /* 柜前/变前统一齐平于 z = -3.7 */
    window.HV_CABINETS.forEach((cab, i) => {
      buildHv(cab, -2.0 + i * 0.8, HV_Z);
    });
    buildTransformer(window.TRANSFORMERS[0], -4.15, TR_Z);
    buildTransformer(window.TRANSFORMERS[1], 4.15, TR_Z);

    /* 低压两列（每列：进线柜 + SVG 补偿柜 + 6 面馈电柜，共 8 面），面对面（中间为巡视通道）；
       母联柜 BUS-LINK 接在Ⅰ段排尾（结构展示，未接入平台）。
       Ⅰ段 x = -5.4 + i×0.85（i=0..7 → -5.4~0.55），母联 x = -5.4 + 8×0.85 = 1.4；
       Ⅱ段 x = 5.4 - i×0.85（i=0..7 → 5.4~-0.55） */
    const LV_GAP = 0.85;
    const s1 = window.LV_SECTIONS[0], s2 = window.LV_SECTIONS[1];
    const busCab = window.BUS_CABINET;
    s1.cabinets.forEach((cab, i) => { cab._x = -5.4 + i * LV_GAP; cab._z = 1.6; buildGcs(cab, cab._x, cab._z, true); });
    if (busCab) { busCab._x = -5.4 + s1.cabinets.length * LV_GAP; busCab._z = 1.6; buildGcs(busCab, busCab._x, busCab._z, true); }
    s2.cabinets.forEach((cab, i) => { cab._x = 5.4 - i * LV_GAP; cab._z = 4.1; buildGcs(cab, cab._x, cab._z, true); });

    /* ---------- 母线桥 / 密集母线（横平竖直的折线走线，禁止斜穿设备） ---------- */
    function busRun(pts, y) {
      for (let i = 0; i < pts.length - 1; i++) {
        const x0 = pts[i][0], z0 = pts[i][1], x1 = pts[i + 1][0], z1 = pts[i + 1][1];
        const len = Math.hypot(x1 - x0, z1 - z0);
        if (len < 0.01) continue;
        const ang = Math.atan2(z1 - z0, x1 - x0);
        const m = new T.Mesh(finish.beveledBox(len, 0.18, 0.25, 0.008), M.busduct);
        m.position.set((x0 + x1) / 2, y, (z0 + z1) / 2);
        m.rotation.y = -ang;
        m.castShadow = true;
        scene.add(m);
        detailJobs.push(() => finish.busJoints(scene, x0, z0, x1, z1, y));
      }
      /* 封闭转角填满直段交接处；仅细化既有路径，不表示现场接头施工尺寸。 */
      pts.slice(1, -1).forEach(([x, z]) => {
        box(0.25, 0.18, 0.25, M.busduct, x, y, z, scene);
      });
    }
    /* 变压器 A → Ⅰ段进线柜顶：先向西平移，再沿西墙下行走线 */
    busRun([[-4.15, -3.55], [-5.4, -3.55], [-5.4, 1.5]], 2.5);
    /* 变压器 B → Ⅱ段进线柜顶（镜像） */
    busRun([[4.15, -3.55], [5.4, -3.55], [5.4, 3.9]], 2.5);
    /* 母联柜（BUS-LINK，x=1.4）→ Ⅱ段排尾（B-CAB-06，x=-0.55）：沿柜后沟向南，
     * 再沿 z=4.1 柜顶上方（y=2.36 高于柜顶 2.2）向西横跨到排尾柜 */
    busRun([[1.4, 2.0], [1.4, 4.1], [-0.55, 4.1]], 2.36);
    /* 接入法兰消除柜顶悬空缝隙；设备布局、锚点和原有连接关系保持一致。 */
    detailJobs.push(() => {
      [-4.15, 4.15].forEach(x => finish.busTerminal(scene, x, -3.55, 2.5, 2.42, -0.31));
      finish.busTerminal(scene, -5.4, 1.5, 2.5, CAB_H, 0);
      finish.busTerminal(scene, 5.4, 3.9, 2.5, CAB_H, 0);
      finish.busTerminal(scene, 1.4, 2.0, 2.36, CAB_H, -0.16);
      finish.busTerminal(scene, -0.55, 4.1, 2.36, CAB_H, 0);
    });

    /* ---------- 区域浮动标签（手写 CSS2D） ---------- */
    const labelLayer = document.createElement('div');
    labelLayer.className = 'label-layer';
    container.appendChild(labelLayer);
    const labels = [];
    function addLabel(text, cls, x, y, z) {
      const el = document.createElement('div');
      el.className = 'label3d ' + (cls || '');
      el.textContent = text;
      labelLayer.appendChild(el);
      labels.push({ el, pos: new T.Vector3(x, y, z) });
    }
    /* 标签 x 取新柜列几何中心：Ⅰ段 8 柜+母联全跨度 -5.8~1.8 → 中点 -2.0；
     * Ⅱ段 8 柜全跨度 -0.95~5.8 → 中点 2.4；母联 BUS-LINK x=1.4；
     * SVG 补偿柜位于进线柜与首面馈电柜之间（A-COMP x=-4.55、B-COMP x=4.55） */
    addLabel('高压区', 'zone', 0, 3.1, -4.2);
    addLabel('变压器 A', 'zone dim', -4.15, 2.9, -4.35);
    addLabel('变压器 B', 'zone dim', 4.15, 2.9, -4.35);
    addLabel('低压柜列 A', 'zone', -2.0, 2.9, 1.6);
    addLabel('低压柜列 B', 'zone', 2.4, 2.9, 4.1);
    addLabel('母联', 'zone dim', 1.4, 2.9, 1.6);
    addLabel('SVG 无功补偿', 'zone dim', -4.55, 2.9, 1.6);
    addLabel('SVG 无功补偿', 'zone dim', 4.55, 2.9, 4.1);

    /* ---------- 预设视角 ---------- */
    const VIEWS = {
      overview: { pos: [12.6, 11.8, 17.2], tgt: [0, 0.9, 0] },
      hv:       { pos: [0, 3.4, -0.2],  tgt: [0, 1.5, -4.2] },
      trans:    { pos: [-4.15, 2.6, -0.6], tgt: [-4.15, 1.4, -4.25] },
      s1:       { pos: [-2.0, 6.2, 7.2], tgt: [-2.0, 1.2, 1.6] },
      s2:       { pos: [2.4, 3.2, 10.3],  tgt: [2.4, 1.2, 4.1] }
    };
    let flight = null;
    function flyTo(view, dur) {
      flight = {
        t0: performance.now(), dur: dur || 1200,
        p0: camera.position.clone(), p1: new T.Vector3().fromArray(view.pos),
        q0: controls.target.clone(), q1: new T.Vector3().fromArray(view.tgt)
      };
    }

    /* ---------- LCD 刷新 ---------- */
    function refreshLcds(states, registryMap) {
      registry.forEach(m => {
        const mm = meterMeshes[m.id];
        if (!mm) return;
        const s = states[m.id];
        /* 真实数据可能部分字段缺失（未上报），缺字段显示 -- 而不是崩溃 */
        const ia = s && s.ia != null ? s.ia.toFixed(1) : '--'; // 电流只取实际读数。
        const p = s && s.p != null ? s.p.toFixed(1) : '--'; // 开关状态不覆盖功率或伪造零值。
        drawLcd(mm.lcd, ia + ' A', p + ' kW', false);
      });
    }

    /* ---------- 渲染循环 ---------- */
    /* 同材质同几何的静态零件批量化；动态 LCD 和拾取组保持独立。 */
    finish.batchStatic(scene);
    let running = true;
    let frameHandle = 0;
    let firstFrame = false;
    let detailTimer = 0;
    let resizePending = true;
    /* requestAnimationFrame 之后让出一次事件循环，确保主体已实际呈现。 */
    function nextDetail() {
      if (!running) return;
      /* 每批至多占用 6ms，避免每个螺栓批次都触发一帧未合批的高成本绘制。 */
      const started = performance.now();
      while (detailJobs.length && performance.now() - started < 6) detailJobs.shift()();
      if (detailJobs.length) {
        detailTimer = setTimeout(nextDetail, 0);
      } else {
        finish.batchStatic(scene);
        finish.environment(renderer, scene).then(dispose => {
          if (!running) { dispose(); return; }
          disposers.push(dispose);
          detailsReady = true;
        }).catch(error => {
          /* 反射图损坏时保留可操作主体，同时明确展示细节失败。 */
          if (!running) return;
          document.getElementById('scene-progress').textContent = '环境反射加载失败，请刷新重试';
          console.error('环境反射加载失败', error);
        });
      }
    }
    const tmpV = new T.Vector3();
    function animate() {
      if (!running) return;
      frameHandle = requestAnimationFrame(animate);
      if (flight) {
        const k = Math.min(1, (performance.now() - flight.t0) / flight.dur);
        const e = k < 0.5 ? 2 * k * k : 1 - Math.pow(-2 * k + 2, 2) / 2; /* easeInOut */
        camera.position.lerpVectors(flight.p0, flight.p1, e);
        controls.target.lerpVectors(flight.q0, flight.q1, e);
        if (k >= 1) flight = null;
      }
      controls.update();
      const w = container.clientWidth, h = container.clientHeight;
      if (resizePending && w && h) {
        camera.aspect = w / h;
        camera.updateProjectionMatrix();
        renderer.setSize(w, h);
        resizePending = false;
      }
      labels.forEach(L => {
        tmpV.copy(L.pos).project(camera);
        const vis = tmpV.z < 1;
        L.el.style.display = vis ? 'block' : 'none';
        if (vis) {
          L.el.style.left = ((tmpV.x * 0.5 + 0.5) * w) + 'px';
          L.el.style.top = ((-tmpV.y * 0.5 + 0.5) * h) + 'px';
        }
      });
      renderer.render(scene, camera);
      if (!firstFrame) {
        firstFrame = true;
        /* 新增五金仅接收阴影；主体阴影首帧已完整，加载期间也无需每帧重算。 */
        renderer.shadowMap.autoUpdate = false;
        performance.mark('station-first-frame');
        container.dataset.stage = 'interactive';
        window.dispatchEvent(new CustomEvent('station:interactive'));
        detailTimer = setTimeout(nextDetail, 0);
      }
      if (detailsReady && container.dataset.stage !== 'ready') {
        container.dataset.stage = 'ready';
        performance.mark('station-ready');
        window.dispatchEvent(new CustomEvent('station:ready'));
        /* 场景静态，镜头移动不必重复计算整张阴影贴图。 */
        renderer.shadowMap.autoUpdate = false;
      }
    }
    animate();

    /* 宿主 iframe 尺寸变化不一定触发 window.resize。 */
    const observer = new ResizeObserver(() => { resizePending = true; });
    observer.observe(container);

    return {
      scene, camera, renderer, controls,
      meterMeshes, registry,
      views: VIEWS,
      flyTo,
      refreshLcds,
      stop() {
        running = false;
        cancelAnimationFrame(frameHandle);
        clearTimeout(detailTimer);
        observer.disconnect();
        controls.dispose();
        disposers.forEach(dispose => dispose());
        finish.dispose(scene);
        renderer.dispose();
        labelLayer.remove();
        renderer.domElement.remove();
      }
    };
  };

  window.StationScene = StationScene;
})();
