'use strict';

var wasm;
var logLine = [];
var memory8;

const decoder = new TextDecoder("utf-8");

const mem = () => {
  return new DataView(wasm.exports.memory.buffer);
}

const loadString = (addr, len) => {
  return decoder.decode(new DataView(wasm.exports.memory.buffer, addr, len));
}

var importObject = {
  env: {
    io_get_stdout: function () {
      return 1;
    },
    resource_write: function (fd, ptr, len) {
      if (fd == 1) {
        for (let i = 0; i < len; i++) {
          let c = memory8[ptr + i];
          if (c == 13) { // CR
            // ignore
          } else if (c == 10) { // LF
            // write line
            let line = new TextDecoder("utf-8").decode(new Uint8Array(logLine));
            logLine = [];
            console.log(line);
          } else {
            logLine.push(c);
          }
        }
      } else {
        console.error('invalid file descriptor:', fd);
      }
    },
    message: function (msg, addr, len) {
      if (msg === 0) {
        const body = loadString(addr, len);
        document.getElementById("app").innerHTML = body;
      }
    }
  },
};



const wasmFile = fetch('wasm.wasm');
WebAssembly.instantiateStreaming(wasmFile, importObject).then(function (obj) {
  wasm = obj.instance;
  memory8 = new Uint8Array(wasm.exports.memory.buffer);
  wasm.exports.cwa_main();
})

document.getElementById("app").onmousemove = function(event) {
  if (wasm) {
    wasm.exports.mouse(event.x, event.y)
  }
}