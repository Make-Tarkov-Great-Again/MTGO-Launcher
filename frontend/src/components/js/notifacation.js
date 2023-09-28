import mitt from 'mitt';

const eventEmitter = mitt();
console.log("bruh moment")

function sendNotification(title, message, type, from) {
    const newNotification = {
        title,
        message,
        type,
        from,
    };

    for (const key in newNotification) {
        if (newNotification[key] === undefined || newNotification[key] === "") {
            return "Missing info";
        }
    }

    const existingNotifications = JSON.parse(localStorage.getItem('notifications')) || [];
    existingNotifications.push(newNotification);

    localStorage.setItem('notifications', JSON.stringify(existingNotifications));

    eventEmitter.emit('notifications-changed', existingNotifications);
}

function checkTimers() {
    const scavTimers = JSON.parse(localStorage.getItem('ScavTimers')) || {};

    for (const username in scavTimers) {
        if (scavTimers.hasOwnProperty(username)) {
            const epochTimestamp = scavTimers[username];

            if (hasTimerExpired(epochTimestamp)) {
                const timeDifference = Math.floor(Date.now() / 1000) - epochTimestamp;

                if (timeDifference < 5) {
                    playSound();

                    delete scavTimers[username];

                    localStorage.setItem('ScavTimers', JSON.stringify(scavTimers));
                } else {
                    delete scavTimers[username];
                    localStorage.setItem('ScavTimers', JSON.stringify(scavTimers));
                }
            }
        }
    }
}
function hasTimerExpired(epochTimestamp) {
    const currentTime = Math.floor(Date.now() / 1000);
    return currentTime >= epochTimestamp;
}

function playSound() {
    var audio = new Audio('/src/assets/sounds/scav_ready.mp3');
    audio.play();
}

eventEmitter.on('new-notification', (notification) => {
    let newestNotificationType;
    console.log(notification)
    if (notification.length > 1) {
        newestNotificationType = notification[notification.length - 1].type;
    } else { newestNotificationType = notification.type }
    console.log(newestNotificationType)
    playSoundBasedOnType(newestNotificationType);
});

function playSoundBasedOnType(notificationType) {
    let soundPath;
    console.log("sex", notificationType)

    switch (notificationType) {
        case 'scavTimers':
            soundPath = '/src/assets/sounds/scav_ready.mp3';
            break;
        case 'normal':
            soundPath = '/src/assets/sounds/noti.mp3';
            break;
        case 'system':
            soundPath = '/src/assets/sounds/system.mp3';
            break;
        case 'sus':
            soundPath = '/src/assets/sounds/sus.mp3';
            break;
        default:
            soundPath = '/src/assets/sounds/noti.mp3';
    }

    const audio = new Audio(soundPath);
    audio.play();
}

export default eventEmitter;
