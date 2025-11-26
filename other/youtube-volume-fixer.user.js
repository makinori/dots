// ==UserScript==
// @name         YouTube Volume Fixer
// @namespace    http://tampermonkey.net/
// @version      1
// @description  Makes the YouTube volume slider exponential for easier adjustment of lower volumes.
// @icon         https://www.gstatic.com/youtube/img/branding/favicon/favicon_144x144.png
// @match        https://www.youtube.com/*
// @match        https://music.youtube.com/*
// @run-at       document-start
// @grant        none
// @license MIT
// @downloadURL https://update.greasyfork.org/scripts/487324/YouTube%20Volume%20Fixer.user.js
// @updateURL https://update.greasyfork.org/scripts/487324/YouTube%20Volume%20Fixer.meta.js
// ==/UserScript==

// original https://greasyfork.org/en/scripts/487324-youtube-volume-fixer

(function () {
	"use strict";

	// was 3, might need to use 2
	const EXPONENT = 2.5;

	const storedOriginalVolumes = new WeakMap();
	const { get, set } = Object.getOwnPropertyDescriptor(
		HTMLMediaElement.prototype,
		"volume",
	);

	Object.defineProperty(HTMLMediaElement.prototype, "volume", {
		get() {
			const lowVolume = get.call(this);
			const calculatedOriginalVolume = lowVolume ** (1 / EXPONENT);

			const storedOriginalVolume = storedOriginalVolumes.get(this);
			const storedDeviation = Math.abs(
				storedOriginalVolume - calculatedOriginalVolume,
			);

			const originalVolume =
				storedDeviation < 0.01
					? storedOriginalVolume
					: calculatedOriginalVolume;

			return originalVolume;
		},
		set(originalVolume) {
			const lowVolume = originalVolume ** EXPONENT;
			storedOriginalVolumes.set(this, originalVolume);
			set.call(this, lowVolume);
		},
	});
})();
