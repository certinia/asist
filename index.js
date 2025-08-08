#!/usr/bin/env node

"use strict";
const { spawn } = require("child_process");
const os = require("os");

let bin = "";
let platarch = `${os.platform()}/${os.arch()}`;
switch(platarch) {
	case "win32/x64":
		bin = require.resolve("./asist.exe");
		break;
	case "win32/arm64":
		bin = require.resolve("./asist.exe");
		break;
	case "darwin/x64":
		bin = require.resolve("./asist_darwin_amd64");
		break;
	case "darwin/arm64":
		bin = require.resolve("./asist_darwin_arm64");
		break;
	case "linux/arm64":
		bin = require.resolve("./asist_linux_arm64");
		break;
	default: // "linux/amd64" or other OS
		bin = require.resolve("./asist_linux_amd64");
		break;
}

var child = spawn(bin, process.argv.slice(2), { shell: false });
let output = "";

child.stdout.on("data", (data) => {
	output += data;
});

child.stderr.on("data", (data) => {
	console.error(`child stderr:\n${data}`);
});

child.on("close", (code) => {
	if(output != "") {
		console.log(output);
	}
	if (code !== 0) {
		console.error(`ps process exited with code ${code}`);
		// to exit parent process
		process.exit(code);
	}
});
