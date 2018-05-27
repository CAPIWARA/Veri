import sleep from '@/helpers/sleep';

const SECOND = 1000;

function get () {
  return new Promise((resolve) => {
    window.navigator.geolocation.getCurrentPosition((position) => resolve({
      time: position.timestamp,
      altitude: position.coords.altitude,
      latitude: position.coords.latitude,
      longitude: position.coords.longitude
    }));
  });
}

async function initialize () {
  await sleep(5 * SECOND);
  const answer = await window.navigator.permissions.query({
    name: 'geolocation'
  });
  return get();
}

function watch (fn, time = 8 * SECOND) {
  const watcher = window.setInterval(() => get().then(function () {
    if (!watcher)
      return;
    fn(...arguments);
  }), time);
  return watcher;
}

function unwatch (watcher) {
  window.clearInterval(watcher);
}

export { get, initialize, watch, unwatch };
