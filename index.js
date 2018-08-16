const Gpio = require('pigpio').Gpio
    , CronJob = require('cron').CronJob
    , Red = new Gpio(27, { mode: Gpio.OUTPUT })
    , Green = new Gpio(17, { mode: Gpio.OUTPUT })
    , Blue = new Gpio(22, { mode: Gpio.OUTPUT });

function tTimeout(timeout) {
    return new Promise((resolve) => {
        setTimeout(resolve, timeout);
    })
}

async function wakeup() {
    for (let i = 0; i < 255; i++) {
        Red.pwmWrite(i);
        Green.pwmWrite(i);
        Blue.pwmWrite(i);

        await tTimeout(14117);
    }
}

new CronJob('00 30 5 * * 1-5', wakeup, () => {}, true, 'America/New_York');
