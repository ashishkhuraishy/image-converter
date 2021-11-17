async function initWebAssembly(params) {
    const go = new Go();
    const response = await fetch("main.wasm");
    const buffer = await response.arrayBuffer();
    const output = await WebAssembly.instantiate(buffer, go.importObject);
    go.run(output.instance)
}

function selectImage(arrayList) {
    console.log(arrayList)
}

function imageIsLoaded(e) {
    console.log("image is loaded")
    loadImage(e.target.result)

    $('#myImg').attr('src', e.target.result);
};

function onInputChange(image) {
    var input = image;
    if (input.files && input.files[0]) {
        var reader = new FileReader()
        reader.onload = imageIsLoaded
        reader.readAsDataURL(input.files[0])
    } 
}



initWebAssembly().then(_ => update())
