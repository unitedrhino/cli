/* 文件说明：配电站轻量写实外观。小纹理与几何缓存、设备五金、空间细节、
 * 室内环境反射及静态实例化；不访问网络，不参与业务数据绑定。 */
(function () {
  const T = THREE; // 包内 Three.js r128。
  const geometries = new Map(); // 按尺寸复用静态零件，避免重复分配顶点。
  const metal = new T.MeshStandardMaterial({ color: 0x9ba6b0, metalness: 0.72, roughness: 0.32 }); // 镀锌五金。
  const rubber = new T.MeshStandardMaterial({ color: 0x1e2928, roughness: 0.84 }); // 门缝与绝缘塑料。
  const edge = new T.MeshStandardMaterial({ color: 0xb7c0c6, metalness: 0.35, roughness: 0.48 }); // 折边与百叶。
  [metal, rubber, edge].forEach(material => material.color.convertSRGBToLinear());
  const screw = new T.CylinderGeometry(0.009, 0.009, 0.006, 6); // 小螺栓只使用六边形截面。
  screw.rotateX(Math.PI / 2);

  /** beveledBox 根据外尺寸和倒角宽度返回缓存几何；尺寸单位为米。 */
  function beveledBox(w, h, d, radius) {
    const r = Math.min(radius, w / 5, h / 5, d / 5); // 限制薄面板倒角，防止自交。
    const key = [w, h, d, r].join('/'); // 含倒角的缓存键。
    if (!geometries.has(key)) {
      const shape = new T.Shape(); // 内轮廓，挤出倒角后仍保持原来的外尺寸。
      shape.moveTo(-w / 2 + r, -h / 2 + r);
      shape.lineTo(w / 2 - r, -h / 2 + r);
      shape.lineTo(w / 2 - r, h / 2 - r);
      shape.lineTo(-w / 2 + r, h / 2 - r);
      shape.closePath();
      const geometry = new T.ExtrudeGeometry(shape, { depth: d - r * 2, bevelEnabled: true,
        bevelSegments: 1, steps: 1, bevelSize: r, bevelThickness: r }); // 单段倒角控制面数。
      geometry.translate(0, 0, -(d - r * 2) / 2);
      /* 默认挤出 UV 使用世界尺寸，负半轴被 Clamp 拉成大片色块；各面归一化后涂层才能连续。 */
      const position = geometry.attributes.position; // 倒角完成后的局部顶点。
      const normal = geometry.attributes.normal; // 用主法向选择正面、侧面或顶底面。
      const uv = geometry.attributes.uv; // 复用原缓冲，不创建额外顶点。
      for (let i = 0; i < position.count; i++) {
        const nx = Math.abs(normal.getX(i)), ny = Math.abs(normal.getY(i)), nz = Math.abs(normal.getZ(i));
        if (nz >= nx && nz >= ny) uv.setXY(i, position.getX(i) / w + 0.5, position.getY(i) / h + 0.5);
        else if (nx >= ny) uv.setXY(i, position.getZ(i) / d + 0.5, position.getY(i) / h + 0.5);
        else uv.setXY(i, position.getX(i) / w + 0.5, position.getZ(i) / d + 0.5);
      }
      geometries.set(key, geometry);
    }
    return geometries.get(key);
  }

  /** part 创建带位置、尺寸和材质的五金盒；返回网格以便调用方设置方向。 */
  function part(parent, w, h, d, x, y, z, material) {
    const key = [w, h, d].join(':'); // 普通盒与倒角缓存键分开。
    if (!geometries.has(key)) geometries.set(key, new T.BoxGeometry(w, h, d));
    const mesh = new T.Mesh(geometries.get(key), material || metal); // 静态零件不单独投影。
    mesh.position.set(x, y, z);
    mesh.receiveShadow = true;
    parent.add(mesh);
    return mesh;
  }

  /** applySurface 构造共享颗粒、粗糙度与切线法线贴图；返回地坪复用资源，全部使用128px循环纹理。 */
  function applySurface(materials) {
    const canvas = document.createElement('canvas'); // 共享灰度粗糙度图。
    canvas.width = canvas.height = 128;
    const context = canvas.getContext('2d'); // 只在初始化绘制一次。
    const pixels = context.createImageData(128, 128); // 固定种子保证截图一致。
    let seed = 716; // 程序材质专用种子，仅影响表面纹理，不生成任何读数。
    for (let i = 0; i < pixels.data.length; i += 4) {
      seed = (1664525 * seed + 1013904223) >>> 0;
      const gray = 240 + (seed >>> 28); // 小幅扰动避免镜面与磨砂之间产生噪点。
      pixels.data[i] = pixels.data[i + 1] = pixels.data[i + 2] = gray;
      pixels.data[i + 3] = 255;
    }
    context.putImageData(pixels, 0, 0);
    const bump = new T.CanvasTexture(canvas); // 粗糙度图为线性数据，不设置 sRGB。
    bump.wrapS = bump.wrapT = T.RepeatWrapping;
    bump.repeat.set(8, 8);
    const normalPixels = new Uint8Array(128 * 128 * 4); // 从循环高度场求法线，边界取模避免铺设接缝。
    for (let y = 0; y < 128; y++) {
      for (let x = 0; x < 128; x++) {
        const index = (y * 128 + x) * 4; // 当前像素和周期相邻像素的索引。
        const dx = pixels.data[(y * 128 + (x + 1) % 128) * 4] - pixels.data[(y * 128 + (x + 127) % 128) * 4];
        const dy = pixels.data[(((y + 1) % 128) * 128 + x) * 4] - pixels.data[(((y + 127) % 128) * 128 + x) * 4];
        normalPixels[index] = 128 - dx * 2;
        normalPixels[index + 1] = 128 - dy * 2;
        normalPixels[index + 2] = 255;
        normalPixels[index + 3] = 255;
      }
    }
    const normal = new T.DataTexture(normalPixels, 128, 128); // 法线是线性数据，不使用sRGB转换。
    normal.wrapS = normal.wrapT = T.RepeatWrapping;
    normal.minFilter = T.LinearMipmapLinearFilter;
    normal.magFilter = T.LinearFilter;
    normal.generateMipmaps = true;
    normal.repeat.set(8, 8);
    normal.needsUpdate = true;
    materials.wall.normalMap = normal;
    materials.wall.normalScale.set(0.34, 0.34);
    materials.wall.roughnessMap = bump;
    const panel = document.createElement('canvas'); // 喷涂面板轻微边缘压暗，与几何折边配合。
    panel.width = panel.height = 64;
    const paint = panel.getContext('2d');
    const shade = paint.createLinearGradient(0, 0, 64, 0);
    shade.addColorStop(0, '#adb2b7'); shade.addColorStop(0.04, '#eaedef');
    shade.addColorStop(0.35, '#ffffff'); shade.addColorStop(0.92, '#f5f6f7');
    shade.addColorStop(1, '#bac1c5');
    paint.fillStyle = shade; paint.fillRect(0, 0, 64, 64);
    materials.seam.map = new T.CanvasTexture(panel);
    materials.seam.map.encoding = T.sRGBEncoding;
    ['cabinet', 'cabinetHV', 'transCase', 'seam'].forEach(key => {
      materials[key].roughnessMap = bump;
      materials[key].normalMap = normal;
      materials[key].normalScale.set(0.16, 0.16);
      materials[key].envMapIntensity = 0.65;
    });
    return { normal, roughness: bump, grain: canvas };
  }

  /** cabinet 细化柜体的五金与通风结构；kind 区分设备外壳，不改变业务台账。 */
  function cabinet(group, w, h, d, kind) {
    const front = d / 2 + 0.055; // 高于原有前门表面，避免共面闪烁。
    [-1, 1].forEach(side => {
      part(group, 0.012, h - 0.24, 0.008, side * (w / 2 - 0.035), h / 2, front - 0.01, rubber);
      [0.35, h - 0.46].forEach(y => {
        part(group, 0.035, 0.085, 0.028, side * (w / 2 - 0.045), y, front, metal);
      });
      [0.17, h - 0.09].forEach(y => {
        const bolt = new T.Mesh(screw, metal); // 四角压紧螺栓。
        bolt.position.set(side * (w / 2 - 0.025), y, front);
        group.add(bolt);
      });
    });
    part(group, w - 0.045, 0.025, d - 0.025, 0, h - 0.01, 0, edge);
    part(group, w - 0.035, 0.012, 0.022, 0, 0.125, front, rubber);
    if (kind === 'feeder') {
      /* 抽屉自身已有独立表计与拉手，仅补立柱和面板折边。 */
      part(group, 0.018, h - 0.58, 0.017, 0, 0.97, front - 0.005, rubber);
    } else {
      part(group, 0.016, 0.155, 0.045, w / 2 - 0.088, h * 0.47, front + 0.014, metal);
      const rows = kind === 'transformer' ? 18 : 5; // 变压器使用真实有倾角的百叶片。
      for (let i = 0; i < rows; i++) {
        const y = kind === 'transformer' ? 0.44 + i * 0.071 : 0.24 + i * 0.039;
        part(group, w * 0.72, 0.033, 0.008, 0, y + 0.002, front - 0.009, rubber);
        const slat = part(group, w * 0.72, 0.023, 0.035, 0, y, front + 0.002, edge);
        slat.rotation.x = -0.3;
      }
    }
    /* 柜背及侧面为结构展示；浅层折边让斜视角可读。 */
    for (let i = 0; i < 6; i++) {
      part(group, 0.008, 0.012, d * 0.46, w / 2 + 0.002, 0.28 + i * 0.04, 0, rubber);
    }
  }

  /** room 补充房间细节；depth/centerZ为墙体深度与中心，踢脚随墙延伸，设备及沟盖板位置保持原布局。 */
  function room(scene, materials, depth, centerZ) {
    const trim = new T.MeshStandardMaterial({ color: 0x354b45, roughness: 0.64 }); // 墙脚防污饰面。
    part(scene, 15, 0.2, 0.025, 0, 0.1, -4.98, trim);
    [-7.48, 7.48].forEach(x => part(scene, 0.025, 0.2, depth, x, 0.1, centerZ, trim));
    /* 盖板沿既有电缆沟排布，表面略高于地坪。 */
    [-5.4, 5.4].forEach(x => {
      const count = x < 0 ? 10 : 14; // Ⅰ、Ⅱ段沟长不同。
      for (let i = 0; i < count; i++) {
        const z = -3.25 + i * 0.5; // 单块半米盖板。
        part(scene, 0.285, 0.012, 0.483, x, 0.009, z, metal);
        part(scene, 0.07, 0.002, 0.012, x, 0.017, z, rubber);
      }
    });
    const lightMat = new T.MeshStandardMaterial({ color: 0xfff5df, emissive: 0xfff1d8,
      emissiveIntensity: 1.6, roughness: 0.7 }); // 灯具可见发光表面，不增加影子光源。
    [-5, 0, 5].forEach(x => {
      part(scene, 1.5, 0.11, 0.17, x, 3.18, -4.94, materials.dark);
      part(scene, 1.36, 0.052, 0.18, x, 3.16, -4.92, lightMat);
    });
    /* 柜底使用共用径向灰度软阴影，补充实时阴影难以覆盖的接触层次。 */
    const canvas = document.createElement('canvas'); // 64px 软阴影可任意拉伸。
    canvas.width = canvas.height = 64;
    const ctx = canvas.getContext('2d');
    const gradient = ctx.createRadialGradient(32, 32, 12, 32, 32, 32);
    gradient.addColorStop(0, 'rgba(8,18,14,0.42)');
    gradient.addColorStop(1, 'rgba(8,18,14,0)');
    ctx.fillStyle = gradient; ctx.fillRect(0, 0, 64, 64);
    const contact = new T.MeshBasicMaterial({ map: new T.CanvasTexture(canvas), transparent: true,
      depthWrite: false, polygonOffset: true, polygonOffsetFactor: -1 }); // 只混合阴影，不遮挡地面。
    const plane = new T.PlaneGeometry(1, 1);
    [[-2, 1.6, 8.5, 1.6], [2.4, 4.1, 7.8, 1.6], [0, -4.1, 6, 1.6],
      [-4.15, -4.25, 2.8, 2], [4.15, -4.25, 2.8, 2]].forEach(([x, z, w, d]) => {
      const shadow = new T.Mesh(plane, contact);
      shadow.rotation.x = -Math.PI / 2;
      shadow.position.set(x, 0.018, z);
      shadow.scale.set(w, d, 1);
      scene.add(shadow);
    });
  }

  /** busJoints 沿两点之间的既有母线路径补接头，参数为世界坐标及安装高度。 */
  function busJoints(scene, x0, z0, x1, z1, height) {
    const length = Math.hypot(x1 - x0, z1 - z0); // 外壳中轴长度，外观尺寸不作为工程选型参数。
    const count = Math.max(1, Math.ceil(length / 1.4)); // 不改变母线走向。
    const shell = new T.Group(); // 连续折边表现薄金属壳，避免整根混凝土梁的观感。
    shell.position.set((x0 + x1) / 2, height, (z0 + z1) / 2);
    shell.rotation.y = -Math.atan2(z1 - z0, x1 - x0);
    scene.add(shell);
    [-1, 1].forEach(side => {
      part(shell, length, 0.018, 0.018, 0, 0.069, side * 0.122, metal);
      part(shell, length, 0.009, 0.006, 0, -0.057, side * 0.127, rubber);
    });
    for (let i = 0; i <= count; i++) {
      const ratio = i / count; // 接头板沿直线段均布，转角两端使用相同封闭截面。
      const group = new T.Group(); // 接头零件最终合并为实例化批次。
      group.position.set(x0 + (x1 - x0) * ratio, height, z0 + (z1 - z0) * ratio);
      group.rotation.y = -Math.atan2(z1 - z0, x1 - x0);
      scene.add(group);
      part(group, 0.09, 0.196, 0.27, 0, 0, 0, edge);
      [-1, 1].forEach(side => {
        [-0.055, 0.055].forEach(y => {
          const bolt = new T.Mesh(screw, metal); // 螺栓复用柜体六角几何，不增加独立几何资源。
          bolt.position.set(0, y, side * 0.139);
          group.add(bolt);
        });
      });
    }
  }

  /** busTerminal 补齐母线至柜顶的封闭接入端；roof 为柜顶高，offset 为图示接口沿 z 的偏移。 */
  function busTerminal(scene, x, z, height, roof, offset) {
    const depth = Math.abs(offset) + 0.27; // 变压器正面接口回接外壳，低压柜直接向下接入。
    const top = height + 0.095; // 接口上沿与水平外壳平齐。
    part(scene, 0.27, top - roof, depth, x, (top + roof) / 2, z + offset / 2, metal);
    part(scene, 0.36, 0.022, 0.35, x, roof + 0.012, z + offset, edge);
    [-0.135, 0.135].forEach(dx => {
      const bolt = new T.Mesh(screw, metal); // 柜顶法兰连接螺栓。
      bolt.rotation.x = Math.PI / 2;
      bolt.position.set(x + dx, roof + 0.027, z + offset);
      scene.add(bolt);
    });
  }

  /** bakeEnvironment 仅供离线生成环境图；返回释放闭包，运行时不调用。 */
  function bakeEnvironment(renderer, scene) {
    /* r128 的 PMREM 输出固定布局，预处理仅在离线烘焙命令中执行。 */
    const panorama = document.createElement('canvas'); // 小全景图定义室内光照分布。
    panorama.width = 256; panorama.height = 128;
    const ctx = panorama.getContext('2d');
    const gradient = ctx.createLinearGradient(0, 0, 0, 128);
    gradient.addColorStop(0, '#d9dde1'); gradient.addColorStop(0.45, '#a5adb5');
    gradient.addColorStop(0.75, '#606c70'); gradient.addColorStop(1, '#343e40');
    ctx.fillStyle = gradient; ctx.fillRect(0, 0, 256, 128);
    ctx.fillStyle = '#fffefb';
    [12, 76, 140, 204].forEach(x => ctx.fillRect(x, 26, 32, 9));
    const texture = new T.CanvasTexture(panorama);
    texture.encoding = T.sRGBEncoding;
    texture.mapping = T.EquirectangularReflectionMapping;
    const generator = new T.PMREMGenerator(renderer);
    const target = generator.fromEquirectangular(texture);
    scene.environment = target.texture;
    generator.dispose(); texture.dispose();
    return () => { scene.environment = null; target.dispose(); };
  }

  /** environment 解码包内反射图，直接使用预滤波布局；返回异步释放闭包。 */
  function environment(renderer, scene) {
    return new Promise((resolve, reject) => {
      /** decode 在包内数据就绪后解码预滤波图；离线单页已内联数据可直接使用。 */
      function decode() {
      const atlas = window.StationEnvironment;
      new T.TextureLoader().load(atlas.image, texture => {
        texture.mapping = T.CubeUVReflectionMapping;
        texture.encoding = atlas.encoding;
        texture.flipY = false;
        texture.generateMipmaps = false;
        texture.minFilter = T.LinearFilter;
        texture.magFilter = T.LinearFilter;
        scene.environment = texture;
        resolve(() => { scene.environment = null; texture.dispose(); });
      }, undefined, reject);
      }
      if (window.StationEnvironment) decode();
      else {
        const script = document.createElement('script'); // 反射数据在主体出现后才请求。
        script.src = 'js/environment.js';
        script.onload = () => { script.remove(); decode(); };
        script.onerror = () => { script.remove(); reject(new Error('环境资源未加载')); };
        document.head.appendChild(script);
      }
    });
  }

  /** batchStatic 将同几何同材质静态网格合成实例；电表组与透明平面保持原对象。 */
  function batchStatic(scene) {
    const batches = new Map(); // 按几何、材质和阴影行为分组。
    scene.updateMatrixWorld(true);
    scene.traverse(mesh => {
      if (!mesh.isMesh || mesh.isInstancedMesh || Array.isArray(mesh.material) || mesh.material.transparent) return;
      let parent = mesh;
      while (parent) {
        if (parent.userData.meterId && !mesh.userData.staticDetail) return; // 避免破坏 raycast 与 LCD 刷新。
        parent = parent.parent;
      }
      const key = [mesh.geometry.uuid, mesh.material.uuid, mesh.castShadow, mesh.receiveShadow].join('/');
      if (!batches.has(key)) batches.set(key, []);
      batches.get(key).push(mesh);
    });
    batches.forEach(meshes => {
      if (meshes.length < 2) return;
      const first = meshes[0];
      const batch = new T.InstancedMesh(first.geometry, first.material, meshes.length);
      batch.castShadow = first.castShadow;
      batch.receiveShadow = first.receiveShadow;
      batch.frustumCulled = false; // r128 无实例包围盒，避免柜列移出原点后被错误剔除。
      meshes.forEach((mesh, index) => {
        batch.setMatrixAt(index, mesh.matrixWorld);
        mesh.parent.remove(mesh);
      });
      scene.add(batch);
    });
  }

  /** dispose 去重释放场景使用的几何、材质与纹理；不释放业务层对象。 */
  function dispose(scene) {
    const resources = new Set();
    scene.traverse(mesh => {
      if (mesh.geometry) resources.add(mesh.geometry);
      if (mesh.shadow && mesh.shadow.map) resources.add(mesh.shadow.map);
      if (mesh.material) {
        const list = Array.isArray(mesh.material) ? mesh.material : [mesh.material];
        list.forEach(material => {
          resources.add(material);
          Object.values(material).forEach(value => { if (value && value.isTexture) resources.add(value); });
        });
      }
    });
    resources.forEach(resource => resource.dispose());
  }

  window.StationFinish = { beveledBox, applySurface, cabinet, room, busJoints, busTerminal, environment, bakeEnvironment, batchStatic, dispose };
})();
