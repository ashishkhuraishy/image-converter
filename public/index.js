async function initWebAssembly(params) {
    const go = new Go();
    const response = await fetch("test.wasm");
    const buffer = await response.arrayBuffer();
    const output = await WebAssembly.instantiate(buffer, go.importObject);
    go.run(output.instance)

    update()

}

function selectImage(arrayList) {
    console.log(arrayList)
}

function imageIsLoaded(e) {
    console.log("image is loaded")

    $('#myImg').attr('src', e.target.result);
};

function onInputChange(image) {
    var input = image;
    if (input.files && input.files[0]) {
        var reader = new FileReader()
        reader.onload = imageIsLoaded
        reader.readAsDataURL(input.files[0])
    } else {
        $('#myImg').attr('src', '/assets/no_preview.png')
    }
}



initWebAssembly()