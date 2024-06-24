const go = new Go();
let wasmFunctions = {};

let currentInput = "0";
let previousInput = "";
let currentOperation = null;

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
    wasmFunctions.add = add;
    wasmFunctions.subtract = subtract;
    wasmFunctions.multiply = multiply;
    wasmFunctions.divide = divide;
  }
);

function updateDisplay() {
  document.getElementById("currentInput").textContent = currentInput;
  document.getElementById("previousInput").textContent = previousInput;
}

function appendNumber(number) {
  if (currentInput === "0") {
    currentInput = number.toString();
  } else {
    currentInput += number.toString();
  }
  updateDisplay();
}

function setOperation(operation) {
  if (currentOperation !== null) calculate();
  currentOperation = operation;
  previousInput = currentInput + " " + getOperationSymbol(operation);
  currentInput = "0";
  updateDisplay();
}

function getOperationSymbol(operation) {
  switch (operation) {
    case "add":
      return "+";
    case "subtract":
      return "-";
    case "multiply":
      return "×";
    case "divide":
      return "÷";
  }
}

function clearCalculator() {
  currentInput = "0";
  previousInput = "";
  currentOperation = null;
  updateDisplay();
}

function calculate() {
  if (currentOperation === null || previousInput === "") return;

  const prev = parseInt(previousInput);
  const current = parseInt(currentInput);
  let result;

  switch (currentOperation) {
    case "add":
      result = wasmFunctions.add(prev, current);
      break;
    case "subtract":
      result = wasmFunctions.subtract(prev, current);
      break;
    case "multiply":
      result = wasmFunctions.multiply(prev, current);
      break;
    case "divide":
      if (current === 0) {
        result = "Error";
      } else {
        result = wasmFunctions.divide(prev, current);
      }
      break;
  }

  currentInput = result.toString();
  previousInput = "";
  currentOperation = null;
  updateDisplay();
  document.getElementById("currentInput").classList.add("animate");
  setTimeout(
    () => document.getElementById("currentInput").classList.remove("animate"),
    500
  );
}
