import "./wasm_exec"

export default async function initWasm() {
    const go = new Go();
    const { instance } = await WebAssembly.instantiateStreaming(
        fetch('mortgage.wasm'),
        go.importObject
    );
    go.run(instance);
}