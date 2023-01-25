const go = new Go();

let mod, wasm;

WebAssembly.instantiateStreaming(fetch("wasm.wasm"),go.importObject).then(
    result => {
    mod = result.module;
    wasm = result.instance;
});


async function run() {
    console.clear();
    await go.run(wasm);
    wasm = await WebAssembly.instantiate(mod,go.importObject);
}