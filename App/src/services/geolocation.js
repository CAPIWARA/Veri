import sleep from '@/helpers/sleep';
import { discal as distance } from 'geo-dist-calc';

const SECOND = 1000;

function get () {
  return new Promise((resolve) => {
    navigator.geolocation.getCurrentPosition((position) => resolve({
      time: position.timestamp,
      altitude: position.coords.altitude,
      latitude: position.coords.latitude,
      longitude: position.coords.longitude
    }));
  });
}

async function initialize () {
  await sleep(5 * SECOND);
  const { state } = await navigator.permissions.query({
    name: 'geolocation'
  });
  return (state === 'granted');
}

function watch (fn, time = 8 * SECOND) {
  const watcher = setInterval(() => get().then((...params) => watcher && fn(...params)), time);
  return watcher;
}

function unwatch (watcher) {
  clearInterval(watcher);
}

function distances (coordinates) {
  let total = 0;

  coordinates.reduce((A, B) => {
    if (A && B) {
      const { kilometers = 0 } = distance(A, B);
      total += kilometers;
      return A;
    }
    return A;
  });

  return total;
}

export { get, initialize, watch, unwatch, distances, distance };
