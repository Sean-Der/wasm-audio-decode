# wasm-audio-decode

This page demonstrates using [pion/opus](https://github.com/pion/opus) to decode in WASM. Those decoded buffers are played using the [Web Audio API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Audio_API)

The source code for this project is available at [github.com/sean-der/wasm-audio-decode](https://github.com/sean-der/wasm-audio-decode)

The files are encoded at ~10K and are telephony quality at best.

Thank you Gerard Webb for suggesting this project in [pion/opus#26](https://github.com/pion/opus/issues/26)

### Why?

Hopefully this project can serve as an interesting building block for others. Building + Running Go (via WASM) in the browser is incredibly easy.

As more audiovisual libraries become available in Go it could unlock things that weren't possible before.

For a long time I have wanted to learn more about audio codecs and sound processing in general. It was fun to tie this all together and finally be able
to have something 'demoable' of pion/opus.

I also hope that having codecs in higher level languages makes it easier for others to learn.

### Building

I used TinyGo, but may not be a requirement. `ran` is a HTTP server that can be used to serve locally.


* `cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .`
* `tinygo build -o wasm.wasm -target wasm  -no-debug --panic trap`
* `go install github.com/m3ng9i/ran@latest`
* `ran`
* `http://localhost:8080`

