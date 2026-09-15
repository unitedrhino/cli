/* 文件说明：轻量 Three.js 建筑建模、分批立面装配、整楼展开与准确拾取。
 * 外观参考照片；尺寸、楼层轮廓和室内区域均为示意。状态由 app.js 驱动，动画始终从当前姿态趋近目标。
 */
(function (root) {
  'use strict';
  /** createBuildingScene 接收画布容器和选择回调，返回模式/选中控制及只读验收信息。 */
  root.createBuildingScene = function (container, onSelect) {
    // 渲染资源全部局部持有，离开页面时统一释放。
    const T = root.THREE;
    const config = root.BuildingConfig;
    const data = root.BuildingData;
    const scene = new T.Scene();
    const camera = new T.PerspectiveCamera(36, 1, .1, 500);
    const renderer = new T.WebGLRenderer({ antialias: true, alpha: true, powerPreference: 'high-performance' });
    renderer.setPixelRatio(Math.min(root.devicePixelRatio || 1, 1.5));
    renderer.outputEncoding = T.sRGBEncoding;
    renderer.toneMapping = T.ACESFilmicToneMapping;
    renderer.toneMappingExposure = .95;
    renderer.shadowMap.enabled = true;
    renderer.shadowMap.type = T.PCFSoftShadowMap;
    renderer.shadowMap.autoUpdate = false;
    container.appendChild(renderer.domElement);
    const controls = new T.OrbitControls(camera, renderer.domElement); // 画布真实边界独立于左右面板。
    controls.enableDamping = true;
    controls.dampingFactor = .075;
    controls.minDistance = 30;
    controls.maxDistance = 220;
    controls.maxPolarAngle = Math.PI * .485;
    controls.minPolarAngle = .14;
    controls.enablePan = true;
    const ambient = new T.HemisphereLight(0xc6dcf3, 0x666756, .65); // 天光与地面反射填充暗部。
    scene.add(ambient);
    const sun = new T.DirectionalLight(0xfff4df, 1.6); // 唯一主要阴影光源。
    sun.position.set(-38, 65, 42);
    sun.castShadow = true;
    sun.shadow.mapSize.set(2048, 2048);
    Object.assign(sun.shadow.camera, { left: -62, right: 62, top: 85, bottom: -42, near: 1, far: 180 });
    sun.shadow.normalBias = .055;
    sun.shadow.bias = -.00015;
    scene.add(sun);
    const fill = new T.DirectionalLight(0xadd3ed, .4); // 无阴影补光，勾勒背面轮廓。
    fill.position.set(25, 30, -40);
    scene.add(fill);
    const box = new T.BoxGeometry(1, 1, 1); // 全局复用基础几何。
    const materials = []; // 含分层透明材质，供释放与环境贴图复用。
    const geometries = new Set([box]);
    const textures = new Set();
    const white = material(0xd8ddda, .43, .22);
    const edgeWhite = material(0xe8e7df, .37, .23);
    const glass = material(0x233b4b, .19, .38);
    const seam = material(0x52636a, .65, .25);
    const paving = material(0x66756e, .95, .03);
    const groundMat = material(0x293e48, .92, .03);
    const lawn = material(0x425c42, .95, 0);
    const trunkMat = material(0x6c5d49, 1, 0);
    const foliageMat = material(0x557952, .94, 0);
    const floorGroups = []; // 每组保存楼板、立面、区域和当前可见度。
    const pickables = [];
    const stage = new T.Group(); // 环境场地在展开时弱化，B1 不被地面遮住。
    scene.add(stage);
    const roofGroup = new T.Group();
    const state = { mode: 'exterior', expanded: false, spacing: 4, floor: null, zone: null };
    const metrics = { bodyMs: null, detailMs: null, ready: false }; // 从 iframe navigationStart（performance 原点）计时。
    let frameId = 0;
    let disposed = false;
    let cameraTween = null;
    let lastTime = performance.now();
    let moved = false; // 记录整个手势是否越过阈值，拖出后回到原位仍属于旋转。
    let pressedPointer = null; // 仅响应本次按下的主指针，取消手势不触发选择。
    let pressedObject = null; // 锁定按下时的对象，轻微抖动或镜头更新不改变点击目标。
    const clickSlop = 6; // CSS 像素阈值，与设备像素比无关。
    const pointer = new T.Vector2();
    const ray = new T.Raycaster();
    const down = new T.Vector2();
    const labelRoot = container.querySelector('#floor-labels');
    const tooltip = container.querySelector('#tooltip');
    const labels = [];
    const zoneLabels = []; // 仅在选层时投影区域名称，与颜色无关也能识别具体空间。
    const outlineGeometry = new T.BufferGeometry().setFromPoints([
      new T.Vector3(-.5, 0, -.5), new T.Vector3(.5, 0, -.5), new T.Vector3(.5, 0, .5), new T.Vector3(-.5, 0, .5),
    ]);
    geometries.add(outlineGeometry);
    const outlineMaterial = new T.LineBasicMaterial({ color: 0xffffff, depthTest: false });
    materials.push(outlineMaterial);
    const selectionOutline = new T.LineLoop(outlineGeometry, outlineMaterial);
    selectionOutline.renderOrder = 100;
    selectionOutline.visible = false;
    scene.add(selectionOutline);
    const jobs = [];
    const drainChannel = new MessageChannel(); // 每一帧后启动下一批，无固定隐藏加载延迟。

    /** material 创建并登记可共享的标准材质，颜色、粗糙度和金属度均为本地参数。 */
    function material(color, roughness, metalness) {
      const value = new T.MeshStandardMaterial({ color, roughness, metalness });
      value.color.convertSRGBToLinear();
      materials.push(value);
      return value;
    }
    /** meshBox 按尺寸缩放共享盒子；返回网格供分层拾取和细节装配。 */
    function meshBox(parent, mat, x, y, z, w, h, d, shadow = true) {
      const mesh = new T.Mesh(box, mat);
      mesh.position.set(x, y, z);
      mesh.scale.set(w, h, d);
      mesh.castShadow = shadow;
      mesh.receiveShadow = true;
      parent.add(mesh);
      return mesh;
    }
    /** instances 将重复构件合为一批绘制，items 是位置和尺寸数组。 */
    function instances(parent, mat, items, geometry = box) {
      const mesh = new T.InstancedMesh(geometry, mat, items.length);
      const dummy = new T.Object3D();
      items.forEach((item, index) => {
        dummy.position.set(item[0], item[1], item[2]);
        dummy.scale.set(item[3], item[4], item[5]);
        dummy.rotation.set(0, item[6] || 0, 0);
        dummy.updateMatrix();
        mesh.setMatrixAt(index, dummy.matrix);
      });
      mesh.castShadow = true;
      mesh.receiveShadow = true;
      parent.add(mesh);
      return mesh;
    }
    /** roundedShape 返回水平圆角矩形的二维轮廓，供裙房与上部楼板共用。 */
    function roundedShape(w, d, r) {
      const shape = new T.Shape();
      shape.moveTo(-w / 2 + r, -d / 2);
      shape.lineTo(w / 2 - r, -d / 2);
      shape.quadraticCurveTo(w / 2, -d / 2, w / 2, -d / 2 + r);
      shape.lineTo(w / 2, d / 2 - r);
      shape.quadraticCurveTo(w / 2, d / 2, w / 2 - r, d / 2);
      shape.lineTo(-w / 2 + r, d / 2);
      shape.quadraticCurveTo(-w / 2, d / 2, -w / 2, d / 2 - r);
      shape.lineTo(-w / 2, -d / 2 + r);
      shape.quadraticCurveTo(-w / 2, -d / 2, -w / 2 + r, -d / 2);
      return shape;
    }
    /** slab 构建薄楼板，圆角几何按楼型缓存，避免每层重复创建。 */
    const slabs = new Map();
    function slab(floor, mat, parent) {
      const key = floor.width + ':' + floor.depth;
      if (!slabs.has(key)) {
        const geometry = new T.ExtrudeGeometry(roundedShape(floor.width, floor.depth, 1.5), { depth: .25, bevelEnabled: false, curveSegments: 6 });
        geometry.rotateX(-Math.PI / 2);
        slabs.set(key, geometry);
        geometries.add(geometry);
      }
      const mesh = new T.Mesh(slabs.get(key), mat);
      mesh.position.z = floor.centerZ;
      mesh.castShadow = true;
      mesh.receiveShadow = true;
      parent.add(mesh);
      return mesh;
    }
    /** buildBody 先建房间范围、玻璃主体和区域地板，保证首帧即可交互。 */
    function buildBody() {
      meshBox(stage, groundMat, 0, -1.5, 0, 66, 2, 46);
      meshBox(stage, paving, 0, -.35, 0, 60, .3, 41);
      config.floors.forEach(floor => {
        const group = new T.Group();
        group.position.y = floor.y;
        scene.add(group);
        const surface = material(0xa7b9b9, .75, .08); // 楼层独立材质使选中弱化不影响其他层。
        const shellMat = glass.clone();
        const finMat = white.clone();
        materials.push(shellMat, finMat);
        const structure = new T.Group();
        const exterior = new T.Group();
        const zonesGroup = new T.Group();
        group.add(structure, exterior, zonesGroup);
        const base = slab(floor, surface, structure);
        base.userData = { floor: floor.id };
        pickables.push(base);
        // 玻璃边界只覆盖外围，展开时隐藏，避免半透明排序遮挡区域。
        const wallDefs = [[0, floor.centerZ + floor.depth / 2 - .32, floor.width - 2, .25], [0, floor.centerZ - floor.depth / 2 + .32, floor.width - 2, .25], [-floor.width / 2 + .32, floor.centerZ, .25, floor.depth - 2], [floor.width / 2 - .32, floor.centerZ, .25, floor.depth - 2]];
        wallDefs.forEach(([x, z, w, d]) => {
          const wall = meshBox(exterior, shellMat, x, 1.9, z, w, 3.55, d);
          wall.userData = { floor: floor.id };
          pickables.push(wall);
        });
        const zoneMeshes = floor.zones.map(zone => {
          const sample = data.values[zone.id];
          const ratio = data.loadRatio(sample);
          const mat = material(data.color(ratio), .62, .1);
          const mesh = meshBox(zonesGroup, mat, zone.x, .41, zone.z, zone.width, .3, zone.depth, false);
          mesh.userData = { floor: floor.id, zone: zone.id };
          pickables.push(mesh);
          const label = document.createElement('button');
          label.className = 'zone-label';
          label.innerHTML = zone.name + '<small>' + (sample.power === null ? '暂无数据' : sample.power.toFixed(2) + ' kW') + '</small>';
          label.style.setProperty('--zone-color', data.color(ratio));
          label.onclick = () => onSelect(floor.id, zone.id);
          label.hidden = true;
          labelRoot.appendChild(label);
          zoneLabels.push({ element: label, zone, group });
          return mesh;
        });
        floorGroups.push({ floor, group, structure, exterior, zonesGroup, surface, shellMat, finMat, zoneMeshes });
        const label = document.createElement('button');
        label.className = 'floor-label';
        label.textContent = floor.id;
        label.setAttribute('aria-label', '选择 ' + floor.id + ' ' + floor.name);
        label.onclick = () => onSelect(floor.id, null);
        labelRoot.appendChild(label);
        labels.push(label);
        jobs.push(() => detailFloor(floorGroups[floor.index]));
      });
      scene.add(roofGroup);
      buildRoof();
      fitCamera(false);
      applyVisibility();
      renderer.shadowMap.needsUpdate = true;
    }
    /** detailFloor 为单层添加实例化竖梃、结构柱和室内隔断，不生成不可见微小零件。 */
    function detailFloor(entry) {
      const { floor, exterior, structure, finMat } = entry;
      const fins = [];
      for (let x = -floor.width / 2 + 1.2; x < floor.width / 2 - .5; x += floor.index <= 2 ? 2.4 : 1.22) {
        [floor.centerZ - floor.depth / 2, floor.centerZ + floor.depth / 2].forEach(z => fins.push([x, 1.93, z, .13, 3.65, .66]));
      }
      for (let z = floor.centerZ - floor.depth / 2 + 1.2; z < floor.centerZ + floor.depth / 2 - .5; z += 1.22) {
        [-floor.width / 2, floor.width / 2].forEach(x => fins.push([x, 1.93, z, .64, 3.65, .13]));
      }
      if (floor.index >= 3) instances(exterior, finMat, fins);
      const columns = [];
      [-1, 1].forEach(x => [-1, 1].forEach(z => columns.push([x * (floor.width / 2 - 2), 1.85, floor.centerZ + z * (floor.depth / 2 - 2), .46, 3.5, .46])));
      instances(structure, finMat, columns);
      // 墙体低于视线，区分区域但不掩盖热力色块。
      // 计量范围不是房间建筑墙体，区域间隙已形成边界。

      if (floor.index === 1) buildPodium(entry);
      if (floor.index === 2) {
        // 裙房顶板闭合外围空腔，功耗模式下随外壳隐藏。
        const deck = slab({ ...floor, width: 44, depth: 28 }, finMat, exterior);
        deck.position.y = 6.85;
        deck.userData = { floor: floor.id };
        pickables.push(deck);
      }
    }
    /** buildPodium 用连续两层高的折面包裹裙房，起伏下缘和收口上缘对应照片的悬挑外皮。 */
    function buildPodium(entry) {
      const { exterior, finMat } = entry; // 外皮随一层挂载，展开时整体隐藏。
      const positions = []; // 板块接缝留白，避免引入纹理请求。
      const points = roundedShape(44, 28, 4.2).getSpacedPoints(48);
      for (let i = 0; i < 48; i++) {
        const a = points[i].clone().lerp(points[i + 1], .006);
        const b = points[i + 1].clone().lerp(points[i], .006);
        const middle = a.clone().lerp(b, .5);
        const normal = new T.Vector2(middle.x / 44, middle.y / 28).normalize();
        const bottom = 1.15; // 下缘波形按空间坐标计算，邻板端点保持连续。
        const rows = [
          { y: bottom, offset: -.35 }, { y: 5.7, offset: 1.35 },
          { y: 8.7, offset: .8 }, { y: 10.9, offset: -.15 },
        ];
        // 每块板的中脊略外凸，形成宽板折面，避免细密波纹抢过建筑轮廓。
        const grid = rows.map(row => [a, middle, b].map((point, column) => [
          point.x + normal.x * (row.offset + (column === 1 ? .22 : 0)),
          row.y === bottom ? bottom + 2.1 * (.5 + .5 * Math.sin(point.x * .15 + point.y * .18)) : row.y,
          point.y + normal.y * (row.offset + (column === 1 ? .22 : 0)),
        ]));
        for (let row = 0; row < grid.length - 1; row++) {
          for (let column = 0; column < 2; column++) {
            [grid[row][column], grid[row + 1][column], grid[row][column + 1],
              grid[row + 1][column], grid[row + 1][column + 1], grid[row][column + 1]].forEach(vertex => positions.push(...vertex));
          }
        }
      }
      const geometry = new T.BufferGeometry(); // 少量三角面覆盖整圈，阴影来自真实折面法线。
      geometry.setAttribute('position', new T.Float32BufferAttribute(positions, 3));
      geometry.computeVertexNormals();
      geometries.add(geometry);
      finMat.side = T.DoubleSide;
      const panel = new T.Mesh(geometry, finMat);
      panel.castShadow = true; panel.receiveShadow = true;
      exterior.add(panel);
      panel.userData = { floor: '1F' };
      pickables.push(panel);
    }
    /** buildRoof 按照片制作向右下弯的连续冠顶与通高幕墙；展开仍使用独立空间楼层。 */
    function buildRoof() {
      const profile = [ // 世界高度：左侧圆肩最高，长边向右逐渐落到裙房。
        ...new T.CubicBezierCurve3(new T.Vector3(-15.4, 30, 0), new T.Vector3(-15.4, 33.4, 0), new T.Vector3(-10.4, 34.6, 0), new T.Vector3(-5.8, 33.6, 0)).getPoints(36),
        ...new T.CubicBezierCurve3(new T.Vector3(-5.8, 33.6, 0), new T.Vector3(4.5, 31.7, 0), new T.Vector3(11.5, 22.5, 0), new T.Vector3(15.4, 8.4, 0)).getPoints(60).slice(1),
      ];
      /** roofHeight 查找幕墙横坐标对应的屋顶高度，截断玻璃与竖梃避免穿插。 */
      function roofHeight(x) {
        const index = profile.findIndex(point => point.x >= x);
        const a = profile[Math.max(0, index - 1)], b = profile[index < 0 ? profile.length - 1 : index];
        return T.MathUtils.lerp(a.y, b.y, T.MathUtils.clamp((x - a.x) / Math.max(.001, b.x - a.x), 0, 1));
      }
      const shape = new T.Shape(); // 薄壳厚度仅 .65，消除原先厚重的帽檐。
      shape.moveTo(profile[0].x, profile[0].y);
      profile.slice(1).forEach(point => shape.lineTo(point.x, point.y));
      profile.slice().reverse().forEach(point => shape.lineTo(point.x, point.y - .65));
      shape.closePath();
      const geometry = new T.ExtrudeGeometry(shape, { depth: 19.4, bevelEnabled: true, bevelSegments: 2, steps: 1, bevelSize: .18, bevelThickness: .18 });
      geometries.add(geometry);
      const roof = new T.Mesh(geometry, edgeWhite);
      roof.position.z = -13.2;
      roof.castShadow = true; roof.receiveShadow = true;
      roof.userData = { floor: '8F' };
      roofGroup.add(roof); pickables.push(roof);
      const fins = [], rails = [], pier = []; // 连续竖梃与弱化的深色横框分别实例化。
      // 玻璃按楼层裁切成拾取条带，外观保持连续，点击仍能定位当前高度的楼层。
      config.floors.filter(floor => floor.index >= 3).forEach(floor => {
        const panels = [];
        for (let x = -14.75; x < 15; x += .74) {
          const top = Math.min(floor.id === '8F' ? 40 : floor.y + 3.8, roofHeight(x + .35) - .72);
          const bottom = Math.max(10.55, floor.y);
          if (top <= bottom) continue;
          [-12.85, 5.85].forEach(z => panels.push([x, (bottom + top) / 2, z, .735, top - bottom, .15]));
        }
        if (panels.length) {
          const mesh = instances(roofGroup, glass, panels);
          mesh.userData = { floor: floor.id }; pickables.push(mesh);
        }
        // 短边幕墙被宽圆角立柱包裹，使用深色横框避免楼板切断竖向节奏。
        const top = Math.min(floor.y + 3.8, 29.3);
        const bottom = Math.max(10.55, floor.y); // 裙房内不再重复构造幕墙。
        if (top > bottom) {
          const wall = meshBox(roofGroup, glass, -15.03, (bottom + top) / 2, -3.5, .15, top - bottom, 16.5);
          wall.userData = { floor: floor.id }; pickables.push(wall);
        }
        for (let x = -14.7; x < 14.9; x += .74) {
          if (roofHeight(x + .35) < floor.y + .8) continue;
          [-12.96, 5.96].forEach(z => rails.push([x, floor.y + .1, z, .74, .055, .06]));
        }
      });
      for (let x = -14.7; x < 15; x += .74) {
        const height = roofHeight(x + .35) - .8 - 10.55;
        if (height <= 0) continue;
        [-13.02, 6.02].forEach(z => fins.push([x, 10.55 + height / 2, z, .085, height, .38]));
      }
      for (let z = -11.6; z < 5.2; z += .8) fins.push([-15.2, 19.9, z, .4, 18.7, .09]);
      // 照片中的宽白色转角从裙房连至圆肩，以圆角立柱柔化两片幕墙交界。
      const columnGeometry = new T.CylinderGeometry(1, 1, 1, 12);
      geometries.add(columnGeometry);
      [-12.3, 5.2].forEach(z => pier.push([-14.7, 20.2, z, 1.15, 19.4, 1.35]));
      const columns = instances(roofGroup, edgeWhite, pier, columnGeometry);
      columns.userData = { floor: '5F' }; pickables.push(columns);
      instances(roofGroup, white, fins);
      instances(roofGroup, seam, rails);
      // 屋面采用细线分格，低成本呈现白色金属板拼缝。
      const joints = [];
      for (let z = -12; z < 6; z += 2.4) {
        for (let i = 1; i < profile.length; i++) joints.push(new T.Vector3(profile[i - 1].x, profile[i - 1].y + .2, z), new T.Vector3(profile[i].x, profile[i].y + .2, z));
      }
      const jointGeometry = new T.BufferGeometry().setFromPoints(joints);
      const jointMaterial = new T.LineBasicMaterial({ color: 0xabb6b8, transparent: true, opacity: .3 });
      geometries.add(jointGeometry); materials.push(jointMaterial);
      roofGroup.add(new T.LineSegments(jointGeometry, jointMaterial));
    }
    /** detailSite 实例化铺装接缝、花池和树木，场地只作为建筑比例参照。 */
    function detailSite() {
      const joints = [];
      for (let x = -29; x <= 29; x += 3) joints.push([x, -.18, 0, .025, .015, 40]);
      for (let z = -19; z <= 19; z += 3) joints.push([0, -.18, z, 60, .015, .025]);
      instances(stage, seam, joints);
      [-27, 27].forEach(x => meshBox(stage, lawn, x, -.12, -1, 4.1, .3, 34));
      meshBox(stage, paving, 0, -.05, 16, 14, .25, 3);
      meshBox(stage, white, 0, 3.4, 14.2, 13.5, .2, 2.8);
      const trunks = [], crowns = [];
      [-27, 27].forEach(x => [-14, -7, 0, 7, 14].forEach((z, i) => {
        trunks.push([x, 1.5, z, .15, 3.1, .15]);
        crowns.push([x, 3.2, z, 1.5 + i % 2 * .35, 2.1, 1.5]);
        crowns.push([x - .65, 3.1, z + .35, .95, 1.4, 1]);
      }));
      instances(stage, trunkMat, trunks);
      const leaf = new T.IcosahedronGeometry(1, 1);
      geometries.add(leaf);
      instances(stage, foliageMat, crowns, leaf);
      const bollards = [];
      [-18, -12, -6, 6, 12, 18].forEach(x => bollards.push([x, .45, 18.8, .13, 1.1, .13]));
      instances(stage, seam, bollards);
    }
    /** environment 以小尺寸本地立方体贴图补充柔和反射，无在线纹理请求。 */
    function environment() {
      const faces = Array.from({ length: 6 }, (_, index) => {
        const canvas = document.createElement('canvas');
        canvas.width = canvas.height = 64;
        const ctx = canvas.getContext('2d');
        const gradient = ctx.createLinearGradient(0, 0, 0, 64);
        gradient.addColorStop(0, index === 2 ? '#e3eef6' : '#8da6ba');
        gradient.addColorStop(.5, '#a3b4bc');
        gradient.addColorStop(1, '#3f5355');
        ctx.fillStyle = gradient;
        ctx.fillRect(0, 0, 64, 64);
        ctx.fillStyle = '#e0e8df';
        ctx.fillRect(8, 0, 6, 36);
        ctx.fillRect(38, 0, 12, 24);
        return canvas;
      });
      const cube = new T.CubeTexture(faces);
      cube.encoding = T.sRGBEncoding;
      cube.needsUpdate = true;
      textures.add(cube);
      scene.environment = cube;
      materials.forEach(mat => { mat.envMapIntensity = .15; mat.needsUpdate = true; });
    }
    /** fitCamera 基于主体包围盒和画布宽高适配镜头，确保展开后顶部和地下层不裁切。 */
    function fitCamera(animate = true) {
      const top = 35.5 + (state.expanded ? 8 * state.spacing : 0);
      const centerY = state.floor ? config.floors.find(floor => floor.id === state.floor).y + (state.expanded ? config.floors.find(floor => floor.id === state.floor).index * state.spacing : 0) + 1.8 : (top - 14) / 2;
      const target = new T.Vector3(0, centerY, -1.5);
      const height = state.floor ? 25 : top + 10;
      const width = state.floor ? 48 : 70;
      const distance = Math.max(height / 2 / Math.tan(T.MathUtils.degToRad(18)), width / 2 / Math.tan(T.MathUtils.degToRad(18)) / camera.aspect) * (state.floor ? 1.42 : 1.56);
      const direction = new T.Vector3(-.86, state.floor ? 1.15 : .43, 1.23).normalize();
      const position = target.clone().addScaledVector(direction, distance);
      // 以场地/楼层边界投影复核构图，防止透视近大远小把底座裁出画布。
      if (!state.floor) {
        const extentX = state.expanded ? 23 : 33;
        const extentZ = state.expanded ? 16 : 23;
        for (let attempt = 0; attempt < 12; attempt++) {
          const fittingCamera = camera.clone();
          fittingCamera.position.copy(position);
          fittingCamera.lookAt(target);
          fittingCamera.updateMatrixWorld();
          let clipped = false;
          [-extentX, extentX].forEach(x => [-4, top + 1].forEach(y => [-extentZ, extentZ].forEach(z => {
            const point = new T.Vector3(x, y, z).project(fittingCamera);
            if (Math.abs(point.x) > .92 || Math.abs(point.y) > .92) clipped = true;
          })));
          if (!clipped) break;
          position.sub(target).multiplyScalar(1.05).add(target);
        }
      }
      controls.maxDistance = Math.max(220, distance * 2);
      if (animate) cameraTween = { target, position };
      else { camera.position.copy(position); controls.target.copy(target); controls.update(); }
    }
    /** applyVisibility 根据展示模式与选层切换外壳/热力地板；B1 在整楼闭合时隐藏。 */
    function applyVisibility() {
      const showZones = state.mode === 'energy' || state.expanded;
      floorGroups.forEach(entry => {
        const { floor, group, exterior, surface, finMat, zonesGroup, zoneMeshes } = entry;
        const sharedSelection = config.floors.find(f => f.id === state.floor)?.sharedGroup; // 同一跨层计量范围统一着色，不复制汇总值。
        const dim = state.floor && state.floor !== floor.id && !(sharedSelection && floor.sharedGroup===sharedSelection);
        group.visible = floor.index !== 0 || state.expanded || state.floor === 'B1';
        exterior.visible = !showZones && floor.index <= 2;
        entry.structure.visible = showZones; // 闭合外观采用连续外壳，避免示意楼板穿出曲面。
        zonesGroup.visible = showZones;
        surface.color.set(showZones ? (dim ? 0x42545d : 0xa1b6b8) : 0xd4dcda).convertSRGBToLinear();
        surface.transparent = !!dim;
        surface.opacity = dim ? .13 : 1;
        surface.depthWrite = !dim;
        finMat.color.set(dim ? 0x677b83 : 0xd8ddda).convertSRGBToLinear();
        finMat.transparent = !!dim;
        finMat.opacity = dim ? .12 : 1;
        finMat.depthWrite = !dim;
        zoneMeshes.forEach(mesh => {
          const sample = data.values[mesh.userData.zone];
          mesh.material.color.set(data.color(data.loadRatio(sample))).convertSRGBToLinear();
          mesh.material.transparent = !!dim;
          mesh.material.opacity = dim ? .1 : 1;
          mesh.material.depthWrite = !dim;
          if (dim) mesh.material.color.lerp(new T.Color(0x233e4c), .74);
          mesh.material.emissive.set(mesh.userData.zone === state.zone ? data.color(data.loadRatio(sample)) : 0);
          mesh.material.emissiveIntensity = .2;
          mesh.scale.y = mesh.userData.zone === state.zone ? .65 : .3;
        });
      });
      zoneLabels.forEach(({ element, zone }) => {
        element.hidden = !showZones || zone.floorId !== state.floor;
        element.classList.toggle('selected', zone.id === state.zone);
      });
      floorGroups.forEach(entry => entry.group.traverse(object => {
        if (!object.isMesh) return;
        if (object.userData.initialShadow === undefined) object.userData.initialShadow = object.castShadow;
        object.castShadow = object.userData.initialShadow && (!state.floor || state.floor === entry.floor.id);
      }));
      selectionOutline.visible = !!state.zone;
      roofGroup.visible = !showZones;
      stage.visible = !state.expanded && state.floor !== 'B1';
      labels.forEach((label, index) => {
        label.hidden = !showZones || !floorGroups[index].group.visible || !!(state.floor && state.floor !== floorGroups[index].floor.id);
        label.classList.toggle('selected', state.floor === floorGroups[index].floor.id);
      });
      renderer.shadowMap.needsUpdate = true;
    }
    /** setState 合并新的模式和选择，并让新动画覆盖旧目标而不重置当前位置。 */
    function setState(next, fit = true) {
      Object.assign(state, next);
      applyVisibility();
      if (fit) fitCamera();
    }
    /** resize 使用真实画布尺寸而非宿主或窗口尺寸，防止拾取坐标偏移。 */
    function resize() {
      const rect = container.getBoundingClientRect();
      renderer.setSize(rect.width, rect.height);
      camera.aspect = rect.width / Math.max(1, rect.height);
      camera.updateProjectionMatrix();
      fitCamera();
    }
    /** hit 优先命中当前楼层；淡化楼层的外露楼板用于切层，隐藏对象仍不可拾取。 */
    function hit(event) {
      const rect = renderer.domElement.getBoundingClientRect();
      pointer.set((event.clientX - rect.left) / rect.width * 2 - 1, -(event.clientY - rect.top) / rect.height * 2 + 1);
      // 输入可发生在两帧之间，先同步矩阵以使用当前的模型与镜头位置。
      scene.updateMatrixWorld(true);
      camera.updateMatrixWorld();
      ray.setFromCamera(pointer, camera);
      const visible = pickables.filter(mesh => {
        // 其他层只保留楼板切换入口，不能隔着淡化层直接选中其区域。
        if (state.floor && mesh.userData.floor !== state.floor && mesh.userData.zone) return false;
        let parent = mesh;
        while (parent) { if (!parent.visible) return false; parent = parent.parent; }
        return true;
      });
      const hits = ray.intersectObjects(visible, false); // 按射线距离排序，外露处选择最近楼板。
      // 当前区域与淡化楼层重叠时，优先当前层，避免上层透明板拦截选区。
      return (state.floor && hits.find(hit => hit.object.userData.floor === state.floor)?.object) || hits[0]?.object || null;
    }
    /** pointerDown 锁定一次主指针点击的起点与对象，不要求鼠标完全静止。 */
    function pointerDown(event) {
      if (pressedPointer !== null || event.isPrimary === false || event.button !== 0) {
        moved = true;
        pressedObject = null;
        return;
      }
      down.set(event.clientX, event.clientY);
      moved = false;
      pressedPointer = event.pointerId;
      pressedObject = hit(event);
    }
    /** trackPointer 持续记录最大位移，离开画布后返回也不会被误认为点击。 */
    function trackPointer(event) {
      if (event.pointerId !== pressedPointer) return;
      if (Math.hypot(event.clientX - down.x, event.clientY - down.y) > clickSlop) moved = true;
    }
    /** cancelPointer 清除取消或失焦的手势，防止后续松手误选。 */
    function cancelPointer() { pressedPointer = null; pressedObject = null; moved = false; }
    /** pointerUp 单击即选中并展开；越过阈值的旋转操作绝不触发选择。 */
    function pointerUp(event) {
      if (event.pointerId !== pressedPointer) return;
      trackPointer(event);
      const object = pressedObject;
      const bounds = renderer.domElement.getBoundingClientRect();
      const inside = event.clientX >= bounds.left && event.clientX <= bounds.right && event.clientY >= bounds.top && event.clientY <= bounds.bottom;
      const select = event.button === 0 && !moved && inside && object;
      cancelPointer();
      if (select) onSelect(object.userData.floor, object.userData.zone || null);
    }
    /** pointerMove 只显示轻量 DOM 提示；按住时由独立手势跟踪决定是否为拖动。 */
    function pointerMove(event) {
      if (event.buttons) { tooltip.hidden = true; return; }
      const object = hit(event);
      renderer.domElement.style.cursor = object ? 'pointer' : 'grab';
      if (!object) { tooltip.hidden = true; return; }
      const floor = config.floors.find(value => value.id === object.userData.floor);
      const zone = floor.zones.find(value => value.id === object.userData.zone);
      const summary = data.aggregate(zone ? [zone] : floor.zones);
      tooltip.innerHTML = floor.id + ' · ' + (zone ? zone.name : floor.name) + '<small>' + (summary.power === null ? '暂无数据' : summary.power.toFixed(2) + ' kW') + '</small>';
      const rect = container.getBoundingClientRect();
      tooltip.style.left = Math.max(4, Math.min(rect.width - 210, event.clientX - rect.left + 15)) + 'px';
      tooltip.style.top = Math.max(4, Math.min(rect.height - 70, event.clientY - rect.top + 12)) + 'px';
      tooltip.hidden = false;
    }
    /** render 更新楼层位置与镜头，完成首帧后才撤下主体加载提示。 */
    function render(now) {
      if (disposed) return;
      const dt = Math.min(.08, (now - lastTime) / 1000);
      lastTime = now;
      const blend = 1 - Math.exp(-dt * 8);
      let changed = false;
      floorGroups.forEach((entry, index) => {
        const target = entry.floor.y + (state.expanded ? index * state.spacing : 0);
        if (Math.abs(entry.group.position.y - target) > .005) {
          entry.group.position.y = T.MathUtils.lerp(entry.group.position.y, target, blend);
          changed = true;
        } else entry.group.position.y = target;
        if (!labels[index].hidden) {
          const position = new T.Vector3(entry.floor.width / 2 + 1.5, entry.group.position.y + .4, entry.floor.centerZ + entry.floor.depth / 2).project(camera);
          labels[index].style.left = Math.min(container.clientWidth - 42, Math.max(16, (position.x * .5 + .5) * container.clientWidth)) + 'px';
          labels[index].style.top = Math.min(container.clientHeight - 36, Math.max(15, (-position.y * .5 + .5) * container.clientHeight)) + 'px';
        }
      });
      zoneLabels.forEach(({ element, zone, group }) => {
        if (element.hidden) return;
        const projected = new T.Vector3(zone.x, group.position.y + .8, zone.z).project(camera);
        element.style.left = (projected.x * .5 + .5) * container.clientWidth + 'px';
        element.style.top = (-projected.y * .5 + .5) * container.clientHeight + 'px';
        if (zone.id === state.zone) {
          selectionOutline.position.set(zone.x, group.position.y + .85, zone.z);
          selectionOutline.scale.set(zone.width, 1, zone.depth);
        }
      });
      if (cameraTween) {
        camera.position.lerp(cameraTween.position, blend);
        controls.target.lerp(cameraTween.target, blend);
        if (camera.position.distanceTo(cameraTween.position) < .02) cameraTween = null;
      }
      controls.update();
      if (changed) renderer.shadowMap.needsUpdate = true;
      renderer.render(scene, camera);
      if (metrics.bodyMs === null) {
        metrics.bodyMs = performance.now();
        container.querySelector('#loading').hidden = true;
        document.getElementById('scene-status').textContent = '主体就绪 · 补充立面细节';
        drainChannel.port2.postMessage(0);
      }
      frameId = requestAnimationFrame(render);
    }
    /** nextBatch 每帧最多完成一项细节任务，让镜头与浏览器输入获得处理时间。 */
    function nextBatch() {
      if (disposed) return;
      const job = jobs.shift();
      if (job) { job(); applyVisibility(); }
      if (jobs.length) requestAnimationFrame(() => drainChannel.port2.postMessage(0));
      else requestAnimationFrame(() => {
        if (disposed) return;
        renderer.shadowMap.needsUpdate = true;
        renderer.render(scene, camera);
        metrics.detailMs = performance.now();
        metrics.ready = true;
        document.getElementById('scene-status').textContent = '细节就绪 · 可自由探索';
      });
    }
    controls.addEventListener('start', () => { cameraTween = null; });
    renderer.domElement.addEventListener('pointerdown', pointerDown);
    renderer.domElement.ownerDocument.addEventListener('pointerup', pointerUp);
    renderer.domElement.ownerDocument.addEventListener('pointermove', trackPointer);
    renderer.domElement.ownerDocument.addEventListener('pointercancel', cancelPointer);
    root.addEventListener('blur', cancelPointer);
    renderer.domElement.addEventListener('pointermove', pointerMove);
    renderer.domElement.addEventListener('pointerleave', () => { tooltip.hidden = true; });
    renderer.domElement.addEventListener('webglcontextlost', event => {
      event.preventDefault();
      container.querySelector('#loading').hidden = false;
      document.getElementById('load-text').textContent = '图形上下文已中断，请刷新页面重试';
      metrics.ready = false;
    });
    const observer = new ResizeObserver(resize);
    observer.observe(container);
    drainChannel.port1.onmessage = nextBatch;
    const rect = container.getBoundingClientRect();
    camera.aspect = rect.width / rect.height;
    camera.updateProjectionMatrix();
    renderer.setSize(rect.width, rect.height);
    buildBody();
    jobs.push(detailSite, environment);
    frameId = requestAnimationFrame(render);
    /** dispose 释放当前原型拥有的渲染器、材质与纹理，避免重复打开累计占用。 */
    function dispose() {
      disposed = true;
      cancelAnimationFrame(frameId);
      observer.disconnect();
      drainChannel.port1.close(); drainChannel.port2.close();
      controls.dispose();
      renderer.domElement.ownerDocument.removeEventListener('pointerup', pointerUp);
      renderer.domElement.ownerDocument.removeEventListener('pointermove', trackPointer);
      renderer.domElement.ownerDocument.removeEventListener('pointercancel', cancelPointer);
      root.removeEventListener('blur', cancelPointer);
      geometries.forEach(value => value.dispose());
      textures.forEach(value => value.dispose());
      materials.forEach(value => value.dispose());
      renderer.dispose();
    }
    root.addEventListener('pagehide', dispose, { once: true });
    return { setState, state, metrics, dispose,
      // 遥测只刷新颜色和标签，不触发镜头动画或重建几何。
      refreshData: () => { applyVisibility(); zoneLabels.forEach(({element, zone}) => { const value=data.values[zone.id]; element.replaceChildren(document.createTextNode(zone.name)); const text=document.createElement('small'); text.textContent=value.power===null?'暂无有效数据':value.power.toFixed(2)+' kW'; element.appendChild(text); element.style.setProperty('--zone-color',data.color(data.loadRatio(value))); }); },
      inspect: () => ({ ...metrics, calls: renderer.info.render.calls, triangles: renderer.info.render.triangles, geometries: renderer.info.memory.geometries, textures: renderer.info.memory.textures, pixelRatio: renderer.getPixelRatio(), groups: floorGroups.length, positions: floorGroups.map(entry => entry.group.position.y), animating: !!cameraTween || floorGroups.some(entry => Math.abs(entry.group.position.y - entry.floor.y - (state.expanded ? entry.floor.index * state.spacing : 0)) > .01), mode: state.mode, expanded: state.expanded, floor: state.floor, zone: state.zone }),
      // 可复现验收坐标仅投影已存在区域；浏览器测试仍以真实鼠标点击触发拾取。
      projectZone: (id, offset = 0) => {
        const entry = floorGroups.find(value => value.floor.zones.some(zone => zone.id === id));
        const zone = entry.floor.zones.find(value => value.id === id);
        const position = new T.Vector3(zone.x + zone.width * offset, entry.group.position.y + .58, zone.z + zone.depth * offset).project(camera);
        const bounds = renderer.domElement.getBoundingClientRect();
        return { x: bounds.left + (position.x * .5 + .5) * bounds.width, y: bounds.top + (-position.y * .5 + .5) * bounds.height };
      },
    };
  };
})(window);
