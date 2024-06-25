const go = new Go();

// WebAssembly コードをコンパイル、インスタンス化するためのAPI
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  }
);

// 入力された数値を表示する関数
function updateDisplay() {
  document.getElementById("base").textContent = getBase();
  document.getElementById("log").textContent = getLog();

  // 出力時のアニメーション
  const baseElement = document.getElementById("base");
  baseElement.classList.add("animate");
  setTimeout(() => baseElement.classList.remove("animate"), 500);
}

// 数値ボタンが押された時の処理（ハンドラ）
function handleClick(funcName, arg) {
  // windowオブジェクトの関数を呼び出す
  if (arg !== undefined) {
    window[funcName](arg);
  } else {
    window[funcName]();
  }
  updateDisplay();
}
