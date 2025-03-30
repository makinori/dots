#!/usr/bin/env -S deno run -A

import $, { CommandBuilder } from "jsr:@david/dax@0.41.0";
import * as path from "jsr:@std/path@0.225.2";
import axios from "npm:axios@1.7.9";

const apiKey = "";
const immichServer = "";

// actions for nautilus
/*
{
	"debug": false,
	"actions": [
		{
			"type": "command",
			"label": "Maki Immich",
			"use_shell": true,
			"command_line": "NAUTILUS=1 ~/maki-immich.ts \"%U\"",
			"min_items": 1,
			"filetypes": ["file"]
		}
	]
}
*/

function getRandom() {
	const buf = new Uint8Array(16 / 2);
	crypto.getRandomValues(buf);
	return Array.from(buf)
		.map(v => v.toString(16).padStart(2, "0"))
		.join("");
}

async function uploadFile(pathToFile: string, albumId: string) {
	// prepare file. strip exif cause immich prefers that

	let fileBuffer = new Uint8Array();
	try {
		fileBuffer = await Deno.readFile(pathToFile);
	} catch (error) {
		throw new Error("Failed to read file");
	}

	const filename = path.basename(pathToFile);

	try {
		// find more using: exiftool -G -a -s -time:all image.jpg

		const toStrip = [
			"AllDates",
			"CreateDate",
			"DateCreated",
			// "DateTimeCreated", // not writeable
			"DateTimeDigitized",
			"DateTimeOriginal",
			"GPSDateStamp",
			"GPSDateTime",
			"GPSTimeStamp",
			"ModifyDate",
			"TimeCreated",
			"DigitalCreationDate",
			"DigitalCreationTime",
			// "DigitalCreationDateTime", // not writeable
		];

		const exifStripDate = await new CommandBuilder()
			.command(["exiftool", ...toStrip.map(s => "-" + s + "="), "-"])
			.stdin(fileBuffer)
			.stdout("piped")
			.stderr("null") // inherit
			.spawn();

		fileBuffer = exifStripDate.stdoutBytes;
	} catch (error) {
		console.error(error);
	}

	// upload file which also handles deduplication

	const dateStr = new Date().toISOString();

	const formData = new FormData();
	formData.set("assetData", new Blob([fileBuffer]), filename);
	formData.set("deviceAssetId", getRandom());
	formData.set("deviceId", "DENO"); // WEB
	formData.set("fileCreatedAt", dateStr);
	formData.set("fileModifiedAt", dateStr);

	const upload = await axios(`${immichServer}/api/assets`, {
		method: "POST",
		data: formData,
		headers: {
			"x-api-key": apiKey,
			"Content-Type": "multipart/form-data",
		},
	});

	// add to album

	const addToAlbum = await axios(
		`${immichServer}/api/albums/${albumId}/assets`,
		{
			method: "PUT",
			data: {
				ids: [upload.data.id],
			},
			headers: {
				"x-api-key": apiKey,
			},
		},
	);

	const addToAlbumRes = addToAlbum.data[0];

	if (addToAlbumRes.success == false && addToAlbumRes.error != "duplicate") {
		throw new Error("Failed to add to album");
	}

	return { duplicate: addToAlbumRes.error == "duplicate" };
}

try {
	// get file and strip date

	if (Deno.args.length == 0) {
		throw new Error("Please select one or more files");
	}

	let filePaths = Deno.args;

	if (Deno.env.get("NAUTILUS") != null) {
		filePaths = filePaths[0]
			.split(" ")
			.reverse()
			.map(url => path.fromFileUrl(url));
	}

	// get album id

	const albums = await axios(`${immichServer}/api/albums`, {
		headers: {
			"x-api-key": apiKey,
		},
	});

	const args = albums.data
		.map((album: any) => {
			album.lastModifiedAsset = new Date(
				album.lastModifiedAssetTimestamp,
			);
			return album;
		})
		.sort((a: any, b: any) => b.lastModifiedAsset - a.lastModifiedAsset)
		.map((album: any) => [album.id, album.albumName, "off"])
		.flat();

	const albumSelect =
		await $`kdialog --radiolist ${"Select an album to upload to"} ${args}`
			.stdout("piped")
			.noThrow(true);

	const albumId = albumSelect.stdout.trim();
	const albumName = args[args.indexOf(albumId) + 1];

	if (albumSelect.code != 0 || albumId == "") {
		Deno.exit(0);
	}

	// upload files

	let completed: string[] = [];
	// let duplicate: string[] = [];
	let failed: string[] = [];

	for (const filePath of filePaths) {
		const filename = path.basename(filePath);

		try {
			const { duplicate } = await uploadFile(filePath, albumId);
			completed.push(filename + (duplicate ? " (duplicate)" : ""));
		} catch (error) {
			console.error(error);
			failed.push(filename);
		}
	}

	let finalMsg = `Added ${completed.length} to: ${albumName}`;

	if (completed.length > 0) {
		finalMsg += `\n${completed.join("\n")}`;
	}

	if (failed.length > 0) {
		finalMsg += `\n\nFailed:\n${failed.join("\n")}`;
	}

	await $`kdialog --msgbox ${finalMsg}`.noThrow(true);
} catch (error) {
	console.log(error);
	await $`kdialog --error ${error.message}`;
	Deno.exit(1);
}
