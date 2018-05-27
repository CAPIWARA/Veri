function toTime (time) {
  const SECOND = 1000;
  const MINUTE = SECOND * 60;
  const HOUR   = MINUTE * 60;

  const zero = (value) => value.toString().padStart(2, '0');

  const HH = zero(~~(time / HOUR));
  const MM = zero(~~((time % HOUR) / MINUTE));
  const SS = zero(~~((time % HOUR % MINUTE) / SECOND));

  return `${HH}:${MM}:${SS}`;
}

function toDistance (distance) {
  return (+distance).toFixed(2);
}

function toCoins (coins) {
  return ~~coins;
}

export { toTime, toDistance, toCoins };


