const path = require("path");
const fs = require("fs-extra");
require("./common/wasm_exec");

class Siegfried {
  constructor() {
    this.instance = null;
    this.module = null;
  }

  async load() {
    const source = fs.readFileSync(path.resolve(__dirname, "../sf.wasm"));
    const go = new globalThis.Go();
    const typedArray = new Uint8Array(source);
    const { instance, module } = await WebAssembly.instantiate(
      typedArray.buffer,
      go.importObject
    );
    this.instance = instance;
    this.module = module;

    go.run(this.instance);
  }

  identify(path) {
    const file = fs.readFileSync(path);
    const buf = new Uint8Array(file);
    const result = globalThis.identify(buf, path);
    return result;
  }
}

module.exports = Siegfried;
