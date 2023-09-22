
async function elapsedTime(fileName) {

    const minutes = Math.floor(elapsedTimeInSeconds / 60);
    const seconds = elapsedTimeInSeconds % 60;
    const elapsedTimeFormatted = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    label.textContent = `Installing & Parsing ${fileName} | ${elapsedTimeFormatted}`;
    elapsedTimeInSeconds++;
    await new Promise(resolve => setTimeout(resolve, 1000));
    elapsedTime(fileName);
}



// Function to check if a timer has expired
function hasTimerExpired(epochTimestamp) {
    const currentTime = Math.floor(Date.now() / 1000); // Convert current time to seconds
    return currentTime >= epochTimestamp;
}

// Function to play a sound when a timer expires
function playSound() {
    var audio = new Audio('/src/assets/sounds/scav_ready.mp3'); // path to file
    audio.play();
}

// Function to periodically check timers
    function checkTimers() {
        // Iterate over the "ScavTimers" array in local storage
        const scavTimers = JSON.parse(localStorage.getItem('ScavTimers')) || {};

        for (const username in scavTimers) {
            if (scavTimers.hasOwnProperty(username)) {
                const epochTimestamp = scavTimers[username];

                if (hasTimerExpired(epochTimestamp)) {
                    // Calculate the time difference in seconds
                    const timeDifference = Math.floor(Date.now() / 1000) - epochTimestamp;

                    // Check if the timer has expired by more than 5 seconds
                    if (timeDifference < 5) {
                        // Timer has expired by more than 5 seconds, play a sound
                        playSound();

                        // Optionally, you can remove the expired timer from the array
                        delete scavTimers[username];

                        // Update the "ScavTimers" array in local storage
                        localStorage.setItem('ScavTimers', JSON.stringify(scavTimers));
                    }
                    else {
                        delete scavTimers[username];
                        localStorage.setItem('ScavTimers', JSON.stringify(scavTimers));
                    }

                }
            }
        }
    }

// Periodically check timers every 5 seconds (5000 milliseconds)
setInterval(checkTimers, 1000);
