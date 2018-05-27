const not = (A) => (B) => !Object.is(A, B);

const SECOND = 1000;

class PositionWatcher {
  interval = null;
  watchers = [];

  broadcast () {
    window.navigator.geolocation.getCurrentPosition((position) => {
      this.watchers.map((watcher) => watcher({
        time: position.timestamp,
        coords: {
          altitude: position.coords.altitude,
          latitude: position.coords.latitude,
          longitude: position.coords.longitude
        }
      }));
    });
  }

  async start () {
    const { state = 'granted' } = await window.navigator.permissions.query({ name: 'geolocation' });

    if (state !== 'granted')
      throw new Error('O usuário rejeitou a solicitação pela API de geolocalização.');

    this.interval = setInterval(this.broadcast, 3 * SECOND);
  }

  register (watcher) {
    this.watchers = this.watchers.concat(watcher);
  }

  unregister (watcher) {
    this.watchers = this.watchers.filter(not(watcher));
  }
}

export default new PositionWatcher();
