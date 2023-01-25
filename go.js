const go = new Go();

let mod, wasm;

WebAssembly.instantiateStreaming(fetch("wasm.wasm"),go.importObject).then(
    async result => {
    mod = result.module;
    wasm = result.instance;
    await go.run(wasm);
});