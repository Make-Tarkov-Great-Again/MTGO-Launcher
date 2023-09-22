import { Check } from "../../../wailsjs/go/launcher/Online";
import { Question } from "../../../wailsjs/go/launcher/UI";
let elapsedTimeInSeconds = 0;
let stopElapsedTime = false;

var downloading = false
let wasOffline;
var count = 0
var failedCount = 0



const downloadProgress = document.querySelector('.download-progress');
const progressBar = document.querySelector('.progress-bar');
const label = document.querySelector('.progress-label')
let online;
online = checkOnlineStatus()
console.log("first online done")

async function OnlineCheck() {
    return Check()
}

async function checkOnlineStatus() {
    online = await OnlineCheck();
    console.log(online)
    if (!online) {
        label.textContent = "Your in offline mode"
        label.style.color = "red"
        wasOffline = true
    }
    else {
        if (downloading) {
            return
        }
        if (wasOffline && online) {
            if (document.hasFocus) {
                label.style.color = "green"
                label.textContent = "Connection restored"
                wasOffline = false
            } else {
                const focus = () => { //well this was a waste, wails dont have support for focus lol. but it was suggested 2 days ago (as of 9/11/23) so leaving this here.
                    label.style.color = "green"
                    label.textContent = "Connection restored"
                    window.removeEventListener("focus", focus);
                    wasOffline = false
                }
                window.addEventListener("focus", focus);

            }


        }

    }
}
setInterval(checkOnlineStatus, 60000);





window.addEventListener('online', () => {
    console.log("Came online")
    checkOnlineStatus()
});

window.addEventListener('offline', () => {
    console.log("Went offline")
    checkOnlineStatus()
});

document.addEventListener('DOMContentLoaded', () => {
    downloadProgress.addEventListener('mouseenter', () => {
        if (online) {
            downloadProgress.style.bottom = '5px';
            label.style.opacity = "100%"
        }
    });


    downloadProgress.addEventListener('mouseleave', () => {
        if (online) {
            if (!downloading) {
                downloadProgress.style.bottom = '-22px';
            } else { downloadProgress.style.bottom = '-21px' }
            label.style.opacity = "15%"
        }
    });

    downloadProgress.addEventListener('mousemove', (e) => {
        if (online) {
            const mouseY = e.clientY;
            const progressRect = progressBar.getBoundingClientRect();
            const progressTop = progressRect.top;
            const progressBottom = progressRect.bottom;
        }
    });
    label.addEventListener('click', () => {
        if (!downloading) {
            label.textContent = "No downlaods"
        }
    });
})

const socket = new WebSocket("ws://localhost:42145/ws"); // Update the URL accordingly

socket.onopen = () => {
    console.log("WebSocket is running.");
};
socket.onmessage = (event) => {
    const data = JSON.parse(event.data);

    if (data.step === 'Error') {
        failedCount++
        label.textContent = `Failed to download ${failedCount} mod(s), Downloaded ${count} mod(s)`;
        return;
    }

    if (data && data.percentage !== undefined) {
        downloading = true;
        if (data.percentage > 1) {
            downloadProgress.style.bottom = "-21px"
        }
        const progressBar = document.querySelector('.progress-bar');
        progressBar.value = data.percentage;

        const label = document.querySelector('.progress-label');
        const convertedsize = data.filesize / 1000000
        if (data.percentage < 100) {
            label.textContent = `Downloading ${data.modName} (${convertedsize} MB)`;
        }
        if (typeof data.filename === 'string' && !['.zip', '.rar', '.tar', '.tar.bz2', '.tar.gz', '.tar.lz4', '.tar.sz', '.tar.xz', '.7z'].some(ext => data.filename.endsWith(ext))) {
            count++
            if (count = 1) {
                label.textContent = `${count} mod downloaded`;
            } else {
                label.textContent = `${count} mods downloaded`;
            }
        }
    } else if (data.step === 'Extracting') {
        label.textContent = `Installing & Parsing ${data.modName}`;
        if (elapsedTimeInSeconds <= 1) {
            elapsedTime(data.filename);
        }
    }
    else if (data.step === 'Done') {
        count++
        stopElapsedTime = true;
        if (count = 1) {
            label.textContent = `${count} mod downloaded`;
        } else {
            label.textContent = `${count} mods downloaded`;
        }
    }
    else if (data.type === 'Error') {
    }
};



socket.onclose = () => {
    console.log("WebSocket is not running.");
};

export function initiateDownload(modID, fileURL, modName, modAuthor) {
    console.log(`Downloading ${modName}`)
    if (!online) {
        return;
    }
    if (localStorage.getItem("Agreed") !== true) {
        const agree = Question("License Agreement", "Whoa, hold up there, buckaroo. This is your first time downloading a mod. While that's great and all, these mods have licenses that you have to agree to.\n\nAll these mod licenses can be found on the mod page, and clicking on them will show you the can-dos and can't-dos with these mods.\n\nBy continuing, you agree to this, and all future mod licenses you may or may not download. Do you understand?")
        if (agree) {
            localStorage.setItem("Agreed", true)
        } else {
            return;
        }
    }
    const message = {
        type: 'download',
        modID: modID,
        fileURL: fileURL,
        modName: modName,
        modAuthor: modAuthor
    };
    socket.send(JSON.stringify(message));
}