/* ============================================================
 * interaction.js — 射线拾取交互
 * 悬停高亮 + tooltip；单击选中（拖拽阈值防误触）；双击镜头聚焦
 * ============================================================ */

(function () {
  const T = THREE;

  window.setupInteraction = function (api, hooks) {
    const { camera, renderer, meterMeshes } = api;
    const raycaster = new T.Raycaster();
    const mouse = new T.Vector2();
    const dom = renderer.domElement;
    const removeListeners = []; // 离开页面时释放本模块的拾取监听。
    /** listen 记录绑定到渲染画布的监听器，便于统一销毁。 */
    function listen(type, handler) {
      dom.addEventListener(type, handler);
      removeListeners.push(() => dom.removeEventListener(type, handler));
    }

    let hoveredId = null, selectedId = null;
    let downPos = null;

    const hoverMats = new Map(); /* group -> [mesh, origMat] */

    function setGroupEmissive(group, on) {
      group.traverse(o => {
        if (o.isMesh && o.material && o.material.emissive) {
          if (on) {
            if (!hoverMats.has(o)) hoverMats.set(o, o.material);
            o.material = o.material.clone();
            o.material.emissive = new T.Color(0x1c6b2a);
            o.material.emissiveIntensity = 0.9;
          }
        }
      });
      if (!on) {
        /* 还原：LCD 贴图材质没有 emissive，不受影响；仅还原 body 材质 */
        group.traverse(o => {
          if (o.isMesh && o.material && o.material.emissive &&
              o.material.emissive.getHex() === 0x1c6b2a) {
            const orig = hoverMats.get(o) || M0();
            o.material.dispose();
            o.material = orig;
            hoverMats.delete(o);
          }
        });
      }
    }
    function M0() { return new T.MeshStandardMaterial({ color: 0x14181d }); }

    /* 选中指示：电表外的脉动边框 */
    let selRing = null;
    function showRing(group) {
      removeRing();
      /* 以电表局部坐标建框，避免柜组平移或旋转被重复应用。 */
      const body = group.children.find(child => child.isMesh); // 外壳定义选择范围。
      if (!body) return;
      body.geometry.computeBoundingBox();
      const bounds = body.geometry.boundingBox; // 几何坐标与电表组一致。
      const size = new T.Vector3(); bounds.getSize(size);
      const geometry = new T.BoxGeometry(size.x + 0.022, size.y + 0.022, size.z + 0.018);
      selRing = new T.LineSegments(new T.EdgesGeometry(geometry),
        new T.LineBasicMaterial({ color: 0xf0c879, toneMapped: false }));
      geometry.dispose();
      bounds.getCenter(selRing.position);
      selRing.name = 'meter-selection';
      group.add(selRing);
    }
    function removeRing() {
      if (selRing) { selRing.parent.remove(selRing); selRing.geometry.dispose(); selRing.material.dispose(); selRing = null; }
    }

    /* 电表清单只计算一次，悬停时避免重复分配数组。 */
    const pickGroups = Object.values(meterMeshes).map(value => value.group);
    function pick(e) {
      const r = dom.getBoundingClientRect();
      mouse.x = ((e.clientX - r.left) / r.width) * 2 - 1;
      mouse.y = -((e.clientY - r.top) / r.height) * 2 + 1;
      raycaster.setFromCamera(mouse, camera);
      const hits = raycaster.intersectObjects(pickGroups, true);
      for (const h of hits) {
        let o = h.object;
        while (o && !o.userData.meterId) o = o.parent;
        if (o) return o;
      }
      return null;
    }

    listen('pointermove', e => {
      const g = pick(e);
      const id = g ? g.userData.meterId : null;
      if (id !== hoveredId) {
        if (hoveredId && meterMeshes[hoveredId]) setGroupEmissive(meterMeshes[hoveredId].group, false);
        hoveredId = id;
        if (id) setGroupEmissive(meterMeshes[id].group, true);
        dom.style.cursor = id ? 'pointer' : 'default';
        hooks.onHover && hooks.onHover(id, e);
      } else if (id) {
        hooks.onHoverMove && hooks.onHoverMove(e);
      }
    });

    listen('pointerdown', e => { downPos = [e.clientX, e.clientY]; });

    listen('pointerup', e => {
      if (!downPos) return;
      const dx = e.clientX - downPos[0], dy = e.clientY - downPos[1];
      downPos = null;
      if (dx * dx + dy * dy > 25) return;      /* 拖拽视角，不算点击 */
      const g = pick(e);
      const id = g ? g.userData.meterId : null;
      if (id) {
        selectedId = id;
        showRing(meterMeshes[id].group);
        hooks.onSelect && hooks.onSelect(id);
      } else if (selectedId) {
        selectedId = null;
        removeRing();
        hooks.onDeselect && hooks.onDeselect();
      }
    });

    listen('dblclick', e => {
      const g = pick(e);
      if (g) {
        const wp = new T.Vector3();
        g.getWorldPosition(wp);
        const dir = new T.Vector3().subVectors(camera.position, api.controls.target).normalize();
        const p = wp.clone().add(dir.multiplyScalar(1.6));
        p.y = Math.max(p.y, wp.y + 0.5);
        api.flyTo({ pos: p.toArray(), tgt: wp.toArray() }, 800);
      }
    });

    return {
      /** dispose 释放边框、悬停材质和鼠标监听。 */
      dispose() {
        removeRing();
        if (hoveredId && meterMeshes[hoveredId]) setGroupEmissive(meterMeshes[hoveredId].group, false);
        removeListeners.forEach(remove => remove());
        hoverMats.clear();
      },
      getSelected() { return selectedId; },
      select(id, fly) {
        selectedId = id;
        if (id && meterMeshes[id]) {
          showRing(meterMeshes[id].group);
          if (fly) {
            const wp = new T.Vector3();
            meterMeshes[id].group.getWorldPosition(wp);
            const dir = new T.Vector3().subVectors(camera.position, api.controls.target).normalize();
            const p = wp.clone().add(dir.multiplyScalar(2.2));
            p.y = Math.max(p.y, wp.y + 0.8);
            api.flyTo({ pos: p.toArray(), tgt: wp.toArray() }, 900);
          }
        } else removeRing();
      },
      deselect() { selectedId = null; removeRing(); }
    };
  };
})();
